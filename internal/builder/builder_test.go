package builder

import (
	"encoding/json"
	"testing"
)

func TestBuild_PopulatesTopLevelFields(t *testing.T) {
	family := Build(minimalPalette())

	if family.Schema != "https://zed.dev/schema/themes/v0.2.0.json" {
		t.Errorf("$schema = %q, want zed v0.2.0 url", family.Schema)
	}
	if family.Name != "Veda Pro" {
		t.Errorf("name = %q, want Veda Pro", family.Name)
	}
	if family.Author != "Manikandan" {
		t.Errorf("author = %q, want Manikandan", family.Author)
	}
	if len(family.Themes) != 1 {
		t.Fatalf("themes count = %d, want 1", len(family.Themes))
	}
}

func TestBuild_EmitsDarkTheme(t *testing.T) {
	theme := Build(minimalPalette()).Themes[0]

	if theme.Name != "Veda Pro" {
		t.Errorf("theme name = %q, want Veda Pro", theme.Name)
	}
	if theme.Appearance != "dark" {
		t.Errorf("appearance = %q, want dark", theme.Appearance)
	}
	if theme.Style == nil {
		t.Fatal("style must not be nil")
	}
}

func TestBuild_IsDeterministic(t *testing.T) {
	a, err := json.Marshal(Build(minimalPalette()))
	if err != nil {
		t.Fatal(err)
	}
	b, err := json.Marshal(Build(minimalPalette()))
	if err != nil {
		t.Fatal(err)
	}
	if string(a) != string(b) {
		t.Error("marshal output is non-deterministic between two builds")
	}
}
