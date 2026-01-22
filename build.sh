#!/bin/bash
# Build script for Hackerminal - cross-platform compilation

set -e

VERSION=${1:-"1.1.0"}
BUILD_DIR="build"
BINARY_NAME="hackerminal"

echo "🔨 Building Hackerminal v${VERSION}"
echo "================================"
echo ""

# Clean previous builds
if [ -d "$BUILD_DIR" ]; then
    echo "🧹 Cleaning previous builds..."
    rm -rf "$BUILD_DIR"
fi

mkdir -p "$BUILD_DIR"

# Function to build for a specific platform
build() {
    local GOOS=$1
    local GOARCH=$2
    local OUTPUT_NAME=$3
    local PLATFORM_NAME=$4

    echo "📦 Building for ${PLATFORM_NAME}..."

    GOOS=$GOOS GOARCH=$GOARCH go build -ldflags="-s -w" -o "${BUILD_DIR}/${OUTPUT_NAME}" .

    if [ $? -eq 0 ]; then
        # Create platform-specific directory
        PLATFORM_DIR="${BUILD_DIR}/${BINARY_NAME}-${VERSION}-${GOOS}-${GOARCH}"
        mkdir -p "$PLATFORM_DIR"

        # Move binary
        mv "${BUILD_DIR}/${OUTPUT_NAME}" "$PLATFORM_DIR/"

        # Copy assets directory
        cp -r assets "$PLATFORM_DIR/"

        # Copy README and LICENSE if they exist
        [ -f README.md ] && cp README.md "$PLATFORM_DIR/"
        [ -f LICENSE ] && cp LICENSE "$PLATFORM_DIR/"

        # Create archive
        cd "$BUILD_DIR"
        if [ "$GOOS" = "windows" ]; then
            zip -r "${BINARY_NAME}-${VERSION}-${GOOS}-${GOARCH}.zip" "$(basename $PLATFORM_DIR)" > /dev/null
            echo "   ✓ Created ${BINARY_NAME}-${VERSION}-${GOOS}-${GOARCH}.zip"
        else
            tar -czf "${BINARY_NAME}-${VERSION}-${GOOS}-${GOARCH}.tar.gz" "$(basename $PLATFORM_DIR)"
            echo "   ✓ Created ${BINARY_NAME}-${VERSION}-${GOOS}-${GOARCH}.tar.gz"
        fi
        cd ..

        # Get file size
        if [ "$GOOS" = "windows" ]; then
            SIZE=$(ls -lh "${BUILD_DIR}/${BINARY_NAME}-${VERSION}-${GOOS}-${GOARCH}.zip" | awk '{print $5}')
        else
            SIZE=$(ls -lh "${BUILD_DIR}/${BINARY_NAME}-${VERSION}-${GOOS}-${GOARCH}.tar.gz" | awk '{print $5}')
        fi
        echo "   📊 Archive size: ${SIZE}"
        echo ""
    else
        echo "   ❌ Build failed for ${PLATFORM_NAME}"
        echo ""
        return 1
    fi
}

# Build for different platforms
echo "🐧 Linux Builds"
echo "---------------"
build "linux" "amd64" "${BINARY_NAME}" "Linux (amd64)"
build "linux" "arm64" "${BINARY_NAME}" "Linux (arm64)"

echo "🍎 macOS Builds"
echo "---------------"
build "darwin" "amd64" "${BINARY_NAME}" "macOS (Intel)"
build "darwin" "arm64" "${BINARY_NAME}" "macOS (Apple Silicon)"

echo "🪟 Windows Builds"
echo "-----------------"
build "windows" "amd64" "${BINARY_NAME}.exe" "Windows (amd64)"
build "windows" "arm64" "${BINARY_NAME}.exe" "Windows (arm64)"

echo "================================"
echo "✅ Build complete!"
echo ""
echo "📂 Binaries available in: ${BUILD_DIR}/"
echo ""
ls -lh "$BUILD_DIR"/*.{tar.gz,zip} 2>/dev/null | awk '{print "   " $9 " (" $5 ")"}'
echo ""
echo "🎉 All builds successful!"
