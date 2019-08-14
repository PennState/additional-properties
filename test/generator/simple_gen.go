package acceptance

import (
	"encoding/json"

	log "github.com/sirupsen/logrus"
)

// func (s *Simple) MarshalJSON() ([]byte, error) {

// }

func (s *Simple) UnmarshalJSON(data []byte) error {
	type Alias Simple
	aux := struct{ *Alias }{Alias: (*Alias)(s)}
	err := json.Unmarshal(data, &aux)
	if err != nil {
		return err
	}
	log.Info("Aux Ptr: ", aux)
	log.Info("Aux: ", aux.Alias)
	var ap map[string]interface{}
	err = json.Unmarshal(data, &ap)
	if err != nil {
		return err
	}
	log.Info("AP: ", ap)
	delete(ap, "a")
	log.Info("AP: ", ap)
	s.AP = ap
	return nil
}
