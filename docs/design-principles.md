# Design principles — Base

## Role

Base is **FastyGo layer 0**:

- Canonical HTML semantics, tag guards (`utils.ResolveTag`), ARIA fragments (`utils.MergeAttrs`, helpers).
- **No** prefixed design classes (`ui-*`), **no** Tailwind literals, **no** variant maps — those belong in [`github.com/fastygo/ui8kit`](https://pkg.go.dev/github.com/fastygo/ui8kit).
- Structural hooks only (`data-grow` on `group`, grid `data-cols`, `data-icon` attribute on icons, etc.).
- Imports are **fine-grained** (`github.com/fastygo/base/button`, …); consumers pull only atoms they compose.

Prototype / brand-only styling should use **`@apply` and theme bundles** outside Base and outside strict `ui8px` policy scopes (handled at app/CSS layer).

## Layering recap

```
Base → UI8Kit (molecules) → Fasty UI (`internal/ui/`) → Views
 ui8px policy applies upward from molecules + app (not Base).
```
