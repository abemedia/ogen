package openapi

import (
	"encoding/json"

	"github.com/ogen-go/ogen/jsonschema"
)

// MediaType is Media Type Object.
type MediaType struct {
	Schema   *jsonschema.Schema
	Example  json.RawMessage
	Examples map[string]*Example
	Encoding map[string]*Encoding
}

// Encoding is Encoding Type Object.
type Encoding struct {
	Style         ParameterStyle
	Explode       bool
	AllowReserved bool
}