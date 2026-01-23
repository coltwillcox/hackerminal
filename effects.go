package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// ProgressBar displays an animated progress bar for the given task and duration
func (h *HackerTerminal) ProgressBar(task string, duration time.Duration) {
	fmt.Printf("%s [", task)
	steps := 30
	stepDuration := duration / time.Duration(steps)

	for range steps {
		fmt.Print("█")
		time.Sleep(stepDuration)
	}
	fmt.Println("] \033[32mDONE\033[0m")
}

// Spinner displays an animated spinner for the given task and duration
func (h *HackerTerminal) Spinner(task string, duration time.Duration) {
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

// FakeIPScan simulates a network IP scan with random results
func (h *HackerTerminal) FakeIPScan() {
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

// MatrixRain displays a Matrix-style character rain effect
func (h *HackerTerminal) MatrixRain() {
	// Track effect
	if h.Stats != nil {
		h.Stats.TrackEffect("matrix_rain")
	}

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

	termWidth := getTerminalWidth()
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
			for range termWidth {
				line += string(glitchChars[rand.Intn(len(glitchChars))])
			}
			fmt.Printf("%s%s\033[0m\n", color, line)
		case 1: // Partial corruption (some spaces)
			for i := range termWidth {
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
			startPos := rand.Intn(termWidth - blockSize)
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
			padding := rand.Intn(termWidth - len(scrambled))
			color := colors[rand.Intn(len(colors))]
			fmt.Printf("%s%s%s\033[0m\n", strings.Repeat(" ", padding), color, scrambled)
		case 4: // Horizontal tear effect
			tearPoint := rand.Intn(termWidth)
			leftPart := strings.Repeat("▓", tearPoint)
			rightPart := strings.Repeat("░", termWidth-tearPoint)
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

func (h *HackerTerminal) crtScanLines() {
	termWidth := getTerminalWidth()
	numLines := 8 + rand.Intn(12) // 8-19 lines of scan effect

	// CRT phosphor green colors with varying intensity
	brightLine := "\033[38;5;46m"   // Bright green (phosphor glow)
	dimLine := "\033[38;5;22m"      // Dim green (scan line shadow)
	veryDimLine := "\033[38;5;234m" // Very dim (almost black scan line)

	// Generate some fake "data" characters
	dataChars := "01█▓▒░│┤┐└┴┬├─┼"

	fmt.Println("\033[36m[*] Initializing CRT display mode...\033[0m")
	time.Sleep(300 * time.Millisecond)

	for i := range numLines {
		line := ""

		// Create scan line pattern - alternating bright and dim
		if i%2 == 0 {
			// Bright scan line with data
			line += brightLine
			dataLen := rand.Intn(termWidth / 2)
			padding := rand.Intn(termWidth - dataLen - 10)
			line += strings.Repeat(" ", padding)

			// Add some "data" on bright lines
			for range dataLen {
				if rand.Float32() > 0.7 {
					line += string(dataChars[rand.Intn(len(dataChars))])
				} else {
					line += " "
				}
			}
		} else {
			// Dim scan line (the dark gap between phosphor lines)
			if rand.Float32() > 0.3 {
				line += dimLine
				// Occasional horizontal line to simulate scan
				line += strings.Repeat("─", termWidth)
			} else {
				line += veryDimLine
				line += strings.Repeat("▁", termWidth)
			}
		}

		fmt.Printf("%s\033[0m\n", line)
		time.Sleep(time.Duration(20+rand.Intn(40)) * time.Millisecond)
	}

	// Simulate CRT refresh flicker
	time.Sleep(200 * time.Millisecond)
	fmt.Print("\033[?5h") // Reverse video (flash)
	time.Sleep(50 * time.Millisecond)
	fmt.Print("\033[?5l") // Normal video
	time.Sleep(100 * time.Millisecond)

	// Clear the scan lines
	for range numLines + 1 { // +1 for the initial message
		fmt.Print("\033[1A\033[2K")
	}
}

func (h *HackerTerminal) asciiSplash() {
	// Collection of hacker-themed ASCII art
	artworks := []struct {
		art   string
		color string
	}{
		{
			art:   skull,
			color: "\033[31m", // Red
		},
		{
			art:   hacker,
			color: "\033[32m", // Green
		},
		{
			art:   breach,
			color: "\033[36m", // Cyan
		},
		{
			art:   lockpick,
			color: "\033[33m", // Yellow
		},
		{
			art:   cybereye,
			color: "\033[35m", // Magenta
		},
		{
			art:   glider,
			color: "\033[37m", // White
		},
		{
			art:   trojan,
			color: "\033[91m", // Bright red
		},
		{
			art:   matrix,
			color: "\033[32m", // Green
		},
		{
			art:   vhs,
			color: "\033[35m", // Magenta
		},
		{
			art:   amiga,
			color: "\033[33m", // Yellow
		},
		{
			art:   keygen,
			color: "\033[96m", // Bright cyan
		},
	}

	// Select random artwork
	selected := artworks[rand.Intn(len(artworks))]
	h.DrawCentered(selected.art, selected.color, 1500, true)
}

func (h *HackerTerminal) networkTopology() {
	// Collection of network topology ASCII art diagrams
	topologies := []struct {
		title        string
		colorTitle   string
		colorDiagram string
		diagram      string
	}{
		{
			title:        "COMPROMISED NETWORK MAP",
			colorDiagram: "\033[36m",
			colorTitle:   "\033[33m",
			diagram:      diagram1,
		},
		{
			title:        "ATTACK VECTOR PATH",
			colorDiagram: "\033[36m",
			colorTitle:   "\033[33m",
			diagram:      diagram2,
		},
		{
			title:        "LATERAL MOVEMENT MAP",
			colorDiagram: "\033[36m",
			colorTitle:   "\033[33m",
			diagram:      diagram3,
		},
		{
			title:        "BOTNET COMMAND & CONTROL",
			colorDiagram: "\033[36m",
			colorTitle:   "\033[33m",
			diagram:      diagram4,
		},
		{
			title:        "INTERNAL NETWORK SCAN",
			colorDiagram: "\033[36m",
			colorTitle:   "\033[33m",
			diagram:      diagram5,
		},
		{
			title:        "VPN TUNNEL ARCHITECTURE",
			colorDiagram: "\033[36m",
			colorTitle:   "\033[33m",
			diagram:      diagram6,
		},
	}

	// Select random topology
	selected := topologies[rand.Intn(len(topologies))]
	h.DrawCentered(selected.title, selected.colorTitle, 0, false)
	h.DrawCentered(selected.diagram, selected.colorDiagram, 1500, false)
}

func (h *HackerTerminal) splitScreen() {
	termWidth := getTerminalWidth()
	splitPos := termWidth / 2

	// Window titles
	leftTitle := "WINDOW 1: EXPLOIT"
	rightTitle := "WINDOW 2: MONITOR"

	// Left and right window activities
	leftActivities := []string{
		"[*] Scanning ports...",
		"[+] Port 22: OPEN",
		"[+] Port 80: OPEN",
		"[*] Attempting SSH...",
		"[+] Credentials found!",
		"[*] Escalating privileges...",
		"[+] ROOT ACCESS GRANTED",
	}

	rightActivities := []string{
		"[SYS] CPU: 94.2%",
		"[NET] Packets: 15248",
		"[MEM] Usage: 2.1GB",
		"[LOG] Login attempt detected",
		"[ALERT] Firewall breach!",
		"[WARN] Unauthorized access",
		"[CRIT] System compromised",
	}

	// Colors for windows
	leftColor := "\033[32m"   // Green
	rightColor := "\033[36m"  // Cyan
	borderColor := "\033[33m" // Yellow

	numLines := 12

	// Draw top border
	leftBorder := strings.Repeat("═", splitPos-2)
	rightBorder := strings.Repeat("═", termWidth-splitPos-2)
	fmt.Printf("%s╔%s╦%s╗\033[0m\n", borderColor, leftBorder, rightBorder)

	// Draw window titles
	leftTitleLeftMargin := (splitPos - len(leftTitle) - 2) / 2
	leftTitleRightMargin := visibleLength(leftBorder) - len(leftTitle) - leftTitleLeftMargin
	rightTitleLeftMargin := (termWidth - splitPos - len(rightTitle) - 2) / 2
	rightTitleRightMargin := visibleLength(rightBorder) - len(rightTitle) - rightTitleLeftMargin
	leftTitleFull := fmt.Sprintf("%s%s%s%s%s", strings.Repeat(" ", leftTitleLeftMargin), leftColor, leftTitle, "\033[0m"+borderColor, strings.Repeat(" ", leftTitleRightMargin))
	rightTitleFull := fmt.Sprintf("%s%s%s%s%s", strings.Repeat(" ", rightTitleLeftMargin), rightColor, rightTitle, "\033[0m"+borderColor, strings.Repeat(" ", rightTitleRightMargin))
	fmt.Printf("%s║%s║%s║\033[0m\n", borderColor, leftTitleFull, rightTitleFull)

	// Draw separator
	fmt.Printf("%s╠%s╬%s╣\033[0m\n", borderColor, leftBorder, rightBorder)

	// Draw content lines
	for i := range numLines {
		// Left window content
		var leftContent string
		if i < len(leftActivities) {
			leftContent = leftActivities[i]
		} else {
			leftContent = ""
		}

		// Right window content
		var rightContent string
		if i < len(rightActivities) {
			rightContent = rightActivities[i]
		} else {
			rightContent = ""
		}

		// Pad content to fit window width
		leftLen := visibleLength(leftContent)
		rightLen := visibleLength(rightContent)

		leftSpace := splitPos - leftLen - 3
		rightSpace := termWidth - splitPos - rightLen - 3

		fmt.Printf("%s║\033[0m %s%s%s%s%s║\033[0m %s%s%s%s%s║\033[0m\n",
			borderColor,
			leftColor, leftContent, "\033[0m",
			strings.Repeat(" ", max(leftSpace, 0)),
			borderColor,
			rightColor, rightContent, "\033[0m",
			strings.Repeat(" ", max(rightSpace, 0)),
			borderColor)

		time.Sleep(150 * time.Millisecond)
	}

	// Draw bottom border
	fmt.Printf("%s╚%s╩%s╝\033[0m\n", borderColor, leftBorder, rightBorder)

	// Hold for viewing
	time.Sleep(2000 * time.Millisecond)
}

// Hacker-themed directory structures
func (h *HackerTerminal) fileTree() {
	title := titles[rand.Intn(len(titles))]
	root := roots[rand.Intn(len(roots))]

	rand.Shuffle(len(dirs), func(i, j int) {
		dirs[i], dirs[j] = dirs[j], dirs[i]
	})
	selectedDirs := dirs[:rand.Intn(len(dirs)-2)+2]

	rand.Shuffle(len(files), func(i, j int) {
		files[i], files[j] = files[j], files[i]
	})
	selectedFiles := files[:rand.Intn(len(files)-2)+2]

	// Display title
	h.DrawCentered(title, "\033[33m", 0, false)
	printSeparator()

	// Display root with tree icon
	fmt.Printf("\033[36m📁 %s\033[0m\n", root)
	time.Sleep(200 * time.Millisecond)

	for _, dir := range selectedDirs {
		icon := "📁"
		color := "\033[36m" // Cyan for directories
		prefix := "├── "

		fmt.Printf("\033[90m%s\033[0m%s %s%s\033[0m", prefix, icon, color, dir)

		printSeparator()
		time.Sleep(time.Duration(100+rand.Intn(150)) * time.Millisecond)
	}

	for i, file := range selectedFiles {
		var icon string
		var color string
		if strings.HasSuffix(file, ".txt") || strings.HasSuffix(file, ".log") {
			icon = "📄"
			color = "\033[37m" // White
		} else if strings.HasSuffix(file, ".pdf") || strings.HasSuffix(file, ".doc") {
			icon = "📋"
			color = "\033[93m" // Yellow
		} else if strings.HasSuffix(file, ".sql") || strings.HasSuffix(file, ".db") {
			icon = "🗄️"
			color = "\033[35m" // Magenta
		} else if strings.HasSuffix(file, ".key") || strings.HasSuffix(file, ".pem") || strings.HasSuffix(file, ".gpg") {
			icon = "🔑"
			color = "\033[31m" // Red
		} else if strings.HasSuffix(file, ".zip") || strings.HasSuffix(file, ".tar") || strings.HasSuffix(file, ".gz") {
			icon = "📦"
			color = "\033[33m" // Yellow
		} else if strings.HasSuffix(file, ".json") || strings.HasSuffix(file, ".xml") || strings.HasSuffix(file, ".conf") {
			icon = "⚙️"
			color = "\033[36m" // Cyan
		} else {
			icon = "📄"
			color = "\033[32m" // Green
		}

		isLast := i == len(selectedFiles)-1

		// Tree characters
		var prefix string
		if isLast {
			prefix = "└── "
		} else {
			prefix = "├── "
		}

		fmt.Printf("\033[90m%s\033[0m%s %s%s\033[0m", prefix, icon, color, file)

		size := rand.Intn(9999) + 1
		unit := []string{"KB", "MB"}[rand.Intn(2)]
		fmt.Printf(" \033[90m(%d %s)\033[0m", size, unit)

		printSeparator()
		time.Sleep(time.Duration(100+rand.Intn(150)) * time.Millisecond)
	}

	printSeparator()

	fmt.Println(summaries[rand.Intn(len(summaries))])

	time.Sleep(1500 * time.Millisecond)
}

// RandomEffect triggers a random visual effect based on configured probabilities
func (h *HackerTerminal) RandomEffect() {
	// Random chance for visual effects
	effectRoll := rand.Float32()
	if effectRoll < chanceGlitch {
		time.Sleep(200 * time.Millisecond)
		if h.Stats != nil {
			h.Stats.TrackEffect("glitch")
		}
		h.screenGlitch()
	} else if effectRoll < chanceGlitch+chanceCrtScan {
		time.Sleep(200 * time.Millisecond)
		if h.Stats != nil {
			h.Stats.TrackEffect("crt_scan")
		}
		h.crtScanLines()
	} else if effectRoll < chanceGlitch+chanceCrtScan+chanceAsciiSplash {
		time.Sleep(200 * time.Millisecond)
		if h.Stats != nil {
			h.Stats.TrackEffect("splash")
		}
		h.asciiSplash()
	} else if effectRoll < chanceGlitch+chanceCrtScan+chanceAsciiSplash+chanceTopology {
		time.Sleep(200 * time.Millisecond)
		if h.Stats != nil {
			h.Stats.TrackEffect("topology")
		}
		h.networkTopology()
	} else if effectRoll < chanceGlitch+chanceCrtScan+chanceAsciiSplash+chanceTopology+chanceSplitScreen {
		time.Sleep(200 * time.Millisecond)
		if h.Stats != nil {
			h.Stats.TrackEffect("split_screen")
		}
		h.splitScreen()
	} else if effectRoll < chanceGlitch+chanceCrtScan+chanceAsciiSplash+chanceTopology+chanceSplitScreen+chanceFileTree {
		time.Sleep(200 * time.Millisecond)
		if h.Stats != nil {
			h.Stats.TrackEffect("file_tree")
		}
		h.fileTree()
	}
}
