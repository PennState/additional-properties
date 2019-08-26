// Code generated by additional-properties DO NOT EDIT.

package acceptance

import (
	"encoding/json"
	"reflect"
	"strings"
)

type Value4bd927e0ba48429499ecb361915bc568 reflect.Value // DO NOT REMOVE (guarantees the reflect package is used)

// MarshalJSON encodes the Simple struct to JSON with additional-properties
func (s Simple) MarshalJSON() ([]byte, error) {
	type Alias Simple
	aux := (Alias)(s)
	if aux.AP == nil {
		aux.AP = map[string]interface{}{}
	}
	aux.AP["fieldA"] = aux.FieldA
	return json.Marshal(aux.AP)
}

// UnmarshalJSON decodes JSON into the Simple struct with additional-properties
func (s *Simple) UnmarshalJSON(data []byte) error {
	type Alias Simple
	aux := (*Alias)(s)
	err := json.Unmarshal(data, &aux)
	if err != nil {
		return err
	}
	_ = json.Unmarshal(data, &s.AP)
	names := map[string]bool{
		"fieldA": true, "fielda": true,
	}
	for k := range s.AP {
		if names[k] {
			delete(s.AP, k)
			continue
		}
		if names[strings.ToLower(k)] {
			delete(s.AP, k)
		}
	}
	if len(s.AP) == 0 {
		s.AP = nil
	}
	return nil
}
