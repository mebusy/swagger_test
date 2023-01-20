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

	"H4sIAAAAAAAC/9RWS4/bNhD+K8Q0t0qW/IjT6LZNiyJogRQteqmrFlxpLDErkVxytBt3of9eDCXbsa3N",
	"NsUemotBi9Lwe8yDD1CY1hqNmjxkD+DQW6M9hj+/oP/eOeN4XRhNqImX0tpGFZKU0cl7bzQ/80WNreTV",
	"C4dbyOCr5Bg4GXZ9EqK9MSVC3/cRlOgLpywHggzQucKUKGKRCrUVvisK9H7bNTNhqEZ3rzyKFqX2AjmO",
	"qKW1qGfQR4z0J1Mp/VlIZdO820K2+deYowewzlh0pAaBSklPcg643nUEzJh2FiEDc/0eC4I+n5BBFred",
	"cijI3KAeuf3m0f2A9D9kx8g+g1yFJDqPTlhntqrB8N0Yi486wsnO0Yzpwcutca0kyEBpWi7gcLLShBU6",
	"Fg2da33FL497npzSVTjPYVC4hGxziJr3EQSn3urLo4taao3N2zKockIoFVvjRIl3kZiH5X1RS4rEIvy5",
	"vYUIUHctZJs0mkeLfAqrsaiH2PhBtrbh7fliiauX61cxfvP6Op4vymUsVy/X8WqxXs9X81erNE0hAiuJ",
	"0DGQPzdX8e8y/juNX8f51y+OouyJR2AbSSxdOGkABVKXzqgSIlC2Nhohjz5CMT6LnhBxJPDRCflFNozy",
	"cq5c6Duk+pNeDa9NheYsnDJOnao6T6OT3FmvIIJWadWyGLx7aY6WLZ5ac+Wo7lz8HVfhYw78NWnBGZ8g",
	"e4j/GKdJtb5cUlzrWHRO0e5XrvmBz5VVP+LuqqM6sONTa5Qlun2gbHT+EE+GL4YZovTWXJalV8xLEHpS",
	"uhL+XlZVCEiKAt829ujuwqM7dH4s5lk6m+8rUloFGSxn6WwstDqgTfinGnoxuxLaL1cvcIPmF51skdD5",
	"0HvPYHXXguNGA83bDt3uyHLYOXbwc6Hz6HQ8L9L0bCIQfqDENlKdzYKjz+ZmwsGLNm1uTsyCbJNH4Lu2",
	"lW4HGV8ZWqMFGdOMdEhWzHfYaSHnr5NmP5Gt8RN6/Ww8DUN7yCD09K0pd8922di38/40Rcl12E8rORXt",
	"8F5yuGN8SprAWfzRpeliLfajXFp1gzuxdaYVh7TbSzaoNCjGo/FT+cU9Af4j9v0dIqA/4HVIndPDTA6l",
	"dAQWwPBcfNy+A57nd29s6c9o3nCVPaX/xqEk9EIKjfdBhdmlAmd+n3asTc51Obg6VfKNKWRzdL1zDfc3",
	"IpslSdirjadsyeO8z/t/AgAA///Nw5zTjwsAAA==",
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
