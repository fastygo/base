#!/usr/bin/env bash
set -euo pipefail

# Usage: ./scripts/release.sh 0.1.0
# Creates a release commit and annotated tag v0.1.0.

VERSION="${1:-}"
RUN_RACE="${RELEASE_RUN_RACE:-0}"

if [ -z "$VERSION" ]; then
  echo "Usage: $0 <version>"
  echo "Example: $0 0.1.1"
  exit 1
fi

if ! echo "$VERSION" | grep -qE '^[0-9]+\.[0-9]+\.[0-9]+$'; then
  echo "Error: version must be semver without v prefix"
  exit 1
fi

TAG="v${VERSION}"

if git rev-parse "$TAG" >/dev/null 2>&1; then
  echo "Error: tag $TAG already exists"
  exit 1
fi

if [ -n "$(git status --porcelain)" ]; then
  echo "Error: working tree is not clean. Commit or stash first."
  exit 1
fi

BRANCH="$(git branch --show-current)"
if [ "$BRANCH" != "main" ]; then
  echo "Warning: releasing from '$BRANCH' (not main). Continue? [y/N]"
  read -r CONFIRM
  if [ "$CONFIRM" != "y" ] && [ "$CONFIRM" != "Y" ]; then
    echo "Aborted."
    exit 1
  fi
fi

echo "Running preflight..."
PREFLIGHT_RUN_RACE="$RUN_RACE" bash ./scripts/preflight.sh

echo "Updating Version constant to ${VERSION}..."
go run ./scripts/cmd/set-version "${VERSION}"

git add doc.go
git commit -m "chore: release ${TAG}"

echo "Creating annotated tag ${TAG}..."
git tag -a "$TAG" -m "${TAG}"

echo ""
echo "Done. Push:"
echo "  git push origin ${BRANCH} --tags"
echo ""
echo "After push:"
echo "  - CI verifies tag/version in doc.go"
echo "  - release.yml publishes GitHub Release + triggers proxy.golang.org indexing"
