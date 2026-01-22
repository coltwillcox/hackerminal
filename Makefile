.PHONY: all build clean test run install help

VERSION ?= 1.1.1
BINARY_NAME = hackerminal
BUILD_DIR = build

# Default target
all: build

# Build for current platform
build:
	@echo "🔨 Building $(BINARY_NAME)..."
	@go build -ldflags="-s -w" -o $(BINARY_NAME) .
	@echo "✅ Build complete: ./$(BINARY_NAME)"

# Build for all platforms
build-all:
	@./build.sh $(VERSION)

# Build for specific platform
build-linux:
	@echo "🐧 Building for Linux..."
	@GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o $(BUILD_DIR)/$(BINARY_NAME)-linux-amd64 .
	@echo "✅ Linux build complete"

build-macos:
	@echo "🍎 Building for macOS..."
	@GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-amd64 .
	@GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w" -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-arm64 .
	@echo "✅ macOS builds complete"

build-windows:
	@echo "🪟 Building for Windows..."
	@GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o $(BUILD_DIR)/$(BINARY_NAME)-windows-amd64.exe .
	@echo "✅ Windows build complete"

# Run the application
run: build
	@./$(BINARY_NAME)

# Clean build artifacts
clean:
	@echo "🧹 Cleaning build artifacts..."
	@rm -f $(BINARY_NAME)
	@rm -rf $(BUILD_DIR)
	@echo "✅ Clean complete"

# Install to system (Unix-like only)
install: build
	@echo "📦 Installing $(BINARY_NAME)..."
	@sudo cp $(BINARY_NAME) /usr/local/bin/
	@sudo mkdir -p /usr/local/share/$(BINARY_NAME)
	@echo "✅ Installed to /usr/local/bin/$(BINARY_NAME)"

# Uninstall from system
uninstall:
	@echo "🗑️  Uninstalling $(BINARY_NAME)..."
	@sudo rm -f /usr/local/bin/$(BINARY_NAME)
	@sudo rm -rf /usr/local/share/$(BINARY_NAME)
	@echo "✅ Uninstalled"

# Run tests (if any exist)
test:
	@echo "🧪 Running tests..."
	@go test -v ./...

# Show help
help:
	@echo "Hackerminal Build System"
	@echo ""
	@echo "Usage:"
	@echo "  make              - Build for current platform"
	@echo "  make build-all    - Build for all platforms (Linux, macOS, Windows)"
	@echo "  make build-linux  - Build for Linux only"
	@echo "  make build-macos  - Build for macOS only"
	@echo "  make build-windows- Build for Windows only"
	@echo "  make run          - Build and run the application"
	@echo "  make clean        - Remove build artifacts"
	@echo "  make install      - Install to /usr/local/bin (requires sudo)"
	@echo "  make uninstall    - Remove from /usr/local/bin (requires sudo)"
	@echo "  make test         - Run tests"
	@echo "  make help         - Show this help message"
	@echo ""
	@echo "Variables:"
	@echo "  VERSION=$(VERSION)  - Set version for build-all (e.g., make build-all VERSION=1.2.0)"
