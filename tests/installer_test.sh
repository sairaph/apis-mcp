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
  DAEMON_LOG="$CASE_DIR/daemon.log"
  EVENT_LOG="$CASE_DIR/events.log"
  DAEMON_STOP_STATUS=0
  DAEMON_STOP_MODE=exit
  : > "$URL_LOG"
  : > "$DAEMON_LOG"
  : > "$EVENT_LOG"

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
printf 'download\n' >> "$EVENT_LOG"
case "$DOWNLOAD_MODE" in
  fail) exit 22 ;;
  zero) : > "$output" ;;
  truncated) printf 'partial download\n' > "$output"; exit 18 ;;
  *) cp "$FIXTURE" "$output" ;;
esac
EOF
  chmod 755 "$FAKE_BIN/uname" "$FAKE_BIN/curl"
}

install_old_binary() {
  mkdir -p "$HOME_DIR/.apis-mcp/bin"
  cat > "$HOME_DIR/.apis-mcp/bin/apis-mcp" <<'EOF'
#!/bin/sh
printf '%s\n' "$*" >> "$DAEMON_LOG"
printf 'daemon-stop\n' >> "$EVENT_LOG"
if [ "$DAEMON_STOP_MODE" = hang ]; then
  sleep 10 &
  daemon_hang_pid=$!
  trap 'kill "$daemon_hang_pid" 2>/dev/null || true; wait "$daemon_hang_pid" 2>/dev/null || true; exit 0' TERM
  wait "$daemon_hang_pid"
fi
exit "$DAEMON_STOP_STATUS"
EOF
  chmod 755 "$HOME_DIR/.apis-mcp/bin/apis-mcp"
}

assert_file_equals() {
  actual=$(cat "$1")
  [ "$actual" = "$2" ] || fail "unexpected contents in $1: $actual"
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
    DAEMON_LOG="$DAEMON_LOG" \
    EVENT_LOG="$EVENT_LOG" \
    DAEMON_STOP_STATUS="$DAEMON_STOP_STATUS" \
    DAEMON_STOP_MODE="$DAEMON_STOP_MODE" \
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

make_case fresh-daemon-compat
: > "$HOME_DIR/.profile"
TEST_UNAME_S=Linux
TEST_UNAME_M=x86_64
DOWNLOAD_MODE=valid
run_installer success
[ ! -s "$DAEMON_LOG" ] || fail 'fresh install tried to stop a daemon'
assert_file_equals "$EVENT_LOG" 'download'
PASSED=$((PASSED + 1))

make_case update-daemon-order
: > "$HOME_DIR/.profile"
install_old_binary
TEST_UNAME_S=Linux
TEST_UNAME_M=x86_64
DOWNLOAD_MODE=valid
run_installer success
assert_file_equals "$DAEMON_LOG" 'daemon --stop'
assert_file_equals "$EVENT_LOG" 'download
daemon-stop'
cmp -s "$FIXTURE" "$HOME_DIR/.apis-mcp/bin/apis-mcp" || fail 'update did not replace the old binary'
PASSED=$((PASSED + 1))

make_case hanging-daemon-command
: > "$HOME_DIR/.profile"
install_old_binary
DAEMON_STOP_MODE=hang
TEST_UNAME_S=Linux
TEST_UNAME_M=x86_64
DOWNLOAD_MODE=valid
started=$(date +%s)
run_installer success
elapsed=$(($(date +%s) - started))
[ "$elapsed" -lt 8 ] || fail "hung daemon stop was not bounded: ${elapsed}s"
assert_file_equals "$DAEMON_LOG" 'daemon --stop'
assert_file_equals "$EVENT_LOG" 'download
daemon-stop'
cmp -s "$FIXTURE" "$HOME_DIR/.apis-mcp/bin/apis-mcp" || fail 'hung daemon stop prevented the update'
PASSED=$((PASSED + 1))

make_case unsupported-daemon-command
: > "$HOME_DIR/.profile"
install_old_binary
DAEMON_STOP_STATUS=64
TEST_UNAME_S=Linux
TEST_UNAME_M=x86_64
DOWNLOAD_MODE=valid
run_installer success
assert_file_equals "$DAEMON_LOG" 'daemon --stop'
assert_file_equals "$EVENT_LOG" 'download
daemon-stop'
cmp -s "$FIXTURE" "$HOME_DIR/.apis-mcp/bin/apis-mcp" || fail 'unsupported daemon command prevented the update'
PASSED=$((PASSED + 1))

for mode in fail zero truncated; do
  make_case "download-$mode"
  install_old_binary
  cp "$HOME_DIR/.apis-mcp/bin/apis-mcp" "$CASE_DIR/old"
  DAEMON_STOP_STATUS=64
  TEST_UNAME_S=Linux
  TEST_UNAME_M=x86_64
  DOWNLOAD_MODE=$mode
  run_installer failure
  [ ! -s "$DAEMON_LOG" ] || fail "$mode download tried to stop the daemon"
  assert_file_equals "$EVENT_LOG" 'download'
  cmp -s "$CASE_DIR/old" "$HOME_DIR/.apis-mcp/bin/apis-mcp" || fail 'failed download replaced the old binary'
  if [ "$mode" = fail ] || [ "$mode" = truncated ]; then
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
    if [ "$configure_status" -eq 0 ]; then
      install_old_binary
    fi
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
      DAEMON_LOG="$DAEMON_LOG" \
      EVENT_LOG="$EVENT_LOG" \
      DAEMON_STOP_STATUS="$DAEMON_STOP_STATUS" \
      DAEMON_STOP_MODE="$DAEMON_STOP_MODE" \
      CONFIGURE_STATUS="$configure_status" \
      script -qefc "sh '$INSTALLER'" /dev/null > "$CASE_DIR/tty-output" 2>&1
    status=$?
    set -e
    [ -e "$CONFIGURE_LOG" ] || fail 'configure was not started on a terminal'
    if [ "$configure_status" -eq 0 ]; then
      [ "$status" -eq 0 ] || fail "successful configure returned $status"
      assert_file_equals "$DAEMON_LOG" 'daemon --stop'
      assert_file_equals "$EVENT_LOG" 'download
daemon-stop
configure'
    else
      [ "$status" -ne 0 ] || fail 'cancelled configure returned success'
      [ ! -s "$DAEMON_LOG" ] || fail 'fresh install tried to stop a daemon before configure'
      assert_file_equals "$EVENT_LOG" 'download
configure'
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
grep -Fq '[System.IO.File]::Create($tempTarget)' "$ROOT/install.ps1" || fail 'PowerShell download is not staged'
grep -Fq '[System.IO.File]::Replace($tempTarget, $Target, $backupTarget, $true)' "$ROOT/install.ps1" || fail 'PowerShell atomic replacement is missing'
grep -Fq '$stopProcess.WaitForExit(3000)' "$ROOT/install.ps1" || fail 'PowerShell daemon stop is not bounded'
if grep -Eq 'Remove-Item[[:space:]]+\$Target([[:space:]]|$)' "$ROOT/install.ps1"; then
  fail 'PowerShell installer removes the existing target before replacement'
fi
if grep -Eq 'Finding the latest|SHA256SUMS|Checksum verified|is ready|\[(ok|\.\.)\]' "$ROOT/install.ps1" "$ROOT/install.sh"; then
  fail 'removed installer UX remains in an installer'
fi
PASSED=$((PASSED + 1))

printf 'installer tests passed: %s\n' "$PASSED"
