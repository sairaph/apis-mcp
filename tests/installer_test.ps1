$ErrorActionPreference = 'Stop'

$Root = Split-Path -Parent $PSScriptRoot
$Installer = Join-Path $Root 'install.ps1'
$ServerScript = Join-Path $PSScriptRoot 'installer_fixture_server.py'
$Temp = Join-Path ([System.IO.Path]::GetTempPath()) "apis-mcp-installer-test.$PID"
$FixtureRoot = Join-Path $Temp 'release'
$LocalAppData = Join-Path $Temp 'local-app-data'
$Target = Join-Path $LocalAppData 'apis-mcp/apis-mcp.exe'
$Asset = 'apis-mcp-windows-amd64.exe'
$AssetPath = Join-Path $FixtureRoot $Asset
$ManifestPath = Join-Path $FixtureRoot 'SHA256SUMS.txt'
$PowerShellExecutable = (Get-Process -Id $PID).Path
$Server = $null
$Passed = 0

function Assert-True([bool]$Condition, [string]$Message) {
  if (-not $Condition) { throw "PowerShell installer test failed: $Message" }
}

function Write-ValidManifest {
  $hash = (Get-FileHash -LiteralPath $AssetPath -Algorithm SHA256).Hash.ToLowerInvariant()
  [System.IO.File]::WriteAllText($ManifestPath, "$hash  $Asset`n")
}

$saved = @{
  CI = $env:CI
  LOCALAPPDATA = $env:LOCALAPPDATA
  NO_COLOR = $env:NO_COLOR
  PATH = $env:PATH
  TestBase = $env:APIS_MCP_INSTALLER_TEST_BASE_URL
  UserPath = [Environment]::GetEnvironmentVariable('PATH', 'User')
  ErrorPreference = $ErrorActionPreference
  ProgressPreference = $ProgressPreference
}

try {
  New-Item -ItemType Directory -Force -Path $FixtureRoot, (Split-Path -Parent $Target) | Out-Null
  [System.IO.File]::WriteAllBytes($AssetPath, [Text.Encoding]::UTF8.GetBytes('new fixture binary'))
  Write-ValidManifest

  $listener = [System.Net.Sockets.TcpListener]::new([Net.IPAddress]::Loopback, 0)
  $listener.Start()
  $port = ([System.Net.IPEndPoint]$listener.LocalEndpoint).Port
  $listener.Stop()

  $python = (Get-Command python -ErrorAction Stop).Source
  $Server = Start-Process -FilePath $python -ArgumentList @($ServerScript, '--root', $FixtureRoot, '--port', $port) -PassThru -WindowStyle Hidden
  $ready = $false
  for ($attempt = 0; $attempt -lt 50; $attempt++) {
    try {
      $client = [System.Net.Sockets.TcpClient]::new()
      $client.Connect('127.0.0.1', $port)
      $client.Dispose()
      $ready = $true
      break
    } catch {
      Start-Sleep -Milliseconds 100
    }
  }
  Assert-True $ready 'fixture server did not start'

  $env:CI = 'true'
  $env:LOCALAPPDATA = $LocalAppData
  $env:NO_COLOR = '1'
  $env:APIS_MCP_INSTALLER_TEST_BASE_URL = "http://127.0.0.1:$port"

  [System.IO.File]::WriteAllText($Target, 'working old binary')
  $successOutput = (& $Installer *>&1 | Out-String)
  Assert-True ($successOutput -match '\[ok\]\s+Release\s+v1\.2\.3') 'release helper output was not scalar and deterministic'
  Assert-True ($successOutput -notmatch 'HttpResponseMessage') 'HTTP response object leaked into installer output'
  Assert-True ($successOutput -match 'No interactive PowerShell host') 'noninteractive configuration guidance was missing'
  Assert-True ([System.IO.File]::ReadAllText($Target) -eq 'new fixture binary') 'successful upgrade did not replace the old binary'
  $Passed++

  $expectedError = $ErrorActionPreference
  $expectedProgress = $ProgressPreference
  [System.IO.File]::WriteAllText($Target, 'working old binary')
  [System.IO.File]::WriteAllText($ManifestPath, "$('0' * 64)  $Asset`n")
  $caught = $null
  try {
    Invoke-Expression ([System.IO.File]::ReadAllText($Installer))
  } catch {
    $caught = $_
  }
  Assert-True ($null -ne $caught) 'Invoke-Expression checksum failure was not terminating'
  Assert-True ($caught.Exception.Message -match '^  \[error\] Checksum mismatch') 'Invoke-Expression failure was not formatted once'
  Assert-True ([System.IO.File]::ReadAllText($Target) -eq 'working old binary') 'checksum failure replaced the old binary'
  Assert-True ($ErrorActionPreference -eq $expectedError) 'ErrorActionPreference was not restored'
  Assert-True ($ProgressPreference -eq $expectedProgress) 'ProgressPreference was not restored'
  $Passed++

  $stderr = Join-Path $Temp 'direct.stderr'
  $null = & $PowerShellExecutable -NoProfile -NonInteractive -File $Installer 2> $stderr
  $directStatus = $LASTEXITCODE
  $directError = [System.IO.File]::ReadAllText($stderr)
  Assert-True ($directStatus -ne 0) 'direct script failure returned success'
  Assert-True ($directError -match '\[error\] Checksum mismatch') 'direct script failure was not visible on stderr'
  Assert-True (([regex]::Matches($directError, '\[error\]')).Count -eq 1) 'direct script emitted duplicate formatted errors'
  $Passed++

  $iexRunner = Join-Path $Temp 'invoke-expression.ps1'
  [System.IO.File]::WriteAllText($iexRunner, 'Invoke-Expression ([System.IO.File]::ReadAllText($env:INSTALLER_UNDER_TEST))')
  $env:INSTALLER_UNDER_TEST = $Installer
  $iexStderr = Join-Path $Temp 'iex.stderr'
  $null = & $PowerShellExecutable -NoProfile -NonInteractive -File $iexRunner 2> $iexStderr
  $iexStatus = $LASTEXITCODE
  $iexError = [System.IO.File]::ReadAllText($iexStderr)
  Assert-True ($iexStatus -ne 0) 'Invoke-Expression process failure returned success'
  Assert-True ($iexError -match '\[error\] Checksum mismatch') 'Invoke-Expression process failure was not visible on stderr'
  $Passed++

  $tokens = $null
  $parseErrors = $null
  [void][System.Management.Automation.Language.Parser]::ParseFile($Installer, [ref]$tokens, [ref]$parseErrors)
  Assert-True ($parseErrors.Count -eq 0) 'install.ps1 did not parse'
  $Passed++
} finally {
  if ($null -ne $Server -and -not $Server.HasExited) { Stop-Process -Id $Server.Id -Force }
  [Environment]::SetEnvironmentVariable('PATH', $saved.UserPath, 'User')
  $env:CI = $saved.CI
  $env:LOCALAPPDATA = $saved.LOCALAPPDATA
  $env:NO_COLOR = $saved.NO_COLOR
  $env:PATH = $saved.PATH
  $env:APIS_MCP_INSTALLER_TEST_BASE_URL = $saved.TestBase
  Remove-Item Env:INSTALLER_UNDER_TEST -ErrorAction SilentlyContinue
  Remove-Item -LiteralPath $Temp -Recurse -Force -ErrorAction SilentlyContinue
}

Write-Host "PowerShell installer tests passed: $Passed"
