package utils

import (
	"encoding/json"
)

func Marshal(object interface{}) (string, error) {
	b, err := json.Marshal(object)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
