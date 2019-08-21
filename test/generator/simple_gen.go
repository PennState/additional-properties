package acceptance

import (
	"encoding/json"
)

// MarshalJSON encodes the Simple struct to JSON with additional-properties
func (s *Simple) MarshalJSON() ([]byte, error) {
	type Alias Simple
	aux := (*Alias)(s)
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
	delete(s.AP, "fieldA")
	return nil
}
