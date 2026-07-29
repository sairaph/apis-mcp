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
$EventLog = Join-Path $Temp 'events.log'
$DownloadModeFile = Join-Path $Temp 'download-mode'
$PowerShellExecutable = (Get-Process -Id $PID).Path
$Server = $null
$Passed = 0

function Assert-True([bool]$Condition, [string]$Message) {
  if (-not $Condition) { throw "PowerShell installer test failed: $Message" }
}

function Assert-Events([string[]]$Expected, [string]$Message) {
  $actual = @([System.IO.File]::ReadAllLines($EventLog))
  Assert-True ($actual.Count -eq $Expected.Count) "$Message (expected '$($Expected -join ', ')', got '$($actual -join ', ')')"
  for ($i = 0; $i -lt $Expected.Count; $i++) {
    Assert-True ($actual[$i] -eq $Expected[$i]) "$Message (expected '$($Expected -join ', ')', got '$($actual -join ', ')')"
  }
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
  DaemonStopStatus = $env:DAEMON_STOP_STATUS
  DaemonStopDelay = $env:DAEMON_STOP_DELAY_MS
  InstallerEventLog = $env:INSTALLER_EVENT_LOG
  InstallerDownloadModeFile = $env:INSTALLER_DOWNLOAD_MODE_FILE
  UserPath = [Environment]::GetEnvironmentVariable('PATH', 'User')
}

try {
  New-Item -ItemType Directory -Force -Path $FixtureRoot, (Split-Path -Parent $Target) | Out-Null

  $fixtureSource = @'
using System;
using System.IO;
using System.Threading;
public static class InstallerFixture {
    public static int Main(string[] args) {
        string log = Environment.GetEnvironmentVariable("INSTALLER_EVENT_LOG");
        int status;
        if (args.Length == 2 && args[0] == "daemon" && args[1] == "--stop") {
            File.AppendAllText(log, "daemon-stop" + Environment.NewLine);
            int delay;
            if (int.TryParse(Environment.GetEnvironmentVariable("DAEMON_STOP_DELAY_MS"), out delay)) {
                Thread.Sleep(delay);
            }
            return int.TryParse(Environment.GetEnvironmentVariable("DAEMON_STOP_STATUS"), out status) ? status : 0;
        }
        if (args.Length == 1 && args[0] == "configure") {
            File.AppendAllText(log, "configure" + Environment.NewLine);
            Console.WriteLine("configure fixture");
            return int.TryParse(Environment.GetEnvironmentVariable("CONFIGURE_STATUS"), out status) ? status : 0;
        }
        if (args.Length == 1 && args[0] == "hold") {
            string executable = System.Reflection.Assembly.GetExecutingAssembly().Location;
            using (FileStream held = new FileStream(executable, FileMode.Open, FileAccess.Read, FileShare.Read)) {
                File.AppendAllText(log, "hold-ready" + Environment.NewLine);
                Thread.Sleep(30000);
            }
            return 0;
        }
        return 64;
    }
}
'@
  $fixtureSourcePath = Join-Path $Temp 'installer-fixture.cs'
  [System.IO.File]::WriteAllText($fixtureSourcePath, $fixtureSource)
  $compiler = Join-Path $env:WINDIR 'Microsoft.NET/Framework64/v4.0.30319/csc.exe'
  Assert-True (Test-Path $compiler) 'Windows C# compiler is unavailable'
  $compilerProcess = Start-Process -FilePath $compiler -ArgumentList @('/nologo', '/target:exe', "/out:$AssetPath", $fixtureSourcePath) -Wait -PassThru -NoNewWindow
  Assert-True ($compilerProcess.ExitCode -eq 0) 'fixture executable did not compile'
  [System.IO.File]::WriteAllText($EventLog, '')
  [System.IO.File]::WriteAllText($DownloadModeFile, 'valid')
  $env:INSTALLER_EVENT_LOG = $EventLog
  $env:INSTALLER_DOWNLOAD_MODE_FILE = $DownloadModeFile

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
  $env:DAEMON_STOP_STATUS = '0'
  $env:DAEMON_STOP_DELAY_MS = '0'
  [Environment]::SetEnvironmentVariable('PATH', $saved.UserPath, 'User')

  [System.IO.File]::WriteAllText($EventLog, '')
  $success = Invoke-InstallerPipeline 'success'
  Assert-True ($success.Status -eq 0) "installer pipeline failed: $($success.Output)"
  Assert-True ($success.Output -match 'apis-mcp installer') 'installer title was missing'
  Assert-True ($success.Output -match 'Downloading apis-mcp-windows-amd64\.exe\.\.\.') 'single download step was missing'
  Assert-True ($success.Output -match 'configure fixture') 'configure did not run after download'
  Assert-True ($success.Output -notmatch 'Finding the latest|Verifying|Checksum|Installed to|is ready|\[(?:ok|\.\.)\]') 'installer contains UX not present in the reference installer'
  Assert-True ((Get-FileHash $Target -Algorithm SHA256).Hash -eq (Get-FileHash $AssetPath -Algorithm SHA256).Hash) 'downloaded executable was not installed'
  Assert-Events -Expected @('download', 'configure') -Message 'fresh install invoked the legacy daemon hook or ran steps out of order'
  $Passed++

  [System.IO.File]::WriteAllText($EventLog, '')
  $updated = Invoke-InstallerPipeline 'update'
  Assert-True ($updated.Status -eq 0) "installer update failed: $($updated.Output)"
  Assert-Events -Expected @('download', 'daemon-stop', 'configure') -Message 'staged download did not precede daemon stop and replacement'
  $Passed++

  $env:DAEMON_STOP_STATUS = '64'
  [System.IO.File]::WriteAllText($EventLog, '')
  $unsupported = Invoke-InstallerPipeline 'unsupported-daemon-command'
  Assert-True ($unsupported.Status -eq 0) "unsupported daemon command prevented update: $($unsupported.Output)"
  Assert-Events -Expected @('download', 'daemon-stop', 'configure') -Message 'unsupported daemon command changed update ordering'
  Assert-True ((Get-FileHash $Target -Algorithm SHA256).Hash -eq (Get-FileHash $AssetPath -Algorithm SHA256).Hash) 'unsupported daemon command prevented executable replacement'
  $Passed++

  $env:DAEMON_STOP_STATUS = '0'
  $env:DAEMON_STOP_DELAY_MS = '10000'
  [System.IO.File]::WriteAllText($EventLog, '')
  $timer = [System.Diagnostics.Stopwatch]::StartNew()
  $bounded = Invoke-InstallerPipeline 'bounded-daemon-stop'
  $timer.Stop()
  Assert-True ($bounded.Status -eq 0) "bounded daemon stop update failed: $($bounded.Output)"
  Assert-True ($timer.Elapsed.TotalSeconds -lt 8) "daemon stop was not bounded: $($timer.Elapsed.TotalSeconds) seconds"
  Assert-Events -Expected @('download', 'daemon-stop', 'configure') -Message 'bounded daemon stop changed update ordering'
  $Passed++

  $env:DAEMON_STOP_DELAY_MS = '0'
  $env:CONFIGURE_STATUS = '7'
  [System.IO.File]::WriteAllText($EventLog, '')
  $cancelled = Invoke-InstallerPipeline 'cancelled'
  Assert-True ($cancelled.Status -eq 7) "configure cancellation returned $($cancelled.Status) instead of 7"
  Assert-True ($cancelled.Output -match "Re-run 'apis-mcp configure' later to finish setup") 'configure cancellation guidance was missing'
  Assert-True ($cancelled.Output -notmatch 'is ready|Run now:|Run .* to browse|In a new window:') 'configure cancellation fell through to completion guidance'
  Assert-Events -Expected @('download', 'daemon-stop', 'configure') -Message 'configure cancellation changed installer ordering'
  $Passed++

  $env:CONFIGURE_STATUS = '0'
  Copy-Item -LiteralPath $AssetPath -Destination $Target -Force
  [System.IO.File]::AppendAllText($Target, 'old-locked')
  $lockedHash = (Get-FileHash $Target -Algorithm SHA256).Hash
  [System.IO.File]::WriteAllText($EventLog, '')
  $holder = Start-Process -FilePath $Target -ArgumentList 'hold' -PassThru -WindowStyle Hidden
  try {
    $holderReady = $false
    for ($attempt = 0; $attempt -lt 50; $attempt++) {
      if ([System.IO.File]::ReadAllText($EventLog) -match 'hold-ready') {
        $holderReady = $true
        break
      }
      Start-Sleep -Milliseconds 100
    }
    Assert-True $holderReady 'locked-target fixture did not start'
    [System.IO.File]::WriteAllText($EventLog, '')
    $locked = Invoke-InstallerPipeline 'locked-target'
  } finally {
    if (-not $holder.HasExited) { Stop-Process -Id $holder.Id -Force }
    $holder.WaitForExit()
  }
  Assert-True ($locked.Status -eq 1) "replacement failure returned $($locked.Status) instead of 1"
  Assert-True ((Get-FileHash $Target -Algorithm SHA256).Hash -eq $lockedHash) 'replacement failure did not preserve the existing executable'
  Assert-True (-not (Test-Path "$Target.new")) 'replacement failure left a staging executable'
  Assert-True ($locked.Output -match 'Could not replace apis-mcp\.exe') 'replacement failure was not reported'
  Assert-True ($locked.Output -match 'existing executable was kept') 'replacement preservation was not reported'
  Assert-True ($locked.Output -match 'Close any running apis-mcp process and re-run the installer') 'replacement recovery guidance was missing'
  Assert-Events -Expected @('download', 'daemon-stop') -Message 'failed replacement configured or ran steps out of order'
  $Passed++

  foreach ($mode in @('fail', 'zero', 'truncated')) {
    Copy-Item -LiteralPath $AssetPath -Destination $Target -Force
    [System.IO.File]::AppendAllText($Target, "old-$mode")
    $oldHash = (Get-FileHash $Target -Algorithm SHA256).Hash
    [System.IO.File]::WriteAllText($DownloadModeFile, $mode)
    [System.IO.File]::WriteAllText($EventLog, '')
    $invalid = Invoke-InstallerPipeline "download-$mode"
    Assert-True ($invalid.Status -eq 1) "$mode download returned $($invalid.Status) instead of 1"
    Assert-True ((Get-FileHash $Target -Algorithm SHA256).Hash -eq $oldHash) "$mode download modified the existing executable"
    Assert-True (-not (Test-Path "$Target.new")) "$mode download left a staging executable"
    Assert-Events -Expected @('download') -Message "$mode download invoked daemon stop or configure"
    if ($mode -eq 'zero') {
      Assert-True ($invalid.Output -match 'Download did not complete; nothing was installed') 'zero-byte download guidance was missing'
    } else {
      Assert-True ($invalid.Output -match 'Download failed') "$mode download failure guidance was missing"
    }
    $Passed++
  }
  [System.IO.File]::WriteAllText($DownloadModeFile, 'valid')

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
  $env:DAEMON_STOP_STATUS = $saved.DaemonStopStatus
  $env:DAEMON_STOP_DELAY_MS = $saved.DaemonStopDelay
  $env:INSTALLER_EVENT_LOG = $saved.InstallerEventLog
  $env:INSTALLER_DOWNLOAD_MODE_FILE = $saved.InstallerDownloadModeFile
  Remove-Item -LiteralPath $Temp -Recurse -Force -ErrorAction SilentlyContinue
}

Write-Host "PowerShell installer tests passed: $Passed"
exit 0
