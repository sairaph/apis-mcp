# apis-mcp installer for Windows PowerShell.
$previousErrorPreference = $ErrorActionPreference
$previousProgressPreference = $ProgressPreference

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
  function Write-Phase([string]$Name, [string]$Text) {
    Write-InstallerLine ("  [..] {0,-10} {1}" -f $Name, $Text)
  }
  function Write-Success([string]$Name, [string]$Text) {
    Write-InstallerLine ("  [ok] {0,-10} {1}" -f $Name, $Text) Green
  }
  function Write-Guidance([string]$Text) {
    Write-InstallerLine "  [--] $Text" Yellow
  }
  function Get-ConcreteReleaseUrl {
    param([System.Net.Http.HttpClient]$Client, [string]$Url)
    $response = $null
    try {
      $response = $Client.GetAsync($Url, [System.Net.Http.HttpCompletionOption]::ResponseHeadersRead).GetAwaiter().GetResult()
      $null = $response.EnsureSuccessStatusCode()
      return $response.RequestMessage.RequestUri.AbsoluteUri.TrimEnd('/')
    } catch {
      throw "Could not resolve the latest release. URL: $Url. $($_.Exception.Message)"
    } finally {
      if ($null -ne $response) { $response.Dispose() }
    }
  }

  function Invoke-InstallerDownload {
    param(
      [System.Net.Http.HttpClient]$Client,
      [string]$Url,
      [string]$Destination,
      [bool]$ShowProgress
    )
    $response = $null
    $stream = $null
    $file = $null
    try {
      $response = $Client.GetAsync($Url, [System.Net.Http.HttpCompletionOption]::ResponseHeadersRead).GetAwaiter().GetResult()
      $null = $response.EnsureSuccessStatusCode()
      $total = if ($null -ne $response.Content.Headers.ContentLength) {
        [long]$response.Content.Headers.ContentLength
      } else {
        [long]-1
      }
      $stream = $response.Content.ReadAsStreamAsync().GetAwaiter().GetResult()
      $file = [System.IO.File]::Open($Destination, [System.IO.FileMode]::CreateNew, [System.IO.FileAccess]::Write, [System.IO.FileShare]::None)
      $buffer = New-Object byte[] 65536
      [long]$downloaded = 0
      $stopwatch = [System.Diagnostics.Stopwatch]::StartNew()
      $lastRender = [TimeSpan]::Zero

      while (($read = $stream.Read($buffer, 0, $buffer.Length)) -gt 0) {
        $file.Write($buffer, 0, $read)
        $downloaded += $read
        if ($ShowProgress -and $interactiveHost -and $total -gt 0 -and
            ($stopwatch.Elapsed - $lastRender).TotalMilliseconds -ge 125) {
          $lastRender = $stopwatch.Elapsed
          $percent = [Math]::Min(100, [Math]::Floor(($downloaded * 100.0) / $total))
          try { $consoleWidth = [Console]::WindowWidth } catch { $consoleWidth = 80 }
          $barWidth = [Math]::Max(10, [Math]::Min(30, $consoleWidth - 48))
          $filled = [Math]::Min($barWidth, [Math]::Floor(($percent * $barWidth) / 100))
          $bar = ('#' * $filled).PadRight($barWidth, '-')
          $line = "  [$bar] {0,3}%  {1:N1}/{2:N1} MB" -f $percent, ($downloaded / 1MB), ($total / 1MB)
          if ($line.Length -ge $consoleWidth) { $line = $line.Substring(0, [Math]::Max(1, $consoleWidth - 1)) }
          Write-Host ("`r{0}" -f $line.PadRight([Math]::Max($line.Length, $consoleWidth - 1))) -NoNewline
        }
      }
      $file.Flush()
      if ($total -ge 0 -and $downloaded -ne $total) {
        throw "The response was truncated: expected $total bytes, received $downloaded."
      }
      if ($ShowProgress -and $interactiveHost -and $total -gt 0) {
        try { $consoleWidth = [Console]::WindowWidth } catch { $consoleWidth = 80 }
        $barWidth = [Math]::Max(10, [Math]::Min(30, $consoleWidth - 48))
        $line = "  [{0}] 100%  {1:N1}/{1:N1} MB" -f ('#' * $barWidth), ($downloaded / 1MB)
        Write-Host ("`r{0}" -f $line.PadRight([Math]::Max($line.Length, $consoleWidth - 1)))
      }
      return $downloaded
    } catch {
      throw "Download failed. URL: $Url. $($_.Exception.Message)"
    } finally {
      # Close every stream before callers inspect, remove, or replace files.
      if ($null -ne $file) { $file.Dispose() }
      if ($null -ne $stream) { $stream.Dispose() }
      if ($null -ne $response) { $response.Dispose() }
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
  $isWindows = $env:OS -eq 'Windows_NT'
  try {
    $isWindows = [Runtime.InteropServices.RuntimeInformation]::IsOSPlatform([Runtime.InteropServices.OSPlatform]::Windows)
    $osDescription = [Runtime.InteropServices.RuntimeInformation]::OSDescription
    $architecture = [Runtime.InteropServices.RuntimeInformation]::OSArchitecture.ToString()
  } catch {
    $architecture = if ($env:PROCESSOR_ARCHITEW6432) { $env:PROCESSOR_ARCHITEW6432 } else { $env:PROCESSOR_ARCHITECTURE }
  }
  if (-not $isWindows) {
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
  Write-Success 'Platform' "$osDescription/$architecture -> $Asset"

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
  $client = $null

  try {
    New-Item -ItemType Directory -Force -Path $InstallDir | Out-Null
    $handler = New-Object System.Net.Http.HttpClientHandler
    $handler.AllowAutoRedirect = $true
    $client = New-Object System.Net.Http.HttpClient -ArgumentList $handler
    $client.DefaultRequestHeaders.UserAgent.ParseAdd('apis-mcp-installer')

    Write-Phase 'Release' 'Resolving the latest concrete release...'
    $releaseUrl = Get-ConcreteReleaseUrl $client $LatestUrl
    if (-not $releaseUrl.StartsWith($ReleasePrefix, [StringComparison]::OrdinalIgnoreCase)) {
      throw "GitHub returned an unexpected latest-release URL: $releaseUrl"
    }
    $tag = $releaseUrl.Substring($ReleasePrefix.Length)
    if ($tag -notmatch '^v[A-Za-z0-9._-]+$') {
      throw "GitHub returned a release tag outside the required v<version> contract: $tag"
    }
    Write-Success 'Release' $tag

    $baseUrl = "$DownloadRoot/$tag"
    $assetUrl = "$baseUrl/$Asset"
    $manifestUrl = "$baseUrl/SHA256SUMS.txt"

    Write-Phase 'Download' $Asset
    $assetBytes = Invoke-InstallerDownload $client $assetUrl $staged $true
    if ($assetBytes -le 0 -or (Get-Item -LiteralPath $staged).Length -le 0) {
      throw "Download was empty or incomplete; the existing binary was not replaced. URL: $assetUrl"
    }

    Write-Phase 'Verify' 'Downloading SHA256SUMS.txt...'
    $manifestBytes = Invoke-InstallerDownload $client $manifestUrl $manifest $false
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
    Write-Success 'Verify' "SHA-256 $actual"

    try {
      if (Test-Path -LiteralPath $Target) {
        [System.IO.File]::Replace($staged, $Target, $null, $true)
      } else {
        [System.IO.File]::Move($staged, $Target)
      }
      $staged = $null
    } catch {
      throw "Could not replace $Target. Close any running apis-mcp process, then re-run the installer. If it remains locked, check antivirus activity. $($_.Exception.Message)"
    }
    Write-Success 'Install' $Target

    $userPath = [Environment]::GetEnvironmentVariable('PATH', 'User')
    $pathPersisted = Test-ExactPathComponent $userPath $InstallDir
    if (-not $pathPersisted) {
      try {
        $newPath = if ([string]::IsNullOrWhiteSpace($userPath)) { $InstallDir } else { "$InstallDir;$userPath" }
        [Environment]::SetEnvironmentVariable('PATH', $newPath, 'User')
        $pathPersisted = $true
        Write-Success 'PATH' "Saved $InstallDir"
      } catch {
        throw "Installed $Target, but could not save the user PATH. Add $InstallDir manually. $($_.Exception.Message)"
      }
    }
    if (-not (Test-ExactPathComponent $env:PATH $InstallDir)) {
      $env:PATH = if ([string]::IsNullOrWhiteSpace($env:PATH)) { $InstallDir } else { "$InstallDir;$env:PATH" }
    }
    if ($pathPersisted) {
      Write-Guidance 'apis-mcp is available in this PowerShell session and in new shells.'
    }

    if ($interactiveHost) {
      Write-Phase 'Configure' 'Starting interactive client setup...'
      $configureError = $null
      try {
        & $Target configure
      } catch {
        $configureError = $_
      }
      $configureStatus = $LASTEXITCODE
      if ($null -eq $configureError -and $configureStatus -eq 0) {
        Write-Success 'Configure' 'Complete'
      } else {
        Write-Guidance "Configure was cancelled or did not complete. Run: & '$Target' configure"
        $detail = if ($null -ne $configureError) { $configureError.Exception.Message } else { "exit status $configureStatus" }
        throw "apis-mcp was installed, but configuration did not complete ($detail)."
      }
    } else {
      Write-Guidance 'No interactive PowerShell host; configuration was not started.'
      Write-Guidance "Finish setup from a terminal: & '$Target' configure"
    }
  } finally {
    # Download helpers dispose streams before control reaches this cleanup.
    $cleanupFailures = @()
    if ($null -ne $client) {
      try { $client.Dispose() } catch { $cleanupFailures += $_.Exception.Message }
    }
    foreach ($temporaryPath in @($staged, $manifest)) {
      if ($temporaryPath -and (Test-Path -LiteralPath $temporaryPath)) {
        try {
          Remove-Item -LiteralPath $temporaryPath -Force -ErrorAction Stop
        } catch {
          $cleanupFailures += "${temporaryPath}: $($_.Exception.Message)"
        }
      }
    }
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
  $formattedFailure = "  [error] $($installerFailure.Exception.Message)"
  throw [System.InvalidOperationException]::new($formattedFailure)
}
