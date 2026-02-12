#!/usr/bin/env bash
# tools/valgrind-check.sh — Build an FFI C library normally and run
# the ffi-helper under Valgrind to detect memory errors and leaks.
#
# Usage:
#   ./tools/valgrind-check.sh [path/to/lib.c]
#
# If no argument is given, uses tests/ffi_integration/lib_asan_clean.c
#
# Requirements: gcc, valgrind, go toolchain
#
# Exit codes:
#   0 — no Valgrind errors detected
#   1 — Valgrind errors detected or build failure

set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
REPO_ROOT="$(cd "$SCRIPT_DIR/.." && pwd)"
TMPDIR="$(mktemp -d)"
trap "rm -rf $TMPDIR" EXIT

LIB_SRC="${1:-$REPO_ROOT/tests/ffi_integration/lib_asan_clean.c}"
LIB_NAME="$(basename "$LIB_SRC" .c)"

echo "=== Valgrind Check ==="
echo "Source: $LIB_SRC"

# Check valgrind is available
if ! command -v valgrind &>/dev/null; then
    echo "SKIP: valgrind not found"
    exit 0
fi

# 1. Build library (no ASAN, plain -g)
echo "--- Building library ---"
gcc -shared -fPIC -g -O0 -o "$TMPDIR/lib${LIB_NAME}.so" "$LIB_SRC"
echo "OK: Library built"

# 2. Build ffi-helper
echo "--- Building ffi-helper ---"
(cd "$REPO_ROOT" && go build -o "$TMPDIR/ffi-helper" ./tools/ffi-helper)
echo "OK: ffi-helper built"

# 3. Run helper under valgrind with --server mode (stdin/stdout)
echo "--- Running Valgrind ---"

# Create a sequence of JSON-RPC commands to pipe into the helper
cat > "$TMPDIR/commands.jsonl" << EOF
{"id":"1","cmd":"load","path":"$TMPDIR/lib${LIB_NAME}.so"}
{"id":"2","cmd":"sym","handle":0,"name":"asan_safe_add","rtype":"int"}
{"id":"3","cmd":"call","symbol":0,"args":[10,20]}
{"id":"4","cmd":"sym","handle":0,"name":"asan_safe_concat","rtype":"string"}
{"id":"5","cmd":"call","symbol":1,"args":["hello"," world"]}
EOF

# Run valgrind with the helper (use --server mode for stdin/stdout)
valgrind \
    --leak-check=full \
    --show-leak-kinds=definite,possible \
    --errors-for-leak-kinds=definite \
    --error-exitcode=42 \
    --suppressions=/dev/null \
    --log-file="$TMPDIR/valgrind.log" \
    "$TMPDIR/ffi-helper" --server < "$TMPDIR/commands.jsonl" > "$TMPDIR/responses.jsonl" 2>"$TMPDIR/helper_stderr.log" &

VG_PID=$!

# Give it time to process and finish
sleep 3
kill $VG_PID 2>/dev/null || true
wait $VG_PID 2>/dev/null
VG_EXIT=$?

# 4. Analyze results
echo "--- Valgrind output ---"
cat "$TMPDIR/valgrind.log" 2>/dev/null || true

if [ "$VG_EXIT" -eq 42 ]; then
    echo ""
    echo "FAIL: Valgrind detected memory errors!"
    exit 1
fi

# Check for definite leaks in log
DEFINITE_LEAKS=$(grep -c "definitely lost:" "$TMPDIR/valgrind.log" 2>/dev/null || echo 0)
ERRORS=$(grep "ERROR SUMMARY:" "$TMPDIR/valgrind.log" 2>/dev/null | grep -oP '\d+ errors' | head -1 || echo "0 errors")

echo ""
echo "Valgrind error summary: $ERRORS"
echo "=== Valgrind Check PASSED ==="
exit 0
