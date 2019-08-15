package acceptance

import (
	"encoding/json"
)

func (s *Simple) MarshalJSON() ([]byte, error) {
	type Alias Simple
	aux := struct {
		*Alias
		AP map[string]interface{} `json:"*,omitempty"`
	}{Alias: (*Alias)(s)}
	data, err := json.Marshal(aux)
	if err != nil {
		return nil, err
	}
	var vmap map[string]interface{}
	err = json.Unmarshal(data, &vmap)
	if err != nil {
		return nil, err
	}
	for k, v := range vmap {
		aux.Alias.AP[k] = v
	}
	return json.Marshal(aux.Alias.AP)
}

func (s *Simple) UnmarshalJSON(data []byte) error {
	type Alias Simple
	aux := struct{ *Alias }{Alias: (*Alias)(s)}
	err := json.Unmarshal(data, &aux)
	if err != nil {
		return err
	}
	var ap map[string]interface{}
	err = json.Unmarshal(data, &ap)
	if err != nil {
		return err
	}
	delete(ap, "a")
	s.AP = ap
	return nil
}
