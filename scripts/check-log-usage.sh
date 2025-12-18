#!/usr/bin/env bash
set -euo pipefail

# Scan the repository for usages of the standard "log" package and fmt.Printf
# We require using the project's common/log package instead.

ROOT_DIR=$(git rev-parse --show-toplevel 2>/dev/null || pwd)
cd "$ROOT_DIR"

echo "Checking for forbidden usages of std log and fmt.Printf..."

# Find Go files, excluding the common/log package and common third-party/build dirs
mapfile -t GO_FILES < <(find . -name '*.go' \
    -not -path './common/log/*' \
    -not -path './third_party/*' \
    -not -path './build/*' \
    -not -path './vendor/*' \
    -not -path './.git/*')

violations=()
for f in "${GO_FILES[@]}"; do
  # Skip if file is empty
  [ -s "$f" ] || continue

  # Detect fmt.Printf specifically
  matches=$(grep -nH -E "fmt\.Printf\(" "$f" || true)
  if [ -n "$matches" ]; then
    while IFS= read -r m; do
      violations+=("$m")
    done <<< "$matches"
  fi

  # Detect import of the standard log package explicitly
  matches=$(grep -nH '"log"' "$f" || true)
  if [ -n "$matches" ]; then
    while IFS= read -r m; do
      violations+=("$m")
    done <<< "$matches"
  fi

done

if [ ${#violations[@]} -ne 0 ]; then
  echo "\nForbidden logging usages detected (file:line:match):"
  for v in "${violations[@]}"; do
    echo " - $v"
  done
  echo "\nPlease use the project's common/log package instead of the standard library 'log' or fmt.Printf."
  exit 2
fi

echo "OK: no forbidden logging usages found."
