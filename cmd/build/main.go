package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/xqsit94/veda-pro-zed/internal/builder"
	"github.com/xqsit94/veda-pro-zed/internal/schema"
)

func main() {
	paletteFile := flag.String("palette", "palette.json", "path to palette source")
	outFile := flag.String("out", "themes/veda-pro.json", "path to generated theme")
	flag.Parse()

	if err := run(*paletteFile, *outFile); err != nil {
		fmt.Fprintln(os.Stderr, "build failed:", err)
		os.Exit(1)
	}
	fmt.Println("wrote", *outFile)
}

func run(paletteFile, outFile string) error {
	raw, err := os.ReadFile(paletteFile)
	if err != nil {
		return fmt.Errorf("read palette: %w", err)
	}
	var p builder.Palette
	if err := json.Unmarshal(raw, &p); err != nil {
		return fmt.Errorf("parse palette: %w", err)
	}

	family := builder.Build(p)

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	enc.SetIndent("", "  ")
	enc.SetEscapeHTML(false)
	if err := enc.Encode(family); err != nil {
		return fmt.Errorf("encode theme: %w", err)
	}

	if err := schema.Validate(buf.Bytes()); err != nil {
		return fmt.Errorf("schema validation: %w", err)
	}

	if err := os.WriteFile(outFile, buf.Bytes(), 0o644); err != nil {
		return fmt.Errorf("write theme: %w", err)
	}
	return nil
}
