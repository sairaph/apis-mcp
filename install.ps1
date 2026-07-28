# apis-mcp installer for Windows PowerShell.
$previousErrorPreference = $ErrorActionPreference
$previousProgressPreference = $ProgressPreference
$installerRunFromFile = (
  -not [string]::IsNullOrWhiteSpace($MyInvocation.MyCommand.Path) -and
  [string]::Equals([System.IO.Path]::GetFileName($MyInvocation.MyCommand.Path), 'install.ps1', [StringComparison]::OrdinalIgnoreCase)
)

function Install-ApisMcp {
  $Owner = 'sairaph'
  $Repo = 'apis-mcp'
  $Binary = 'apis-mcp'
  $LatestUrl = "https://github.com/$Owner/$Repo/releases/latest"
  $ReleasePrefix = "https://github.com/$Owner/$Repo/releases/tag/"
  $DownloadRoot = "https://github.com/$Owner/$Repo/releases/download"

  $interactiveHost = [Environment]::UserInteractive
  try {
    $interactiveHost = $interactiveHost -and -not [Console]::IsOutputRedirected
    $null = $Host.UI.RawUI.WindowSize
  } catch {
    $interactiveHost = $false
  }
  $useColor = $interactiveHost -and -not (Test-Path Env:NO_COLOR)

  function Write-InstallerLine {
    param([string]$Text, [ConsoleColor]$Color = [ConsoleColor]::Gray)
    if ($useColor) { Write-Host $Text -ForegroundColor $Color } else { Write-Host $Text }
  }
  function Write-ProgressLine {
    param([string]$Text, [bool]$Complete)
    try { $width = $Host.UI.RawUI.WindowSize.Width } catch { $width = 80 }
    $width = [Math]::Max(1, $width - 1)
    if ($Text.Length -gt $width) { $Text = $Text.Substring(0, $width) }
    $rendered = "`r$($Text.PadRight($width))"
    if ($Complete) { Write-Host $rendered } else { Write-Host $rendered -NoNewline }
  }
  function Close-InstallerErrorResponse {
    param([System.Management.Automation.ErrorRecord]$ErrorRecord)
    $exception = $ErrorRecord.Exception
    while ($null -ne $exception) {
      $property = $exception.PSObject.Properties['Response']
      if ($null -ne $property -and $null -ne $property.Value) {
        try { $property.Value.Close() } catch {}
        return
      }
      $exception = $exception.InnerException
    }
  }
  function Get-ConcreteReleaseUrl {
    param([string]$Url)
    $request = $null
    $response = $null
    try {
      $request = [System.Net.HttpWebRequest]::Create($Url)
      $request.Method = 'GET'
      $request.AllowAutoRedirect = $true
      $request.UserAgent = 'apis-mcp-installer'
      $response = $request.GetResponse()
      return $response.ResponseUri.AbsoluteUri.TrimEnd('/')
    } catch {
      Close-InstallerErrorResponse $_
      throw "Could not resolve the latest release. URL: $Url. $($_.Exception.Message)"
    } finally {
      if ($null -ne $response) { $response.Close() }
    }
  }

  function Invoke-InstallerDownload {
    param(
      [string]$Url,
      [string]$Destination,
      [bool]$ShowProgress
    )
    $request = $null
    $response = $null
    $stream = $null
    $file = $null
    try {
      $request = [System.Net.HttpWebRequest]::Create($Url)
      $request.Method = 'GET'
      $request.AllowAutoRedirect = $true
      $request.UserAgent = 'apis-mcp-installer'
      $response = $request.GetResponse()
      [long]$total = $response.ContentLength
      $stream = $response.GetResponseStream()
      $file = [System.IO.File]::Open($Destination, [System.IO.FileMode]::CreateNew, [System.IO.FileAccess]::Write, [System.IO.FileShare]::None)
      $buffer = New-Object byte[] 65536
      [long]$downloaded = 0
      $stopwatch = [System.Diagnostics.Stopwatch]::StartNew()
      [long]$lastRender = 0

      while (($read = $stream.Read($buffer, 0, $buffer.Length)) -gt 0) {
        $file.Write($buffer, 0, $read)
        $downloaded += $read
        $elapsedMilliseconds = $stopwatch.ElapsedMilliseconds
        if ($ShowProgress -and $interactiveHost -and $total -gt 0 -and $elapsedMilliseconds - $lastRender -ge 150) {
          $lastRender = $elapsedMilliseconds
          $percent = [Math]::Min(100, [Math]::Round(($downloaded * 100.0) / $total))
          $filled = [Math]::Min(20, [Math]::Floor($percent / 5))
          $bar = ('#' * $filled).PadRight(20, '-')
          $downloadedMB = [Math]::Round($downloaded / 1MB, 1)
          $totalMB = [Math]::Round($total / 1MB, 1)
          $speed = if ($stopwatch.Elapsed.TotalSeconds -gt 0) { [Math]::Round($downloadedMB / $stopwatch.Elapsed.TotalSeconds, 1) } else { 0 }
          $eta = if ($speed -gt 0) { [Math]::Max(0, [Math]::Round(($totalMB - $downloadedMB) / $speed)) } else { 0 }
          $line = "  [$bar] {0,3}%  {1}/{2} MB  {3} MB/s  ETA {4:00}s" -f $percent, $downloadedMB, $totalMB, $speed, $eta
          Write-ProgressLine $line $false
        }
      }
      $file.Flush()
      if ($total -ge 0 -and $downloaded -ne $total) {
        throw "The response was truncated: expected $total bytes, received $downloaded."
      }
      if ($ShowProgress -and $interactiveHost -and $total -gt 0) {
        $downloadedMB = [Math]::Round($downloaded / 1MB, 1)
        $speed = if ($stopwatch.Elapsed.TotalSeconds -gt 0) { [Math]::Round($downloadedMB / $stopwatch.Elapsed.TotalSeconds, 1) } else { 0 }
        $line = "  [####################] 100%  {0}/{0} MB  {1} MB/s  ETA 00s" -f $downloadedMB, $speed
        Write-ProgressLine $line $true
      }
      return $downloaded
    } catch {
      Close-InstallerErrorResponse $_
      throw "Download failed. URL: $Url. $($_.Exception.Message)"
    } finally {
      # Close every stream before callers inspect, remove, or replace files.
      if ($null -ne $file) { $file.Dispose() }
      if ($null -ne $stream) { $stream.Dispose() }
      if ($null -ne $response) { $response.Close() }
    }
  }

  function Test-ExactPathComponent {
    param([string]$PathValue, [string]$Component)
    if ([string]::IsNullOrWhiteSpace($PathValue)) { return $false }
    $comparison = [StringComparison]::OrdinalIgnoreCase
    $wanted = [System.IO.Path]::GetFullPath($Component).TrimEnd([char[]]'\/')
    foreach ($entry in ($PathValue -split [regex]::Escape([System.IO.Path]::PathSeparator))) {
      $entry = $entry.Trim().Trim('"')
      if ([string]::IsNullOrWhiteSpace($entry)) { continue }
      try {
        $candidate = [System.IO.Path]::GetFullPath($entry).TrimEnd([char[]]'\/')
        if ([string]::Equals($candidate, $wanted, $comparison)) { return $true }
      } catch {}
    }
    return $false
  }

  Write-Host ''
  Write-InstallerLine '  apis-mcp installer' Cyan
  Write-Host ''

  if (-not [string]::IsNullOrWhiteSpace($env:APIS_MCP_INSTALLER_TEST_BASE_URL)) {
    if ($env:CI -ne 'true') {
      throw 'APIS_MCP_INSTALLER_TEST_BASE_URL is restricted to CI test runs.'
    }
    try { $testUri = [Uri]$env:APIS_MCP_INSTALLER_TEST_BASE_URL } catch { throw 'The installer test base URL is invalid.' }
    if (-not $testUri.IsAbsoluteUri -or $testUri.Scheme -ne 'http' -or
        $testUri.Host -notin @('127.0.0.1', 'localhost', '::1')) {
      throw 'The installer test base URL must be an absolute loopback HTTP URL.'
    }
    $testBase = $testUri.AbsoluteUri.TrimEnd('/')
    $LatestUrl = "$testBase/releases/latest"
    $ReleasePrefix = "$testBase/releases/tag/"
    $DownloadRoot = "$testBase/releases/download"
    $interactiveHost = $false
    $useColor = $false
  }

  $osDescription = [Environment]::OSVersion.VersionString
  $runningOnWindows = $env:OS -eq 'Windows_NT'
  try {
    $runningOnWindows = [Runtime.InteropServices.RuntimeInformation]::IsOSPlatform([Runtime.InteropServices.OSPlatform]::Windows)
    $osDescription = [Runtime.InteropServices.RuntimeInformation]::OSDescription
    $architecture = [Runtime.InteropServices.RuntimeInformation]::OSArchitecture.ToString()
  } catch {
    $architecture = if ($env:PROCESSOR_ARCHITEW6432) { $env:PROCESSOR_ARCHITEW6432 } else { $env:PROCESSOR_ARCHITECTURE }
  }
  if (-not $runningOnWindows) {
    throw "Unsupported platform: OS=$osDescription, architecture=$architecture. This installer supports Windows."
  }
  $arch = switch -Regex ($architecture) {
    '^(AMD64|x64|X64)$' { 'amd64'; break }
    '^ARM64$' { 'arm64'; break }
    default {
      throw "Unsupported platform: OS=$osDescription, architecture=$architecture. Supported architectures are amd64 and arm64."
    }
  }
  $Asset = "$Binary-windows-$arch.exe"

  if (-not (Get-Command Get-FileHash -ErrorAction SilentlyContinue)) {
    throw 'Missing required tool: Get-FileHash is needed to verify SHA-256 checksums.'
  }
  if ([string]::IsNullOrWhiteSpace($env:LOCALAPPDATA)) {
    throw 'LOCALAPPDATA is not set; cannot choose a per-user install directory.'
  }

  $InstallDir = Join-Path $env:LOCALAPPDATA $Repo
  $Target = Join-Path $InstallDir "$Binary.exe"
  $suffix = "$PID.$([Guid]::NewGuid().ToString('N'))"
  $staged = Join-Path $InstallDir ".$Binary.new.$suffix"
  $manifest = Join-Path $InstallDir ".SHA256SUMS.txt.$suffix"
  $backup = Join-Path $InstallDir ".$Binary.backup.$suffix"
  $previousSecurityProtocol = [Net.ServicePointManager]::SecurityProtocol

  try {
    New-Item -ItemType Directory -Force -Path $InstallDir | Out-Null
    try {
      [Net.ServicePointManager]::SecurityProtocol = [Net.ServicePointManager]::SecurityProtocol -bor [Net.SecurityProtocolType]::Tls12
    } catch {}

    Write-InstallerLine '  Finding the latest release...' DarkGray
    $releaseUrl = Get-ConcreteReleaseUrl $LatestUrl
    if (-not $releaseUrl.StartsWith($ReleasePrefix, [StringComparison]::OrdinalIgnoreCase)) {
      throw "GitHub returned an unexpected latest-release URL: $releaseUrl"
    }
    $tag = $releaseUrl.Substring($ReleasePrefix.Length)
    if ($tag -notmatch '^v[A-Za-z0-9._-]+$') {
      throw "GitHub returned a release tag outside the required v<version> contract: $tag"
    }
    $baseUrl = "$DownloadRoot/$tag"
    $assetUrl = "$baseUrl/$Asset"
    $manifestUrl = "$baseUrl/SHA256SUMS.txt"

    Write-InstallerLine "  Downloading $Asset from $tag..." Cyan
    Write-Host ''
    $assetBytes = Invoke-InstallerDownload $assetUrl $staged $true
    if ($assetBytes -le 0 -or (Get-Item -LiteralPath $staged).Length -le 0) {
      throw "Download was empty or incomplete; the existing binary was not replaced. URL: $assetUrl"
    }

    Write-InstallerLine '  Verifying SHA-256 checksum...' DarkGray
    $manifestBytes = Invoke-InstallerDownload $manifestUrl $manifest $false
    if ($manifestBytes -le 0) { throw "Checksum manifest was empty. URL: $manifestUrl" }

    $expectedHashes = @()
    foreach ($line in [System.IO.File]::ReadAllLines($manifest)) {
      if ($line -match '^([0-9A-Fa-f]{64})[ \t]+\*?(.+)$' -and
          [string]::Equals($Matches[2], $Asset, [StringComparison]::Ordinal)) {
        $expectedHashes += $Matches[1].ToLowerInvariant()
      }
    }
    if ($expectedHashes.Count -ne 1) {
      throw "SHA256SUMS.txt does not contain exactly one valid checksum for $Asset. URL: $manifestUrl"
    }
    $actual = (Get-FileHash -LiteralPath $staged -Algorithm SHA256).Hash.ToLowerInvariant()
    if ($actual -ne $expectedHashes[0]) {
      throw "Checksum mismatch for $Asset; the existing binary was not replaced."
    }
    Write-InstallerLine '  Checksum verified.' Green

    try {
      if (Test-Path -LiteralPath $Target) {
        [System.IO.File]::Replace($staged, $Target, $backup, $true)
      } else {
        [System.IO.File]::Move($staged, $Target)
      }
      $staged = $null
    } catch {
      throw "Could not replace $Target. Close any running apis-mcp process, then re-run the installer. If it remains locked, check antivirus activity. $($_.Exception.Message)"
    }
    Write-InstallerLine "  Installed to $Target" Green

    $pathChanged = $false
    $userPath = [Environment]::GetEnvironmentVariable('PATH', 'User')
    $pathPersisted = Test-ExactPathComponent $userPath $InstallDir
    if (-not $pathPersisted) {
      try {
        $newPath = if ([string]::IsNullOrWhiteSpace($userPath)) { $InstallDir } else { "$InstallDir;$userPath" }
        [Environment]::SetEnvironmentVariable('PATH', $newPath, 'User')
        $pathPersisted = $true
        $pathChanged = $true
      } catch {
        throw "Installed $Target, but could not save the user PATH. Add $InstallDir manually. $($_.Exception.Message)"
      }
    }
    if (-not (Test-ExactPathComponent $env:PATH $InstallDir)) {
      $env:PATH = if ([string]::IsNullOrWhiteSpace($env:PATH)) { $InstallDir } else { "$InstallDir;$env:PATH" }
    }
    Write-Host ''
    if ($interactiveHost) {
      Write-InstallerLine '  Opening setup...' Cyan
      $configureError = $null
      try {
        & $Target configure
      } catch {
        $configureError = $_
      }
      $configureStatus = $LASTEXITCODE
      if ($null -eq $configureError -and $configureStatus -eq 0) {
        Write-Host ''
      } else {
        $detail = if ($null -ne $configureError) { $configureError.Exception.Message } else { "exit status $configureStatus" }
        Write-InstallerLine "  Setup did not complete ($detail)." Yellow
        Write-InstallerLine "  Run '$Binary configure' whenever you are ready." Yellow
      }
    } else {
      Write-InstallerLine '  Setup needs an interactive PowerShell window.' Yellow
      Write-InstallerLine "  Finish later with: & '$Target' configure" Yellow
    }

    Write-Host ''
    Write-InstallerLine '  apis-mcp is ready.' Green
    if ($pathChanged) {
      Write-InstallerLine "  Added to your PATH: $InstallDir" Cyan
    }
    Write-InstallerLine "  Run now: $Binary" Gray
  } finally {
    # Download helpers dispose streams before control reaches this cleanup.
    $cleanupFailures = @()
    foreach ($temporaryPath in @($staged, $manifest, $backup)) {
      if ($temporaryPath -and (Test-Path -LiteralPath $temporaryPath)) {
        try {
          Remove-Item -LiteralPath $temporaryPath -Force -ErrorAction Stop
        } catch {
          $cleanupFailures += "${temporaryPath}: $($_.Exception.Message)"
        }
      }
    }
    [Net.ServicePointManager]::SecurityProtocol = $previousSecurityProtocol
    if ($cleanupFailures.Count -gt 0) {
      throw "Could not clean installer temporary resources. Remove them manually, then retry: $($cleanupFailures -join '; ')"
    }
  }
}

$installerFailure = $null
try {
  $ErrorActionPreference = 'Stop'
  $ProgressPreference = 'SilentlyContinue'
  Install-ApisMcp
} catch {
  $installerFailure = $_
} finally {
  $ErrorActionPreference = $previousErrorPreference
  $ProgressPreference = $previousProgressPreference
}
if ($null -ne $installerFailure) {
  Write-Host ''
  if (Test-Path Env:NO_COLOR) {
    Write-Host "  Installation failed: $($installerFailure.Exception.Message)"
    Write-Host '  Fix the issue above and run the installer again.'
  } else {
    Write-Host "  Installation failed: $($installerFailure.Exception.Message)" -ForegroundColor Red
    Write-Host '  Fix the issue above and run the installer again.' -ForegroundColor Yellow
  }
  $global:LASTEXITCODE = 1
  if ($installerRunFromFile) { exit 1 }
  Write-Error 'apis-mcp installation failed' -ErrorAction SilentlyContinue
} else {
  $global:LASTEXITCODE = 0
}
