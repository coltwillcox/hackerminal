package main

import (
	"fmt"
	"time"
)

func main() {
	// Clear terminal screen
	fmt.Print("\033[H\033[2J")

	terminal := NewHackerTerminal()
	terminal.showBanner()

	fmt.Println("\033[33m[!] Warning: This is a parody. Real hacking is illegal and boring.\033[0m")
	fmt.Println("\033[33m[!] Press Ctrl+C to exit this Hollywood nonsense\033[0m")
	fmt.Println()
	time.Sleep(2 * time.Second)

	for {
		terminal.runSequence()
	}
}
