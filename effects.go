package main

import (
	"fmt"
	"math/rand"
	"time"
)

func (h *HackerTerminal) progressBar(task string, duration time.Duration) {
	fmt.Printf("%s [", task)
	steps := 30
	stepDuration := duration / time.Duration(steps)

	for range steps {
		fmt.Print("█")
		time.Sleep(stepDuration)
	}
	fmt.Println("] \033[32mDONE\033[0m")
}

func (h *HackerTerminal) spinner(task string, duration time.Duration) {
	spinChars := []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
	// Alternative spinners for variety
	altSpinners := [][]string{
		{"◐", "◓", "◑", "◒"},
		{"◴", "◷", "◶", "◵"},
		{"▹▹▹▹▹", "▸▹▹▹▹", "▹▸▹▹▹", "▹▹▸▹▹", "▹▹▹▸▹", "▹▹▹▹▸"},
		{"[    ]", "[=   ]", "[==  ]", "[=== ]", "[ ===]", "[  ==]", "[   =]"},
	}

	var chars []string
	if rand.Float32() > 0.5 {
		chars = spinChars
	} else {
		chars = altSpinners[rand.Intn(len(altSpinners))]
	}

	iterations := int(duration.Milliseconds() / 100)
	for i := range iterations {
		char := chars[i%len(chars)]
		fmt.Printf("\r\033[36m%s\033[0m %s", char, task)
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Printf("\r\033[32m✓\033[0m %s \033[32mDONE\033[0m\n", task)
}

func (h *HackerTerminal) fakeIPScan() {
	ips := []string{
		"192.168.1.1", "10.0.0.1", "172.16.0.1", "8.8.8.8", "1.1.1.1", "127.0.0.1", "192.168.0.255",
	}

	fmt.Println("\033[36m[*] Scanning network...\033[0m")
	time.Sleep(500 * time.Millisecond)

	for i := 0; i < 5; i++ {
		ip := ips[rand.Intn(len(ips))]
		status := "OPEN"
		if rand.Float32() > 0.7 {
			status = "FILTERED"
		}
		fmt.Printf("    %s - Port %d: %s\n", ip, 20+rand.Intn(9000), status)
		time.Sleep(time.Duration(100+rand.Intn(300)) * time.Millisecond)
	}
}

func (h *HackerTerminal) matrixRain() {
	chars := "01アイウエオカキクケコサシスセソタチツテト"
	lines := 3
	width := min(getTerminalWidth(), 120) // Cap at reasonable size

	fmt.Println("\033[32m")
	for range lines {
		output := ""
		for i := 0; i < width-2; i++ {
			output += string(chars[rand.Intn(len(chars))])
		}
		fmt.Println(output)
		time.Sleep(50 * time.Millisecond)
	}
	fmt.Print("\033[0m")
}
