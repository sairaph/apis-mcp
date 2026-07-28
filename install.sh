#!/bin/sh
# apis-mcp installer for Linux and macOS.
set -eu

OWNER="sairaph"
REPO="apis-mcp"
BINARY="apis-mcp"
LATEST_URL="https://github.com/${OWNER}/${REPO}/releases/latest"

INSTALL_DIR="${HOME}/.${REPO}/bin"
TARGET="${INSTALL_DIR}/${BINARY}"
STAGED=""
MANIFEST=""
PROFILE_TEMP=""
REPLACEMENT_STATE="staging"

if [ -t 1 ] && [ "${NO_COLOR+x}" != x ]; then
  GREEN=$(printf '\033[32m')
  YELLOW=$(printf '\033[33m')
  CYAN=$(printf '\033[36m')
  RESET=$(printf '\033[0m')
else
  GREEN=""
  YELLOW=""
  CYAN=""
  RESET=""
fi
if [ -t 2 ] && [ "${NO_COLOR+x}" != x ]; then
  RED=$(printf '\033[31m')
  ERROR_RESET=$(printf '\033[0m')
else
  RED=""
  ERROR_RESET=""
fi

title() {
  printf '\n  %s%s installer%s\n\n' "$CYAN" "$BINARY" "$RESET"
}

phase() {
  printf '  [..] %-10s %s\n' "$1" "$2"
}

success() {
  printf '  %s[ok]%s %-10s %s\n' "$GREEN" "$RESET" "$1" "$2"
}

guidance() {
  printf '  %s[--]%s %s\n' "$YELLOW" "$RESET" "$1"
}

fail() {
  printf '  %s[error]%s %s\n' "$RED" "$ERROR_RESET" "$1" >&2
  exit 1
}

cleanup() {
  cleanup_failed=0
  if [ -n "$STAGED" ] && ! rm -f -- "$STAGED"; then cleanup_failed=1; fi
  if [ -n "$MANIFEST" ] && ! rm -f -- "$MANIFEST"; then cleanup_failed=1; fi
  if [ -n "$PROFILE_TEMP" ] && ! rm -f -- "$PROFILE_TEMP"; then cleanup_failed=1; fi
  if [ "$cleanup_failed" -ne 0 ]; then
    printf '  %s[error]%s Could not remove installer temporary files; remove files ending in .new.*, .SHA256SUMS.txt.*, or .apis-mcp.* from the reported install/profile directories.\n' "$RED" "$ERROR_RESET" >&2
  fi
  return 0
}

interrupted() {
  code=$1
  cleanup
  case "$REPLACEMENT_STATE" in
    staging)
      message="Installation interrupted; the existing binary was not replaced."
      ;;
    replacing)
      message="Installation interrupted during final replacement; verify ${TARGET} and re-run the installer."
      ;;
    replaced)
      message="Installation interrupted after ${TARGET} was installed. Finish setup with: ${TARGET} configure"
      ;;
  esac
  printf '\n  %s[error]%s %s\n' "$RED" "$ERROR_RESET" "$message" >&2
  exit "$code"
}

trap cleanup EXIT
trap 'interrupted 129' HUP
trap 'interrupted 130' INT
trap 'interrupted 143' TERM

title

OS_ACTUAL=$(uname -s 2>/dev/null || printf unknown)
ARCH_ACTUAL=$(uname -m 2>/dev/null || printf unknown)
case "$OS_ACTUAL" in
  Linux*) os="linux" ;;
  Darwin*) os="darwin" ;;
  *) fail "Unsupported platform: OS=${OS_ACTUAL}, architecture=${ARCH_ACTUAL}. Supported OS values are Linux and Darwin." ;;
esac
case "$ARCH_ACTUAL" in
  x86_64|amd64) arch="amd64" ;;
  arm64|aarch64) arch="arm64" ;;
  *) fail "Unsupported platform: OS=${OS_ACTUAL}, architecture=${ARCH_ACTUAL}. Supported architectures are amd64 and arm64." ;;
esac

ASSET="${BINARY}-${os}-${arch}"
success "Platform" "${OS_ACTUAL}/${ARCH_ACTUAL} -> ${ASSET}"

if command -v curl >/dev/null 2>&1; then
  DOWNLOADER="curl"
elif command -v wget >/dev/null 2>&1; then
  DOWNLOADER="wget"
else
  fail "Missing download tool: install curl or wget and try again."
fi
if ! command -v awk >/dev/null 2>&1; then
  fail "Missing required tool: awk is needed to read the checksum manifest."
fi
if command -v sha256sum >/dev/null 2>&1; then
  HASH_TOOL="sha256sum"
elif command -v shasum >/dev/null 2>&1; then
  HASH_TOOL="shasum"
else
  fail "Missing checksum tool: install sha256sum or shasum and try again."
fi

phase "Release" "Resolving the latest concrete release..."
if [ "$DOWNLOADER" = "curl" ]; then
  if ! RELEASE_URL=$(curl -fsSL -o /dev/null -w '%{url_effective}' "$LATEST_URL"); then
    fail "Could not resolve the latest release. URL: ${LATEST_URL}"
  fi
else
  if ! RESPONSE=$(wget -S -O /dev/null "$LATEST_URL" 2>&1); then
    fail "Could not resolve the latest release. URL: ${LATEST_URL}"
  fi
  RELEASE_URL=$(printf '%s\n' "$RESPONSE" | awk '
    tolower($1) == "location:" { location=$2 }
    END { sub(/\r$/, "", location); print location }
  ')
fi

RELEASE_PREFIX="https://github.com/${OWNER}/${REPO}/releases/tag/"
case "$RELEASE_URL" in
  "${RELEASE_PREFIX}"*) TAG=${RELEASE_URL#"$RELEASE_PREFIX"}; TAG=${TAG%/} ;;
  *) fail "GitHub returned an unexpected latest-release URL: ${RELEASE_URL:-<empty>}" ;;
esac
case "$TAG" in
  v?*) ;;
  *) fail "GitHub returned a release tag outside the required v<version> contract: ${TAG:-<empty>}" ;;
esac
case "${TAG#v}" in
  *[!A-Za-z0-9._-]*) fail "GitHub returned an unsafe release tag: ${TAG}" ;;
esac
success "Release" "$TAG"

BASE_URL="https://github.com/${OWNER}/${REPO}/releases/download/${TAG}"
ASSET_URL="${BASE_URL}/${ASSET}"
MANIFEST_URL="${BASE_URL}/SHA256SUMS.txt"

if ! mkdir -p "$INSTALL_DIR"; then
  fail "Could not create install directory: ${INSTALL_DIR}"
fi
STAGED="${TARGET}.new.$$"
MANIFEST="${INSTALL_DIR}/.SHA256SUMS.txt.$$"

download() {
  source_url=$1
  destination=$2
  show_progress=$3

  if [ "$DOWNLOADER" = "curl" ]; then
    if [ "$show_progress" = "yes" ] && [ -t 2 ]; then
      curl -fSL --progress-bar "$source_url" -o "$destination"
    else
      curl -fsSL "$source_url" -o "$destination"
    fi
  else
    wget -q -O "$destination" "$source_url"
  fi
}

phase "Download" "$ASSET"
if ! download "$ASSET_URL" "$STAGED" yes; then
  fail "Download failed. URL: ${ASSET_URL}"
fi
if [ ! -s "$STAGED" ]; then
  fail "Download was empty or incomplete; the existing binary was not replaced. URL: ${ASSET_URL}"
fi

phase "Verify" "Downloading SHA256SUMS.txt..."
if ! download "$MANIFEST_URL" "$MANIFEST" no; then
  fail "Checksum manifest download failed. URL: ${MANIFEST_URL}"
fi
if [ ! -s "$MANIFEST" ]; then
  fail "Checksum manifest was empty. URL: ${MANIFEST_URL}"
fi

EXPECTED=$(awk -v asset="$ASSET" '
  length($1) == 64 && $1 !~ /[^0-9A-Fa-f]/ {
    name=$2
    sub(/^\*/, "", name)
    if (name == asset) { count++; hash=tolower($1) }
  }
  END { if (count == 1) print hash }
' "$MANIFEST")
if [ -z "$EXPECTED" ]; then
  fail "SHA256SUMS.txt does not contain exactly one valid checksum for ${ASSET}. URL: ${MANIFEST_URL}"
fi

if [ "$HASH_TOOL" = "sha256sum" ]; then
  ACTUAL=$(sha256sum "$STAGED" | awk '{ print tolower($1) }')
else
  ACTUAL=$(shasum -a 256 "$STAGED" | awk '{ print tolower($1) }')
fi
if [ "$ACTUAL" != "$EXPECTED" ]; then
  fail "Checksum mismatch for ${ASSET}; the existing binary was not replaced."
fi
success "Verify" "SHA-256 ${ACTUAL}"

if ! chmod 755 "$STAGED"; then
  fail "Could not make the staged binary executable; the existing binary was not replaced."
fi
REPLACEMENT_STATE="replacing"
if ! mv -f "$STAGED" "$TARGET"; then
  REPLACEMENT_STATE="staging"
  fail "Could not replace ${TARGET}; check directory permissions and try again."
fi
REPLACEMENT_STATE="replaced"
STAGED=""
rm -f -- "$MANIFEST"
MANIFEST=""
success "Install" "$TARGET"

on_path=0
remaining_path=${PATH:-}
while :; do
  case "$remaining_path" in
    *:*) path_component=${remaining_path%%:*}; remaining_path=${remaining_path#*:}; path_done=0 ;;
    *) path_component=$remaining_path; path_done=1 ;;
  esac
  if [ "$path_component" = "$INSTALL_DIR" ]; then
    on_path=1
    break
  fi
  [ "$path_done" -eq 0 ] || break
done

profile_seen=0
profile_write_failed=0
legacy_path_line="export PATH=\"${INSTALL_DIR}:\$PATH\""
path_line="case \":\$PATH:\" in *\":${INSTALL_DIR}:\"*) ;; *) export PATH=\"${INSTALL_DIR}:\$PATH\" ;; esac"
current_path_command=$path_line

update_profile() {
  candidate=$1
  line=$2
  [ -f "$candidate" ] || return 0
  profile_seen=1
  if grep -Fqx "$legacy_path_line" "$candidate" 2>/dev/null; then
    PROFILE_TEMP="${candidate}.apis-mcp.$$"
    if ! awk -v legacy="$legacy_path_line" '$0 != legacy' "$candidate" > "$PROFILE_TEMP" || ! cat "$PROFILE_TEMP" > "$candidate"; then
      profile_write_failed=1
      guidance "Could not migrate the previous PATH entry in $candidate. Remove this line and retry: $legacy_path_line"
      return 0
    fi
    if rm -f -- "$PROFILE_TEMP"; then PROFILE_TEMP=""; fi
  fi
  if grep -Fqx "$line" "$candidate" 2>/dev/null; then
    success "PATH" "Already configured in $candidate"
  elif printf '\n# added by %s installer\n%s\n' "$BINARY" "$line" >> "$candidate"; then
    success "PATH" "Updated $candidate"
  else
    profile_write_failed=1
    guidance "Could not update $candidate. Add this line yourself: $line"
  fi
}

if [ "$on_path" -eq 0 ]; then
  case "${SHELL:-}" in
    */bash)
      update_profile "$HOME/.bashrc" "$path_line"
      for candidate in "$HOME/.bash_profile" "$HOME/.bash_login" "$HOME/.profile"; do
        if [ -f "$candidate" ]; then
          update_profile "$candidate" "$path_line"
          break
        fi
      done
      ;;
    */zsh)
      update_profile "$HOME/.zshrc" "$path_line"
      update_profile "$HOME/.zprofile" "$path_line"
      ;;
    */fish)
      path_line="fish_add_path \"${INSTALL_DIR}\""
      current_path_command=$path_line
      update_profile "$HOME/.config/fish/config.fish" "$path_line"
      ;;
    */sh|*/dash|*/ksh)
      update_profile "$HOME/.profile" "$path_line"
      ;;
    *)
      if [ -n "${SHELL:-}" ]; then
        guidance "Automatic PATH setup is not available for ${SHELL}."
      fi
      ;;
  esac
fi

if [ "$on_path" -eq 0 ]; then
  if [ "$profile_seen" -eq 1 ]; then
    if [ "$profile_write_failed" -eq 0 ]; then
      guidance "Shells that load the updated startup file(s) will include ${BINARY}. For this shell: $current_path_command"
    else
      guidance "For this shell: $current_path_command"
    fi
  else
    guidance "PATH was not changed because no shell profile exists."
    guidance "Add this to the profile you use: $current_path_command"
  fi
fi

if ( : </dev/tty >/dev/tty 2>/dev/tty ) 2>/dev/null; then
  phase "Configure" "Starting interactive client setup..."
  if "$TARGET" configure </dev/tty >/dev/tty 2>/dev/tty; then
    success "Configure" "Complete"
  else
    configure_status=$?
    guidance "Configure was cancelled or did not complete. Run: ${TARGET} configure"
    exit "$configure_status"
  fi
else
  guidance "No usable terminal; configuration was not started."
  guidance "Finish setup from a terminal: ${TARGET} configure"
fi
