#!/usr/bin/env bash
# tools/asan-check.sh — Build an FFI C library with AddressSanitizer and
# run a simple load test via the ffi-helper to detect memory errors.
#
# Usage:
#   ./tools/asan-check.sh [path/to/lib.c]
#
# If no argument is given, uses tests/ffi_integration/lib_asan_clean.c
#
# Requirements: gcc with ASAN support, go toolchain
#
# Exit codes:
#   0 — no ASAN errors detected
#   1 — ASAN errors detected or build failure

set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
REPO_ROOT="$(cd "$SCRIPT_DIR/.." && pwd)"
TMPDIR="$(mktemp -d)"
trap "rm -rf $TMPDIR" EXIT

LIB_SRC="${1:-$REPO_ROOT/tests/ffi_integration/lib_asan_clean.c}"
LIB_NAME="$(basename "$LIB_SRC" .c)"

echo "=== ASAN Check ==="
echo "Source: $LIB_SRC"
echo "Temp dir: $TMPDIR"

# 1. Build the C library with ASAN
echo "--- Building library with AddressSanitizer ---"
gcc -shared -fPIC -fsanitize=address -g -O1 \
    -o "$TMPDIR/lib${LIB_NAME}.so" "$LIB_SRC"

if [ $? -ne 0 ]; then
    echo "FAIL: gcc ASAN build failed"
    exit 1
fi
echo "OK: Library built with ASAN"

# 2. Build the ffi-helper
echo "--- Building ffi-helper ---"
(cd "$REPO_ROOT" && go build -o "$TMPDIR/ffi-helper" ./tools/ffi-helper)
if [ $? -ne 0 ]; then
    echo "FAIL: ffi-helper build failed"
    exit 1
fi
echo "OK: ffi-helper built"

# 3. Create a minimal project
mkdir -p "$TMPDIR/project"
cat > "$TMPDIR/project/fig.toml" << EOF
[ffi]
enabled = true
helper = "$TMPDIR/ffi-helper"
EOF

# 4. Run the helper with the ASAN library and capture stderr for ASAN output
echo "--- Running ASAN test via helper ---"
SOCK="$TMPDIR/project/.fig/ffi/ffi.sock"
mkdir -p "$(dirname "$SOCK")"

# Start helper in background with socket mode
export ASAN_OPTIONS="detect_leaks=0:halt_on_error=0:print_stats=1"
export LD_PRELOAD="$(gcc -print-file-name=libasan.so)"

"$TMPDIR/ffi-helper" --socket "$SOCK" 2>"$TMPDIR/helper_stderr.log" &
HELPER_PID=$!

# Wait for socket
for i in $(seq 1 40); do
    if [ -S "$SOCK" ]; then break; fi
    sleep 0.05
done

if [ ! -S "$SOCK" ]; then
    echo "FAIL: helper socket not created"
    kill $HELPER_PID 2>/dev/null || true
    cat "$TMPDIR/helper_stderr.log"
    exit 1
fi

# Send load command
LOAD_REQ='{"id":"1","cmd":"load","path":"'"$TMPDIR/lib${LIB_NAME}.so"'"}'
echo "$LOAD_REQ" | socat - UNIX-CONNECT:"$SOCK" > "$TMPDIR/load_resp.json" 2>/dev/null || true

# Send a sym + call if function exists
SYM_REQ='{"id":"2","cmd":"sym","handle":0,"name":"asan_safe_add","rtype":"int"}'
echo "$SYM_REQ" | socat - UNIX-CONNECT:"$SOCK" > "$TMPDIR/sym_resp.json" 2>/dev/null || true

CALL_REQ='{"id":"3","cmd":"call","symbol":0,"args":[10,20]}'
echo "$CALL_REQ" | socat - UNIX-CONNECT:"$SOCK" > "$TMPDIR/call_resp.json" 2>/dev/null || true

# Give small time for any deferred ASAN output
sleep 0.2

# Kill helper
kill $HELPER_PID 2>/dev/null || true
wait $HELPER_PID 2>/dev/null || true

# 5. Check for ASAN errors in stderr
echo "--- Checking ASAN output ---"
ASAN_ERRORS=0
if grep -qi "ERROR: AddressSanitizer" "$TMPDIR/helper_stderr.log" 2>/dev/null; then
    ASAN_ERRORS=1
fi
if grep -qi "ERROR: LeakSanitizer" "$TMPDIR/helper_stderr.log" 2>/dev/null; then
    ASAN_ERRORS=1
fi

if [ $ASAN_ERRORS -eq 1 ]; then
    echo "FAIL: ASAN errors detected!"
    echo "--- ASAN output ---"
    cat "$TMPDIR/helper_stderr.log"
    exit 1
fi

echo "OK: No ASAN errors detected"
echo "--- Helper stderr (info) ---"
cat "$TMPDIR/helper_stderr.log" 2>/dev/null || true
echo ""
echo "=== ASAN Check PASSED ==="
exit 0
