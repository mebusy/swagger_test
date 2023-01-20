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

// ResErrorCode is also a custom error type
func (e ResErrorCode) Error() string {
	var b strings.Builder
	json.NewEncoder(&b).Encode(e)
	return b.String()
}

func NewResError(code int32, msg string) ResErrorCode {
	return ResErrorCode{
		Errcode: code,
		Errmsg:  &msg,
	}
}

func Json2ResErrorCode(s string) ResErrorCode {
	var e ResErrorCode
	json.NewDecoder(strings.NewReader(s)).Decode(&e)
	return e
}

var NoError = ResErrorCode{Errcode: ERR_NO_ERROR}
