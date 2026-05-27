package builder

import "testing"

func minimalPalette() Palette {
	p := Palette{}
	p.Meta.Name = "Veda Pro"
	p.Meta.Author = "Manikandan"
	p.Meta.Appearance = "dark"
	p.Background.Editor = "#0f0f1a"
	p.Foreground.Editor = "#ededfe"
	p.Syntax = map[string]SyntaxToken{"keyword": {Color: "#FF5C6C"}}
	return p
}

func fullPalette() Palette {
	p := Palette{}
	p.Meta.Name = "Veda Pro"
	p.Meta.Author = "Manikandan"
	p.Meta.Appearance = "dark"

	p.Background.Editor = "#0f0f1a"
	p.Background.Panel = "#0a0a13"
	p.Background.Elevated = "#1a1a2a"

	p.Foreground.Editor = "#ededfe"
	p.Foreground.Muted = "#9a9ab0"
	p.Foreground.Placeholder = "#5e5e76"

	p.Accent = "#57B9FF"
	p.Cursor = "#ededfe"
	p.Selection = "#202036"

	p.Border.Default = "#3e3e4b"
	p.Border.Variant = "#2a2a36"
	p.Border.Selected = "#293b5b"

	p.Search.Match = "#57B9FF66"
	p.Search.ActiveMatch = "#F8DC7E66"

	p.Diagnostic.Error = "#FF5C6C"
	p.Diagnostic.Warning = "#F8DC7E"
	p.Diagnostic.Info = "#57B9FF"
	p.Diagnostic.Hint = "#9a9ab0"

	p.Syntax = map[string]SyntaxToken{
		"keyword":          {Color: "#FF5C6C"},
		"string":           {Color: "#A6E36F"},
		"function":         {Color: "#57B9FF"},
		"type":             {Color: "#F8DC7E"},
		"comment":          {Color: "#5C6370", Italic: true},
		"variable.special": {Color: "#D362FF", Italic: true},
	}
	return p
}

func assertStyleEqual(t *testing.T, style map[string]any, expectations map[string]string) {
	t.Helper()
	for key, want := range expectations {
		got, ok := style[key]
		if !ok {
			t.Errorf("style.%s missing, want %q", key, want)
			continue
		}
		if got != want {
			t.Errorf("style.%s = %v, want %v", key, got, want)
		}
	}
}
