[![Go Reference](https://pkg.go.dev/badge/github.com/fastygo/base.svg)](https://pkg.go.dev/github.com/fastygo/base)
[![MIT License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

# Base (`github.com/fastygo/base`)

Semantic **headless** [templ](https://github.com/a-h/templ) primitives — **no** `ui-*` presets and **no** Tailwind/class opinions in-kit. Caller supplies styling via `Class` and `Attrs` only. Styled molecules (`Variant`, `ui-*`, `ui8px` policy) belong in [`github.com/fastygo/ui8kit`](https://pkg.go.dev/github.com/fastygo/ui8kit); product compositions live in apps (`internal/ui/...`, `fastygo add` roadmap).

Imports are **one package per atom** so `go mod vendor` trims to what your graph uses.

Semver tagged releases track `doc.go`'s [`Version`](./doc.go) (mirrors annotated tags). Pre-1.0 policy: MINOR may add atoms or props ([docs/versioning.md](./docs/versioning.md)).

## Install

```bash
go get github.com/fastygo/base@v0.1.0
```

## Minimal example

```go
package main

import (
  "context"
  "os"

  "github.com/fastygo/base/button"
)

func main() {
  _ = button.Button(button.ButtonProps{Class: "my-trigger"}, "OK").Render(context.Background(), os.Stdout)
}
```

Generate templ output after editing `.templ` files:

```bash
go generate ./...
```

## Packages

See [docs/packages.md](docs/packages.md) for links and notes.

## Tooling & CI

```bash
bash ./scripts/preflight.sh           # templ generate via go generate, vet, build, test
PREFLIGHT_RUN_RACE=1 bash ./scripts/preflight.sh
bash ./scripts/release.sh <semver>    # bumps doc.go Version + tag vX.Y.Z
```

## Docs

| Doc | Topic |
|-----|-------|
| [design-principles.md](docs/design-principles.md) | Headless contract |
| [packages.md](docs/packages.md) | Atom catalogue |
| [versioning.md](docs/versioning.md) | Releases & semver |
| [integration.md](docs/integration.md) | Layering Base → UI8Kit → app |

## License

MIT — see [LICENSE](LICENSE).
