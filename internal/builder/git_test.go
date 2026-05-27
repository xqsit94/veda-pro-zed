package builder

import "testing"

func TestGit_DiffColors(t *testing.T) {
	p := fullPalette()
	p.Git.Added = "#A6E36F"
	p.Git.Modified = "#F8DC7E"
	p.Git.Removed = "#FF5C6C"

	style := Build(p).Themes[0].Style
	assertStyleEqual(t, style, map[string]string{
		"created":                  "#A6E36Fff",
		"deleted":                  "#FF5C6Cff",
		"version_control.added":    "#A6E36Fff",
		"version_control.modified": "#F8DC7Eff",
		"version_control.deleted":  "#FF5C6Cff",
	})
}
