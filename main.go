package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	terminal := NewHackerTerminal()

	// Load statistics
	stats, err := LoadStats()
	if err != nil {
		fmt.Printf("Warning: Could not load stats: %v\n", err)
		stats, _ = LoadStats() // Try to create fresh stats
	}

	// Start new session
	terminal.stats = stats

	// Set up notification callback
	terminal.stats.OnNotification = terminal.PrintNotification

	terminal.stats.StartNewSession()

	// Set up signal handling for graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigChan
		fmt.Println("\n\n\033[33m[!] Shutting down...\033[0m")
		terminal.stats.PrintStats()
		if err := terminal.stats.Save(); err != nil {
			fmt.Printf("Warning: Could not save stats: %v\n", err)
		}
		os.Exit(0)
	}()

	// Clear terminal screen
	fmt.Print("\033[H\033[2J")

	terminal.showBanner()

	fmt.Println("\033[33m[!] Warning: This is a parody. Real hacking is illegal and boring.\033[0m")
	fmt.Println("\033[33m[!] Press Ctrl+C to exit this Hollywood nonsense\033[0m")
	printSeparator()
	time.Sleep(2 * time.Second)

	for {
		terminal.RunSequence()
	}
}
