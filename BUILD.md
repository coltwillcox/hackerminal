# Build Documentation

This document describes how to build Hackerminal from source for various platforms.

## Quick Start

```bash
# Build for current platform
make

# Build for all platforms
make build-all

# Run the application
make run
```

## Prerequisites

- Go 1.16 or higher
- Make (optional, but recommended)
- Bash (for build scripts)

## Build Methods

### 1. Using Make (Recommended)

The Makefile provides convenient commands for building:

```bash
# Show all available commands
make help

# Build for current platform
make build

# Build for all platforms (creates archives)
make build-all

# Build for specific platforms
make build-linux
make build-macos
make build-windows

# Clean build artifacts
make clean

# Build and run
make run

# Test audio system
make audio-test

# Install to system (Unix-like, requires sudo)
make install

# Uninstall from system
make uninstall
```

### 2. Using Build Script

The `build.sh` script creates optimized binaries for all platforms:

```bash
# Build version 1.0.0 for all platforms
./build.sh 1.0.0

# Build with custom version
./build.sh 2.0.0-beta
```

Output structure:
```
build/
├── hackerminal-1.0.0-linux-amd64.tar.gz
├── hackerminal-1.0.0-linux-arm64.tar.gz
├── hackerminal-1.0.0-darwin-amd64.tar.gz
├── hackerminal-1.0.0-darwin-arm64.tar.gz
├── hackerminal-1.0.0-windows-amd64.zip
└── hackerminal-1.0.0-windows-arm64.zip
```

### 3. Manual Build

#### Simple Build
```bash
go build -o hackerminal .
```

#### Optimized Build (smaller binary)
```bash
go build -ldflags="-s -w" -o hackerminal .
```

#### Cross-Platform Build
```bash
# Linux AMD64
GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o hackerminal-linux-amd64 .

# Linux ARM64
GOOS=linux GOARCH=arm64 go build -ldflags="-s -w" -o hackerminal-linux-arm64 .

# macOS Intel
GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o hackerminal-darwin-amd64 .

# macOS Apple Silicon
GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w" -o hackerminal-darwin-arm64 .

# Windows AMD64
GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o hackerminal-windows-amd64.exe .

# Windows ARM64
GOOS=windows GOARCH=arm64 go build -ldflags="-s -w" -o hackerminal-windows-arm64.exe .
```

## Supported Platforms

| Platform | Architecture | Tested | Archive Format |
|----------|-------------|---------|----------------|
| Linux    | amd64 (x86_64) | ✅ | tar.gz |
| Linux    | arm64 (aarch64) | ✅ | tar.gz |
| macOS    | amd64 (Intel) | ✅ | tar.gz |
| macOS    | arm64 (Apple Silicon) | ✅ | tar.gz |
| Windows  | amd64 (x86_64) | ✅ | zip |
| Windows  | arm64 | ✅ | zip |

## Build Flags Explained

### `-ldflags="-s -w"`
- `-s`: Omit symbol table and debug information
- `-w`: Omit DWARF symbol table
- **Result**: Reduces binary size by ~30-40%

Example sizes:
- Without flags: ~3.5 MB
- With flags: ~2.3 MB (includes embedded 72KB audio file)

### Embedded Assets

The `assets/beep.wav` file is embedded directly in the binary using Go's `embed` package:
- **Self-contained**: No external files required to run
- **Automatic fallback**: External files in `assets/` directory take priority if present
- **Customizable**: Users can place custom audio files alongside the binary
- **Zero dependencies**: Complete functionality in a single executable

### Common GOOS/GOARCH Values

```go
GOOS=linux   GOARCH=amd64   // Linux 64-bit
GOOS=linux   GOARCH=arm64   // Linux ARM 64-bit
GOOS=darwin  GOARCH=amd64   // macOS Intel
GOOS=darwin  GOARCH=arm64   // macOS Apple Silicon
GOOS=windows GOARCH=amd64   // Windows 64-bit
GOOS=windows GOARCH=arm64   // Windows ARM 64-bit
```

## Build Output

### Archive Contents
Each platform archive contains:
```
hackerminal-1.0.0-{os}-{arch}/
├── hackerminal (or hackerminal.exe on Windows)  [self-contained with embedded audio]
├── assets/                                       [optional for customization]
│   └── beep.wav
└── README.md
```

**Note**: The binary is fully functional without the `assets/` directory. The audio file is embedded in the executable. External assets are only needed if you want to customize the sound.

### Size Comparison
```
Platform                    Binary Size    Archive Size
──────────────────────────────────────────────────────
Linux AMD64                 2.3 MB         1.1 MB
Linux ARM64                 2.2 MB         1.0 MB
macOS Intel                 2.4 MB         1.1 MB
macOS Apple Silicon         2.3 MB         1.0 MB
Windows AMD64               2.4 MB         1.1 MB
Windows ARM64               2.2 MB         1.0 MB
```

All binaries include the embedded 72KB audio file.

## CI/CD Pipeline

### GitHub Actions

The repository includes a GitHub Actions workflow (`.github/workflows/release.yml`) that automatically:

1. **Triggers on**:
   - Git tags matching `v*` (e.g., `v1.0.0`)
   - Manual workflow dispatch

2. **Builds**:
   - All 6 platform binaries in parallel
   - Creates optimized archives

3. **Releases**:
   - Creates GitHub release
   - Uploads all archives
   - Generates release notes

### Creating a Release

```bash
# Tag the release
git tag -a v1.0.0 -m "Release version 1.0.0"

# Push the tag to trigger CI/CD
git push origin v1.0.0

# GitHub Actions will automatically:
# - Build all platform binaries
# - Create archives
# - Create GitHub release
# - Upload artifacts
```

## Development Builds

For development, use the standard Go workflow:

```bash
# Run without building
go run .

# Build and run
make run

# Build with race detector (for development)
go build -race -o hackerminal .

# Run tests (if any)
go test -v ./...
```

## Troubleshooting

### Build Fails on Cross-Compilation

If cross-compilation fails, ensure you have Go 1.16+ which includes built-in cross-compilation support.

```bash
# Verify Go version
go version

# Should output: go version go1.16 or higher
```

### Binary Doesn't Run on Target Platform

1. **Linux**: Ensure GLIBC compatibility
   - Built binaries are statically linked where possible
   - Should work on most modern Linux distributions

2. **macOS**: Code signing may be required for distribution
   - For personal use, right-click → Open on first run
   - For distribution, sign with Apple Developer certificate

3. **Windows**: May be blocked by SmartScreen
   - Right-click → Properties → Unblock
   - Or sign with a code signing certificate

### Audio Files Missing

The build scripts automatically include the `assets/` directory. If audio doesn't work:

```bash
# Verify assets are included
ls -la hackerminal-*/assets/

# Should show: beep.wav
```

## Clean Build

To ensure a completely clean build:

```bash
# Clean all build artifacts
make clean

# Clean Go cache
go clean -cache -modcache

# Rebuild
make build
```

## Build Statistics

- **Build Time**: ~10-15 seconds for all platforms (parallel)
- **Binary Size**: ~2.5 MB uncompressed per platform
- **Archive Size**: ~1 MB compressed per platform
- **Total Release Size**: ~6 MB (all platforms)
- **Dependencies**: Zero external dependencies (100% Go stdlib)

## Contributing

When adding new features that affect builds:

1. Update this BUILD.md
2. Test cross-compilation: `make build-all`
3. Verify all archives extract correctly
4. Update CHANGELOG.md
5. Update version in build scripts if needed

## License

See LICENSE file for details.
