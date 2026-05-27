package builder

import "testing"

func TestSyntax_MapsOneDarkProHues(t *testing.T) {
	style := Build(fullPalette()).Themes[0].Style
	syntax, ok := style["syntax"].(map[string]Highlight)
	if !ok {
		t.Fatalf("style.syntax type = %T, want map[string]Highlight", style["syntax"])
	}

	expectations := map[string]string{
		"keyword":          "#FF5C6Cff",
		"string":           "#A6E36Fff",
		"function":         "#57B9FFff",
		"type":             "#F8DC7Eff",
		"comment":          "#5C6370ff",
		"variable.special": "#D362FFff",
	}
	for token, want := range expectations {
		h, ok := syntax[token]
		if !ok {
			t.Errorf("syntax.%s missing", token)
			continue
		}
		if h.Color != want {
			t.Errorf("syntax.%s.color = %q, want %q", token, h.Color, want)
		}
	}
}

func TestSyntax_CommentIsItalic(t *testing.T) {
	style := Build(fullPalette()).Themes[0].Style
	syntax := style["syntax"].(map[string]Highlight)

	comment, ok := syntax["comment"]
	if !ok {
		t.Fatal("syntax.comment missing")
	}
	if comment.FontStyle == nil {
		t.Fatal("syntax.comment.font_style is nil, want italic")
	}
	if *comment.FontStyle != "italic" {
		t.Errorf("syntax.comment.font_style = %q, want italic", *comment.FontStyle)
	}
}

func TestSyntax_VariableSpecialIsItalic(t *testing.T) {
	style := Build(fullPalette()).Themes[0].Style
	syntax := style["syntax"].(map[string]Highlight)

	vs, ok := syntax["variable.special"]
	if !ok {
		t.Fatal("syntax.variable.special missing")
	}
	if vs.FontStyle == nil {
		t.Fatal("syntax.variable.special.font_style is nil, want italic")
	}
	if *vs.FontStyle != "italic" {
		t.Errorf("syntax.variable.special.font_style = %q, want italic", *vs.FontStyle)
	}
}

func TestSyntax_NonItalicTokensHaveNilFontStyle(t *testing.T) {
	style := Build(fullPalette()).Themes[0].Style
	syntax := style["syntax"].(map[string]Highlight)

	for _, token := range []string{"keyword", "string", "function", "type"} {
		h, ok := syntax[token]
		if !ok {
			continue
		}
		if h.FontStyle != nil {
			t.Errorf("syntax.%s.font_style = %q, want nil", token, *h.FontStyle)
		}
	}
}
