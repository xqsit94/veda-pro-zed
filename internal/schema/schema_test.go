package schema

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/xqsit94/veda-pro-zed/internal/builder"
)

func builtThemeJSON(t *testing.T) []byte {
	t.Helper()
	p := builder.Palette{}
	p.Meta.Name = "Veda Pro"
	p.Meta.Author = "Manikandan"
	p.Meta.Appearance = "dark"
	p.Background.Editor = "#0f0f1a"
	p.Foreground.Editor = "#ededfe"
	p.Syntax = map[string]builder.SyntaxToken{"keyword": {Color: "#FF5C6C"}}

	data, err := json.Marshal(builder.Build(p))
	if err != nil {
		t.Fatal(err)
	}
	return data
}

func TestValidate_AcceptsBuiltTheme(t *testing.T) {
	if err := Validate(builtThemeJSON(t)); err != nil {
		t.Errorf("built theme should validate, got: %v", err)
	}
}

func TestValidate_RejectsMissingAuthor(t *testing.T) {
	var obj map[string]any
	if err := json.Unmarshal(builtThemeJSON(t), &obj); err != nil {
		t.Fatal(err)
	}
	delete(obj, "author")
	bad, _ := json.Marshal(obj)

	err := Validate(bad)
	if err == nil {
		t.Fatal("expected error for missing author, got nil")
	}
	if !strings.Contains(strings.ToLower(err.Error()), "author") {
		t.Errorf("error should mention 'author', got: %v", err)
	}
}

func TestValidate_RejectsInvalidAppearance(t *testing.T) {
	var obj map[string]any
	if err := json.Unmarshal(builtThemeJSON(t), &obj); err != nil {
		t.Fatal(err)
	}
	themes := obj["themes"].([]any)
	theme := themes[0].(map[string]any)
	theme["appearance"] = "twilight"
	bad, _ := json.Marshal(obj)

	if err := Validate(bad); err == nil {
		t.Fatal("expected error for invalid appearance, got nil")
	}
}

func TestValidate_RejectsMalformedJSON(t *testing.T) {
	if err := Validate([]byte("{not json")); err == nil {
		t.Error("expected error for malformed JSON, got nil")
	}
}
