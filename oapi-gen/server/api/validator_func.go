package api

import (
	"context"
	"encoding/json"
	"errors"
	// "fmt"

	"net/http"
	"strings"

	"github.com/getkin/kin-openapi/openapi3filter"
)

// authenticator
func CustomAuthenticationFunc(c context.Context, input *openapi3filter.AuthenticationInput) error {
	// error returned here, will be wrapped and passed to the customer error handler

	// Type:apiKey Description: Name:token In:header
	if input.SecurityScheme == nil || input.SecurityScheme.Type != "apiKey" || input.SecurityScheme.In != "header" || input.SecurityScheme.Name != "token" {

		return errors.New("apiKey not found, or not in header")
	}
	token := input.RequestValidationInput.Request.Header.Get(input.SecurityScheme.Name)
	_ = token
	// fmt.Printf("token: %s\n", token)

	return nil
}

// validator error
func CustomErrorHandler(w http.ResponseWriter, message string, statusCode int) {
	// validator error
	err := Json2ResErrorCode(message)
	if err.Errcode > 0 {
		// custom error
		json.NewEncoder(w).Encode(err)
		return
	} else {
		// shink validator error
		idx := strings.Index(message, "Error at")
		if idx > 0 {
			message = message[idx:]
		}

		var res = ResErrorCode{Errcode: ERR_VALIDATOR_ERROR, Errmsg: &message}
		json.NewEncoder(w).Encode(res)
		return
	}

}
