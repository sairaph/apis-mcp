#!/bin/sh
set -eu

ROOT=$(CDPATH= cd -- "$(dirname -- "$0")/.." && pwd)
INSTALLER="$ROOT/install.sh"
FIXTURE="$ROOT/tests/fixtures/installer/downloaded-binary"
TMP=${TMPDIR:-/tmp}/apis-mcp-installer-test.$$
PASSED=0

cleanup() {
  rm -rf "$TMP"
}
trap cleanup EXIT HUP INT TERM
mkdir -p "$TMP"

fail() {
  printf 'installer test failed: %s\n' "$1" >&2
  exit 1
}

assert_contains() {
  case "$1" in
    *"$2"*) ;;
    *) fail "expected output to contain: $2" ;;
  esac
}

assert_not_contains() {
  case "$1" in
    *"$2"*) fail "expected output not to contain: $2" ;;
    *) ;;
  esac
}

make_case() {
  name=$1
  CASE_DIR="$TMP/$name"
  HOME_DIR="$CASE_DIR/home"
  FAKE_BIN="$CASE_DIR/bin"
  mkdir -p "$HOME_DIR" "$FAKE_BIN"
  TEST_PATH="$FAKE_BIN:/usr/bin:/bin"
  URL_LOG="$CASE_DIR/urls.log"
  CONFIGURE_LOG="$CASE_DIR/configure.log"
  : > "$URL_LOG"

  cat > "$FAKE_BIN/uname" <<'EOF'
#!/bin/sh
case "${1-}" in
  -s) printf '%s\n' "$TEST_UNAME_S" ;;
  -m) printf '%s\n' "$TEST_UNAME_M" ;;
  *) exit 1 ;;
esac
EOF

  cat > "$FAKE_BIN/curl" <<'EOF'
#!/bin/sh
output=""
url=""
while [ "$#" -gt 0 ]; do
  case "$1" in
    -o) output=$2; shift 2 ;;
    -*) shift ;;
    *) url=$1; shift ;;
  esac
done
printf '%s\n' "$url" >> "$URL_LOG"
case "$DOWNLOAD_MODE" in
  fail) exit 22 ;;
  zero) : > "$output" ;;
  *) cp "$FIXTURE" "$output" ;;
esac
EOF
  chmod 755 "$FAKE_BIN/uname" "$FAKE_BIN/curl"
}

run_installer() {
  expected_status=$1
  set +e
  HOME="$HOME_DIR" \
    PATH="$TEST_PATH" \
    SHELL=/bin/sh \
    TEST_UNAME_S="$TEST_UNAME_S" \
    TEST_UNAME_M="$TEST_UNAME_M" \
    DOWNLOAD_MODE="$DOWNLOAD_MODE" \
    FIXTURE="$FIXTURE" \
    URL_LOG="$URL_LOG" \
    CONFIGURE_LOG="$CONFIGURE_LOG" \
    sh "$INSTALLER" > "$CASE_DIR/stdout" 2> "$CASE_DIR/stderr"
  status=$?
  set -e
  OUTPUT=$(cat "$CASE_DIR/stdout"; cat "$CASE_DIR/stderr")
  if [ "$expected_status" = success ] && [ "$status" -ne 0 ]; then
    printf '%s\n' "$OUTPUT" >&2
    fail "installer unexpectedly exited $status"
  fi
  if [ "$expected_status" = failure ] && [ "$status" -eq 0 ]; then
    fail 'installer unexpectedly succeeded'
  fi
}

for mapping in \
  'Linux x86_64 apis-mcp-linux-amd64' \
  'Linux aarch64 apis-mcp-linux-arm64' \
  'Darwin amd64 apis-mcp-darwin-amd64' \
  'Darwin arm64 apis-mcp-darwin-arm64'
do
  set -- $mapping
  make_case "mapping-$1-$2"
  : > "$HOME_DIR/.profile"
  TEST_UNAME_S=$1
  TEST_UNAME_M=$2
  EXPECTED_ASSET=$3
  DOWNLOAD_MODE=valid
  run_installer success
  assert_contains "$(cat "$URL_LOG")" "/releases/latest/download/$EXPECTED_ASSET"
  assert_contains "$OUTPUT" 'apis-mcp installer'
  assert_contains "$OUTPUT" 'Downloading...'
  assert_contains "$OUTPUT" 'Not running on a terminal. Finish setup with:'
  assert_not_contains "$OUTPUT" '[ok]'
  assert_not_contains "$OUTPUT" 'Checksum'
  assert_not_contains "$OUTPUT" 'is ready'
  cmp -s "$FIXTURE" "$HOME_DIR/.apis-mcp/bin/apis-mcp" || fail 'downloaded binary was not installed'
  PASSED=$((PASSED + 1))
done

for mode in fail zero; do
  make_case "download-$mode"
  mkdir -p "$HOME_DIR/.apis-mcp/bin"
  printf 'working old binary\n' > "$CASE_DIR/old"
  cp "$CASE_DIR/old" "$HOME_DIR/.apis-mcp/bin/apis-mcp"
  TEST_UNAME_S=Linux
  TEST_UNAME_M=x86_64
  DOWNLOAD_MODE=$mode
  run_installer failure
  cmp -s "$CASE_DIR/old" "$HOME_DIR/.apis-mcp/bin/apis-mcp" || fail 'failed download replaced the old binary'
  if [ "$mode" = fail ]; then
    assert_contains "$OUTPUT" 'Download failed. Please check your connection and try again.'
  else
    assert_contains "$OUTPUT" 'Download did not complete; nothing was installed.'
  fi
  [ ! -e "$HOME_DIR/.apis-mcp/bin/apis-mcp.new" ] || fail 'temporary download was not cleaned'
  PASSED=$((PASSED + 1))
done

make_case path-setup
: > "$HOME_DIR/.profile"
TEST_UNAME_S=Linux
TEST_UNAME_M=x86_64
DOWNLOAD_MODE=valid
run_installer success
assert_contains "$(cat "$HOME_DIR/.profile")" "export PATH=\"$HOME_DIR/.apis-mcp/bin:\$PATH\""
assert_contains "$OUTPUT" 'Open a new terminal so `apis-mcp` is on your PATH.'
PASSED=$((PASSED + 1))

make_case no-profile
TEST_UNAME_S=Linux
TEST_UNAME_M=x86_64
DOWNLOAD_MODE=valid
run_installer success
assert_contains "$OUTPUT" 'Add this to your shell profile:'
[ ! -e "$HOME_DIR/.profile" ] || fail 'installer created a shell profile'
PASSED=$((PASSED + 1))

if [ "$(uname -s)" != Darwin ] && command -v script >/dev/null 2>&1; then
  for configure_status in 0 7; do
    make_case "configure-$configure_status"
    : > "$HOME_DIR/.profile"
    TEST_UNAME_S=Linux
    TEST_UNAME_M=x86_64
    DOWNLOAD_MODE=valid
    set +e
    HOME="$HOME_DIR" \
      PATH="$TEST_PATH" \
      SHELL=/bin/sh \
      TEST_UNAME_S="$TEST_UNAME_S" \
      TEST_UNAME_M="$TEST_UNAME_M" \
      DOWNLOAD_MODE="$DOWNLOAD_MODE" \
      FIXTURE="$FIXTURE" \
      URL_LOG="$URL_LOG" \
      CONFIGURE_LOG="$CONFIGURE_LOG" \
      CONFIGURE_STATUS="$configure_status" \
      script -qefc "sh '$INSTALLER'" /dev/null > "$CASE_DIR/tty-output" 2>&1
    status=$?
    set -e
    [ -e "$CONFIGURE_LOG" ] || fail 'configure was not started on a terminal'
    if [ "$configure_status" -eq 0 ]; then
      [ "$status" -eq 0 ] || fail "successful configure returned $status"
    else
      [ "$status" -ne 0 ] || fail 'cancelled configure returned success'
      tty_output=$(cat "$CASE_DIR/tty-output")
      assert_contains "$tty_output" 'configure skipped or failed'
      assert_not_contains "$tty_output" 'Open a new terminal'
    fi
    PASSED=$((PASSED + 1))
  done
fi

make_case unsupported
TEST_UNAME_S=FreeBSD
TEST_UNAME_M=riscv64
DOWNLOAD_MODE=valid
run_installer failure
assert_contains "$OUTPUT" 'Unsupported OS: FreeBSD'
PASSED=$((PASSED + 1))

grep -Fq '[System.Net.HttpWebRequest]::Create($Url)' "$ROOT/install.ps1" || fail 'PowerShell 5.1-compatible transport is missing'
grep -Fq 'Downloading $Asset...' "$ROOT/install.ps1" || fail 'PowerShell reference download step is missing'
if grep -Eq 'Finding the latest|SHA256SUMS|Checksum verified|is ready|\[(ok|\.\.)\]' "$ROOT/install.ps1" "$ROOT/install.sh"; then
  fail 'removed installer UX remains in an installer'
fi
PASSED=$((PASSED + 1))

printf 'installer tests passed: %s\n' "$PASSED"
