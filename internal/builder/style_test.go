package builder

import "testing"

func TestStyle_Surfaces(t *testing.T) {
	style := Build(fullPalette()).Themes[0].Style
	assertStyleEqual(t, style, map[string]string{
		"background":                  "#0f0f1aff",
		"editor.background":           "#0f0f1aff",
		"editor.gutter.background":    "#0f0f1aff",
		"toolbar.background":          "#0f0f1aff",
		"tab.active_background":       "#0f0f1aff",
		"surface.background":          "#0a0a13ff",
		"panel.background":            "#0a0a13ff",
		"tab_bar.background":          "#0a0a13ff",
		"tab.inactive_background":     "#0a0a13ff",
		"status_bar.background":       "#0a0a13ff",
		"title_bar.background":        "#0a0a13ff",
		"elevated_surface.background": "#0a0a13ff",
	})
}

func TestStyle_TextHierarchy(t *testing.T) {
	style := Build(fullPalette()).Themes[0].Style
	assertStyleEqual(t, style, map[string]string{
		"editor.foreground": "#ededfeff",
		"text":              "#ededfeff",
		"text.muted":        "#9a9ab0ff",
		"text.placeholder":  "#5e5e76ff",
		"text.accent":       "#57B9FFff",
	})
}

func TestStyle_AccentOnFocusedElements(t *testing.T) {
	style := Build(fullPalette()).Themes[0].Style
	accent := "#57B9FFff"
	assertStyleEqual(t, style, map[string]string{
		"border.focused":  accent,
		"text.accent":     accent,
		"link_text.hover": accent,
	})
}

func TestStyle_Borders(t *testing.T) {
	style := Build(fullPalette()).Themes[0].Style
	assertStyleEqual(t, style, map[string]string{
		"border":             "#3e3e4bff",
		"border.variant":     "#2a2a36ff",
		"border.selected":    "#293b5bff",
		"border.transparent": "#00000000",
	})
}

func TestStyle_CursorAndSelection(t *testing.T) {
	style := Build(fullPalette()).Themes[0].Style
	players, ok := style["players"].([]Player)
	if !ok {
		t.Fatalf("style.players type = %T, want []Player", style["players"])
	}
	if len(players) == 0 {
		t.Fatal("players slice is empty")
	}
	if players[0].Cursor != "#ededfeff" {
		t.Errorf("players[0].cursor = %q, want #ededfeff", players[0].Cursor)
	}
	if players[0].Selection != "#202036ff" {
		t.Errorf("players[0].selection = %q, want #202036ff", players[0].Selection)
	}
}

func TestStyle_Diagnostics(t *testing.T) {
	style := Build(fullPalette()).Themes[0].Style
	assertStyleEqual(t, style, map[string]string{
		"error":   "#FF5C6Cff",
		"warning": "#F8DC7Eff",
		"info":    "#57B9FFff",
		"hint":    "#9a9ab0ff",
	})
}

func TestStyle_SelectedElementsUseSelection(t *testing.T) {
	style := Build(fullPalette()).Themes[0].Style
	want := "#202036ff"
	assertStyleEqual(t, style, map[string]string{
		"element.selected":       want,
		"element.active":         want,
		"ghost_element.selected": want,
		"ghost_element.active":   want,
	})
}

func TestStyle_SearchHighlights(t *testing.T) {
	style := Build(fullPalette()).Themes[0].Style
	assertStyleEqual(t, style, map[string]string{
		"search.match_background":        "#57B9FF66",
		"search.active_match_background": "#F8DC7E66",
	})
}
