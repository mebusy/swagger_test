package api

import (
	"context"
	"encoding/json"
	"fmt"

	// "fmt"
	"errors"
	"net/http"

	"github.com/getkin/kin-openapi/openapi3filter"
)

func CustomAuthenticationFunc(c context.Context, input *openapi3filter.AuthenticationInput) error {

	// Type:apiKey Description: Name:token In:header
	if input.SecurityScheme == nil || input.SecurityScheme.Type != "apiKey" || input.SecurityScheme.In != "header" || input.SecurityScheme.Name != "token" {
		return errors.New("api key authentication not found")
	}
	token := input.RequestValidationInput.Request.Header.Get(input.SecurityScheme.Name)
	fmt.Printf("token: %s\n", token)

	return nil
}

func CustomErrorHandler(w http.ResponseWriter, message string, statusCode int) {
	// fmt.Println(">>>> INSIDE ErrorHandler")
	// valid error
	var res = ResErrorCode{Errmsg: &message}
	json.NewEncoder(w).Encode(res)
}
