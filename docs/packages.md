# Package map

Canonical import pattern:

```go
import "github.com/fastygo/base/<atom>"
```

| Path | Responsibility |
|------|----------------|
| `github.com/fastygo/base` | Root module docs + semver `Version`; no runtime API. |
| `github.com/fastygo/base/utils` | `Cn`, mergeable ARIA attrs, semantic tag whitelist. |
| `github.com/fastygo/base/box` | Polymorphic block-level container primitive. |
| `github.com/fastygo/base/block` | Public anchor primitive (`ID`, tag, Class). |
| `github.com/fastygo/base/container` | Narrow container tags (`main`/`section`/`div`). |
| `github.com/fastygo/base/stack` | Vertical flow stack with allowed tags. |
| `github.com/fastygo/base/group` | Horizontal group + `data-grow` hook (no presets). |
| `github.com/fastygo/base/grid` | Structural grid wrappers + data hints (`data-cols`, `data-span`, …). |
| `github.com/fastygo/base/title` | Heading order primitive (`Order` maps to `<hN>`). |
| `github.com/fastygo/base/text` | Typography tag selector + Class slot (no typography map). |
| `github.com/fastygo/base/button` | `<button>` / `<a>` swap, disabled-safe link attrs. |
| `github.com/fastygo/base/badge` | Span primitive for labels/counts/chips with Class slot. |
| `github.com/fastygo/base/field` | Wrapped field control + labeling via Class slots (`WrapperClass`, …). |
| `github.com/fastygo/base/form` | `<form>` attributes (action/method/target/attrs). |
| `github.com/fastygo/base/fieldset` … `legend` … `datalist` … `optgroup` … `output` … `meter` … `progress` | Native control bundles (see source). |
| `github.com/fastygo/base/image` … `picture` … `source` … `figure` | Picture pipeline without object-fit presets. |
| `github.com/fastygo/base/icon` | `data-icon` name slot — latty/font layering lives in ui8kit. |
| `github.com/fastygo/base/table` | Semantic table primitives. |
| `github.com/fastygo/base/list` … `description` … `disclosure` | Native list + dl + disclosure. |
