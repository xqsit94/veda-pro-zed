package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/xqsit94/veda-pro-zed/internal/schema"
)

func writeFixturePalette(t *testing.T, dir string) string {
	t.Helper()
	palette := `{
  "meta": {"name": "Veda Pro", "author": "Manikandan", "appearance": "dark"},
  "background": {"editor": "#0f0f1a"},
  "foreground": {"editor": "#ededfe"},
  "syntax": {"keyword": {"color": "#FF5C6C"}}
}`
	path := filepath.Join(dir, "palette.json")
	if err := os.WriteFile(path, []byte(palette), 0o644); err != nil {
		t.Fatal(err)
	}
	return path
}

func TestRun_ProducesSchemaValidTheme(t *testing.T) {
	dir := t.TempDir()
	paletteFile := writeFixturePalette(t, dir)
	outFile := filepath.Join(dir, "veda-pro.json")

	if err := run(paletteFile, outFile); err != nil {
		t.Fatalf("run failed: %v", err)
	}

	data, err := os.ReadFile(outFile)
	if err != nil {
		t.Fatalf("output not written: %v", err)
	}
	if err := schema.Validate(data); err != nil {
		t.Fatalf("output failed schema validation: %v", err)
	}

	var out map[string]any
	if err := json.Unmarshal(data, &out); err != nil {
		t.Fatal(err)
	}
	if out["name"] != "Veda Pro" {
		t.Errorf("output name = %v, want Veda Pro", out["name"])
	}
}

func TestRun_IsIdempotent(t *testing.T) {
	dir := t.TempDir()
	paletteFile := writeFixturePalette(t, dir)
	outFile := filepath.Join(dir, "veda-pro.json")

	if err := run(paletteFile, outFile); err != nil {
		t.Fatal(err)
	}
	first, _ := os.ReadFile(outFile)

	if err := run(paletteFile, outFile); err != nil {
		t.Fatal(err)
	}
	second, _ := os.ReadFile(outFile)

	if string(first) != string(second) {
		t.Error("two consecutive builds produced different output bytes")
	}
}

func TestRun_FailsOnMissingPalette(t *testing.T) {
	err := run(filepath.Join(t.TempDir(), "nope.json"), filepath.Join(t.TempDir(), "out.json"))
	if err == nil {
		t.Fatal("expected error for missing palette, got nil")
	}
}
