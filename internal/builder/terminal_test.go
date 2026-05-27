package builder

import "testing"

// terminalPalette is the canonical 16-color ANSI terminal palette.
// Tests reference it directly to enforce the exact-value contract.
var terminalPalette = struct {
	background, foreground string
	ansi                   map[string]string
}{
	background: "#0f0f1a",
	foreground: "#ededfe",
	ansi: map[string]string{
		"black":          "#0a0a13",
		"red":            "#ff0063",
		"green":          "#00f992",
		"yellow":         "#ffe900",
		"blue":           "#00a4ff",
		"magenta":        "#ff57fd",
		"cyan":           "#00ffed",
		"white":          "#ededff",
		"bright_black":   "#3e3e4b",
		"bright_red":     "#ff948b",
		"bright_green":   "#00fcb9",
		"bright_yellow":  "#ffff68",
		"bright_blue":    "#3ea0ff",
		"bright_magenta": "#ff93ff",
		"bright_cyan":    "#38ffff",
		"bright_white":   "#f8f8ff",
	},
}

func paletteWithTerminal() Palette {
	p := fullPalette()
	p.Terminal.Background = terminalPalette.background
	p.Terminal.Foreground = terminalPalette.foreground
	p.Terminal.Bright = terminalPalette.ansi["bright_white"]
	p.Terminal.Ansi.Black = terminalPalette.ansi["black"]
	p.Terminal.Ansi.Red = terminalPalette.ansi["red"]
	p.Terminal.Ansi.Green = terminalPalette.ansi["green"]
	p.Terminal.Ansi.Yellow = terminalPalette.ansi["yellow"]
	p.Terminal.Ansi.Blue = terminalPalette.ansi["blue"]
	p.Terminal.Ansi.Magenta = terminalPalette.ansi["magenta"]
	p.Terminal.Ansi.Cyan = terminalPalette.ansi["cyan"]
	p.Terminal.Ansi.White = terminalPalette.ansi["white"]
	p.Terminal.Ansi.BrightBlack = terminalPalette.ansi["bright_black"]
	p.Terminal.Ansi.BrightRed = terminalPalette.ansi["bright_red"]
	p.Terminal.Ansi.BrightGreen = terminalPalette.ansi["bright_green"]
	p.Terminal.Ansi.BrightYellow = terminalPalette.ansi["bright_yellow"]
	p.Terminal.Ansi.BrightBlue = terminalPalette.ansi["bright_blue"]
	p.Terminal.Ansi.BrightMagenta = terminalPalette.ansi["bright_magenta"]
	p.Terminal.Ansi.BrightCyan = terminalPalette.ansi["bright_cyan"]
	p.Terminal.Ansi.BrightWhite = terminalPalette.ansi["bright_white"]
	return p
}

func TestTerminal_BackgroundAndForeground(t *testing.T) {
	style := Build(paletteWithTerminal()).Themes[0].Style
	assertStyleEqual(t, style, map[string]string{
		"terminal.background":        "#0f0f1aff",
		"terminal.foreground":        "#ededfeff",
		"terminal.bright_foreground": "#f8f8ffff",
	})
}

func TestTerminal_MirrorsAnsiPalette(t *testing.T) {
	style := Build(paletteWithTerminal()).Themes[0].Style
	for name, hex := range terminalPalette.ansi {
		key := "terminal.ansi." + name
		want := hex + "ff"
		got, ok := style[key]
		if !ok {
			t.Errorf("%s missing, want %s", key, want)
			continue
		}
		if got != want {
			t.Errorf("%s = %v, want %v (ANSI palette mirror)", key, got, want)
		}
	}
}
