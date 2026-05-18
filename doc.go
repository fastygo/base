// Package base is the root module identity for github.com/fastygo/base.
//
// This module provides per-atom subpackages — import them directly instead of importing base:
//
//	import "github.com/fastygo/base/button"
//
// Base is semantics-only headless primitives: slots for Class / Attrs, no presets.
// Styled molecules live in github.com/fastygo/ui8kit.
//
// Prefer github.com/fastygo/ui8kit when you need Variant, ui-* classes on top of primitives.
//
// # Packages
//
// See [docs/packages.md](/docs/packages.md) or [README](README.md) for the atom list.
//
// # Compatibility
//
// Pre-1.0: APIs are semver-advertised with tags (vX.Y.Z) but MINOR may add atoms or props;
// deprecations noted in changelog.
package base

// Version matches annotated git tags (v0.M.P). Updated by scripts/release.sh via scripts/cmd/set-version.
const Version = "0.1.0"
