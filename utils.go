package main

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
)

var ansi = regexp.MustCompile(`\x1b\[[0-9;]*m`)

func getTerminalWidth() int {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		return 80 // Default width if unable to detect
	}

	// Output format is "rows columns"
	parts := strings.Fields(string(out))
	if len(parts) >= 2 {
		width, err := strconv.Atoi(parts[1])
		if err == nil && width > 0 {
			return width
		}
	}

	return 80 // Default width
}

func visibleLength(s string) int {
	clean := ansi.ReplaceAllString(s, "")
	return utf8.RuneCountInString(clean)
}

func printSeparator() {
	fmt.Println()
}
