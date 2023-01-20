package api

import (
	"encoding/json"
	"strings"
)

const (
	ERR_NO_ERROR      = 0
	ERR_UNIMPLEMENTED = 1

	ERR_APIKEY_NOT_PROVIDED = 10
	ERR_VALIDATOR_ERROR     = 20
)

// ResError is also a custom error type
func (e ResError) Error() string {
	var b strings.Builder
	json.NewEncoder(&b).Encode(e)
	return b.String()
}

func NewResError(code int32, msg string) ResError {
	return ResError{
		Errcode: code,
		Errmsg:  &msg,
	}
}

func Json2ResError(s string) ResError {
	var e ResError
	json.NewDecoder(strings.NewReader(s)).Decode(&e)
	return e
}

var NoError = ResError{Errcode: ERR_NO_ERROR}
