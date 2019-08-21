package generator

import (
	"errors"
	"go/ast"
	"go/token"
	"os"
	"strings"
	"text/template"

	"github.com/PennState/additional-properties/pkg/aputil"
	"github.com/selesy/genutil/pkg/genutil"
	log "github.com/sirupsen/logrus"
)

const apTag = "*"

var apFilter = func(n ast.Node) bool {
	_, st, ok := getStructType(n)
	if !ok {
		return false
	}
	_, ok = getApField(st)
	if !ok {
		return false
	}
	return true
}

func getStructType(n ast.Node) (*ast.TypeSpec, *ast.StructType, bool) {
	ts := &ast.TypeSpec{}
	st := &ast.StructType{}
	genDecl, ok := n.(*ast.GenDecl)
	if !ok {
		return ts, st, false
	}
	if genDecl.Tok != token.TYPE {
		return ts, st, false
	}
	if len(genDecl.Specs) != 1 {
		return ts, st, false
	}
	ts, ok = genDecl.Specs[0].(*ast.TypeSpec)
	if !ok {
		return ts, st, false
	}
	st, ok = ts.Type.(*ast.StructType)
	return ts, st, ok
}

func getApField(st *ast.StructType) (*ast.Field, bool) {
	for _, f := range st.Fields.List {
		t, ok := aputil.NewTagFromField(*f)
		if !ok {
			continue
		}
		if t.Name == apTag {
			return f, true
		}
	}
	return nil, false
}

type FileSpec struct {
	GoFile string
	Pkg    string
	Code   []CodeSpec
}

type CodeSpec struct {
	TypeName string
	APName   string
	VarName  string
	Fields   []FieldSpec
}

type FieldSpec struct {
	FieldName string
	JsonName  string
}

func Run() error {
	aps, err := findTargets()
	if err != nil {
		return err
	}
	for _, ap := range aps {
		err := generate(ap)
		if err != nil {
			log.Error(err)
		}
	}
	return nil
}

func findTargets() ([]FileSpec, error) {
	matches, err := genutil.FilterAstNodesFromArgs(apFilter)
	if err != nil {
		return nil, err
	}
	if len(matches) == 0 {
		return nil, errors.New("No targets found")
	}
	specs := map[string]FileSpec{}
	for _, match := range matches {
		fs, ok := specs[match.GoFile]
		if !ok {
			fs = FileSpec{
				GoFile: match.GoFile,
				Pkg:    match.File.Name.Name,
				Code:   []CodeSpec{},
			}
		}
		ts, st, ok := getStructType(match.Node)
		if !ok {
			return nil, errors.New("This shouldn't happen if the filter is working")
		}
		log.Info("Target found in file \"", match.GoFile, "\" - struct named: ", ts.Name.Name)
		ap, ok := getApField(st)
		if !ok {
			return nil, errors.New("This shouldn't happen if the filter is working")
		}
		cs := CodeSpec{
			TypeName: ts.Name.Name,
			APName:   ap.Names[0].Name,
			VarName:  strings.ToLower(ts.Name.Name[0:1]),
			Fields:   []FieldSpec{},
		}
		for _, f := range st.Fields.List {
			name := f.Names[0].Name
			if name == cs.APName {
				continue
			}
			cs.Fields = append(cs.Fields, FieldSpec{
				FieldName: name,
				JsonName:  aputil.GetJSONName(f),
			})
		}
		fs.Code = append(fs.Code, cs)
		specs[match.GoFile] = fs
	}
	s := []FileSpec{}
	for _, v := range specs {
		s = append(s, v)
	}
	return s, nil
}

func generate(spec FileSpec) error {
	tmpl := template.New("file")
	tmpl.Funcs(template.FuncMap{
		"lower": strings.ToLower,
	})
	err := genutil.LoadTemplates(
		tmpl,
		fileTmpl,
		codeTmpl,
		marshalTmpl,
		unmarshalTmpl,
	)
	if err != nil {
		return err
	}

	genFileName, err := genutil.GeneratedGoFileName(spec.GoFile)
	if err != nil {
		return errors.New("This shouldn't happen since AST handed us a .go file")
	}
	log.Info("Generating file : ", genFileName)
	genFile, err := os.Create(genFileName)
	if err != nil {
		return err
	}
	defer genFile.Close()
	return tmpl.ExecuteTemplate(genFile, "file", spec)
}
