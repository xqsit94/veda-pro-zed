# Veda Pro

A vivid One Dark Pro-inspired Zed theme on a deep violet-indigo
background. Designed for developers who want vivid syntax highlighting
against a calm dark surface, with a coordinated 16-color ANSI palette
for the integrated terminal.

![Veda Pro theme preview](preview.png)

## Highlights

- **Background**: Deep violet-indigo `#15131f`. Three coordinated shades
  give visual depth without breaking the aesthetic: `#0d0b17` for
  sidebars and popovers (editor −8 per channel) and `#1d1b27` for the
  active line and hover lifts (editor +8 per channel).
- **Syntax**: One Dark Pro **Vivid** palette — recognizable hue assignments
  (purple keywords, electric-blue functions, green strings, yellow types,
  cyan properties/operators) with saturation cranked up for the dark
  background.
- **Italics**: Comments, `comment.doc`, `variable.special`, `predictive`,
  and `link_text` render italic — One Dark Pro convention.
- **Integrated terminal**: A vivid neon 16-color ANSI palette, so the
  in-editor terminal stays visually consistent with the theme.
- **Accent**: One Dark Pro vivid blue `#57B9FF` for focused borders, links,
  and active tab indicators.

## Install

### From the Zed extension registry

Once published, search **Veda Pro** in `zed: extensions` and click install.

### Local development install

From a clone of this repo:

```bash
zed --install-dev-extension .
```

Zed will load the theme and you can switch to it via `zed: open theme
selector` → **Veda Pro**.

## Color provenance

| Surface | Source | Value |
|---|---|---|
| Editor background | Deep violet-indigo | `#15131f` |
| Panel / sidebar / popovers | Editor −8 per channel | `#0d0b17` |
| Active line / hover lift | Editor +8 per channel | `#1d1b27` |
| Selection | Deep violet-indigo selection | `#202036` |
| Cursor | Muted off-white | `#c9c9d8` |
| Terminal ANSI 0–15 | Vivid neon ANSI palette | — |
| Accent | One Dark Pro Vivid blue | `#57B9FF` |
| Keywords (red) | One Dark Pro Vivid | `#FF5C6C` |
| Strings (green) | GitHub fresh green | `#7EE787` |
| Functions | Sparky electric blue (CSS deep sky blue) | `#00BFFF` |
| Types (yellow) | One Dark Pro Vivid | `#F8DC7E` |
| Constants/numbers (orange) | One Dark Pro Vivid | `#FFB454` |
| `this`/`self`/regex (purple) | One Dark Pro Vivid | `#D362FF` |
| Escapes/support (cyan) | One Dark Pro Vivid | `#56DCE4` |
| Comments (gray, italic) | One Dark Pro | `#5C6370` |

References:
- One Dark reference: <https://github.com/zed-industries/zed/blob/main/assets/themes/one/one.json>
- Zed theme schema v0.2.0: <https://zed.dev/schema/themes/v0.2.0.json>

## Development

Veda Pro is generated. Every color in `themes/veda-pro.json` originates in
`palette.json` — never edit the generated theme file directly. The Go build
program reads `palette.json`, transforms it into the Zed schema shape, and
validates the output against the v0.2.0 schema before writing it.

```bash
make test       # run the test suite (builder + schema + integration)
make build      # regenerate themes/veda-pro.json from palette.json
make            # both
```

To change a color: edit `palette.json`, then `make`. The build step fails if
the result doesn't validate against the Zed schema.

## Project structure

```
veda-pro/
├── extension.toml            # Zed extension manifest
├── palette.json              # semantic source of truth (edit this)
├── themes/veda-pro.json      # generated artifact (do not edit)
├── cmd/build/                # Go program: palette → theme + validate
├── internal/builder/         # palette → ThemeFamily transformer
├── internal/schema/          # vendored v0.2.0 schema + validator
├── README.md
├── LICENSE
└── Makefile
```

## License

[MIT](LICENSE) © 2026 Manikandan
