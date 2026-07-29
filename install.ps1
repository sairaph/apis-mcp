# apis-mcp installer (Windows). Run in PowerShell:
#   irm https://github.com/sairaph/apis-mcp/raw/main/install.ps1 | iex
$ErrorActionPreference = 'Stop'
$global:LASTEXITCODE = 0
# Invoke-WebRequest's built-in progress bar is slow and flickers badly in
# Windows PowerShell; the download below renders its own.
$ProgressPreference    = 'SilentlyContinue'

$Owner  = 'sairaph'
$Repo   = 'apis-mcp'
$Binary = 'apis-mcp'

# --- detect arch -----------------------------------------------------------
$arch = $env:PROCESSOR_ARCHITECTURE
switch ($arch) {
  { $_ -in 'AMD64','x64' } { $target = 'windows-amd64' }
  'ARM64'                  { $target = 'windows-arm64' }
  default                  { Write-Host "  Unsupported architecture: $arch" -ForegroundColor Red; $global:LASTEXITCODE = 1; return }
}

$Asset = "$Binary-$target.exe"
$Url   = "https://github.com/$Owner/$Repo/releases/latest/download/$Asset"

# --- install location ------------------------------------------------------
$InstallDir = Join-Path $env:LOCALAPPDATA $Repo
$Target     = Join-Path $InstallDir "$Binary.exe"
New-Item -ItemType Directory -Force -Path $InstallDir | Out-Null

Write-Host ""
Write-Host "  apis-mcp installer" -ForegroundColor Cyan
Write-Host ""

Write-Host "  Downloading $Asset..."

# Always stage beside the target. The existing executable remains untouched
# until the download is complete, flushed, closed, and validated.
$tempTarget = "$Target.new"
$backupTarget = "$Target.old"
Remove-Item $tempTarget -Force -ErrorAction SilentlyContinue

$request = $null
$response = $null
$stream = $null
$fs = $null
$downloadError = $null
try {
  $request = [System.Net.HttpWebRequest]::Create($Url)
  $request.Method = "GET"
  $response = $request.GetResponse()
  $total = [long]$response.ContentLength
  $stream = $response.GetResponseStream()
  $fs = [System.IO.File]::Create($tempTarget)
  $buffer = New-Object byte[] 65536
  $downloaded = [long]0
  $sw = [System.Diagnostics.Stopwatch]::StartNew()
  $lastUpdate = 0

  $renderBar = {
    param($dl, $tot, $el)
    $pct = if ($tot -gt 0) { [math]::Round($dl / $tot * 100) } else { 0 }
    if ($pct -gt 100) { $pct = 100 }
    $filled = [math]::Floor($pct / 5)
    if ($filled -gt 20) { $filled = 20 }
    $bar = ('#' * $filled).PadRight(20)
    $dlMB = [math]::Round($dl / 1048576, 1)
    $totMB = [math]::Round($tot / 1048576, 1)
    $speed = if ($el -gt 0) { [math]::Round($dlMB / $el, 1) } else { 0 }
    $eta = if ($speed -gt 0) { [math]::Round((($totMB - $dlMB)) / $speed) } else { 0 }
    $line = ("  [{0}] {1,3}%  {2}/{3} MB  {4} MB/s  ETA {5:00}s" -f $bar, $pct, $dlMB, $totMB, $speed, $eta)
    # Pad so shorter lines fully overwrite the previous render (no stale chars).
    Write-Host ("`r{0}" -f $line.PadRight(72)) -NoNewline
  }

  while (($read = $stream.Read($buffer, 0, $buffer.Length)) -gt 0) {
    $fs.Write($buffer, 0, $read)
    $downloaded += $read
    $now = [System.Environment]::TickCount
    if ($now - $lastUpdate -gt 200) {
      $lastUpdate = $now
      & $renderBar $downloaded $total $sw.Elapsed.TotalSeconds
    }
  }
  # Force a final 100% render.
  & $renderBar $downloaded $total $sw.Elapsed.TotalSeconds
  Write-Host ""
  if ($total -ge 0 -and $downloaded -ne $total) {
    throw "Download was incomplete: received $downloaded of $total bytes."
  }
} catch {
  $downloadError = $_
} finally {
  if ($fs -ne $null) {
    try { $fs.Flush() } catch { if ($null -eq $downloadError) { $downloadError = $_ } }
    try { $fs.Close() } catch { if ($null -eq $downloadError) { $downloadError = $_ } }
  }
  if ($stream -ne $null) { try { $stream.Close() } catch {} }
  if ($response -ne $null) { try { $response.Close() } catch {} }
}

if ($null -ne $downloadError) {
  Write-Host ""
  Write-Host "  Download failed: $downloadError" -ForegroundColor Red
  Write-Host "  URL: $Url" -ForegroundColor Red
  Remove-Item $tempTarget -Force -ErrorAction SilentlyContinue
  $global:LASTEXITCODE = 1
  return
}

if (-not (Test-Path -LiteralPath $tempTarget -PathType Leaf) -or (Get-Item -LiteralPath $tempTarget).Length -eq 0) {
  Remove-Item $tempTarget -Force -ErrorAction SilentlyContinue
  Write-Host "  Download did not complete; nothing was installed." -ForegroundColor Red
  $global:LASTEXITCODE = 1
  return
}

$replacingExisting = Test-Path -LiteralPath $Target -PathType Leaf
if ($replacingExisting) {
  # Bound the compatibility hook so an unresponsive legacy binary cannot hold
  # the installer indefinitely. Output is drained and discarded.
  $stopProcess = $null
  try {
    $stopInfo = New-Object System.Diagnostics.ProcessStartInfo
    $stopInfo.FileName = $Target
    $stopInfo.Arguments = 'daemon --stop'
    $stopInfo.UseShellExecute = $false
    $stopInfo.CreateNoWindow = $true
    $stopInfo.RedirectStandardOutput = $true
    $stopInfo.RedirectStandardError = $true
    $stopProcess = New-Object System.Diagnostics.Process
    $stopProcess.StartInfo = $stopInfo
    $null = $stopProcess.Start()
    $stopProcess.BeginOutputReadLine()
    $stopProcess.BeginErrorReadLine()
    if (-not $stopProcess.WaitForExit(3000)) {
      $stopProcess.Kill()
      $stopProcess.WaitForExit()
    } else {
      $stopProcess.WaitForExit()
    }
  } catch {
  } finally {
    if ($null -ne $stopProcess) { $stopProcess.Dispose() }
  }
  Start-Sleep -Milliseconds 300
}

Remove-Item $backupTarget -Force -ErrorAction SilentlyContinue
try {
  if (Test-Path -LiteralPath $Target -PathType Leaf) {
    [System.IO.File]::Replace($tempTarget, $Target, $backupTarget, $true)
  } else {
    Move-Item -LiteralPath $tempTarget -Destination $Target
  }
  Remove-Item $backupTarget -Force -ErrorAction SilentlyContinue
} catch {
  $replaceError = $_
  if (-not (Test-Path -LiteralPath $Target) -and (Test-Path -LiteralPath $backupTarget -PathType Leaf)) {
    try { Move-Item -LiteralPath $backupTarget -Destination $Target -Force } catch {}
  }
  Write-Host ""
  Write-Host "  Could not replace $Binary.exe: $replaceError" -ForegroundColor Red
  if (Test-Path -LiteralPath $Target -PathType Leaf) {
    Write-Host "  The existing executable was kept at $Target." -ForegroundColor Yellow
  } elseif (Test-Path -LiteralPath $backupTarget -PathType Leaf) {
    Write-Host "  The existing executable is backed up at $backupTarget." -ForegroundColor Yellow
  }
  Write-Host "  Close any running $Binary process and re-run the installer." -ForegroundColor Yellow
  Remove-Item $tempTarget -Force -ErrorAction SilentlyContinue
  $global:LASTEXITCODE = 1
  return
}

# --- add to user PATH if missing -------------------------------------------
# SetEnvironmentVariable updates the stored user PATH, which only new processes
# read. This shell keeps the PATH it started with, so the command is not found
# here no matter what we do -- the user has to be told, or the very next thing
# they type fails with "not recognized".
$pathChanged = $false
$userPath = [Environment]::GetEnvironmentVariable('PATH', 'User')
if ($userPath -notlike "*$InstallDir*") {
  $newPath = if ($userPath) { "$InstallDir;$userPath" } else { $InstallDir }
  [Environment]::SetEnvironmentVariable('PATH', $newPath, 'User')
  $pathChanged = $true
}
# Make it work in this session too, so `configure` below and anything the user
# tries immediately afterwards can find the binary without reopening a console.
if ($env:PATH -notlike "*$InstallDir*") {
  $env:PATH = "$InstallDir;$env:PATH"
}

# --- launch the configurer --------------------------------------------------
Write-Host ""
try {
  & $Target configure
  if ($LASTEXITCODE -ne 0) {
    Write-Host "  Re-run '$Binary configure' later to finish setup." -ForegroundColor Yellow
    return
  }
} catch {
  Write-Host "  configure did not complete: $_" -ForegroundColor Red
  Write-Host "  Re-run '$Binary configure' later to finish setup." -ForegroundColor Yellow
  $global:LASTEXITCODE = 1
  return
}

# --- tell the user what to do next -----------------------------------------
Write-Host ""
if ($pathChanged) {
  Write-Host "  Added to your PATH: $InstallDir" -ForegroundColor Cyan
  Write-Host ""
  Write-Host "  Open a NEW PowerShell window before running $Binary." -ForegroundColor Yellow
  Write-Host "  This window still has the PATH it started with, so the command" -ForegroundColor Yellow
  Write-Host "  will not be found here." -ForegroundColor Yellow
  Write-Host ""
  Write-Host "  In a new window:  $Binary"
  Write-Host "  Or right now:     & '$Target'"
} else {
  Write-Host "  Run '$Binary' to browse and use your APIs."
}
$global:LASTEXITCODE = 0
