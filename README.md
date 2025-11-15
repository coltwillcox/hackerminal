# Hackerminal

A parody terminal application that simulates Hollywood-style "hacking" scenes with theatrical command sequences and visual effects.

## Description

Hackerminal is an entertainment program that recreates the unrealistic and dramatic hacking scenes commonly seen in movies and TV shows. It displays randomized command sequences with fancy terminal effects, complete with progress bars, spinners, and references to popular sci-fi and hacker movies.

**Warning:** This is a parody. Real hacking is illegal and boring.

## Features

- **Fancy Terminal Prompt**: CRT phosphor green/amber monochrome terminal style with user, host, directory, git branch, and timestamp segments
- **Typewriter Effect**: Commands and outputs appear character-by-character with realistic typing delays
- **Random Thinking Pauses**: Random delays (200ms-2000ms) after prompts to simulate human thinking
- **60+ Parody Sequences**: Including references to:
  - Classic hacker movies (The Matrix, Hackers, WarGames)
  - Sci-fi franchises (Alien, Predator, Terminator)
  - Popular TV shows (CSI: Cyber, Mr. Robot)
  - Internet culture and memes

## Visual Effects

- Progress bars with ASCII characters
- Multiple spinner animations
- Matrix-style rain effect
- Fake IP scanning
- ANSI color codes for dramatic output

## Installation

### Prerequisites

- Go 1.16 or higher
- Unix-like terminal (Linux, macOS)

### Build

```bash
go build -o hackerminal main.go
```

## Usage

Simply run the program:

```bash
./hackerminal
```

The program will continuously display random hacking sequences. Press `Ctrl+C` to exit.

## Example Sequences

- SSH into fictional systems like "pentagon.gov" or "cyberdyne.sys"
- "Crack" passwords and find weak credentials
- SQL injection demonstrations
- "Enhance" images to impossible levels
- Deploy viruses that rickroll targets
- Many more Hollywood hacking tropes

## How It Works

1. Generates a random "hacker" username and target system
2. Displays a fancy terminal prompt
3. Pauses randomly to simulate thinking
4. Types out a command with typewriter effect
5. Shows dramatic output with animations
6. Repeats with a new random sequence

## Configuration

The program uses:
- **Command typing speed**: 50ms per character
- **Output typing speed**: 30ms per character
- **Thinking pause**: 200-2000ms random delay
- **Terminal width**: Auto-detected (defaults to 80 columns)

## Movie References

Hackerminal includes humorous references to:
- The Matrix
- Hackers (1995)
- WarGames
- Alien franchise
- Predator franchise
- Terminator franchise
- CSI: Cyber
- Mr. Robot
- And many more!

## Disclaimer

This program is purely for entertainment and educational purposes. It does not perform any actual hacking, network operations, or security testing. All "hacking" activities are simulated text output with sleep delays.

**DO NOT** use this to intimidate, deceive, or misrepresent actual hacking capabilities.

## License

This is a parody/entertainment project. Use responsibly and ethically.

## Contributing

Feel free to add more parody sequences, improve visual effects, or add references to your favorite hacker movies!

## Fun Facts

- The password "hunter2" reference comes from a classic IRC joke
- "Bobby Tables" is a reference to the famous XKCD comic about SQL injection
- Many sequences quote iconic movie lines
- The program name is intentionally misspelled as "1337-h4x0r" in the prompt

Enjoy the Hollywood hacking experience! 🎬🖥️
