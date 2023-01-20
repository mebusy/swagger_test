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

	"H4sIAAAAAAAC/9RWS4/bNhD+K8Q0t0qW/IjT6LZNiyJogRQteqmrFlxpLDErkdzhaDfuQv+9ICXbsa3N",
	"NsUemotBi9Lwe8yDD1CY1hqNmh1kD0DorNEOw59f0H1PZMivC6MZNfultLZRhWRldPLeGe2fuaLGVvrV",
	"C8ItZPBVcgycDLsuCdHemBKh7/sISnQFKesDQQZIVJgSRSxSobbCdUWBzm27ZiYM10j3yqFoUWon0McR",
	"tbQW9Qz6yCP9yVRKfxZS2TTvtpBt/jXm6AEsGYvEahColPwk54DrXcfgGfPOImRgrt9jwdDnEzLI4rZT",
	"hILNDeqR228O6Qfk/yE7j+wzyFXIonNIwpLZqgbDd2Msf9QRTnaOZkwPv9waaiVDBkrzcgGHk5VmrJC8",
	"aEjUusq/PO45JqWrcB5hULiEbHOImvcRBKfe6suji1pqjc3bMqhyQigVW0OixLtIzMPyvqglR2IR/tze",
	"QgSouxayTRrNo0U+hdVY1ENs/CBb2/jt+WKJq5frVzF+8/o6ni/KZSxXL9fxarFez1fzV6s0TSECK5mR",
	"PJA/N1fx7zL+O41fx/nXL46i7IlHYBvJXrpw0gAKpC7JqBIiULY2GiGPPkIxPoueEHEk8NEJ+UU2jPL6",
	"XLnQd0j1J70aXpsK7bNwyjh1quo8Okmd9QoiaJVWrddiPmWNli2eGnNFXHcUf+dr8DH9/5o04IxNED3E",
	"f4zRpFZfKiVf51h0pHj3q6/3gc2VVT/i7qrjOnDzp9YoS6R9oGx0/RBPhi+G+aH01lyWpFOel2B0rHQl",
	"3L2sqhCQFQe+beyQ7sKjOyQ3FvIsnc331SitggyWs3Q2Flkd0Cb+pxr6sPcktF5fueCbs3+RZIuM5ELf",
	"PYPVXQsfNxpo3nZIuyPLYefYvc+FzqPT0bxI07NpwPiBE9tIdTYHjj6bmwkHL1q0uTkxC7JNHoHr2lbS",
	"DjJ/XWiNFmxMM9JhWXm+w04Luf86afbT2Bo3odfPxvEwsIcMQsffmnL3bBeNfSvvT1OUqcN+WsmpaIf3",
	"ksP94lPSBM7ijy5NF2uxH+PSqhvciS2ZVhzSbi/ZoNKgmB+Ln8ov3xHgP2Lf3x8C+gNeQu5ID/M4lNIR",
	"WADjZ+Lj9h3wPL97Yzt/RvOGa+wp/TeEktEJKTTeBxVmlwqc+X3asTa5r8vB1amSb0whm6PrHTW+vzHb",
	"LEnCXm0cZ0s/yvu8/ycAAP//8xtwSIsLAAA=",
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
