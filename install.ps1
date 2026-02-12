# PowerShell install script for Windows
param(
    [switch]$All
)

$ErrorActionPreference = 'Stop'
$here = Split-Path -Parent $MyInvocation.MyCommand.Definition
$bin = Join-Path $here 'bin'
if (-not (Test-Path $bin)) { New-Item -ItemType Directory -Path $bin | Out-Null }

Write-Output "Checking Go..."
if (-not (Get-Command go -ErrorAction SilentlyContinue)) {
    Write-Error "Go not found in PATH. Install Go 1.21+ and rerun."
    exit 1
}

Write-Output "Building ffi-gen..."
go build -o (Join-Path $bin 'ffi-gen.exe') ./tools/ffi-gen

Write-Output "Building ffi-helper..."
try {
    go build -o (Join-Path $bin 'ffi-helper.exe') ./tools/ffi-helper
} catch {
    Write-Warning "ffi-helper build failed. Ensure you have a C toolchain (gcc/clang) available for cgo builds."
}

Write-Output "Building fig..."
go build -o (Join-Path $bin 'fig.exe') .

Write-Output "Install complete. Binaries are in: $bin"
if ($All) {
    Write-Output "--all specified: cross-builds are not implemented in this script. Consider setting GOOS/GOARCH manually and ensuring cross C toolchains for ffi-helper."
}
