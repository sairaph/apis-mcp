$ErrorActionPreference = "Stop"
$owner = "sairaph"
$repo = "apis-mcp"

$arch = switch ($env:PROCESSOR_ARCHITECTURE) {
  "AMD64" { "amd64" }
  "ARM64" { "arm64" }
  default { throw "Unsupported architecture: $env:PROCESSOR_ARCHITECTURE" }
}

$installDir = Join-Path $env:LOCALAPPDATA "apis-mcp"
$target = Join-Path $installDir "apis-mcp.exe"
$temporary = "$target.new"
$asset = "apis-mcp-windows-$arch.exe"
$url = "https://github.com/$owner/$repo/releases/latest/download/$asset"

New-Item -ItemType Directory -Force -Path $installDir | Out-Null
Write-Host "Downloading $asset..."
Invoke-WebRequest -Uri $url -OutFile $temporary
Move-Item -Force $temporary $target

$userPath = [Environment]::GetEnvironmentVariable("PATH", "User")
if (($userPath -split ";") -notcontains $installDir) {
  $nextPath = if ($userPath) { "$installDir;$userPath" } else { $installDir }
  [Environment]::SetEnvironmentVariable("PATH", $nextPath, "User")
}

Write-Host "Installed $target"
try {
  & $target configure
} catch {
  Write-Warning "Run 'apis-mcp configure' later to configure clients: $_"
}
