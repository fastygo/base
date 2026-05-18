# Versioning and releases

Base follows semver with annotated git tags as the canonical release surface (`vMAJOR.MINOR.PATCH`). The string in root [`doc.go`](../doc.go)'s `const Version = "..."` MUST match `v` prefix-stripped tags when publishing releases.

Release flow:

1. `bash ./scripts/preflight.sh` (optionally `PREFLIGHT_RUN_RACE=1` locally).
2. `bash ./scripts/release.sh 0.Y.Z`.
3. `git push origin <branch> --tags`.

[`release.yml`](../.github/workflows/release.yml) asserts tag ↔ `doc.go`, runs full preflight, opens a GitHub Release, and primes `proxy.golang.org` via `go list -m github.com/fastygo/base@${TAG}`.

## Semver cues

- **PATCH** — bugfixes, regenerated templ output churn, wording/docs.
- **MINOR** — additive primitives or props preserving existing symbol behavior contract.
- **MAJOR** — removed/renamed component props/templates (rare prior to GA).
