package api

import "embed"

//go:embed v3/*.json
var v3 embed.FS

var OpenAPITemplate string

func init() {
	file, err := v3.ReadFile("v3/openapi.json")
	if err != nil {
		panic(err)
	}

	OpenAPITemplate = string(file)
}
