package builder

import "strings"

const SchemaURL = "https://zed.dev/schema/themes/v0.2.0.json"

type Palette struct {
	Meta struct {
		Name       string `json:"name"`
		Author     string `json:"author"`
		Appearance string `json:"appearance"`
	} `json:"meta"`
	Background struct {
		Editor   string `json:"editor"`
		Panel    string `json:"panel"`
		Elevated string `json:"elevated"`
	} `json:"background"`
	Foreground struct {
		Editor      string `json:"editor"`
		Muted       string `json:"muted"`
		Placeholder string `json:"placeholder"`
	} `json:"foreground"`
	Accent    string `json:"accent"`
	Cursor    string `json:"cursor"`
	Selection string `json:"selection"`
	Border    struct {
		Default  string `json:"default"`
		Variant  string `json:"variant"`
		Selected string `json:"selected"`
	} `json:"border"`
	Search struct {
		Match       string `json:"match"`
		ActiveMatch string `json:"active_match"`
	} `json:"search"`
	Diagnostic struct {
		Error   string `json:"error"`
		Warning string `json:"warning"`
		Info    string `json:"info"`
		Hint    string `json:"hint"`
	} `json:"diagnostic"`
	Terminal struct {
		Background string `json:"background"`
		Foreground string `json:"foreground"`
		Bright     string `json:"bright"`
		Ansi       struct {
			Black         string `json:"black"`
			Red           string `json:"red"`
			Green         string `json:"green"`
			Yellow        string `json:"yellow"`
			Blue          string `json:"blue"`
			Magenta       string `json:"magenta"`
			Cyan          string `json:"cyan"`
			White         string `json:"white"`
			BrightBlack   string `json:"bright_black"`
			BrightRed     string `json:"bright_red"`
			BrightGreen   string `json:"bright_green"`
			BrightYellow  string `json:"bright_yellow"`
			BrightBlue    string `json:"bright_blue"`
			BrightMagenta string `json:"bright_magenta"`
			BrightCyan    string `json:"bright_cyan"`
			BrightWhite   string `json:"bright_white"`
		} `json:"ansi"`
	} `json:"terminal"`
	Git struct {
		Added    string `json:"added"`
		Modified string `json:"modified"`
		Removed  string `json:"removed"`
	} `json:"git"`
	Syntax map[string]SyntaxToken `json:"syntax"`
}

type SyntaxToken struct {
	Color  string `json:"color"`
	Italic bool   `json:"italic,omitempty"`
}

type Player struct {
	Cursor     string `json:"cursor"`
	Background string `json:"background"`
	Selection  string `json:"selection"`
}

type ThemeFamily struct {
	Schema string  `json:"$schema"`
	Name   string  `json:"name"`
	Author string  `json:"author"`
	Themes []Theme `json:"themes"`
}

type Theme struct {
	Name       string         `json:"name"`
	Appearance string         `json:"appearance"`
	Style      map[string]any `json:"style"`
}

type Highlight struct {
	Color      string  `json:"color"`
	FontStyle  *string `json:"font_style"`
	FontWeight *int    `json:"font_weight"`
}

func Build(p Palette) ThemeFamily {
	style := map[string]any{}
	applySurfaces(style, p)
	applyText(style, p)
	applyBorders(style, p)
	applySearch(style, p)
	applyDiagnostics(style, p)
	applyPlayers(style, p)
	applyTerminal(style, p)
	applyGit(style, p)
	applySyntax(style, p)

	return ThemeFamily{
		Schema: SchemaURL,
		Name:   p.Meta.Name,
		Author: p.Meta.Author,
		Themes: []Theme{{
			Name:       p.Meta.Name,
			Appearance: p.Meta.Appearance,
			Style:      style,
		}},
	}
}

func applySurfaces(style map[string]any, p Palette) {
	editor := normalizeHex(p.Background.Editor)
	panel := normalizeHex(p.Background.Panel)
	elevated := normalizeHex(p.Background.Elevated)

	setIf(style, "background", editor)
	setIf(style, "editor.background", editor)
	setIf(style, "editor.gutter.background", editor)
	setIf(style, "toolbar.background", editor)
	setIf(style, "tab.active_background", editor)

	setIf(style, "surface.background", panel)
	setIf(style, "panel.background", panel)
	setIf(style, "tab_bar.background", panel)
	setIf(style, "tab.inactive_background", panel)
	setIf(style, "status_bar.background", panel)
	setIf(style, "title_bar.background", panel)
	setIf(style, "title_bar.inactive_background", panel)
	setIf(style, "editor.subheader.background", panel)

	setIf(style, "elevated_surface.background", panel)
	setIf(style, "editor.active_line.background", elevated)
	setIf(style, "editor.highlighted_line.background", elevated)
	setIf(style, "element.background", elevated)
	setIf(style, "element.hover", elevated)
	setIf(style, "ghost_element.hover", elevated)

	selected := normalizeHex(p.Selection)
	setIf(style, "element.active", selected)
	setIf(style, "element.selected", selected)
	setIf(style, "ghost_element.active", selected)
	setIf(style, "ghost_element.selected", selected)
}

func applyText(style map[string]any, p Palette) {
	fg := normalizeHex(p.Foreground.Editor)
	muted := normalizeHex(p.Foreground.Muted)
	placeholder := normalizeHex(p.Foreground.Placeholder)
	accent := normalizeHex(p.Accent)

	setIf(style, "editor.foreground", fg)
	setIf(style, "text", fg)
	setIf(style, "icon", fg)
	setIf(style, "text.muted", muted)
	setIf(style, "icon.muted", muted)
	setIf(style, "text.placeholder", placeholder)
	setIf(style, "text.disabled", placeholder)
	setIf(style, "icon.disabled", placeholder)
	setIf(style, "icon.placeholder", muted)
	setIf(style, "editor.line_number", placeholder)
	setIf(style, "editor.active_line_number", fg)

	setIf(style, "text.accent", accent)
	setIf(style, "icon.accent", accent)
	setIf(style, "link_text.hover", accent)
}

func applyBorders(style map[string]any, p Palette) {
	setIf(style, "border", normalizeHex(p.Border.Default))
	setIf(style, "border.variant", normalizeHex(p.Border.Variant))
	setIf(style, "border.disabled", normalizeHex(p.Border.Variant))
	setIf(style, "border.selected", normalizeHex(p.Border.Selected))
	setIf(style, "border.focused", normalizeHex(p.Accent))
	style["border.transparent"] = "#00000000"
}

func applySearch(style map[string]any, p Palette) {
	setIf(style, "search.match_background", normalizeHex(p.Search.Match))
	setIf(style, "search.active_match_background", normalizeHex(p.Search.ActiveMatch))
}

func applyDiagnostics(style map[string]any, p Palette) {
	setIf(style, "error", normalizeHex(p.Diagnostic.Error))
	setIf(style, "warning", normalizeHex(p.Diagnostic.Warning))
	setIf(style, "info", normalizeHex(p.Diagnostic.Info))
	setIf(style, "hint", normalizeHex(p.Diagnostic.Hint))
}

func applyPlayers(style map[string]any, p Palette) {
	if p.Cursor == "" && p.Selection == "" {
		return
	}
	cursor := normalizeHex(p.Cursor)
	selection := normalizeHex(p.Selection)
	style["players"] = []Player{{
		Cursor:     cursor,
		Background: cursor,
		Selection:  selection,
	}}
}

func applyTerminal(style map[string]any, p Palette) {
	setIf(style, "terminal.background", normalizeHex(p.Terminal.Background))
	setIf(style, "terminal.foreground", normalizeHex(p.Terminal.Foreground))
	setIf(style, "terminal.bright_foreground", normalizeHex(p.Terminal.Bright))
	setIf(style, "terminal.dim_foreground", normalizeHex(p.Foreground.Muted))

	a := p.Terminal.Ansi
	setIf(style, "terminal.ansi.black", normalizeHex(a.Black))
	setIf(style, "terminal.ansi.red", normalizeHex(a.Red))
	setIf(style, "terminal.ansi.green", normalizeHex(a.Green))
	setIf(style, "terminal.ansi.yellow", normalizeHex(a.Yellow))
	setIf(style, "terminal.ansi.blue", normalizeHex(a.Blue))
	setIf(style, "terminal.ansi.magenta", normalizeHex(a.Magenta))
	setIf(style, "terminal.ansi.cyan", normalizeHex(a.Cyan))
	setIf(style, "terminal.ansi.white", normalizeHex(a.White))
	setIf(style, "terminal.ansi.bright_black", normalizeHex(a.BrightBlack))
	setIf(style, "terminal.ansi.bright_red", normalizeHex(a.BrightRed))
	setIf(style, "terminal.ansi.bright_green", normalizeHex(a.BrightGreen))
	setIf(style, "terminal.ansi.bright_yellow", normalizeHex(a.BrightYellow))
	setIf(style, "terminal.ansi.bright_blue", normalizeHex(a.BrightBlue))
	setIf(style, "terminal.ansi.bright_magenta", normalizeHex(a.BrightMagenta))
	setIf(style, "terminal.ansi.bright_cyan", normalizeHex(a.BrightCyan))
	setIf(style, "terminal.ansi.bright_white", normalizeHex(a.BrightWhite))
}

func applyGit(style map[string]any, p Palette) {
	added := normalizeHex(p.Git.Added)
	modified := normalizeHex(p.Git.Modified)
	removed := normalizeHex(p.Git.Removed)

	setIf(style, "created", added)
	setIf(style, "deleted", removed)
	setIf(style, "version_control.added", added)
	setIf(style, "version_control.modified", modified)
	setIf(style, "version_control.deleted", removed)
}

func applySyntax(style map[string]any, p Palette) {
	syntax := make(map[string]Highlight, len(p.Syntax))
	for token, tok := range p.Syntax {
		h := Highlight{Color: normalizeHex(tok.Color)}
		if tok.Italic {
			s := "italic"
			h.FontStyle = &s
		}
		syntax[token] = h
	}
	style["syntax"] = syntax
}

func setIf(style map[string]any, key, value string) {
	if value == "" {
		return
	}
	style[key] = value
}

func normalizeHex(hex string) string {
	hex = strings.TrimSpace(hex)
	if len(hex) == 7 {
		return hex + "ff"
	}
	return hex
}
