package schema

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"fmt"

	"github.com/santhosh-tekuri/jsonschema/v6"
)

//go:embed zed-theme-v0.2.0.json
var zedThemeSchema []byte

const schemaURI = "https://zed.dev/schema/themes/v0.2.0.json"

var compiled = mustCompile()

func mustCompile() *jsonschema.Schema {
	var raw any
	if err := json.Unmarshal(zedThemeSchema, &raw); err != nil {
		panic(fmt.Errorf("decode embedded schema: %w", err))
	}
	c := jsonschema.NewCompiler()
	if err := c.AddResource(schemaURI, raw); err != nil {
		panic(fmt.Errorf("add schema resource: %w", err))
	}
	s, err := c.Compile(schemaURI)
	if err != nil {
		panic(fmt.Errorf("compile schema: %w", err))
	}
	return s
}

// Validate checks the given theme JSON bytes against the Zed theme v0.2.0 schema.
func Validate(themeJSON []byte) error {
	var instance any
	if err := json.NewDecoder(bytes.NewReader(themeJSON)).Decode(&instance); err != nil {
		return fmt.Errorf("invalid JSON: %w", err)
	}
	return compiled.Validate(instance)
}
