# Changelog

All notable changes to Hackerminal will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.3/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.2.2] - 2026-01-23

### Added
- New ASCII art splash screens: VHS cassette, Amiga computer, keygen, Doom II

---

## [1.2.1] - 2026-01-23

### Fixed
- Replaced Nerd Font icons with standard Unicode emojis for better terminal compatibility
- Stats now saved after every sequence (prevents data loss when terminal is closed abruptly)
- Version field now updates correctly when loading existing stats

---

## [1.2.0] - 2026-01-22

### Added
- **Achievements/Statistics System** - Gamified tracking of hacking sessions
  - Persistent stats stored in `~/.config/hackerminal/stats.json`
  - Session tracking: commands typed, sequences triggered, uptime, and more
  - All-time statistics: total sessions, total commands, longest session, current streak
  - 15 unlockable achievements across Rookie, Intermediate, Advanced, and Rare tiers
  - Achievement notifications with visual celebration
  - Stats displayed on exit (Ctrl+C)
  - Configurable stats directory for future config file support

### Changed
- Refactored function naming to follow Go conventions (exported vs unexported)
- Fixed duplicate Sequence struct definition
- Added comprehensive documentation comments to all exported functions and types
- Improved code organization and separation of concerns
- Terminal stats now tracked per command typed
- Effects tracking for glitches, CRT scans, Matrix rain, ASCII splashes, and more

---

## [1.1.1] - 2026-01-22

### Changed
- Removed assets directory from build archives (binary is now fully self-contained)
- Updated build scripts to exclude assets from distribution packages

---

## [1.1.0] - 2026-01-22

### Changed
- Updated binary size from ~2.3 MB to ~2.2 MB
- Improved build documentation with accurate size information

---

## [1.0.3] - 2026-01-21

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
  - README and documentation bundled

- 🧪 **Testing Utilities**
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
- **Code Quality**: Clean separation of concerns, well-documented, Go idioms

### File Structure
```
├── main.go           - Entry point
├── config.go         - Configuration constants
├── terminal.go       - Terminal management
├── sequences.go      - 60+ parody sequences
├── effects.go        - Visual effects
├── assets.go         - ASCII art
├── utils.go          - Utility functions
├── build.sh          - Cross-platform build script
├── Makefile          - Build automation
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
- Achievements/statistics system

---

[1.2.2]: https://github.com/coltwillcox/hackerminal/releases/tag/v1.2.2
[1.2.1]: https://github.com/coltwillcox/hackerminal/releases/tag/v1.2.1
[1.2.0]: https://github.com/coltwillcox/hackerminal/releases/tag/v1.2.0
[1.1.1]: https://github.com/coltwillcox/hackerminal/releases/tag/v1.1.1
[1.1.0]: https://github.com/coltwillcox/hackerminal/releases/tag/v1.1.0
[1.0.3]: https://github.com/coltwillcox/hackerminal/releases/tag/v1.0.3
