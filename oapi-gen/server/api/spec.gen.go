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

	"H4sIAAAAAAAC/6xTzW7cPAx8FYHfd3Rs5Qc96JamRVH0UrToKdiDanNtpaufkFTSReB3LyRvsjU2veUm",
	"mdTMcDh+gj76FAMGYTBPQMgpBsZ6+Yb8kSjSTRyw3PsYBIOUo01p53orLobujmMo37if0Nty+p9wCwb+",
	"647g3VLl7og4z3MDA3JPLhUgMIBEfRxQnSmt3FZx7ntk3uZdq6JMSI+OUXm0gRUWHDXZlDC0UKAODEXA",
	"SnaimJDELUMdKMpxG8lbAQMuyOUFNCD7hMsVRySYm9LtwjaW7kORhVwYF/E/GOmUwQ2V57f1aYdgrpoV",
	"0bsraMC74Hz2YM61fo02WI8rELgmmTKpD8X/5kRLA4T32REOYG6LgAPE5qU1/rzDXhafsM/kZP+9+LVI",
	"vk7uC+6vs0x1gLKKCe2A9AxkQOIvDEdqW18sNjw7tN4luyJdCbK4MCp+tONYAcVJHcmfMdJD/fSAxMsr",
	"3er2vFgQEwabHBi4bHWroYFkZapqu3zwfcQaxmJ+jeLnAQx8Qql7adZhvtD6X9F86etWia9eZe8t7cEA",
	"oWQKqlCrOnADYkcuflc5m7mBFPkVQV8jHxXdZ2R5H4f9m/1QFXpeZ0Ao4/zGBtwQWkFWVgV8rD60px78",
	"nS8wt+tk3W7mTSmXpXOtrhOzi73dqZdQZNqVHIok03W1NkUWc6m1hnkz/wkAAP//3xYSk7sEAAA=",
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
