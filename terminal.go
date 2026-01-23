package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var (
	usernames = []string{
		"cyb3rn1nja",
		"h4ck3rm4n",
		"zero_cool",
		"acidburn",
		"crash_override",
		"phantom_phreak",
	}
	targets = []string{
		"mainframe",
		"pentagon.gov",
		"cyberdyne.sys",
		"oscorp.net",
		"umbrella.corp",
		"weyland.industries",
		"nostromo.ship",
		"sulaco.vessel",
		"mother.ai",
		"predator.net",
	}
)

// HackerTerminal manages the terminal display and sequence execution
type HackerTerminal struct {
	Username        string
	Target          string
	Sequences       []Sequence
	CurrentSequence Sequence
	Stats           *Stats
}

// NewHackerTerminal creates a new terminal with random username and target
func NewHackerTerminal() *HackerTerminal {
	hackerTerminal := &HackerTerminal{
		Username: usernames[rand.Intn(len(usernames))],
		Target:   targets[rand.Intn(len(targets))],
	}
	hackerTerminal.CreateSequences()

	return hackerTerminal
}

// TypeText prints text with a typing effect
func (h *HackerTerminal) TypeText(text string, delayMs int) {
	for _, char := range text {
		fmt.Print(string(char))
		time.Sleep(time.Duration(delayMs) * time.Millisecond)
	}
	printSeparator()
}

// TypeCommand prints a command with typing effect, realistic typos, and command tracking
func (h *HackerTerminal) TypeCommand(text string, delayMs int) {
	// Adjacent keys on QWERTY keyboard for realistic typos
	adjacentKeys := map[rune][]rune{
		'a': {'s', 'q', 'z', 'w'},
		'b': {'v', 'g', 'h', 'n'},
		'c': {'x', 'd', 'f', 'v'},
		'd': {'s', 'e', 'r', 'f', 'c', 'x'},
		'e': {'w', 'r', 'd', 's'},
		'f': {'d', 'r', 't', 'g', 'v', 'c'},
		'g': {'f', 't', 'y', 'h', 'b', 'v'},
		'h': {'g', 'y', 'u', 'j', 'n', 'b'},
		'i': {'u', 'o', 'k', 'j'},
		'j': {'h', 'u', 'i', 'k', 'm', 'n'},
		'k': {'j', 'i', 'o', 'l', 'm'},
		'l': {'k', 'o', 'p'},
		'm': {'n', 'j', 'k'},
		'n': {'b', 'h', 'j', 'm'},
		'o': {'i', 'p', 'l', 'k'},
		'p': {'o', 'l'},
		'q': {'w', 'a'},
		'r': {'e', 't', 'f', 'd'},
		's': {'a', 'w', 'e', 'd', 'x', 'z'},
		't': {'r', 'y', 'g', 'f'},
		'u': {'y', 'i', 'j', 'h'},
		'v': {'c', 'f', 'g', 'b'},
		'w': {'q', 'e', 's', 'a'},
		'x': {'z', 's', 'd', 'c'},
		'y': {'t', 'u', 'h', 'g'},
		'z': {'a', 's', 'x'},
	}

	// Calculate variable typing speed based on character complexity
	getCharDelay := func(char rune) int {
		baseDelay := delayMs

		switch {
		// Common lowercase letters (home row and frequent) - fastest
		case char == 'a' || char == 's' || char == 'd' || char == 'f' ||
			char == 'e' || char == 't' || char == 'n' || char == 'o' ||
			char == 'i' || char == 'r' || char == 'l':
			return baseDelay - rand.Intn(typeSpeed1)

		// Common punctuation (space, dash, underscore, dot) - fairly fast
		case char == ' ' || char == '-' || char == '_' || char == '.':
			return baseDelay + rand.Intn(typeSpeed2)

		// Other lowercase letters - normal speed
		case char >= 'a' && char <= 'z':
			return baseDelay + rand.Intn(typeSpeed3)

		// Numbers - moderate slowdown
		case char >= '0' && char <= '9':
			return baseDelay + 15 + rand.Intn(typeSpeed4)

		// Uppercase letters - slower (need shift key)
		case char >= 'A' && char <= 'Z':
			return baseDelay + 20 + rand.Intn(typeSpeed5)

		// Special characters requiring shift or hard to reach - slowest
		case char == '!' || char == '@' || char == '#' || char == '$' ||
			char == '%' || char == '^' || char == '&' || char == '*' ||
			char == '(' || char == ')' || char == '+' || char == '=' ||
			char == '{' || char == '}' || char == '[' || char == ']' ||
			char == '|' || char == '\\' || char == ':' || char == ';' ||
			char == '"' || char == '\'' || char == '<' || char == '>' ||
			char == '?' || char == '/' || char == '~' || char == '`':
			return baseDelay + 30 + rand.Intn(typeSpeed6)

		default:
			return baseDelay + rand.Intn(typeSpeed0)
		}
	}

	for i, char := range text {
		charDelay := getCharDelay(char)

		if rand.Float32() < chanceTypo && i < len(text)-1 {
			lowerChar := char
			if char >= 'A' && char <= 'Z' {
				lowerChar = char + 32 // Convert to lowercase for lookup
			}

			if adjacent, ok := adjacentKeys[lowerChar]; ok {
				// Type the wrong character
				typoChar := adjacent[rand.Intn(len(adjacent))]
				// Preserve case
				if char >= 'A' && char <= 'Z' {
					typoChar = typoChar - 32
				}
				fmt.Print(string(typoChar))
				time.Sleep(time.Duration(charDelay) * time.Millisecond)

				// Brief pause before noticing the mistake
				time.Sleep(time.Duration(100+rand.Intn(200)) * time.Millisecond)

				// Backspace to delete the typo
				fmt.Print("\b \b")
				time.Sleep(time.Duration(50+rand.Intn(100)) * time.Millisecond)

				// Now type the correct character
				fmt.Print(string(char))
				time.Sleep(time.Duration(charDelay) * time.Millisecond)
				continue
			}
		}

		fmt.Print(string(char))
		time.Sleep(time.Duration(charDelay) * time.Millisecond)
	}
	printSeparator()

	// Track command
	if h.Stats != nil {
		h.Stats.TrackCommand()
	}
}

func (h *HackerTerminal) RandomPause() {
	// Random delay between 200ms and 2000ms to simulate thinking
	delay := 200 + rand.Intn(1800)
	time.Sleep(time.Duration(delay) * time.Millisecond)
}

// ShowPrompt displays the terminal prompt with user, host, directory, git branch, and timestamp
func (h *HackerTerminal) ShowPrompt() {
	// Classic CRT phosphor green/amber monochrome terminal style

	// Left side segments
	leftPrompt := "\033[38;5;22m\033[0m"

	// User segment - Bright phosphor green
	leftPrompt += "\033[48;5;22m\033[38;5;46m"
	leftPrompt += " 󰀄 " + h.Username
	leftPrompt += " \033[0m\033[38;5;22m\033[0m"

	// Host segment - Medium green
	leftPrompt += "\033[48;5;28m\033[38;5;118m"
	leftPrompt += " 󰒋 l33t-h4x0r"
	leftPrompt += " \033[0m\033[38;5;28m\033[48;5;22m\033[0m"

	// Directory segment - Dark green background with bright text
	leftPrompt += "\033[48;5;22m\033[38;5;82m"
	leftPrompt += " ~/top_secret"
	leftPrompt += " \033[0m\033[38;5;22m\033[0m"
	leftPrompt += "\033[38;5;22m\033[0m"

	// Right side segments
	rightPrompt := ""

	// Git-like segment (random branch name) - Slightly dimmed phosphor
	branches := []string{"master", "main", "hack-branch", "exploit-dev", "zero-day"}
	if rand.Float32() > 0.3 {
		branch := branches[rand.Intn(len(branches))]
		rightPrompt += "\033[38;5;58m\033[0m"
		rightPrompt += "\033[38;5;58m\033[0m"
		rightPrompt += "\033[48;5;58m\033[38;5;154m"
		rightPrompt += " " + branch
		rightPrompt += " \033[0m"
	}

	// Time segment - Amber phosphor variant (like old amber terminals)
	if rightPrompt != "" {
		rightPrompt += "\033[38;5;94m\033[48;5;58m\033[0m"
	} else {
		rightPrompt += "\033[38;5;94m\033[0m"
		rightPrompt += "\033[38;5;94m\033[0m"
	}
	rightPrompt += "\033[48;5;94m\033[38;5;220m"
	rightPrompt += " 󱑎 " + time.Now().Format("15:04:05")
	rightPrompt += " \033[0m"
	rightPrompt += "\033[38;5;94m\033[0m"

	// Print the prompt - all left-justified, no spacing between segments
	promptWidth := visibleLength(leftPrompt) + visibleLength(rightPrompt)

	fmt.Print(leftPrompt)

	// Get terminal width and draw a line to the end
	termWidth := getTerminalWidth()
	remainingWidth := termWidth - promptWidth
	if remainingWidth > 0 {
		// Draw a thin line using box drawing character
		fmt.Print("\033[38;5;22m")
		fmt.Print(strings.Repeat("─", remainingWidth))
		fmt.Print("\033[0m")
	}

	fmt.Print(rightPrompt)

	// New line with prompt character - Classic bright green
	fmt.Print("\n")
	fmt.Print("\033[38;5;46m❯\033[0m ")
}

// ShowBanner displays the application banner with dashboard statistics
func (h *HackerTerminal) ShowBanner() {
	// Display dashboard stats before banner
	h.showDashboard()

	// Display main banner
	h.DrawCentered(banner, "\033[38;5;46m", 1000, false)
}

func (h *HackerTerminal) showDashboard() {
	termWidth := getTerminalWidth()

	// Generate random stats
	cpuUsage := 85.0 + rand.Float64()*14.0 // 85-99%
	memUsage := 2.0 + rand.Float64()*30.0  // 2-32 GB
	uptime := 1 + rand.Intn(99)            // 1-99 hours
	threats := rand.Intn(150)              // 0-149
	connActive := 100 + rand.Intn(999)     // 100-1098
	pktsPerSec := 1000 + rand.Intn(9000)   // 1000-9999

	// Left side stats
	leftStats := []string{
		"\033[32m●\033[0m SYS: ONLINE",
		"\033[33m▲\033[0m NET: BREACH",
		"\033[31m◆\033[0m FW: BYPASSED",
		"\033[36m◉\033[0m PORT: 4444",
	}

	// Right side stats
	rightStats := []string{
		fmt.Sprintf("CPU: \033[33m%.1f %%\033[0m", cpuUsage),
		fmt.Sprintf("MEM: \033[36m%.1f GB\033[0m", memUsage),
		fmt.Sprintf("UP: \033[32m%d h\033[0m", uptime),
		fmt.Sprintf("THR: \033[31m%d\033[0m", threats),
	}

	// Bottom stats bar
	bottomStats := fmt.Sprintf("PROTO: \033[36mTCP/IP\033[0m | CONN: \033[36m%d\033[0m | PKT/s: \033[36m%d\033[0m | TARGET: \033[36m%s\033[0m", connActive, pktsPerSec, h.Target)

	// Display top stats (left and right aligned)
	for i := range leftStats {
		leftLen := visibleLength(leftStats[i])
		rightLen := visibleLength(rightStats[i])
		spacing := max(termWidth-leftLen-rightLen, 1)
		fmt.Printf("%s%s%s\n", leftStats[i], strings.Repeat(" ", spacing), rightStats[i])
	}

	// Display bottom stats bar (centered)
	bottomLen := visibleLength(bottomStats)
	padding := max((termWidth-bottomLen)/2, 0)
	fmt.Print(strings.Repeat(" ", padding))
	fmt.Println(bottomStats)
}

// DrawCentered displays text centered on screen with optional hold time and clearing
func (h *HackerTerminal) DrawCentered(image, color string, hold int64, clear bool) {
	// Get terminal width for centering
	termWidth := getTerminalWidth()

	// Split art into lines and display centered
	lines := strings.Split(image, "\n")

	// Find maximum line width for consistent centering
	maxLineLen := 0
	for _, line := range lines {
		maxLineLen = max(maxLineLen, visibleLength(line))
	}

	padding := max((termWidth-maxLineLen)/2, 0)

	for _, line := range lines {
		if line == "" {
			printSeparator()
			continue
		}
		fmt.Print(strings.Repeat(" ", padding))
		fmt.Printf("%s%s\033[0m\n", color, line)
	}

	// Hold the splash for a moment
	time.Sleep(time.Duration(hold) * time.Millisecond)

	// Clear the splash screen
	if clear {
		for range len(lines) {
			fmt.Print("\033[1A\033[2K")
		}
	}
}

// PrintNotification displays a notification message using centered formatting
func (h *HackerTerminal) PrintNotification(notification, color string, hold int64) {
	h.DrawCentered(notification, color, hold, false)
}

// RandomizeSequence selects a random sequence and sets it as the current sequence
func (h *HackerTerminal) RandomizeSequence() {
	h.CurrentSequence = h.Sequences[rand.Intn(len(h.Sequences))]
}

// RunCurrentSequence executes the currently selected sequence
func (h *HackerTerminal) RunCurrentSequence() {
	if h.CurrentSequence.fn != nil {
		h.CurrentSequence.fn()
	}
}

// TrackSequence records the current sequence in statistics
func (h *HackerTerminal) TrackSequence() {
	if h.Stats != nil && h.CurrentSequence.name != "" {
		h.Stats.TrackSequence(h.CurrentSequence.name)
	}
}

// SaveStats persists statistics to disk
func (h *HackerTerminal) SaveStats() {
	if h.Stats != nil {
		h.Stats.Save()
	}
}
