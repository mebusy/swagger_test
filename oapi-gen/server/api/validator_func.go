package api

import (
	"context"
	"encoding/json"
	// "errors"

	"net/http"
	"strings"

	"github.com/getkin/kin-openapi/openapi3filter"
)

// authenticator
func CustomAuthenticationFunc(c context.Context, input *openapi3filter.AuthenticationInput) error {
	// error returned here, will be wrapped and passed to the customer error handler

	// Type:apiKey Description: Name:token In:header
	if input.SecurityScheme != nil && input.SecurityScheme.Type == "apiKey" && input.SecurityScheme.In == "header" && input.SecurityScheme.Name == "token" {
		// we currently use apiKey in header
		token := input.RequestValidationInput.Request.Header.Get(input.SecurityScheme.Name)
		if token == "" {
			// token not privided
			return NewResError(ERR_APIKEY_NOT_PROVIDED, "apikey not privided")
		}
		// check token whether valid
		// TODO
		return nil
	}

	return NewResError(ERR_UNIMPLEMENTED, "unimplemented security scheme")
}

// validator will add some wrapper to the original error message
var OAPI_ERROR_WRAPPER = []string{
	"security requirements failed: ",
}

// validator error
func CustomErrorHandler(w http.ResponseWriter, message string, statusCode int) {
	// remote the oapi error wrapper
	for _, v := range OAPI_ERROR_WRAPPER {
		message = strings.ReplaceAll(message, v, "")
	}

	// validator error
	err := Json2ResError(message)
	if err.Errcode > 0 {
		// custom error
		json.NewEncoder(w).Encode(err)
		return
	} else {
		// validator error, e.g. number outof range, is too long...
		// shink it to leave only the key part
		idx := strings.Index(message, "Error at")
		if idx > 0 {
			message = message[idx:]
		}
		var res = ResError{Errcode: ERR_VALIDATOR_ERROR, Errmsg: &message}
		json.NewEncoder(w).Encode(res)
		return
	}

}
