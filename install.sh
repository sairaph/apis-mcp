#!/bin/sh
set -eu

owner="sairaph"
repo="apis-mcp"

case "$(uname -s)" in
  Linux*) os="linux" ;;
  Darwin*) os="darwin" ;;
  *) printf 'Unsupported operating system.\n' >&2; exit 1 ;;
esac

case "$(uname -m)" in
  x86_64|amd64) arch="amd64" ;;
  arm64|aarch64) arch="arm64" ;;
  *) printf 'Unsupported architecture.\n' >&2; exit 1 ;;
esac

install_dir="${HOME}/.apis-mcp/bin"
target="${install_dir}/apis-mcp"
asset="apis-mcp-${os}-${arch}"
url="https://github.com/${owner}/${repo}/releases/latest/download/${asset}"

mkdir -p "$install_dir"
temporary="${target}.new"
trap 'rm -f "$temporary"' EXIT HUP INT TERM

printf 'Downloading %s...\n' "$asset"
if command -v curl >/dev/null 2>&1; then
  curl -fL --progress-bar "$url" -o "$temporary"
elif command -v wget >/dev/null 2>&1; then
  wget -O "$temporary" "$url"
else
  printf 'curl or wget is required.\n' >&2
  exit 1
fi
chmod 755 "$temporary"
mv -f "$temporary" "$target"
trap - EXIT HUP INT TERM

case ":${PATH}:" in
  *":${install_dir}:"*) ;;
  *)
    profile="${HOME}/.profile"
    [ -f "${HOME}/.zshrc" ] && profile="${HOME}/.zshrc"
    [ -f "${HOME}/.bashrc" ] && profile="${HOME}/.bashrc"
    if ! grep -qF "$install_dir" "$profile" 2>/dev/null; then
      printf '\n# added by apis-mcp installer\nexport PATH="%s:$PATH"\n' "$install_dir" >> "$profile"
    fi
    ;;
esac

printf 'Installed %s\n' "$target"
if [ -e /dev/tty ]; then
  "$target" configure </dev/tty || printf 'Run `apis-mcp configure` later to configure clients.\n'
else
  printf 'Run `apis-mcp configure` to configure clients.\n'
fi
