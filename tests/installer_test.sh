#!/bin/sh
set -eu

ROOT=$(CDPATH= cd -- "$(dirname -- "$0")/.." && pwd)
INSTALLER="$ROOT/install.sh"
FIXTURE="$ROOT/tests/fixtures/installer/downloaded-binary"
if command -v sha256sum >/dev/null 2>&1; then
  FIXTURE_HASH=$(sha256sum "$FIXTURE" | awk '{print $1}')
else
  FIXTURE_HASH=$(shasum -a 256 "$FIXTURE" | awk '{print $1}')
fi
TMP=${TMPDIR:-/tmp}/apis-mcp-installer-test.$$
ORIGINAL_PATH=$PATH
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

assert_file_equals() {
  expected=$1
  actual=$2
  cmp -s "$expected" "$actual" || fail "$actual did not preserve expected contents"
}

make_case() {
  name=$1
  CASE_DIR="$TMP/$name"
  HOME_DIR="$CASE_DIR/home"
  FAKE_BIN="$CASE_DIR/bin"
  mkdir -p "$HOME_DIR" "$FAKE_BIN"
  TEST_PATH="$FAKE_BIN:/usr/bin:/bin"
  TEST_SHELL=/bin/sh

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
    -w) shift 2 ;;
    -*) shift ;;
    *) url=$1; shift ;;
  esac
done
printf '%s\n' "$url" >> "$URL_LOG"
case "$url" in
  */releases/latest)
    printf '%s' 'https://github.com/sairaph/apis-mcp/releases/tag/v1.2.3'
    ;;
  */SHA256SUMS.txt)
    if [ "$DOWNLOAD_MODE" = manifest-fail ]; then exit 22; fi
    if [ "$DOWNLOAD_MODE" = mismatch ]; then
      hash=0000000000000000000000000000000000000000000000000000000000000000
    else
      hash=$FIXTURE_HASH
    fi
    printf '%s  %s\n' "$hash" "$EXPECTED_ASSET" > "$output"
    ;;
  */apis-mcp-*)
    case "$DOWNLOAD_MODE" in
      fail) exit 22 ;;
      zero) : > "$output" ;;
      truncated) printf 'truncated' > "$output" ;;
      *) cp "$FIXTURE" "$output" ;;
    esac
    ;;
  *) exit 22 ;;
esac
EOF
  chmod 755 "$FAKE_BIN/uname" "$FAKE_BIN/curl"
  URL_LOG="$CASE_DIR/urls.log"
  CONFIGURE_LOG="$CASE_DIR/configure.log"
  : > "$URL_LOG"
}

run_installer() {
  expected_status=$1
  output_file="$CASE_DIR/output"
  stdout_file="$CASE_DIR/stdout"
  stderr_file="$CASE_DIR/stderr"
  set +e
  HOME="$HOME_DIR" \
    PATH="$TEST_PATH" \
    NO_COLOR=1 \
    SHELL="$TEST_SHELL" \
    TEST_UNAME_S="$TEST_UNAME_S" \
    TEST_UNAME_M="$TEST_UNAME_M" \
    DOWNLOAD_MODE="$DOWNLOAD_MODE" \
    EXPECTED_ASSET="$EXPECTED_ASSET" \
    FIXTURE="$FIXTURE" \
    FIXTURE_HASH="$FIXTURE_HASH" \
    URL_LOG="$URL_LOG" \
    CONFIGURE_LOG="$CONFIGURE_LOG" \
    sh "$INSTALLER" > "$stdout_file" 2> "$stderr_file"
  status=$?
  set -e
  STDOUT=$(cat "$stdout_file")
  STDERR=$(cat "$stderr_file")
  printf '%s\n%s\n' "$STDOUT" "$STDERR" > "$output_file"
  OUTPUT=$(cat "$output_file")
  if [ "$expected_status" = success ] && [ "$status" -ne 0 ]; then
    printf '%s\n' "$OUTPUT" >&2
    fail "installer unexpectedly exited $status"
  fi
  if [ "$expected_status" = failure ] && [ "$status" -eq 0 ]; then
    fail "installer unexpectedly succeeded"
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
  assert_contains "$(cat "$URL_LOG")" "/releases/download/v1.2.3/$EXPECTED_ASSET"
  assert_contains "$OUTPUT" "[ok] Platform"
  assert_contains "$OUTPUT" "[ok] Verify"
  assert_contains "$OUTPUT" "No usable terminal"
  assert_contains "$OUTPUT" "$HOME_DIR/.apis-mcp/bin/apis-mcp configure"
  [ ! -e "$CONFIGURE_LOG" ] || fail 'configure ran without a controlling input/output terminal'
  case "$STDOUT" in *'[?1049'*) fail 'redirected stdout received alternate-screen UI' ;; esac
  case "$OUTPUT" in *"$(printf '\033')"*) fail 'NO_COLOR output contained ANSI escapes' ;; esac
  PASSED=$((PASSED + 1))
done

make_case successful-upgrade
mkdir -p "$HOME_DIR/.apis-mcp/bin"
printf 'working old binary\n' > "$HOME_DIR/.apis-mcp/bin/apis-mcp"
: > "$HOME_DIR/.profile"
TEST_UNAME_S=Linux
TEST_UNAME_M=x86_64
EXPECTED_ASSET=apis-mcp-linux-amd64
DOWNLOAD_MODE=valid
run_installer success
assert_file_equals "$FIXTURE" "$HOME_DIR/.apis-mcp/bin/apis-mcp"
assert_contains "$OUTPUT" '[ok] Install'
PASSED=$((PASSED + 1))

make_case download-failure
mkdir -p "$HOME_DIR/.apis-mcp/bin"
printf 'working old binary\n' > "$CASE_DIR/old"
cp "$CASE_DIR/old" "$HOME_DIR/.apis-mcp/bin/apis-mcp"
TEST_UNAME_S=Linux
TEST_UNAME_M=x86_64
EXPECTED_ASSET=apis-mcp-linux-amd64
DOWNLOAD_MODE=fail
run_installer failure
assert_file_equals "$CASE_DIR/old" "$HOME_DIR/.apis-mcp/bin/apis-mcp"
assert_contains "$OUTPUT" "Download failed. URL: https://github.com/sairaph/apis-mcp/releases/download/v1.2.3/$EXPECTED_ASSET"
PASSED=$((PASSED + 1))

for mode in mismatch truncated zero; do
  make_case "$mode"
  mkdir -p "$HOME_DIR/.apis-mcp/bin"
  printf 'working old binary\n' > "$CASE_DIR/old"
  cp "$CASE_DIR/old" "$HOME_DIR/.apis-mcp/bin/apis-mcp"
  TEST_UNAME_S=Linux
  TEST_UNAME_M=x86_64
  EXPECTED_ASSET=apis-mcp-linux-amd64
  DOWNLOAD_MODE=$mode
  run_installer failure
  assert_file_equals "$CASE_DIR/old" "$HOME_DIR/.apis-mcp/bin/apis-mcp"
  if [ "$mode" = zero ]; then
    assert_contains "$OUTPUT" 'Download was empty or incomplete'
  else
    assert_contains "$OUTPUT" 'Checksum mismatch'
  fi
  for leftover in "$HOME_DIR/.apis-mcp/bin/"*.new.* "$HOME_DIR/.apis-mcp/bin/.SHA256SUMS.txt."*; do
    [ ! -e "$leftover" ] || fail "temporary file was not cleaned: $leftover"
  done
  PASSED=$((PASSED + 1))
done

make_case wget-download
rm "$FAKE_BIN/curl"
cat > "$FAKE_BIN/wget" <<'EOF'
#!/bin/sh
output=""
url=""
while [ "$#" -gt 0 ]; do
  case "$1" in
    -q|-S) shift ;;
    -O) output=$2; shift 2 ;;
    -*) shift ;;
    *) url=$1; shift ;;
  esac
done
printf '%s\n' "$url" >> "$URL_LOG"
if [ "${url##*/}" = latest ]; then
  printf '  Location: https://github.com/sairaph/apis-mcp/releases/tag/v1.2.3\n' >&2
elif [ "${url##*/}" = SHA256SUMS.txt ]; then
  hash=$FIXTURE_HASH
  printf '%s  %s\n' "$hash" "$EXPECTED_ASSET" > "$output"
else
  cp "$FIXTURE" "$output"
fi
EOF
chmod 755 "$FAKE_BIN/wget"
TOOL_BIN="$CASE_DIR/tools"
mkdir -p "$TOOL_BIN"
for tool in sh awk mkdir rm chmod mv grep cp; do
  tool_path=$(command -v "$tool")
  ln -s "$tool_path" "$TOOL_BIN/$tool"
done
if command -v sha256sum >/dev/null 2>&1; then
  ln -s "$(command -v sha256sum)" "$TOOL_BIN/sha256sum"
else
  ln -s "$(command -v shasum)" "$TOOL_BIN/shasum"
fi
TEST_PATH="$FAKE_BIN:$TOOL_BIN"
: > "$HOME_DIR/.profile"
TEST_UNAME_S=Linux
TEST_UNAME_M=x86_64
EXPECTED_ASSET=apis-mcp-linux-amd64
DOWNLOAD_MODE=valid
run_installer success
assert_contains "$(cat "$URL_LOG")" '/releases/latest'
assert_contains "$(cat "$URL_LOG")" '/releases/download/v1.2.3/apis-mcp-linux-amd64'
assert_contains "$OUTPUT" '[ok] Verify'
if grep -Fq -- '--show-progress' "$ROOT/install.sh"; then fail 'wget fallback uses GNU-only progress flags'; fi
PASSED=$((PASSED + 1))

make_case path-idempotence
install_dir="$HOME_DIR/.apis-mcp/bin"
printf 'export PATH="%s:$PATH"\nexport PATH="%s-tools:$PATH"\n' "$install_dir" "$install_dir" > "$HOME_DIR/.profile"
TEST_UNAME_S=Linux
TEST_UNAME_M=x86_64
EXPECTED_ASSET=apis-mcp-linux-amd64
DOWNLOAD_MODE=valid
run_installer success
run_installer success
path_line="case \":\$PATH:\" in *\":$install_dir:\"*) ;; *) export PATH=\"$install_dir:\$PATH\" ;; esac"
count=$(grep -Fxc "$path_line" "$HOME_DIR/.profile" || true)
[ "$count" -eq 1 ] || fail "PATH line appeared $count times"
legacy_line="export PATH=\"$install_dir:\$PATH\""
legacy_count=$(grep -Fxc "$legacy_line" "$HOME_DIR/.profile" || true)
[ "$legacy_count" -eq 0 ] || fail 'legacy PATH line was not migrated'
component_count=$(PROFILE="$HOME_DIR/.profile" INSTALL_DIR="$install_dir" PATH=/usr/bin:/bin sh -c '
  . "$PROFILE"
  . "$PROFILE"
  count=0
  old_ifs=$IFS
  IFS=:
  for component in $PATH; do
    [ "$component" != "$INSTALL_DIR" ] || count=$((count + 1))
  done
  IFS=$old_ifs
  printf "%s" "$count"
')
[ "$component_count" -eq 1 ] || fail "sourcing updated profiles duplicated the install directory $component_count times"
PASSED=$((PASSED + 1))

make_case no-profile
TEST_UNAME_S=Linux
TEST_UNAME_M=x86_64
EXPECTED_ASSET=apis-mcp-linux-amd64
DOWNLOAD_MODE=valid
run_installer success
assert_contains "$OUTPUT" 'PATH was not changed because no shell profile exists.'
[ ! -e "$HOME_DIR/.profile" ] || fail 'installer silently created .profile'
PASSED=$((PASSED + 1))

make_case literal-path-component
: > "$HOME_DIR/.profile"
TEST_UNAME_S=Linux
TEST_UNAME_M=x86_64
EXPECTED_ASSET=apis-mcp-linux-amd64
DOWNLOAD_MODE=valid
TEST_PATH="$FAKE_BIN:$HOME_DIR/.apis-mcp/b*:/usr/bin:/bin"
run_installer success
path_line="case \":\$PATH:\" in *\":$HOME_DIR/.apis-mcp/bin:\"*) ;; *) export PATH=\"$HOME_DIR/.apis-mcp/bin:\$PATH\" ;; esac"
grep -Fqx "$path_line" "$HOME_DIR/.profile" || fail 'PATH glob was incorrectly treated as the exact install directory'
PASSED=$((PASSED + 1))

make_case fish-profile
mkdir -p "$HOME_DIR/.config/fish"
: > "$HOME_DIR/.config/fish/config.fish"
TEST_UNAME_S=Linux
TEST_UNAME_M=x86_64
EXPECTED_ASSET=apis-mcp-linux-amd64
DOWNLOAD_MODE=valid
TEST_SHELL=/usr/bin/fish
run_installer success
grep -Fqx "fish_add_path \"$HOME_DIR/.apis-mcp/bin\"" "$HOME_DIR/.config/fish/config.fish" || fail 'Fish profile received invalid or missing PATH syntax'
PASSED=$((PASSED + 1))

if [ "$(uname -s)" != Darwin ] && command -v script >/dev/null 2>&1; then
  for configure_status in 0 7; do
    make_case "configure-status-$configure_status"
    : > "$HOME_DIR/.profile"
    TEST_UNAME_S=Linux
    TEST_UNAME_M=x86_64
    EXPECTED_ASSET=apis-mcp-linux-amd64
    DOWNLOAD_MODE=valid
    tty_output="$CASE_DIR/tty-output"
    set +e
    HOME="$HOME_DIR" \
      PATH="$TEST_PATH" \
      NO_COLOR=1 \
      SHELL=/bin/sh \
      TEST_UNAME_S="$TEST_UNAME_S" \
      TEST_UNAME_M="$TEST_UNAME_M" \
      DOWNLOAD_MODE="$DOWNLOAD_MODE" \
      EXPECTED_ASSET="$EXPECTED_ASSET" \
      FIXTURE="$FIXTURE" \
      FIXTURE_HASH="$FIXTURE_HASH" \
      URL_LOG="$URL_LOG" \
      CONFIGURE_LOG="$CONFIGURE_LOG" \
      CONFIGURE_STATUS="$configure_status" \
      script -qefc "sh '$INSTALLER'" /dev/null > "$tty_output" 2>&1
    status=$?
    set -e
    tty_text=$(cat "$tty_output")
    [ -e "$CONFIGURE_LOG" ] || fail "configure was not started on a usable terminal"
    assert_file_equals "$FIXTURE" "$HOME_DIR/.apis-mcp/bin/apis-mcp"
    if [ "$configure_status" -eq 0 ]; then
      [ "$status" -eq 0 ] || fail "successful configure returned $status"
      assert_contains "$tty_text" '[ok] Configure  Complete'
    else
      [ "$status" -eq "$configure_status" ] || fail "configure status $configure_status became $status"
      assert_contains "$tty_text" 'Configure was cancelled or did not complete.'
      case "$tty_text" in *'[ok] Configure  Complete'*) fail 'cancelled configure was reported complete' ;; esac
    fi
    PASSED=$((PASSED + 1))
  done

  make_case configure-redirected-logs
  : > "$HOME_DIR/.profile"
  TEST_UNAME_S=Linux
  TEST_UNAME_M=x86_64
  EXPECTED_ASSET=apis-mcp-linux-amd64
  DOWNLOAD_MODE=valid
  redirected_stdout="$CASE_DIR/installer.stdout"
  redirected_stderr="$CASE_DIR/installer.stderr"
  tty_output="$CASE_DIR/tty-output"
  HOME="$HOME_DIR" \
    PATH="$TEST_PATH" \
    NO_COLOR=1 \
    SHELL=/bin/sh \
    TEST_UNAME_S="$TEST_UNAME_S" \
    TEST_UNAME_M="$TEST_UNAME_M" \
    DOWNLOAD_MODE="$DOWNLOAD_MODE" \
    EXPECTED_ASSET="$EXPECTED_ASSET" \
    FIXTURE="$FIXTURE" \
    FIXTURE_HASH="$FIXTURE_HASH" \
    URL_LOG="$URL_LOG" \
    CONFIGURE_LOG="$CONFIGURE_LOG" \
    CONFIGURE_STATUS=0 \
    script -qefc "sh '$INSTALLER' > '$redirected_stdout' 2> '$redirected_stderr'" /dev/null > "$tty_output" 2>&1
  [ -e "$CONFIGURE_LOG" ] || fail 'configure did not use the available controlling terminal'
  assert_contains "$(cat "$tty_output")" 'configure-ui-stdout'
  assert_contains "$(cat "$tty_output")" 'configure-ui-stderr'
  case "$(cat "$redirected_stdout")$(cat "$redirected_stderr")" in
    *'[?1049'*|*'configure-ui-'*) fail 'redirected installer logs received configure UI output' ;;
  esac
  assert_contains "$(cat "$redirected_stdout")" '[ok] Configure  Complete'
  PASSED=$((PASSED + 1))
fi

make_case unsupported
TEST_UNAME_S=FreeBSD
TEST_UNAME_M=riscv64
EXPECTED_ASSET=unused
DOWNLOAD_MODE=valid
run_installer failure
assert_contains "$OUTPUT" 'Unsupported platform: OS=FreeBSD, architecture=riscv64.'
PASSED=$((PASSED + 1))

grep -Fq '[System.IO.File]::Replace($staged, $Target' "$ROOT/install.ps1" || fail 'PowerShell does not use staged File.Replace'
grep -Fq 'Get-FileHash -LiteralPath $staged -Algorithm SHA256' "$ROOT/install.ps1" || fail 'PowerShell checksum verification is missing'
grep -Fq '$response.Close()' "$ROOT/install.ps1" || fail 'PowerShell response closure is missing'
grep -Fq '$stream.Dispose()' "$ROOT/install.ps1" || fail 'PowerShell stream disposal is missing'
grep -Fq '$file.Dispose()' "$ROOT/install.ps1" || fail 'PowerShell file disposal is missing'
grep -Fq '$ProgressPreference = $previousProgressPreference' "$ROOT/install.ps1" || fail 'PowerShell progress preference is not restored'
count=$(grep -Fc '[System.Net.HttpWebRequest]::Create($Url)' "$ROOT/install.ps1" || true)
[ "$count" -eq 2 ] || fail 'PowerShell 5.1-compatible HTTP transport is missing'
grep -Fq 'Installation failed: $($installerFailure.Exception.Message)' "$ROOT/install.ps1" || fail 'PowerShell failures are not rendered concisely'
grep -Fq '$global:LASTEXITCODE = 1' "$ROOT/install.ps1" || fail 'PowerShell failures do not set a failing status'
grep -Fq "'^(AMD64|x64|X64)$' { 'amd64'" "$ROOT/install.ps1" || fail 'PowerShell amd64 mapping is missing'
grep -Fq "'^ARM64$' { 'arm64'" "$ROOT/install.ps1" || fail 'PowerShell arm64 mapping is missing'
grep -Fq 'Test-ExactPathComponent $userPath $InstallDir' "$ROOT/install.ps1" || fail 'PowerShell saved PATH check is not component-based'
grep -Fq 'Test-ExactPathComponent $env:PATH $InstallDir' "$ROOT/install.ps1" || fail 'PowerShell current PATH check is not component-based'
grep -Fq 'SHA256SUMS.txt' "$ROOT/.github/workflows/release.yml" || fail 'release checksum asset is missing'
grep -Fq 'sha256sum apis-mcp-* > SHA256SUMS.txt' "$ROOT/.github/workflows/release.yml" || fail 'release checksums do not cover binaries'
grep -Fq 'expected v followed by a nonempty version' "$ROOT/.github/workflows/release.yml" || fail 'release tag contract validation is missing'
grep -Fq 'installer_test.ps1' "$ROOT/.github/workflows/ci.yml" || fail 'Windows installer CI job is missing'
PASSED=$((PASSED + 1))

if command -v pwsh >/dev/null 2>&1; then
  PS_FILE="$ROOT/install.ps1" pwsh -NoProfile -NonInteractive -Command '
    $tokens = $null
    $errors = $null
    [void][System.Management.Automation.Language.Parser]::ParseFile($env:PS_FILE, [ref]$tokens, [ref]$errors)
    if ($errors.Count -ne 0) { $errors | ForEach-Object { Write-Error $_ }; exit 1 }
  '
  PASSED=$((PASSED + 1))
fi

PATH=$ORIGINAL_PATH
printf 'installer tests passed: %s\n' "$PASSED"
