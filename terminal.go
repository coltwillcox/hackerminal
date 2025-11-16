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

type HackerTerminal struct {
	username  string
	target    string
	sequences []func()
}

func NewHackerTerminal() *HackerTerminal {
	hackerTerminal := &HackerTerminal{
		username: usernames[rand.Intn(len(usernames))],
		target:   targets[rand.Intn(len(targets))],
	}
	hackerTerminal.createSequences()

	return hackerTerminal
}

func (h *HackerTerminal) typeText(text string, delayMs int) {
	for _, char := range text {
		fmt.Print(string(char))
		time.Sleep(time.Duration(delayMs) * time.Millisecond)
	}
	fmt.Println()
}

func (h *HackerTerminal) randomPause() {
	// Random delay between 200ms and 2000ms to simulate thinking
	delay := 200 + rand.Intn(1800)
	time.Sleep(time.Duration(delay) * time.Millisecond)
}

func (h *HackerTerminal) showPrompt() {
	// Classic CRT phosphor green/amber monochrome terminal style

	// Left side segments
	leftPrompt := "\033[38;5;22m\033[0m"

	// User segment - Bright phosphor green
	leftPrompt += "\033[48;5;22m\033[38;5;46m"
	leftPrompt += " 󰀄 " + h.username
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

func (h *HackerTerminal) showBanner() {
	banner := `
╔═══════════════════════════════════════════════════════════════╗
║                                                               ║
║    █   █  █████  █████  █  █  █████  █████    ████   ████     ║
║    █   █  █   █  █      █ █   █      █   █    █   █  █   █    ║
║    █████  █████  █      ██    ████   ████     █   █  ████     ║
║    █   █  █   █  █      █ █   █      █   █    █   █  █   █    ║
║    █   █  █   █  █████  █  █  █████  █   █    ████   ████     ║
║                                                               ║
║             "I'm in!" - Every movie hacker ever               ║
║                                                               ║
╚═══════════════════════════════════════════════════════════════╝
`

	termWidth := getTerminalWidth()
	lines := strings.Split(banner, "\n")
	maxLineLen := 0
	for _, line := range lines {
		maxLineLen = max(maxLineLen, visibleLength(line))
	}
	padding := max((termWidth-maxLineLen)/2, 0)

	for _, line := range lines {
		if line == "" {
			fmt.Println()
			continue
		}
		fmt.Print(strings.Repeat(" ", padding))
		fmt.Println("\033[38;5;46m" + line)
	}

	time.Sleep(1 * time.Second)
}
