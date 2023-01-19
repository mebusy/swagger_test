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

	"H4sIAAAAAAAC/3SRwY7UMAyGX6Uy3CjTsHPLbYUQQtxAnGbnELWeNtDGke0uGo3y7sjpakY9cLPj+LP9",
	"/zfoacmUMKmAvwGjZEqCNfmB8oWZ+DMNaHlPSTGphSHnOfZBI6Xut1CyN+knXIJF7xkv4OFd94B3W1W6",
	"B7GU0sKA0nPMBgIPN2TuaUDfvLwMHwrYj7dG4+62yUwZWeO261ujhRfiJSh4iEmPT9CCXjNuKY7INtew",
	"2K8c9frT8BvjOcfveH1edbIs2kIThgEZWkhhMYbSH0wPZKgd2yUxXcj69hdJXPKMjaJoTGMjf8M4VqBG",
	"nQ2xfBTk1/r0iixblzu4wycoLVDGFHIED8eDOzhoIQed6rbdKsgWjFgtMTWqId8G8PAV9ZdU6s7SJ+f+",
	"Z9D9X7fzvWq1LkvgK3hg1JVTY6ObenALGkYBf4K6znmnLfjTXtXTuZytbAdLre7VmqkPc3MXZOXZPFDN",
	"vutqbSJRf3TOQTmXfwEAAP//FT8bqL0CAAA=",
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
