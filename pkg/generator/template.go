package generator

const fileTmpl = `
{{ define "file" -}}
// Code generated by additional-properties DO NOT EDIT.

package {{ .Pkg }}

import (
	"encoding/json"
	"strings"
)
{{ range .Code -}}
    {{ template "code" . -}}
{{ end -}}
{{ end -}}
`

const codeTmpl = `
{{ define "code" -}}
{{ template "marshal" . -}}
{{ template "unmarshal" . -}}
{{ end -}}
`

const marshalTmpl = `
{{ define "marshal" }}
// MarshalJSON encodes the {{ .TypeName }} struct to JSON with additional-properties
func ({{ .VarName }} {{ .TypeName }}) MarshalJSON() ([]byte, error) {
	type Alias {{ .TypeName }}
	aux := (Alias)({{ .VarName }})
	if aux.{{ .APName }} == nil {
		aux.{{ .APName }} = map[string]interface{}{}
	}
	{{ $ap := .APName -}}
	{{ range .Fields -}}
	aux.{{ $ap }}["{{ .JsonName }}"] = aux.{{ .FieldName }}
	{{ end -}}
	return json.Marshal(aux.{{ .APName }})
}
{{ end }}
`

const unmarshalTmpl = `
{{ define "unmarshal" }}
// UnmarshalJSON decodes JSON into the {{ .TypeName }} struct with additional-properties
func ({{ .VarName }} *{{ .TypeName }}) UnmarshalJSON(data []byte) error {
	type Alias {{ .TypeName }}
	aux := (*Alias)({{ .VarName }})
	err := json.Unmarshal(data, &aux)
	if err != nil {
		return err
	}
	_ = json.Unmarshal(data, &{{ .VarName }}.{{ .APName }})
	names := map[string]bool{
		{{- range .Fields }}
		{{ $lowerName := .JsonName | lower -}}
		"{{ .JsonName }}": true,{{ if ne .JsonName $lowerName }} "{{ $lowerName }}": true,{{ end -}}
	{{/* DO NOT REMOVE */ -}}
	{{ end }}
	}
	for k := range {{ .VarName }}.{{ .APName }} {
		if names[k] {
			delete({{ .VarName }}.{{ .APName }}, k)
			continue
		}
		if names[strings.ToLower(k)] {
			delete({{ .VarName }}.{{ .APName }}, k)
		}
	}
	if len({{ .VarName }}.{{ .APName }}) == 0 {
		{{ .VarName }}.{{ .APName }} = nil
	}
	return nil
}
{{ end }}
`