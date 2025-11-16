<div align="center">

# 🔓 H A C K E R M I N A L 💀

<img src="screenshot.png" alt="Hackerminal Screenshot" width="800"/>

### ⚡ A parody terminal application that simulates Hollywood-style "hacking" scenes ⚡

```ascii
╔═══════════════════════════════════════════════════════════════╗
║  "I'm in!" - Every movie hacker ever                          ║
╚═══════════════════════════════════════════════════════════════╝
```

[![Go](https://img.shields.io/badge/Go-1.16+-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://golang.org/)
[![Terminal](https://img.shields.io/badge/Terminal-Required-4EAA25?style=for-the-badge&logo=gnubash&logoColor=white)](https://www.gnu.org/software/bash/)
[![License](https://img.shields.io/badge/License-Parody-red?style=for-the-badge)](LICENSE)
[![Status](https://img.shields.io/badge/Status-Hacking_The_Planet-brightgreen?style=for-the-badge)](README.md)

</div>

---

## 📖 Description

**Hackerminal** is an entertainment program that recreates the unrealistic and dramatic hacking scenes commonly seen in movies and TV shows. It displays randomized command sequences with fancy terminal effects, complete with progress bars, spinners, and references to popular sci-fi and hacker movies.

> ⚠️ **Warning:** This is a parody. Real hacking is illegal and boring.

## ✨ Features

- 🖥️ **Fancy Terminal Prompt**: CRT phosphor green/amber monochrome terminal style with user, host, directory, git branch, and timestamp segments
- ⌨️ **Typewriter Effect**: Commands and outputs appear character-by-character with realistic typing delays
- 🤔 **Random Thinking Pauses**: Random delays (200ms-2000ms) after prompts to simulate human thinking
- 🎬 **60+ Parody Sequences**: Including references to:
  - 🕶️ Classic hacker movies (The Matrix, Hackers, WarGames)
  - 👽 Sci-fi franchises (Alien, Predator, Terminator)
  - 📺 Popular TV shows (CSI: Cyber, Mr. Robot)
  - 🌐 Internet culture and memes

## 🎨 Visual Effects

- 📊 Progress bars with ASCII characters
- 🔄 Multiple spinner animations
- 🟢 Matrix-style rain effect
- 🌐 Fake IP scanning
- 🎨 ANSI color codes for dramatic output
- 📺 Random screen glitch effects (static, corruption artifacts, horizontal tears)

## 🚀 Installation

### 📋 Prerequisites

**Required:**
- 🔧 Go 1.16 or higher
- 💻 Unix-like terminal (Linux, macOS)

**Optional:**
- 🔤 [Nerd Fonts](https://www.nerdfonts.com/) - For proper display of icons in the terminal prompt (recommended: JetBrainsMono Nerd Font, FiraCode Nerd Font, or Hack Nerd Font)

### 🔨 Build

```bash
go build -o hackerminal .
```

Or to build and run in one step:
```bash
go run .
```

## 🎮 Usage

### Basic Usage

Simply run the program:

```bash
./hackerminal
```

The program will continuously display random hacking sequences. Press `Ctrl+C` to exit.

> 💡 **Tip:** Run this in front of non-technical people for maximum effect!

### Advanced Usage Ideas

#### 🖥️ As a Screensaver

Use Hackerminal as a screensaver to make your computer look like it's doing important hacking work:

**Linux (with xscreensaver):**
```bash
# Add to ~/.xscreensaver
programs: /path/to/hackerminal -root \n\
```

**macOS (with custom script):**
```bash
# Create a wrapper script that runs hackerminal in fullscreen
#!/bin/bash
clear && /path/to/hackerminal
```

#### 🔒 As a Lock Screen Effect

Impress (or confuse) anyone passing by your desk:

**Linux (using i3lock or similar):**
```bash
# Run before locking
hackerminal & sleep 2 && i3lock
```

**tmux/screen session:**
```bash
# Leave it running in a detached session
tmux new-session -d -s hacker './hackerminal'
# Reattach when you want to show off
tmux attach -t hacker
```

#### 🎬 Display Mode

For presentations or background displays:
```bash
# Run in fullscreen terminal (F11 in most terminals)
./hackerminal
```

> 🎭 **Pro Tip:** Combine with a green or amber terminal color scheme for authentic CRT vibes!

## 🎭 Example Sequences

- 🔐 SSH into fictional systems like "pentagon.gov" or "cyberdyne.sys"
- 🔓 "Crack" passwords and find weak credentials
- 💉 SQL injection demonstrations
- 🔍 "Enhance" images to impossible levels
- 🦠 Deploy viruses that rickroll targets
- 🎪 Many more Hollywood hacking tropes

## ⚙️ How It Works

1. 👤 Generates a random "hacker" username and target system
2. 🖥️ Displays a fancy terminal prompt
3. ⏸️ Pauses randomly to simulate thinking
4. ⌨️ Types out a command with typewriter effect
5. 🎬 Shows dramatic output with animations
6. 🔁 Repeats with a new random sequence

## ⚡ Configuration

The program uses:
- ⌨️ **Command typing speed**: 50ms per character
- 📝 **Output typing speed**: 30ms per character
- 🤔 **Thinking pause**: 200-2000ms random delay
- 📏 **Terminal width**: Auto-detected (defaults to 80 columns)

## 🎬 Movie References

Hackerminal includes humorous references to:
- 🕶️ The Matrix
- 💾 Hackers (1995)
- 🎮 WarGames
- 👽 Alien franchise
- 🦎 Predator franchise
- 🤖 Terminator franchise
- 🔍 CSI: Cyber
- 🎭 Mr. Robot
- 🌟 And many more!

## ⚖️ Disclaimer

This program is purely for entertainment and educational purposes. It does not perform any actual hacking, network operations, or security testing. All "hacking" activities are simulated text output with sleep delays.

> ⛔ **DO NOT** use this to intimidate, deceive, or misrepresent actual hacking capabilities.

## 📜 License

This is a parody/entertainment project. Use responsibly and ethically.

> 🎭 For entertainment purposes only!

## 🤝 Contributing

Feel free to add more parody sequences, improve visual effects, or add references to your favorite hacker movies!

> 💻 Pull requests are welcome! Let's make this even more ridiculous!

## 🎉 Fun Facts

- 🔑 The password "hunter2" reference comes from a classic IRC joke
- 💉 "Bobby Tables" is a reference to the famous XKCD comic about SQL injection
- 🎬 Many sequences quote iconic movie lines
- 🔤 The program name is intentionally misspelled as "l33t-h4x0r" in the prompt

---

<div align="center">

### 🎬 Enjoy the Hollywood hacking experience! 💻

**Made with 💚 (phosphor green) and 🧡 (amber)**

*"Hack the planet!" - Hackers (1995)*

</div>
