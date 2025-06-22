package utils

import (
	"encoding/json"
	"net/http"
)

func DecodeStrictJson(r *http.Request, target interface{}) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	return decoder.Decode(target)
}
