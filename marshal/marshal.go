package marshal

import (
	"bytes"
	"encoding/json"
)

// Marshal _
func Marshal(val interface{}) (string, error) {
	b := bytes.NewBufferString("")
	encoder := json.NewEncoder(b)
	if err := encoder.Encode(val); err != nil {
		return "", err
	}
	return b.String(), nil
}

// Unmarshal _
func Unmarshal(s string) (interface{}, error) {
	if s == "" {
		return nil, nil
	}
	data := make([]interface{}, 0)
	decoder := json.NewDecoder(bytes.NewBufferString(s))
	err := decoder.Decode(&data)
	return data, err
}
