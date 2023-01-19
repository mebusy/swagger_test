// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.4 DO NOT EDIT.
package api

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/6xTzW7cPAx8FYHfd3Rs5Qc96JamRVH0UrToKdiDanNtpaufkFTSReB3LyRvd2skRS+5",
	"SSI5HA5HT9BHn2LAIAzmCQg5xcBYL1+Q3xNFuokDlnsfg2CQcrQp7VxvxcXQ3XEM5Y37Cb0tp/8Jt2Dg",
	"v+4E3i1R7k6I8zw3MCD35FIBAgNI1McB1ZnSym0V575H5m3etSrKhPToGJVHG1hhwVGTTQlDCwXq0KEQ",
	"WNFOFBOSuGWoQ4ty3EbyVsCAC3J5AQ3IPuFyxREJ5qZkex5L8iHGQi6MtR/hfXaEA5jbI+qxYjM38I2R",
	"nhNwQ6Xx0/q0QzBXzYrHmytowLvgfPZgzrV+iVWwHlcgcE0yZVLvynqaf3B1AxwgNsfU+P0Oe1lkxD6T",
	"k/3XIudC+Tq5T7i/zjLVAcqmJrQD0m8gAxJ/YDi1trViWbEL21jq1qtmV6grQRYXRsWPdhwroDipI/kz",
	"RnqoTw9IvFTpVrfnRYKYMNjkwMBlq1sNDSQrU2Xb5YPuI1avFvGrUz8OYOADSt1Ls/b6hdZ/c+4xr1t9",
	"iKpV9t7SHgwQSqagSmtVB25A7MhF70qn+CFFfoHQ58gnRvcZWd7GYf9q/61Cz2sPCGWcX1mAG0IryMqq",
	"gI9Vh/a5Bn/6C8zt2lm3m3lTwmXpXKNrx+xib3fqaIpMu+JDkWS6rsamyGIutdYwb+ZfAQAA//91osND",
	"2gQAAA==",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
