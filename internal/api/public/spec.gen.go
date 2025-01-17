// Package public provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.14.0 DO NOT EDIT.
package public

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

	"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",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
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
	res := make(map[string]func() ([]byte, error))
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
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
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
