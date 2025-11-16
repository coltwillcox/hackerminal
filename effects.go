package main

import (
	"fmt"
	"math/rand"
	"strings"
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

func (h *HackerTerminal) screenGlitch() {
	// Glitch characters for corruption effect
	glitchChars := "█▓▒░▄▀▌▐╳╱╲◤◥◢◣▲▼◀▶●○□■"
	corruptChars := "!@#$%^&*(){}[]|\\/<>?~`"
	blockChars := "▀▁▂▃▄▅▆▇█▉▊▋▌▍▎▏"

	width := getTerminalWidth()
	glitchLines := 2 + rand.Intn(14) // 2-15 lines of glitch

	// Random glitch colors (bright, eye-catching)
	colors := []string{
		"\033[31m", // Red
		"\033[35m", // Magenta
		"\033[36m", // Cyan
		"\033[37m", // White
		"\033[91m", // Bright red
		"\033[95m", // Bright magenta
		"\033[97m", // Bright white
	}

	for range glitchLines {
		line := ""

		// Choose glitch pattern type
		glitchType := rand.Intn(5)

		switch glitchType {
		case 0: // Full static line
			color := colors[rand.Intn(len(colors))]
			for range width {
				line += string(glitchChars[rand.Intn(len(glitchChars))])
			}
			fmt.Printf("%s%s\033[0m\n", color, line)
		case 1: // Partial corruption (some spaces)
			for i := range width {
				if rand.Float32() > 0.3 {
					color := colors[rand.Intn(len(colors))]
					line += color + string(corruptChars[rand.Intn(len(corruptChars))]) + "\033[0m"
				} else {
					line += " "
				}
				if i%20 == 0 && rand.Float32() > 0.7 {
					line += "\033[7m" // Reverse video
				}
			}
			fmt.Println(line)
		case 2: // Block corruption
			color := colors[rand.Intn(len(colors))]
			blockSize := 5 + rand.Intn(15)
			startPos := rand.Intn(width - blockSize)
			line = strings.Repeat(" ", startPos)
			for range blockSize {
				line += string(blockChars[rand.Intn(len(blockChars))])
			}
			fmt.Printf("%s%s\033[0m\n", color, line)
		case 3: // Scrambled text effect
			words := []string{"ERROR", "CORRUPT", "BREACH", "FATAL", "SYSTEM", "MEMORY", "BUFFER", "OVERFLOW"}
			word := words[rand.Intn(len(words))]
			scrambled := ""
			for _, ch := range word {
				if rand.Float32() > 0.5 {
					scrambled += string(glitchChars[rand.Intn(len(glitchChars))])
				} else {
					scrambled += string(ch)
				}
			}
			padding := rand.Intn(width - len(scrambled))
			color := colors[rand.Intn(len(colors))]
			fmt.Printf("%s%s%s\033[0m\n", strings.Repeat(" ", padding), color, scrambled)
		case 4: // Horizontal tear effect
			tearPoint := rand.Intn(width)
			leftPart := strings.Repeat("▓", tearPoint)
			rightPart := strings.Repeat("░", width-tearPoint)
			color1 := colors[rand.Intn(len(colors))]
			color2 := colors[rand.Intn(len(colors))]
			fmt.Printf("%s%s%s%s\033[0m\n", color1, leftPart, color2, rightPart)
		}

		time.Sleep(time.Duration(30+rand.Intn(70)) * time.Millisecond)
	}

	// Brief pause then clear the glitch (move cursor up and clear lines)
	time.Sleep(500 * time.Millisecond)
	for range glitchLines {
		fmt.Print("\033[1A\033[2K") // Move up and clear line
	}
}
