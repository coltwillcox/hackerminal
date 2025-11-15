package main

import (
	"fmt"
	"time"
)

func main() {
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
