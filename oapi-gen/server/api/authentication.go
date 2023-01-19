package api

import (
	"context"
	"fmt"
	"github.com/getkin/kin-openapi/openapi3filter"
	"net/http"
)

func CustomAuthenticationFunc(c context.Context, input *openapi3filter.AuthenticationInput) error {
	// fmt.Println(">>>> INSIDE AuthenticationFunc")
	return nil
}

func CustomErrorHandler(w http.ResponseWriter, message string, statusCode int) {
	// fmt.Println(">>>> INSIDE ErrorHandler")
	// valid error
	fmt.Fprintf(w, "Error: %s", message)
}
