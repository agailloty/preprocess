# Building with GoReleaser

This document explains how to build and release the `preprocess` application using GoReleaser.

## Prerequisites

### 1. Install GoReleaser

**macOS:**
```bash
brew install goreleaser
```

**Linux (using snap):**
```bash
snap install --classic goreleaser
```

**Other methods:**
Download from [goreleaser.com/install](https://goreleaser.com/install/)

### 2. GitHub Token

The configuration expects a GitHub token for creating releases. Place your token in a file:
```bash
# Create token file one directory up from the project
echo "your_github_token" > ../github_token
```

Or set it as an environment variable:
```bash
export GITHUB_TOKEN="your_github_token"
```

## Build Configuration

The `.goreleaser.yaml` configuration includes:

### Supported Platforms
- **Linux** (amd64, 386, arm, arm64)
- **Windows** (amd64, 386, arm, arm64)
- **macOS/Darwin** (amd64, 386, arm, arm64)

### Build Settings
- **CGO**: Disabled (`CGO_ENABLED=0`) for static binaries
- **Go Modules**: Automatically tidied before build via `go mod tidy` hook

### Archive Formats
- `.tar.gz` for Linux and macOS
- `.zip` for Windows
- Naming convention: `preprocess_<OS>_<Arch>.<format>`
  - Example: `preprocess_Linux_x86_64.tar.gz`

### Changelog
- Sorted in ascending order
- Excludes commits starting with:
  - `docs:` (documentation changes)
  - `test:` (test-related changes)

## Build Commands

### 1. Test Build Locally (No Release)

Build binaries without creating a release or publishing:

```bash
goreleaser build --snapshot --clean
```

This command:
- Creates binaries for all platforms
- Does not require a git tag
- Places output in `./dist/` directory
- `--clean` removes previous builds before starting
- `--snapshot` skips validation and versioning

### 2. Full Release Test (Local)

Test the complete release process without publishing:

```bash
goreleaser release --snapshot --clean --skip=publish
```

This command:
- Runs the complete release pipeline
- Creates archives and checksums
- Generates changelog
- Skips GitHub release creation
- Useful for testing before actual release

### 3. Create an Actual Release

To create and publish a release to GitHub:

```bash
# 1. Create a git tag
git tag -a v1.0.0 -m "Release version 1.0.0"

# 2. Push the tag to GitHub
git push origin v1.0.0

# 3. Run goreleaser
goreleaser release --clean
```

This command:
- Validates that you're on a tagged commit
- Builds binaries for all platforms
- Creates archives
- Generates changelog from git commits
- Creates a GitHub release
- Uploads all artifacts to the release

## Build Output

After running GoReleaser, build artifacts are located in the `./dist/` directory:

```
dist/
├── preprocess_Linux_x86_64.tar.gz
├── preprocess_Darwin_x86_64.tar.gz
├── preprocess_Windows_x86_64.zip
├── checksums.txt
└── ...
```

## GitHub Release

Releases are published to:
- Repository: `agailloty/preprocess`
- Each release includes:
  - Binary archives for all platforms
  - Checksums file
  - Auto-generated changelog
  - Footer: "Released by [GoReleaser](https://github.com/goreleaser/goreleaser)"

## Troubleshooting

### Missing GitHub Token
If you see authentication errors:
```bash
export GITHUB_TOKEN="your_github_token"
```

### Not on a Git Tag
For actual releases, you must be on a tagged commit:
```bash
git tag -a v1.0.0 -m "Release v1.0.0"
```

### Build Fails
Ensure Go modules are valid:
```bash
go mod tidy
go build
```

## Additional Resources

- [GoReleaser Documentation](https://goreleaser.com)
- [Configuration Reference](https://goreleaser.com/customization/)
- [GitHub Actions Integration](https://goreleaser.com/ci/actions/)
