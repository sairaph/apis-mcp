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
  SecurityProtocol = [Net.ServicePointManager]::SecurityProtocol
}

try {
  New-Item -ItemType Directory -Force -Path $FixtureRoot, (Split-Path -Parent $Target) | Out-Null
  [System.IO.File]::WriteAllBytes($AssetPath, [Text.Encoding]::UTF8.GetBytes('new fixture binary'))
  [System.IO.File]::Copy($Installer, (Join-Path $FixtureRoot 'install.ps1'))
  Write-ValidManifest

  $listener = [System.Net.Sockets.TcpListener]::new([Net.IPAddress]::Loopback, 0)
  $listener.Start()
  $port = ([System.Net.IPEndPoint]$listener.LocalEndpoint).Port
  $listener.Stop()

  $python = (Get-Command python -ErrorAction Stop).Source
  $Server = Start-Process -FilePath $python -ArgumentList @($ServerScript, '--root', $FixtureRoot, '--port', $port) -PassThru -WindowStyle Hidden
  $ready = $false
  for ($attempt = 0; $attempt -lt 50; $attempt++) {
    $client = $null
    try {
      $client = [System.Net.Sockets.TcpClient]::new()
      $client.Connect('127.0.0.1', $port)
      $ready = $true
      break
    } catch {
      Start-Sleep -Milliseconds 100
    } finally {
      if ($null -ne $client) { $client.Dispose() }
    }
  }
  Assert-True $ready 'fixture server did not start'

  $env:CI = 'true'
  $env:LOCALAPPDATA = $LocalAppData
  $env:NO_COLOR = '1'
  $env:APIS_MCP_INSTALLER_TEST_BASE_URL = "http://127.0.0.1:$port"

  [System.IO.File]::WriteAllText($Target, 'working old binary')
  $successOutput = (& $Installer *>&1 | Out-String)
  Assert-True ($successOutput -match 'Downloading apis-mcp-windows-amd64\.exe from v1\.2\.3') 'release and download were not presented clearly'
  Assert-True ($successOutput -match 'Checksum verified\.' -and $successOutput -match 'apis-mcp is ready\.') 'successful install summary was incomplete'
  Assert-True ($successOutput -notmatch 'HttpResponseMessage|\[(?:ok|\.\.)\]') 'installer leaked transport objects or status-table noise'
  Assert-True ($successOutput -match 'Setup needs an interactive PowerShell window') 'noninteractive setup guidance was missing'
  Assert-True ([System.IO.File]::ReadAllText($Target) -eq 'new fixture binary') 'successful upgrade did not replace the old binary'
  Assert-True (-not (Get-ChildItem -LiteralPath (Split-Path -Parent $Target) -Filter '.apis-mcp.backup.*')) 'successful upgrade left a backup file behind'
  Assert-True (($env:PATH -split ';') -contains (Split-Path -Parent $Target)) 'current process PATH was not updated exactly'
  Assert-True (([Environment]::GetEnvironmentVariable('PATH', 'User') -split ';') -contains (Split-Path -Parent $Target)) 'user PATH was not persisted exactly'
  Assert-True ([Net.ServicePointManager]::SecurityProtocol -eq $saved.SecurityProtocol) 'successful install leaked its TLS setting'
  $Passed++

  $pipelineRunner = Join-Path $Temp 'pipeline.ps1'
  $pipelineURL = "http://127.0.0.1:$port/install.ps1"
  [System.IO.File]::WriteAllText($pipelineRunner, "Invoke-RestMethod '$pipelineURL' | Invoke-Expression`nexit `$global:LASTEXITCODE`n")
  $pipelineStdout = Join-Path $Temp 'pipeline.stdout'
  $pipelineStderr = Join-Path $Temp 'pipeline.stderr'
  $null = & $PowerShellExecutable -NoProfile -NonInteractive -File $pipelineRunner > $pipelineStdout 2> $pipelineStderr
  $pipelineStatus = $LASTEXITCODE
  $pipelineOutput = [System.IO.File]::ReadAllText($pipelineStdout) + [System.IO.File]::ReadAllText($pipelineStderr)
  Assert-True ($pipelineStatus -eq 0) "Invoke-RestMethod pipeline failed: $pipelineOutput"
  Assert-True ($pipelineOutput -match 'Downloading apis-mcp-windows-amd64\.exe from v1\.2\.3') 'real installer pipeline did not run'
  Assert-True ([System.IO.File]::ReadAllText($Target) -eq 'new fixture binary') 'pipeline upgrade did not preserve the verified binary'
  $Passed++

  $expectedError = $ErrorActionPreference
  $expectedProgress = $ProgressPreference
  [System.IO.File]::WriteAllText($Target, 'working old binary')
  [System.IO.File]::WriteAllText($ManifestPath, "$('0' * 64)  $Asset`n")
  $global:LASTEXITCODE = 0
  $failureLog = Join-Path $Temp 'failure.log'
  Invoke-Expression ([System.IO.File]::ReadAllText($Installer)) *> $failureLog
  $failureOutput = [System.IO.File]::ReadAllText($failureLog)
  Assert-True ($global:LASTEXITCODE -eq 1) 'Invoke-Expression checksum failure did not set a failing status'
  Assert-True ($failureOutput -match 'Installation failed: Checksum mismatch') 'Invoke-Expression failure was not concise'
  Assert-True ($failureOutput -notmatch 'InvalidOperationException|FullyQualifiedErrorId|At line:') 'Invoke-Expression failure emitted a PowerShell exception trace'
  Assert-True ([System.IO.File]::ReadAllText($Target) -eq 'working old binary') 'checksum failure replaced the old binary'
  Assert-True ($ErrorActionPreference -eq $expectedError) 'ErrorActionPreference was not restored'
  Assert-True ($ProgressPreference -eq $expectedProgress) 'ProgressPreference was not restored'
  Assert-True ([Net.ServicePointManager]::SecurityProtocol -eq $saved.SecurityProtocol) 'failed install leaked its TLS setting'
  $Passed++

  $stderr = Join-Path $Temp 'direct.stderr'
  $stdout = Join-Path $Temp 'direct.stdout'
  $null = & $PowerShellExecutable -NoProfile -NonInteractive -File $Installer > $stdout 2> $stderr
  $directStatus = $LASTEXITCODE
  $directError = [System.IO.File]::ReadAllText($stdout) + [System.IO.File]::ReadAllText($stderr)
  Assert-True ($directStatus -ne 0) 'direct script failure returned success'
  Assert-True ($directError -match 'Installation failed: Checksum mismatch') 'direct script failure was not visible'
  Assert-True ($directError -notmatch 'InvalidOperationException|FullyQualifiedErrorId|At line:') 'direct script emitted an exception trace'
  $Passed++

  $env:INSTALLER_UNDER_TEST = $Installer
  $iexRunner = Join-Path $Temp 'invoke-expression.ps1'
  [System.IO.File]::WriteAllText($iexRunner, "Invoke-Expression ([System.IO.File]::ReadAllText(`$env:INSTALLER_UNDER_TEST))`nexit `$global:LASTEXITCODE`n")
  $iexStderr = Join-Path $Temp 'iex.stderr'
  $iexStdout = Join-Path $Temp 'iex.stdout'
  $null = & $PowerShellExecutable -NoProfile -NonInteractive -File $iexRunner > $iexStdout 2> $iexStderr
  $iexStatus = $LASTEXITCODE
  $iexError = [System.IO.File]::ReadAllText($iexStdout) + [System.IO.File]::ReadAllText($iexStderr)
  Assert-True ($iexStatus -ne 0) 'Invoke-Expression process failure returned success'
  Assert-True ($iexError -match 'Installation failed: Checksum mismatch') 'Invoke-Expression process failure was not visible'
  $Passed++

  $tokens = $null
  $parseErrors = $null
  [void][System.Management.Automation.Language.Parser]::ParseFile($Installer, [ref]$tokens, [ref]$parseErrors)
  Assert-True ($parseErrors.Count -eq 0) 'install.ps1 did not parse'
  $installerSource = [System.IO.File]::ReadAllText($Installer)
  Assert-True ($installerSource -notmatch 'System\.Net\.Http\.') 'installer still requires the PowerShell 7 HTTP assembly'
  Assert-True ($installerSource -match 'System\.Net\.HttpWebRequest') 'Windows PowerShell-compatible transport is missing'
  $Passed++
} finally {
  if ($null -ne $Server -and -not $Server.HasExited) { Stop-Process -Id $Server.Id -Force }
  [Environment]::SetEnvironmentVariable('PATH', $saved.UserPath, 'User')
  $env:CI = $saved.CI
  $env:LOCALAPPDATA = $saved.LOCALAPPDATA
  $env:NO_COLOR = $saved.NO_COLOR
  $env:PATH = $saved.PATH
  $env:APIS_MCP_INSTALLER_TEST_BASE_URL = $saved.TestBase
  [Net.ServicePointManager]::SecurityProtocol = $saved.SecurityProtocol
  Remove-Item Env:INSTALLER_UNDER_TEST -ErrorAction SilentlyContinue
  Remove-Item -LiteralPath $Temp -Recurse -Force -ErrorAction SilentlyContinue
}

Write-Host "PowerShell installer tests passed: $Passed"
exit 0
