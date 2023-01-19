package api

import (
	"context"
	"encoding/json"
	// "fmt"
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
	var res = ResErrorCode{Errmsg: &message}
	json.NewEncoder(w).Encode(res)
}
