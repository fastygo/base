# Integrating UI8Kit on top of Base

> Roadmap snippet — [`github.com/fastygo/ui8kit`](https://github.com/fastygo/ui8kit) is not rewired onto Base inside this milestone; this sketches the intended layering.

Wrap Base atoms from UI8Kit molecules:

```go
// Future ui8kit
templ StyledButton(props ui.ButtonProps, label string) {
  @basebtn.Button(basebtn.ButtonProps{
    Class: moleculeButtonClasses(props), // merges ui-button presets + callers Class
    Href:  props.Href,
    Type:  props.Type,
    Attrs: props.Attrs,
    // ...
  }, label)
}
```

- **Base**: semantic HTML + guards + Attrs merges.
- **UI8Kit**: `utils/*.go` retains `TypographyClasses`, Tailwind-heavy `variants.go`, emits `ui-*`.
- **App / Fasty UI**: `internal/ui/{elements,widgets,blocks}` via `fastygo add`; optional vendor + Tailwind `@source` coverage.
