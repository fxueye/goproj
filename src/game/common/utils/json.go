package utils

import (
	"encoding/json"
	"strings"
)

func JsonDecode(str string, v interface{}) error {
	dec := json.NewDecoder(strings.NewReader(str))
	dec.UseNumber()
	err := dec.Decode(v)
	if err != nil {
		return err
	}
	return nil
}
