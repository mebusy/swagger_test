package api

import (
	"context"
	"github.com/getkin/kin-openapi/openapi3filter"
)

func CustomAuthenticationFunc(c context.Context, input *openapi3filter.AuthenticationInput) error {
	// fmt.Println(">>>> INSIDE AuthenticationFunc")
	return nil
}
