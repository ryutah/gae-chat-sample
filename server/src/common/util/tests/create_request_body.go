package tests

import (
	"bytes"
	"encoding/json"
	"io"
)

func CreateJSONRequestBody(v interface{}) (io.Reader, error) {
	b, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	return bytes.NewBuffer(b), nil
}
