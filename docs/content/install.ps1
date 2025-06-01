$ErrorActionPreference = "Stop"

$RepoOwner = "agailloty"
$RepoName = "preprocess"
$Version = "0.1.0"
$BinaryName = "preprocess.exe"
$InstallDir = "$env:USERPROFILE\preprocess"

# Detect architecture
$arch = if ([System.Environment]::Is64BitOperatingSystem) {
    if ([System.Runtime.InteropServices.RuntimeInformation]::OSArchitecture -eq "Arm64") {
        "arm64"
    } else {
        "x86_64"
    }
} else {
    "i386"
}

$OS = "Windows"
$FileName = "${RepoName}_${OS}_${arch}.zip"
$Url = "https://github.com/$RepoOwner/$RepoName/releases/download/$Version/$FileName"

# Create a proper temporary .zip file path
$TempDir = [System.IO.Path]::GetTempPath()
$TempZipPath = [System.IO.Path]::Combine($TempDir, "$(New-Guid).zip")

# Download the archive
Write-Host "➡️ Downloading: $Url"
Invoke-WebRequest -Uri $Url -OutFile $TempZipPath

# Create installation directory
Write-Host "📁 Creating install directory: $InstallDir"
New-Item -ItemType Directory -Force -Path $InstallDir | Out-Null

# Extract directly into the installation directory
Write-Host "📦 Extracting to install directory..."
Expand-Archive -LiteralPath $TempZipPath -DestinationPath $InstallDir -Force

# Clean up
Remove-Item $TempZipPath -Force

# Add to PATH if needed
$envPath = [System.Environment]::GetEnvironmentVariable("Path", "User")
if ($envPath -notlike "*$InstallDir*") {
    Write-Host "➕ Adding install directory to PATH"
    $newPath = "$envPath;$InstallDir"
    [System.Environment]::SetEnvironmentVariable("Path", $newPath, "User")
    Write-Host "✅ PATH updated. Please restart your terminal to apply changes."
}

Write-Host "`n✅ Installation complete. You can now run: preprocess"
