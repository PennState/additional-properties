package json

import (
	"encoding/json"
	"errors"
	log "github.com/sirupsen/logrus"
	"reflect"
	"strings"
)

const (
	TagKey                                    = "json"
	NotAStructMessage                         = "Only struct is accepted as an argument to this method"
	NotAMapStringJsonRawMessageMessage        = "Additional properties field must be map[string]json.RawMessage"
	AdditionalPropertiesMustBeExportedMessage = "Additional properties field must be exported"
)

func Marshal(v interface{}) ([]byte, error) {
	//Types that do not contain elements can be directly handled by the
	//standard library's JSON marshaler.
	k := dereferencedKind(v)
	if !hasElem(k) {
		return json.Marshal(v)
	}

	if k == reflect.Struct {
		return marshalStruct(v)
	}

	//TODO: Marshal the additional properties field as defined by the
	//struct's field tags

	// 	//Marshal the map that now contains all the struct's fields plus the
	// 	//original additional properties
	// 	return json.Marshal(r.getAdditionalProperties())

	//TODO: Get rid of this when the rest of the library is converted and
	//tested
	return nil, errors.New("Not yet implemented")
}

// func Unmarshal(data []byte, v resource) error {
// 	t := dereference(reflect.TypeOf(v))
// 	if t.Kind() != reflect.Struct {
// 		return json.Unmarshal(data, v)
// 	}
// 	return unmarshalResource(data, v)
// }

func dereferencedKind(v interface{}) reflect.Kind {
	return dereferencedType(v).Kind()
}

func dereferencedType(v interface{}) reflect.Type {
	t := reflect.TypeOf(v)
	return dereferencedTypeRecursion(t)
}

//Array, Chan, Map, Ptr, or Slice
func dereferencedTypeRecursion(t reflect.Type) reflect.Type {
	k := t.Kind()
	if hasElem(k) {
		return dereferencedTypeRecursion(t.Elem())
	}
	return t
}

//hasElem indicates that the Kind passed as a parameter has an element.
//See https://golang.org/pkg/reflect/#Type Elem()
func hasElem(k reflect.Kind) bool {
	if k == reflect.Array || k == reflect.Chan || k == reflect.Map || k == reflect.Ptr || k == reflect.Slice {
		return true
	}
	return false
}

// Invalid Kind = iota
// Bool sl
// Int sl
// Int8 sl
// Int16 sl
// Int32 sl
// Int64 sl
// Uint sl
// Uint8 sl
// Uint16 sl
// Uint32 sl
// Uint64 sl
// Uintptr sl
// Float32 sl
// Float64 sl
// Complex64 sl
// Complex128 sl
// Array ap for struct or interface elements sl for others
// Chan panic
// Func panic
// Interface
// Map
// Ptr
// Slice
// String
// Struct
// UnsafePointer

func marshalStruct(v interface{}) ([]byte, error) {
	ap, err := additionalPropertiesField(v)
	if err != nil {
		return nil, err
	}

	//Iterate over the individual fields
	st := reflect.TypeOf(v)  //.Elem()
	sv := reflect.ValueOf(v) //.Elem()
	for i := 0; i < st.NumField(); i++ {
		ft := st.Field(i)

		n, ok := jsonName(ft)
		if !ok {
			continue
		}

		//Unexported fields should be skipped
		fv := sv.Field(i)
		//if !fv.CanAddr() || !fv.CanInterface() {
		if !fv.CanInterface() {
			continue
		}

		//Marshal all the other fields
		//TODO: If we can't marshal a struct field that can be interfaced
		//should this throw an error?
		m, err := Marshal(fv.Interface())
		if err != nil {
			log.Error(err)
			continue
		}

		//Add them to the additional properties map as json.RawMessages
		ap[n] = json.RawMessage(m)
	}

	return json.Marshal(ap)
}

//jsonName gets the effective JSON name of the passed StructField and
//provides a flag indicating that the field should be processed.
func jsonName(sf reflect.StructField) (string, bool) {
	t := sf.Tag.Get(TagKey)

	if t == "" {
		return sf.Name, true
	}

	if t == "-" || t == "*" {
		return "", false
	}

	if idx := strings.Index(t, ","); idx != -1 {
		return t[:idx], true
	}

	return t, true
}

//additionalPropertiesField finds the "wild-card" JSON tag if it exists
//and returns the associated map[string]json.RawMessage.  If no "wild-card"
//field is provided, a new map is returned.  This method panics if the passed
//parameter is not a struct.
func additionalPropertiesField(v interface{}) (map[string]json.RawMessage, error) {
	t := reflect.TypeOf(v)
	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).Tag.Get(TagKey) != "*" {
			continue
		}

		fv := reflect.ValueOf(v).Field(i)
		if !fv.CanInterface() {
			return nil, errors.New(AdditionalPropertiesMustBeExportedMessage)
		}

		if m, ok := fv.Interface().(map[string]json.RawMessage); ok {
			return m, nil
		}

		return nil, errors.New(NotAMapStringJsonRawMessageMessage)
	}

	return make(map[string]json.RawMessage), nil
}

// func unmarshalResource(data []byte, resource resource) error {
// 	var ap map[string]json.RawMessage
// 	err := json.Unmarshal(data, &ap)
// 	if err != nil {
// 		return err
// 	}

// 	err = unmarshalStruct(data, resource, ap)
// 	if err != nil {
// 		return err
// 	}

// 	resource.addAdditionalProperties(ap)
// 	return nil
// }

// func unmarshalStruct(data []byte, v interface{}, ap map[string]json.RawMessage) error {
// 	st := reflect.TypeOf(v).Elem()
// 	sv := reflect.ValueOf(v).Elem()

// 	//Iterate through the struct's fields
// 	for i := 0; i < st.NumField(); i++ {
// 		log.Debug("RawMessage count: ", len(ap))

// 		//Get the field's JSON name
// 		ft := st.Field(i)
// 		log.Debug("Field type: ", ft)
// 		n := name(ft)

// 		//Fields tagged with "-" should not be marshaled/unmarshaled so go
// 		//to the next field
// 		if n == "-" {
// 			continue
// 		}

// 		//Fields that can't be addressed or interfaced can't be set through
// 		//reflection so go to the next field
// 		fv := sv.Field(i)
// 		log.Debug("CanAddr: ", fv.CanAddr())
// 		log.Debug("CanInterface: ", fv.CanInterface())
// 		if !fv.CanAddr() || !fv.CanInterface() {
// 			continue
// 		}

// 		//Get a pointer to the value to pass so that we can either recurse
// 		//into anonymous structures or unmarshal it outright
// 		log.Debug("Field value before: ", fv)
// 		pv := fv.Addr().Interface()
// 		log.Debug("Field pointer: ", reflect.TypeOf(pv).Elem())

// 		//Anonymous struct's fields are part of the current JSON object
// 		//but we have to recurse into them to set their fields
// 		if ft.Anonymous {
// 			err := unmarshalStruct(data, pv, ap)
// 			if err != nil {
// 				return err
// 			}
// 			log.Debug("Anonymous field value after: ", fv)
// 			continue
// 		}

// 		//If a RawMessage doesn't exist for a given field name there's no
// 		//point wasting resources trying to unmarshal it
// 		rm, ok := ap[n]
// 		if !ok {
// 			log.Debug("No raw message found: ", n)
// 			continue
// 		}

// 		//Use the encoding/json version of Unmarshal to turn each
// 		//RawMessage into the individual fields
// 		err := json.Unmarshal(rm, pv)
// 		log.Debug("Field value after: ", fv)
// 		if err != nil {
// 			return err
// 		}

// 		//As fields are unmarshaled the struct's values will be filled
// 		//in and the RawMessageCount will decrease
// 		log.Debug("Struct value now: ", sv)
// 		delete(ap, n)
// 	}

// 	return nil
// }
