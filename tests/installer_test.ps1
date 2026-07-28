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
$PowerShellExecutable = (Get-Process -Id $PID).Path
$Server = $null
$Passed = 0

function Assert-True([bool]$Condition, [string]$Message) {
  if (-not $Condition) { throw "PowerShell installer test failed: $Message" }
}

function Invoke-InstallerPipeline([string]$Name) {
  $runner = Join-Path $Temp "$Name.ps1"
  $stdout = Join-Path $Temp "$Name.stdout"
  $stderr = Join-Path $Temp "$Name.stderr"
  [System.IO.File]::WriteAllText($runner, "Invoke-RestMethod 'http://127.0.0.1:$port/install.ps1' | Invoke-Expression`nexit `$LASTEXITCODE`n")
  $null = & $PowerShellExecutable -NoProfile -NonInteractive -File $runner > $stdout 2> $stderr
  return @{
    Status = $LASTEXITCODE
    Output = [System.IO.File]::ReadAllText($stdout) + [System.IO.File]::ReadAllText($stderr)
  }
}

$saved = @{
  LocalAppData = $env:LOCALAPPDATA
  Path = $env:PATH
  ProcessorArchitecture = $env:PROCESSOR_ARCHITECTURE
  ConfigureStatus = $env:CONFIGURE_STATUS
  UserPath = [Environment]::GetEnvironmentVariable('PATH', 'User')
}

try {
  New-Item -ItemType Directory -Force -Path $FixtureRoot, (Split-Path -Parent $Target) | Out-Null

  $fixtureSource = @'
using System;
public static class InstallerFixture {
    public static int Main(string[] args) {
        Console.WriteLine("configure fixture");
        int status;
        return int.TryParse(Environment.GetEnvironmentVariable("CONFIGURE_STATUS"), out status) ? status : 0;
    }
}
'@
  $fixtureSourcePath = Join-Path $Temp 'installer-fixture.cs'
  [System.IO.File]::WriteAllText($fixtureSourcePath, $fixtureSource)
  $compiler = Join-Path $env:WINDIR 'Microsoft.NET/Framework64/v4.0.30319/csc.exe'
  Assert-True (Test-Path $compiler) 'Windows C# compiler is unavailable'
  $compilerProcess = Start-Process -FilePath $compiler -ArgumentList @('/nologo', '/target:exe', "/out:$AssetPath", $fixtureSourcePath) -Wait -PassThru -NoNewWindow
  Assert-True ($compilerProcess.ExitCode -eq 0) 'fixture executable did not compile'

  $listener = [System.Net.Sockets.TcpListener]::new([Net.IPAddress]::Loopback, 0)
  $listener.Start()
  $port = ([System.Net.IPEndPoint]$listener.LocalEndpoint).Port
  $listener.Stop()

  $installerSource = [System.IO.File]::ReadAllText($Installer)
  $replacement = '$Url   = "http://127.0.0.1:{0}/releases/download/v1.2.3/$Asset"' -f $port
  $fixtureInstaller = $installerSource -replace '(?m)^\$Url\s+=.*$', $replacement
  [System.IO.File]::WriteAllText((Join-Path $FixtureRoot 'install.ps1'), $fixtureInstaller)

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

  $env:LOCALAPPDATA = $LocalAppData
  $env:PROCESSOR_ARCHITECTURE = 'AMD64'
  $env:CONFIGURE_STATUS = '0'
  [Environment]::SetEnvironmentVariable('PATH', $saved.UserPath, 'User')

  $success = Invoke-InstallerPipeline 'success'
  Assert-True ($success.Status -eq 0) "installer pipeline failed: $($success.Output)"
  Assert-True ($success.Output -match 'apis-mcp installer') 'installer title was missing'
  Assert-True ($success.Output -match 'Downloading apis-mcp-windows-amd64\.exe\.\.\.') 'single download step was missing'
  Assert-True ($success.Output -match 'configure fixture') 'configure did not run after download'
  Assert-True ($success.Output -notmatch 'Finding the latest|Verifying|Checksum|Installed to|is ready|\[(?:ok|\.\.)\]') 'installer contains UX not present in the reference installer'
  Assert-True ((Get-FileHash $Target -Algorithm SHA256).Hash -eq (Get-FileHash $AssetPath -Algorithm SHA256).Hash) 'downloaded executable was not installed'
  $Passed++

  $env:CONFIGURE_STATUS = '7'
  $cancelled = Invoke-InstallerPipeline 'cancelled'
  Assert-True ($cancelled.Status -eq 7) "configure cancellation returned $($cancelled.Status) instead of 7"
  Assert-True ($cancelled.Output -match "Re-run 'apis-mcp configure' later to finish setup") 'configure cancellation guidance was missing'
  Assert-True ($cancelled.Output -notmatch 'is ready|Run now:|Run .* to browse|In a new window:') 'configure cancellation fell through to completion guidance'
  $Passed++

  $tokens = $null
  $parseErrors = $null
  [void][System.Management.Automation.Language.Parser]::ParseFile($Installer, [ref]$tokens, [ref]$parseErrors)
  Assert-True ($parseErrors.Count -eq 0) 'install.ps1 did not parse'
  Assert-True ($installerSource -match 'System\.Net\.HttpWebRequest') 'Windows PowerShell-compatible transport is missing'
  Assert-True ($installerSource -notmatch 'System\.Net\.Http\.|SHA256SUMS|Finding the latest|is ready') 'removed installer redesign remains in install.ps1'
  $Passed++
} finally {
  if ($null -ne $Server -and -not $Server.HasExited) { Stop-Process -Id $Server.Id -Force }
  [Environment]::SetEnvironmentVariable('PATH', $saved.UserPath, 'User')
  $env:LOCALAPPDATA = $saved.LocalAppData
  $env:PATH = $saved.Path
  $env:PROCESSOR_ARCHITECTURE = $saved.ProcessorArchitecture
  $env:CONFIGURE_STATUS = $saved.ConfigureStatus
  Remove-Item -LiteralPath $Temp -Recurse -Force -ErrorAction SilentlyContinue
}

Write-Host "PowerShell installer tests passed: $Passed"
exit 0
