# Changelog

All notable changes to Hackerminal will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.0.0] - 2026-01-21

### Added
- 🎬 **60+ Hollywood Hacking Sequences** - Parodies of movie/TV hacking scenes
  - The Matrix, Hackers, WarGames references
  - Alien, Predator, Terminator franchises
  - CSI: Cyber, Mr. Robot, and more
  - Classic internet memes (hunter2, Bobby Tables)

- 🖥️ **CRT Phosphor Terminal Style** - Authentic retro aesthetic
  - Green/amber monochrome color scheme
  - Fancy prompt with user, host, directory, git branch, timestamp segments
  - Terminal width auto-detection

- ⌨️ **Realistic Typing Simulation**
  - Variable speed based on key position (home row vs special chars)
  - Typo simulation with QWERTY adjacency (5% chance)
  - Automatic backspace correction
  - Random thinking pauses (200-2000ms)

- 🎨 **Visual Effects System**
  - Progress bars with ASCII characters
  - Multiple spinner animations (4 different styles)
  - Matrix-style rain effect
  - Screen glitch effects (10% chance)
  - CRT scan lines (1% chance)
  - ASCII art splash screens (10% chance)
  - Network topology diagrams (10% chance)
  - Split screen displays (10% chance)
  - File tree visualizations (10% chance)

- 🔊 **Audio Effects** - Dramatic sound system
  - Custom beep sounds at critical moments (15% chance)
  - Smart audio player detection (paplay, aplay, ffplay, mpv, mplayer, cvlc, afplay)
  - Cross-platform support (Linux, macOS)
  - Automatic fallback to terminal bell
  - Single beep for warnings/alerts
  - Triple beep for dramatic moments
  - Custom `assets/beep.wav` audio file
  - **Embedded audio**: Audio file embedded in binary using Go's `embed` package (fully self-contained)
  - External audio files prioritized for customization

- 🎭 **Hollywood Parody Elements**
  - Fake IP scanning
  - Password cracking with ridiculous results
  - Image enhancement beyond pixel limits
  - Two-person keyboard jokes (implicit in sequences)
  - "I'm in!" moments
  - Self-destruct countdowns
  - AI self-awareness warnings
  - Nuclear code access attempts

- 🛠️ **Build System**
  - Cross-platform compilation (Linux, macOS, Windows)
  - Multi-architecture support (amd64, arm64)
  - Automated build script (`build.sh`)
  - Makefile with common commands
  - GitHub Actions CI/CD workflow
  - Pre-built binary releases

- 📦 **Distribution**
  - Optimized binaries with stripped symbols (~2.3MB per platform)
  - Archive packages (tar.gz for Unix, zip for Windows)
  - **Self-contained executables**: Audio embedded, no external dependencies
  - Assets folder included in releases for reference
  - README and documentation bundled

- 🧪 **Testing Utilities**
  - Audio test script (`test_audio.sh`)
  - Makefile test commands

- ⚙️ **Configuration**
  - Centralized constants in `config.go`
  - Easy adjustment of probabilities
  - Typing speed configuration
  - Effect frequency tuning

### Technical Details
- **Language**: Go 1.16+
- **Dependencies**: Zero external dependencies (100% stdlib)
- **Build Size**: ~1MB compressed per platform
- **Platforms**: Linux (amd64/arm64), macOS (Intel/Apple Silicon), Windows (amd64/arm64)
- **Audio Players Supported**: 7 different players with automatic detection
- **Code Quality**: Clean separation of concerns, well-documented, Go idioms

### File Structure
```
├── main.go           - Entry point
├── config.go         - Configuration constants
├── terminal.go       - Terminal management
├── sequences.go      - 60+ parody sequences
├── effects.go        - Visual effects
├── assets.go         - ASCII art
├── utils.go          - Utility functions including audio
├── assets/
│   └── beep.wav     - Custom beep sound
├── build.sh          - Cross-platform build script
├── Makefile          - Build automation
├── test_audio.sh     - Audio testing utility
└── .github/
    └── workflows/
        └── release.yml - GitHub Actions CI/CD
```

---

## [Unreleased]

### Planned Features
- Interactive mode with keyboard input
- More movie/TV show references
- Additional visual effects (3D ASCII, hex dumps, etc.)
- Configuration file support
- Command-line flags for customization
- More audio effects
- Achievements/statistics system

---

[1.0.0]: https://github.com/yourusername/hackerminal/releases/tag/v1.0.0
