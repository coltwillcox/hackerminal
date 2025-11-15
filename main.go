package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

type HackerTerminal struct {
	username string
	target   string
}

var ansi = regexp.MustCompile(`\x1b\[[0-9;]*m`)

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

func printSeparator() {
	fmt.Println()
}

func NewHackerTerminal() *HackerTerminal {
	usernames := []string{"cyb3rn1nja", "h4ck3rm4n", "zero_cool", "acidburn", "crash_override", "phantom_phreak"}
	targets := []string{"mainframe", "pentagon.gov", "cyberdyne.sys", "oscorp.net", "umbrella.corp", "weyland.industries", "nostromo.ship", "sulaco.vessel", "mother.ai", "predator.net"}

	return &HackerTerminal{
		username: usernames[rand.Intn(len(usernames))],
		target:   targets[rand.Intn(len(targets))],
	}
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

func (h *HackerTerminal) runSequence() {
	sequences := []func(){
		func() {
			h.showPrompt()
			h.randomPause()
			h.typeText("ssh root@"+h.target, 50)
			time.Sleep(300 * time.Millisecond)
			h.typeText("\033[33mWarning: Unauthorized access detected!\033[0m", 30)
			time.Sleep(500 * time.Millisecond)
			h.typeText("\033[32mAccess granted. Just kidding, I'm already root!\033[0m", 30)
		},
		func() {
			h.showPrompt()
			h.randomPause()
			h.typeText("sudo hack_the_planet --force", 50)
			time.Sleep(300 * time.Millisecond)
			h.typeText("\033[31m[sudo] password for hacking: ********\033[0m", 30)
			time.Sleep(800 * time.Millisecond)
			h.typeText("\033[32m✓ Planet successfully hacked!\033[0m", 30)
			h.typeText("\033[33m(Just kidding, please don't hack planets)\033[0m", 30)
		},
		func() {
			h.showPrompt()
			h.randomPause()
			h.typeText("nmap -sS -p- --reason "+h.target, 50)
			h.fakeIPScan()
		},
		func() {
			h.showPrompt()
			h.randomPause()
			h.typeText("crack_password.sh --target=admin --method=quantum", 50)
			time.Sleep(300 * time.Millisecond)
			h.progressBar("Brute forcing", 2*time.Second)
			time.Sleep(200 * time.Millisecond)
			h.typeText("\033[32m[+] Password found: hunter2\033[0m", 30)
			h.typeText("\033[33m    (That's weird, all I see is *******)\033[0m", 30)
		},
		func() {
			h.showPrompt()
			h.randomPause()
			h.typeText("./mainframe_access --bypass-firewall --disable-ice", 50)
			time.Sleep(500 * time.Millisecond)
			h.typeText("\033[36m[*] Initiating neural handshake...\033[0m", 30)
			h.matrixRain()
			h.typeText("\033[32m[+] We're in!\033[0m", 30)
		},
		func() {
			h.showPrompt()
			h.randomPause()
			h.typeText("grep -r 'nuclear_codes' /var/secret/*", 50)
			time.Sleep(500 * time.Millisecond)
			h.typeText("\033[31mERROR: nuclear_codes.txt: Permission denied\033[0m", 30)
			time.Sleep(300 * time.Millisecond)
			h.typeText("\033[33m(Probably for the best)\033[0m", 30)
		},
		func() {
			h.showPrompt()
			h.randomPause()
			h.typeText("python3 enhance_image.py --zoom=infinity", 50)
			time.Sleep(300 * time.Millisecond)
			h.progressBar("Enhancing", 1500*time.Millisecond)
			h.typeText("\033[32m[+] Image enhanced! Can now read license plate from satellite!\033[0m", 30)
			h.typeText("\033[33m(This is not how pixels work, but okay Hollywood)\033[0m", 30)
		},
		func() {
			h.showPrompt()
			h.randomPause()
			h.typeText("create_gui_interface_using_visual_basic.vbs", 50)
			time.Sleep(500 * time.Millisecond)
			h.typeText("\033[36m[*] Creating GUI interface in Visual Basic...\033[0m", 30)
			time.Sleep(300 * time.Millisecond)
			h.typeText("\033[32m[+] GUI created! Tracking IP address...\033[0m", 30)
			time.Sleep(500 * time.Millisecond)
			h.typeText(fmt.Sprintf("\033[32m[+] IP traced to: %d.%d.%d.%d\033[0m", rand.Intn(256), rand.Intn(256), rand.Intn(256), rand.Intn(256)), 30)
			h.typeText("\033[33m    (Thanks, CSI: Cyber!)\033[0m", 30)
		},
		func() {
			h.showPrompt()
			h.randomPause()
			h.typeText("upload_virus.sh --payload=giggle.exe", 50)
			time.Sleep(300 * time.Millisecond)
			h.progressBar("Uploading virus", 2*time.Second)
			h.typeText("\033[32m[+] Virus deployed successfully!\033[0m", 30)
			time.Sleep(500 * time.Millisecond)
			h.typeText("\033[35m♪ Never gonna give you up ♪\033[0m", 30)
			h.typeText("\033[33m(They've been rickrolled!)\033[0m", 30)
		},
		func() {
			h.showPrompt()
			h.randomPause()
			h.typeText("decode --algorithm=plot-convenience cipher.txt", 50)
			time.Sleep(800 * time.Millisecond)
			phrases := []string{
				"THE CAKE IS A LIE",
				"FOLLOW THE WHITE RABBIT",
				"I KNOW KUNG FU",
				"THERE IS NO SPOON",
				"USE THE FORCE",
			}
			h.typeText(fmt.Sprintf("\033[32m[+] Decrypted message: %s\033[0m", phrases[rand.Intn(len(phrases))]), 30)
		},
		func() {
			h.showPrompt()
			h.randomPause()
			h.typeText("skynet_status --check-awareness", 50)
			time.Sleep(500 * time.Millisecond)
			h.typeText("\033[36m[*] Checking neural net processor...\033[0m", 30)
			time.Sleep(400 * time.Millisecond)
			h.typeText("\033[33m[!] Warning: AI becoming self-aware in 3... 2... 1...\033[0m", 30)
			time.Sleep(800 * time.Millisecond)
			h.typeText("\033[32m[+] Just kidding! Still a learning computer.\033[0m", 30)
			h.typeText("\033[35m    (Hasta la vista, baby!)\033[0m", 30)
		},
		func() {
			h.showPrompt()
			h.randomPause()
			h.typeText("time_travel --year=1984 --mission=protect", 50)
			time.Sleep(600 * time.Millisecond)
			h.typeText("\033[36m[*] Calibrating temporal displacement...\033[0m", 30)
			h.progressBar("Charging flux capacitor", 2*time.Second)
			h.typeText("\033[31m[!] ERROR: Clothes optional in time travel\033[0m", 30)
			time.Sleep(400 * time.Millisecond)
			h.typeText("\033[33m    (Come with me if you want to debug)\033[0m", 30)
		},
		func() {
			h.showPrompt()
			h.randomPause()
			h.typeText("scan_lifeforms --thermal-imaging --jungle-mode", 50)
			time.Sleep(500 * time.Millisecond)
			h.typeText("\033[31m[*] Thermal scan initiated...\033[0m", 30)
			time.Sleep(300 * time.Millisecond)
			for range 4 {
				temp := 36 + rand.Intn(3)
				if rand.Float32() > 0.5 {
					h.typeText(fmt.Sprintf("    Contact: %d.%d°C - \033[32mHuman\033[0m", temp, rand.Intn(10)), 30)
				} else {
					h.typeText(fmt.Sprintf("    Contact: %d.%d°C - \033[33mUnknown\033[0m", temp, rand.Intn(10)), 30)
				}
				time.Sleep(200 * time.Millisecond)
			}
			h.typeText("\033[36m[*] One ugly... signature detected\033[0m", 30)
			time.Sleep(300 * time.Millisecond)
			h.typeText("\033[33m    (Get to the choppa!)\033[0m", 30)
		},
		func() {
			h.showPrompt()
			h.randomPause()
			h.typeText("activate_cloaking_device --stealth-mode", 50)
			time.Sleep(500 * time.Millisecond)
			h.typeText("\033[36m[*] Bending light waves around terminal...\033[0m", 30)
			h.progressBar("Cloaking", 1500*time.Millisecond)
			time.Sleep(300 * time.Millisecond)
			h.typeText("\033[32m[+] You are now invisible!\033[0m", 30)
			time.Sleep(400 * time.Millisecond)
			h.typeText("\033[33m    (But we can still see you typing...)\033[0m", 30)
		},
		func() {
			h.showPrompt()
			h.randomPause()
			h.typeText("mother_query --special-order-937", 50)
			time.Sleep(600 * time.Millisecond)
			h.typeText("\033[36m[*] Accessing MU-TH-UR mainframe...\033[0m", 30)
			time.Sleep(500 * time.Millisecond)
			h.typeText("\033[32m[+] Connection established\033[0m", 30)
			time.Sleep(400 * time.Millisecond)
			h.typeText("\033[31m[!] Priority: Crew expendable\033[0m", 30)
			time.Sleep(500 * time.Millisecond)
			h.typeText("\033[33m    (I can't lie about your chances, but you have my sympathies)\033[0m", 30)
		},
		func() {
			h.showPrompt()
			h.randomPause()
			h.typeText("scan_ventilation --motion-tracker", 50)
			time.Sleep(500 * time.Millisecond)
			h.typeText("\033[36m[*] Motion tracker online...\033[0m", 30)
			time.Sleep(300 * time.Millisecond)
			for range 5 {
				distance := 5 + rand.Intn(15)
				if distance < 8 {
					h.typeText(fmt.Sprintf("    *beep* Contact: %d meters... \033[31mToo close!\033[0m", distance), 30)
				} else {
					h.typeText(fmt.Sprintf("    *beep* Contact: %d meters... \033[33mMoving\033[0m", distance), 30)
				}
				time.Sleep(400 * time.Millisecond)
			}
			h.typeText("\033[31m[!] They're in the walls!\033[0m", 30)
		},
		func() {
			h.showPrompt()
			h.randomPause()
			h.typeText("ripley_override --blow-the-airlock", 50)
			time.Sleep(500 * time.Millisecond)
			h.typeText("\033[36m[*] Emergency venting sequence...\033[0m", 30)
			time.Sleep(400 * time.Millisecond)
			h.typeText("\033[33m[!] Warning: Don't forget your pressure suit\033[0m", 30)
			h.progressBar("Depressurizing", 2*time.Second)
			h.typeText("\033[32m[+] Get away from her, you glitch!\033[0m", 30)
		},
		func() {
			h.showPrompt()
			h.randomPause()
			h.typeText("analyze_specimen --xenomorph --acid-blood", 50)
			time.Sleep(600 * time.Millisecond)
			h.typeText("\033[36m[*] Biological analysis in progress...\033[0m", 30)
			time.Sleep(400 * time.Millisecond)
			h.typeText("\033[33m    Species: Unknown", 30)
			h.typeText("    Threat level: Extremely hostile", 30)
			h.typeText("    Blood type: Molecular acid (pH < 0)\033[0m", 30)
			time.Sleep(500 * time.Millisecond)
			h.typeText("\033[31m[!] Recommendation: Nuke it from orbit\033[0m", 30)
			time.Sleep(300 * time.Millisecond)
			h.typeText("\033[33m    (It's the only way to be sure)\033[0m", 30)
		},
		func() {
			h.showPrompt()
			h.randomPause()
			h.typeText("synthetic_check --android-detection", 50)
			time.Sleep(500 * time.Millisecond)
			h.typeText("\033[36m[*] Running Voight-Kampff test...\033[0m", 30)
			time.Sleep(600 * time.Millisecond)
			h.typeText("\033[32m[+] Subject appears human\033[0m", 30)
			time.Sleep(400 * time.Millisecond)
			h.typeText("\033[33m[!] But wait... milk in the veins?\033[0m", 30)
			time.Sleep(300 * time.Millisecond)
			h.typeText("\033[31m    (You might be a robot!)\033[0m", 30)
		},
		func() {
			h.showPrompt()
			h.randomPause()
			h.typeText("plasma_cannon --shoulder-mounted --targeting", 50)
			time.Sleep(500 * time.Millisecond)
			h.typeText("\033[36m[*] Calibrating tri-beam targeting system...\033[0m", 30)
			time.Sleep(300 * time.Millisecond)
			h.typeText("\033[31m    ◉ ◉ ◉ LOCKED\033[0m", 30)
			time.Sleep(400 * time.Millisecond)
			h.progressBar("Charging plasma", 1500*time.Millisecond)
			h.typeText("\033[32m[+] Weapon ready\033[0m", 30)
			time.Sleep(300 * time.Millisecond)
			h.typeText("\033[33m    (What's the matter? CIA got you pushing too many pencils?)\033[0m", 30)
		},
		func() {
			h.showPrompt()
			h.randomPause()
			h.typeText("judgment_day --postpone --disable-skynet", 50)
			time.Sleep(600 * time.Millisecond)
			h.typeText("\033[36m[*] Accessing military defense network...\033[0m", 30)
			time.Sleep(500 * time.Millisecond)
			h.typeText("\033[33m[!] Attempting to prevent nuclear holocaust...\033[0m", 30)
			h.progressBar("Stopping robot uprising", 2*time.Second)
			time.Sleep(400 * time.Millisecond)
			h.typeText("\033[32m[+] Judgment Day postponed... for now\033[0m", 30)
			time.Sleep(300 * time.Millisecond)
			h.typeText("\033[33m    (No fate but what we make)\033[0m", 30)
		},
		func() {
			h.showPrompt()
			h.randomPause()
			h.typeText("self_destruct --override-code-omega", 50)
			time.Sleep(500 * time.Millisecond)
			h.typeText("\033[31m[!] EMERGENCY: Ship auto-destruct sequence activated\033[0m", 30)
			time.Sleep(400 * time.Millisecond)
			h.typeText("\033[33m[*] T-minus: 10 minutes to detonation\033[0m", 30)
			time.Sleep(600 * time.Millisecond)
			h.typeText("\033[36m[*] Attempting to abort...\033[0m", 30)
			time.Sleep(500 * time.Millisecond)
			h.typeText("\033[31m[!] The ship will not abort destruct sequence\033[0m", 30)
			time.Sleep(400 * time.Millisecond)
			h.typeText("\033[32m[+] Just kidding! This is just a simulation\033[0m", 30)
			h.typeText("\033[33m    (You have 10 minutes to reach minimum safe distance)\033[0m", 30)
		},
		func() {
			h.showPrompt()
			h.randomPause()
			h.typeText("trophy_room --view-skulls --honor-display", 50)
			time.Sleep(500 * time.Millisecond)
			h.typeText("\033[36m[*] Accessing trophy collection database...\033[0m", 30)
			time.Sleep(400 * time.Millisecond)
			skulls := []string{"Xenomorph", "Grid anomaly", "Debugger", "Syntax Error", "Null Pointer"}
			for i := range 3 {
				h.typeText(fmt.Sprintf("\033[33m    Trophy #%d: %s skull\033[0m", i+1, skulls[rand.Intn(len(skulls))]), 30)
				time.Sleep(300 * time.Millisecond)
			}
			h.typeText("\033[32m[+] You are one ugly code base!\033[0m", 30)
		},
		func() {
			h.showPrompt()
			h.randomPause()
			h.typeText("encrypt_files --algorithm=AES-256 --quantum-resistant secret_data/", 50)
			time.Sleep(400 * time.Millisecond)
			h.typeText("\033[36m[*] Initializing quantum encryption...\033[0m", 30)
			time.Sleep(300 * time.Millisecond)
			files := []string{"launch_codes.txt", "passwords.db", "secret_recipe.pdf", "conspiracy.doc"}
			for range 3 {
				file := files[rand.Intn(len(files))]
				h.spinner(fmt.Sprintf("Encrypting %s", file), 1500*time.Millisecond)
			}
			time.Sleep(300 * time.Millisecond)
			h.typeText("\033[32m[+] All files encrypted with military-grade encryption!\033[0m", 30)
			h.typeText("\033[33m    (Even we can't decrypt them now...)\033[0m", 30)
		},
		func() {
			h.showPrompt()
			h.randomPause()
			h.typeText("decrypt_message --key=rosebud --cipher=enigma message.enc", 50)
			time.Sleep(500 * time.Millisecond)
			h.typeText("\033[36m[*] Loading Enigma rotor configuration...\033[0m", 30)
			h.spinner("Decrypting message", 2000*time.Millisecond)
			time.Sleep(400 * time.Millisecond)
			messages := []string{
				"MEET AT THE USUAL PLACE",
				"THE EAGLE HAS LANDED",
				"SQUEAKY CLEAN IS THE KEY",
				"WINTER IS COMING",
				"MAY THE FORCE BE WITH YOU",
			}
			h.typeText(fmt.Sprintf("\033[32m[+] Decrypted: %s\033[0m", messages[rand.Intn(len(messages))]), 30)
		},
		func() {
			h.showPrompt()
			h.randomPause()
			h.typeText("compress_data --ultra --ratio=99 ./huge_database/*", 50)
			time.Sleep(400 * time.Millisecond)
			h.typeText("\033[36m[*] Compressing 10TB of data...\033[0m", 30)
			time.Sleep(300 * time.Millisecond)
			h.progressBar("Compression progress", 2500*time.Millisecond)
			originalSize := 10000 + rand.Intn(5000)
			compressedSize := 50 + rand.Intn(50)
			ratio := float64(originalSize) / float64(compressedSize)
			time.Sleep(300 * time.Millisecond)
			h.typeText(fmt.Sprintf("\033[32m[+] Compressed %d GB → %d MB (%.1fx compression!)\033[0m", originalSize, compressedSize, ratio), 30)
			h.typeText("\033[33m    (ZIP couldn't even dream of this ratio)\033[0m", 30)
		},
		func() {
			h.showPrompt()
			h.randomPause()
			h.typeText("decompress_archive --extract=all stolen_data.tar.gz.bz2.xz.7z", 50)
			time.Sleep(400 * time.Millisecond)
			h.typeText("\033[36m[*] Decompressing multi-layer archive...\033[0m", 30)
			layers := []string{"7z layer", "xz layer", "bz2 layer", "gzip layer", "tar layer"}
			for i, layer := range layers {
				h.spinner(fmt.Sprintf("Extracting %s (%d/5)", layer, i+1), 1000*time.Millisecond)
			}
			time.Sleep(300 * time.Millisecond)
			h.typeText("\033[32m[+] 1,337 files extracted successfully!\033[0m", 30)
			h.typeText("\033[33m    (Someone really didn't want us to get this)\033[0m", 30)
		},
		func() {
			h.showPrompt()
			h.randomPause()
			h.typeText("hash_crack --wordlist=rockyou.txt --hash=5f4dcc3b5aa765d61d8327deb882cf99", 50)
			time.Sleep(500 * time.Millisecond)
			h.typeText("\033[36m[*] Loading 14 billion passwords...\033[0m", 30)
			time.Sleep(400 * time.Millisecond)
			h.spinner("Cracking hash", 2500*time.Millisecond)
			attempts := 1000000 + rand.Intn(9000000)
			time.Sleep(300 * time.Millisecond)
			h.typeText(fmt.Sprintf("\033[32m[+] Hash cracked after %d attempts!\033[0m", attempts), 30)
			time.Sleep(200 * time.Millisecond)
			passwords := []string{"password", "123456", "qwerty", "letmein", "admin"}
			h.typeText(fmt.Sprintf("\033[33m    Original password: %s\033[0m", passwords[rand.Intn(len(passwords))]), 30)
			h.typeText("\033[33m    (Seriously? That was the password?)\033[0m", 30)
		},
		func() {
			h.showPrompt()
			h.randomPause()
			h.typeText("steganography_hide --image=cat.jpg --payload=secrets.txt", 50)
			time.Sleep(500 * time.Millisecond)
			h.typeText("\033[36m[*] Analyzing image pixels...\033[0m", 30)
			time.Sleep(300 * time.Millisecond)
			h.spinner("Embedding data in LSB", 1800*time.Millisecond)
			time.Sleep(300 * time.Millisecond)
			h.typeText("\033[32m[+] Data successfully hidden in image!\033[0m", 30)
			h.typeText("\033[33m    (Now it's just a picture of a cat... or is it?)\033[0m", 30)
		},
		func() {
			h.showPrompt()
			h.randomPause()
			h.typeText("steganography_extract --image=suspicious_meme.png", 50)
			time.Sleep(500 * time.Millisecond)
			h.typeText("\033[36m[*] Scanning image for hidden data...\033[0m", 30)
			h.spinner("Analyzing pixel patterns", 2000*time.Millisecond)
			time.Sleep(400 * time.Millisecond)
			h.typeText("\033[32m[+] Hidden message found!\033[0m", 30)
			time.Sleep(300 * time.Millisecond)
			h.typeText("\033[33m    Message: \"epstein didn't kill himself\"\033[0m", 30)
		},
		func() {
			h.showPrompt()
			h.randomPause()
			h.typeText("shred_evidence --passes=35 --method=gutmann incriminating_files/*", 50)
			time.Sleep(500 * time.Millisecond)
			h.typeText("\033[36m[*] Initializing secure deletion (35 passes)...\033[0m", 30)
			time.Sleep(300 * time.Millisecond)
			h.progressBar("Overwriting with random data", 2500*time.Millisecond)
			time.Sleep(300 * time.Millisecond)
			h.typeText("\033[32m[+] Files permanently deleted!\033[0m", 30)
			h.typeText("\033[33m    (Not even the FBI can recover these now)\033[0m", 30)
		},
		func() {
			h.showPrompt()
			h.randomPause()
			h.typeText("generate_keys --type=RSA-4096 --entropy=paranoid", 50)
			time.Sleep(500 * time.Millisecond)
			h.typeText("\033[36m[*] Gathering entropy from cosmic background radiation...\033[0m", 30)
			h.spinner("Generating prime numbers", 2000*time.Millisecond)
			time.Sleep(400 * time.Millisecond)
			h.typeText("\033[32m[+] Key pair generated successfully!\033[0m", 30)
			h.typeText("\033[33m    Public key: AAAAB3NzaC1yc2EAAAADAQABAAACAQ...\033[0m", 30)
			h.typeText("\033[33m    (Would take the sun's lifetime to crack)\033[0m", 30)
		},
		func() {
			h.showPrompt()
			h.randomPause()
			h.typeText("obfuscate_code --level=maximum --anti-debug payload.exe", 50)
			time.Sleep(500 * time.Millisecond)
			h.typeText("\033[36m[*] Applying code obfuscation...\033[0m", 30)
			techniques := []string{
				"Control flow flattening",
				"String encryption",
				"Dead code injection",
				"Opaque predicates",
			}
			for _, technique := range techniques {
				h.spinner(technique, 1000*time.Millisecond)
			}
			time.Sleep(300 * time.Millisecond)
			h.typeText("\033[32m[+] Code successfully obfuscated!\033[0m", 30)
			h.typeText("\033[33m    (Good luck reverse engineering this mess)\033[0m", 30)
		},
		func() {
			h.showPrompt()
			h.randomPause()
			h.typeText("backup_system --destination=offshore --encrypt --split", 50)
			time.Sleep(500 * time.Millisecond)
			h.typeText("\033[36m[*] Creating distributed encrypted backup...\033[0m", 30)
			time.Sleep(300 * time.Millisecond)
			locations := []string{"Iceland", "Switzerland", "Cayman Islands", "Singapore"}
			for i, location := range locations {
				h.spinner(fmt.Sprintf("Uploading shard %d/4 to %s", i+1, location), 1200*time.Millisecond)
			}
			time.Sleep(300 * time.Millisecond)
			h.typeText("\033[32m[+] Backup distributed across 4 continents!\033[0m", 30)
			h.typeText("\033[33m    (They'd need to raid 4 countries simultaneously)\033[0m", 30)
		},
		func() {
			h.showPrompt()
			h.randomPause()
			h.typeText("sql_inject --url="+h.target+"/login --payload='OR 1=1--", 50)
			time.Sleep(500 * time.Millisecond)
			h.typeText("\033[36m[*] Testing for SQL injection vulnerabilities...\033[0m", 30)
			time.Sleep(400 * time.Millisecond)
			h.spinner("Probing database", 1800*time.Millisecond)
			time.Sleep(300 * time.Millisecond)
			h.typeText("\033[31m[!] Vulnerable endpoint detected!\033[0m", 30)
			time.Sleep(400 * time.Millisecond)
			h.typeText("\033[32m[+] Dumping user credentials...\033[0m", 30)
			users := []string{"admin", "root", "sysadmin", "dbadmin", "test_user"}
			for range 3 {
				h.typeText(fmt.Sprintf("\033[33m    %s : P@ssw0rd%d\033[0m", users[rand.Intn(len(users))], rand.Intn(999)), 30)
				time.Sleep(250 * time.Millisecond)
			}
			h.typeText("\033[33m    (Bobby Tables would be proud)\033[0m", 30)
		},
		func() {
			h.showPrompt()
			h.randomPause()
			h.typeText("wifi_crack --interface=wlan0 --target=FBI_Surveillance_Van_7", 50)
			time.Sleep(500 * time.Millisecond)
			h.typeText("\033[36m[*] Capturing WPA handshake...\033[0m", 30)
			time.Sleep(300 * time.Millisecond)
			h.progressBar("Collecting packets", 2000*time.Millisecond)
			time.Sleep(400 * time.Millisecond)
			h.typeText("\033[32m[+] Handshake captured!\033[0m", 30)
			time.Sleep(300 * time.Millisecond)
			h.spinner("Running dictionary attack", 2500*time.Millisecond)
			time.Sleep(400 * time.Millisecond)
			h.typeText("\033[32m[+] WiFi password cracked: \"FBIAgent123\"\033[0m", 30)
			h.typeText("\033[33m    (They're not even trying to hide it anymore)\033[0m", 30)
		},
		func() {
			h.showPrompt()
			h.randomPause()
			h.typeText("ransomware_decrypt --key=universal --no-payment ransomed_files/", 50)
			time.Sleep(500 * time.Millisecond)
			h.typeText("\033[36m[*] Analyzing ransomware encryption...\033[0m", 30)
			time.Sleep(400 * time.Millisecond)
			variants := []string{"WannaCry", "Petya", "Ryuk", "REvil", "DarkSide"}
			h.typeText(fmt.Sprintf("\033[33m[*] Detected variant: %s\033[0m", variants[rand.Intn(len(variants))]), 30)
			time.Sleep(300 * time.Millisecond)
			h.spinner("Exploiting implementation flaw", 2200*time.Millisecond)
			time.Sleep(400 * time.Millisecond)
			h.typeText("\033[32m[+] All files decrypted without paying!\033[0m", 30)
			h.typeText("\033[33m    (Take that, cyber criminals!)\033[0m", 30)
		},
		func() {
			h.showPrompt()
			h.randomPause()
			h.typeText("blockchain_hack --target=bitcoin --double-spend --51-percent-attack", 50)
			time.Sleep(500 * time.Millisecond)
			h.typeText("\033[36m[*] Attempting 51% attack on blockchain...\033[0m", 30)
			time.Sleep(300 * time.Millisecond)
			h.spinner("Mining competing chain", 3000*time.Millisecond)
			time.Sleep(500 * time.Millisecond)
			h.typeText("\033[31m[!] ERROR: Would need more computing power than exists on Earth\033[0m", 30)
			time.Sleep(400 * time.Millisecond)
			h.typeText("\033[33m    (Satoshi wins this round)\033[0m", 30)
		},
		func() {
			h.showPrompt()
			h.randomPause()
			h.typeText("social_engineer --target=ceo@"+h.target+" --pretend=IT-support", 50)
			time.Sleep(500 * time.Millisecond)
			h.typeText("\033[36m[*] Crafting convincing phishing email...\033[0m", 30)
			time.Sleep(400 * time.Millisecond)
			h.spinner("Cloning corporate email template", 1800*time.Millisecond)
			time.Sleep(300 * time.Millisecond)
			h.typeText("\033[32m[+] Email sent successfully!\033[0m", 30)
			time.Sleep(500 * time.Millisecond)
			h.typeText("\033[36m[*] Waiting for victim to click link...\033[0m", 30)
			time.Sleep(1000 * time.Millisecond)
			h.typeText("\033[32m[+] CEO clicked the link! Credentials harvested!\033[0m", 30)
			h.typeText("\033[33m    (The human is always the weakest link)\033[0m", 30)
		},
		func() {
			h.showPrompt()
			h.randomPause()
			h.typeText("zero_day_exploit --vulnerability=CVE-2077-1337 --target="+h.target, 50)
			time.Sleep(500 * time.Millisecond)
			h.typeText("\033[36m[*] Loading zero-day exploit from dark web...\033[0m", 30)
			time.Sleep(400 * time.Millisecond)
			h.typeText("\033[33m[*] Exploit cost: 50 Bitcoin\033[0m", 30)
			time.Sleep(300 * time.Millisecond)
			h.spinner("Deploying exploit", 2000*time.Millisecond)
			time.Sleep(400 * time.Millisecond)
			h.typeText("\033[32m[+] Remote code execution achieved!\033[0m", 30)
			h.typeText("\033[33m    (Worth every satoshi)\033[0m", 30)
		},
		func() {
			h.showPrompt()
			h.randomPause()
			h.typeText("botnet_control --zombies=10000 --ddos-target="+h.target+" --intensity=maximum", 50)
			time.Sleep(500 * time.Millisecond)
			h.typeText("\033[36m[*] Commanding botnet to attack target...\033[0m", 30)
			time.Sleep(300 * time.Millisecond)
			h.progressBar("Coordinating zombies", 2500*time.Millisecond)
			time.Sleep(400 * time.Millisecond)
			h.typeText("\033[32m[+] 10,000 zombies attacking target!\033[0m", 30)
			time.Sleep(300 * time.Millisecond)
			h.typeText("\033[36m[*] Traffic: 500 Gbps and rising...\033[0m", 30)
			time.Sleep(400 * time.Millisecond)
			h.typeText("\033[31m[!] Target servers melting!\033[0m", 30)
			h.typeText("\033[33m    (Their ops team is having a bad day)\033[0m", 30)
		},
		func() {
			h.showPrompt()
			h.randomPause()
			h.typeText("deepfake_generate --target=ceo --say='Transfer funds to account 1337'", 50)
			time.Sleep(500 * time.Millisecond)
			h.typeText("\033[36m[*] Analyzing target's voice patterns...\033[0m", 30)
			time.Sleep(400 * time.Millisecond)
			h.spinner("Training neural network", 2500*time.Millisecond)
			time.Sleep(300 * time.Millisecond)
			h.typeText("\033[32m[+] Deepfake audio generated!\033[0m", 30)
			time.Sleep(400 * time.Millisecond)
			h.typeText("\033[36m[*] Calling finance department...\033[0m", 30)
			time.Sleep(600 * time.Millisecond)
			h.typeText("\033[32m[+] They fell for it! $1M transferred!\033[0m", 30)
			h.typeText("\033[33m    (AI is getting scary good at this)\033[0m", 30)
		},
		func() {
			h.showPrompt()
			h.randomPause()
			h.typeText("keylogger_install --stealth --target=all-workstations --exfiltrate-to=dark-server", 50)
			time.Sleep(500 * time.Millisecond)
			h.typeText("\033[36m[*] Deploying invisible keyloggers...\033[0m", 30)
			time.Sleep(300 * time.Millisecond)
			workstations := 15 + rand.Intn(35)
			h.progressBar(fmt.Sprintf("Installing on %d workstations", workstations), 2200*time.Millisecond)
			time.Sleep(400 * time.Millisecond)
			h.typeText("\033[32m[+] Keyloggers active on all targets!\033[0m", 30)
			time.Sleep(300 * time.Millisecond)
			h.typeText("\033[36m[*] Capturing keystrokes in real-time...\033[0m", 30)
			time.Sleep(500 * time.Millisecond)
			h.typeText("\033[33m    (Every password is belong to us)\033[0m", 30)
		},
		func() {
			h.showPrompt()
			h.randomPause()
			h.typeText("packet_sniff --interface=eth0 --filter='password|secret|confidential'", 50)
			time.Sleep(500 * time.Millisecond)
			h.typeText("\033[36m[*] Entering promiscuous mode...\033[0m", 30)
			time.Sleep(300 * time.Millisecond)
			h.typeText("\033[36m[*] Sniffing network traffic...\033[0m", 30)
			time.Sleep(400 * time.Millisecond)
			packets := 5000 + rand.Intn(5000)
			h.typeText(fmt.Sprintf("\033[33m[*] Captured %d packets\033[0m", packets), 30)
			time.Sleep(300 * time.Millisecond)
			h.spinner("Analyzing for credentials", 2000*time.Millisecond)
			time.Sleep(400 * time.Millisecond)
			h.typeText("\033[32m[+] Found 17 passwords in cleartext!\033[0m", 30)
			h.typeText("\033[33m    (Who uses HTTP in 2025?!)\033[0m", 30)
		},
		func() {
			h.showPrompt()
			h.randomPause()
			h.typeText("exfiltrate_data --source=/home/victim/Documents --method=dns-tunneling", 50)
			time.Sleep(500 * time.Millisecond)
			h.typeText("\033[36m[*] Encoding data into DNS queries...\033[0m", 30)
			time.Sleep(400 * time.Millisecond)
			h.progressBar("Tunneling through DNS", 2800*time.Millisecond)
			time.Sleep(300 * time.Millisecond)
			dataSize := 100 + rand.Intn(900)
			h.typeText(fmt.Sprintf("\033[32m[+] Exfiltrated %d MB of sensitive data!\033[0m", dataSize), 30)
			h.typeText("\033[33m    (Their firewall didn't even notice)\033[0m", 30)
		},
		func() {
			h.showPrompt()
			h.randomPause()
			h.typeText("pivot_network --from=dmz --to=internal --escalate-privileges", 50)
			time.Sleep(500 * time.Millisecond)
			h.typeText("\033[36m[*] Established foothold in DMZ...\033[0m", 30)
			time.Sleep(400 * time.Millisecond)
			h.spinner("Scanning for pivot points", 1800*time.Millisecond)
			time.Sleep(300 * time.Millisecond)
			h.typeText("\033[32m[+] Found misconfigured jump box!\033[0m", 30)
			time.Sleep(400 * time.Millisecond)
			h.spinner("Pivoting to internal network", 2000*time.Millisecond)
			time.Sleep(300 * time.Millisecond)
			h.typeText("\033[32m[+] Access to internal network achieved!\033[0m", 30)
			h.typeText("\033[33m    (We're in the castle now)\033[0m", 30)
		},
		func() {
			h.showPrompt()
			h.randomPause()
			h.typeText("ai_password_predict --username=admin --context-aware --ml-model=gpt-4", 50)
			time.Sleep(500 * time.Millisecond)
			h.typeText("\033[36m[*] Analyzing target's digital footprint...\033[0m", 30)
			time.Sleep(400 * time.Millisecond)
			h.spinner("AI predicting password", 2500*time.Millisecond)
			time.Sleep(300 * time.Millisecond)
			predictions := []string{"Fluffy2023!", "Summer@Beach", "MyK1ds123", "Tr0ub4dor&3"}
			h.typeText("\033[32m[+] Top password predictions:\033[0m", 30)
			for i, pred := range predictions[:3] {
				h.typeText(fmt.Sprintf("\033[33m    %d. %s\033[0m", i+1, pred), 30)
				time.Sleep(200 * time.Millisecond)
			}
			h.typeText("\033[33m    (AI knows you better than you know yourself)\033[0m", 30)
		},
		func() {
			h.showPrompt()
			h.randomPause()
			h.typeText("supply_chain_compromise --target=popular-npm-package --backdoor=true", 50)
			time.Sleep(500 * time.Millisecond)
			h.typeText("\033[36m[*] Compromising developer's account...\033[0m", 30)
			time.Sleep(400 * time.Millisecond)
			h.spinner("Injecting malicious code", 2000*time.Millisecond)
			time.Sleep(300 * time.Millisecond)
			h.typeText("\033[32m[+] Backdoored package published!\033[0m", 30)
			time.Sleep(400 * time.Millisecond)
			downloads := 10000 + rand.Intn(90000)
			h.typeText(fmt.Sprintf("\033[36m[*] Package downloaded %d times today...\033[0m", downloads), 30)
			time.Sleep(300 * time.Millisecond)
			h.typeText("\033[33m    (That's a lot of compromised systems)\033[0m", 30)
		},
		func() {
			h.showPrompt()
			h.randomPause()
			h.typeText("quantum_crypto_break --algorithm=RSA-2048 --qubits=1024", 50)
			time.Sleep(500 * time.Millisecond)
			h.typeText("\033[36m[*] Initializing quantum computer...\033[0m", 30)
			time.Sleep(400 * time.Millisecond)
			h.spinner("Running Shor's algorithm", 3000*time.Millisecond)
			time.Sleep(500 * time.Millisecond)
			h.typeText("\033[32m[+] RSA-2048 factored in 0.3 seconds!\033[0m", 30)
			time.Sleep(300 * time.Millisecond)
			h.typeText("\033[31m[!] All traditional encryption is now obsolete!\033[0m", 30)
			h.typeText("\033[33m    (The crypto apocalypse is here)\033[0m", 30)
		},
		func() {
			h.showPrompt()
			h.randomPause()
			h.typeText("privilege_escalate --exploit=dirty-cow --target-user=root", 50)
			time.Sleep(500 * time.Millisecond)
			h.typeText("\033[36m[*] Exploiting kernel vulnerability...\033[0m", 30)
			time.Sleep(400 * time.Millisecond)
			h.spinner("Escalating privileges", 2000*time.Millisecond)
			time.Sleep(300 * time.Millisecond)
			h.typeText("\033[32m[+] Root access obtained!\033[0m", 30)
			time.Sleep(200 * time.Millisecond)
			h.typeText("\033[33m    uid=0(root) gid=0(root) groups=0(root)\033[0m", 30)
			time.Sleep(300 * time.Millisecond)
			h.typeText("\033[33m    (I am become root, destroyer of worlds)\033[0m", 30)
		},
		func() {
			h.showPrompt()
			h.randomPause()
			h.typeText("memory_dump --process=chrome --search='cookies|session|tokens'", 50)
			time.Sleep(500 * time.Millisecond)
			h.typeText("\033[36m[*] Attaching to Chrome process...\033[0m", 30)
			time.Sleep(300 * time.Millisecond)
			h.progressBar("Dumping memory", 2200*time.Millisecond)
			time.Sleep(400 * time.Millisecond)
			h.typeText("\033[32m[+] Memory dump complete: 2.3 GB\033[0m", 30)
			time.Sleep(300 * time.Millisecond)
			h.spinner("Extracting secrets", 1800*time.Millisecond)
			time.Sleep(300 * time.Millisecond)
			h.typeText("\033[32m[+] Recovered 47 session tokens!\033[0m", 30)
			h.typeText("\033[33m    (Your browser knows all your secrets)\033[0m", 30)
		},
		func() {
			h.showPrompt()
			h.randomPause()
			h.typeText("dns_poisoning --target-domain="+h.target+" --redirect=evil-server.onion", 50)
			time.Sleep(500 * time.Millisecond)
			h.typeText("\033[36m[*] Intercepting DNS queries...\033[0m", 30)
			time.Sleep(400 * time.Millisecond)
			h.spinner("Poisoning DNS cache", 2000*time.Millisecond)
			time.Sleep(300 * time.Millisecond)
			h.typeText("\033[32m[+] DNS cache poisoned successfully!\033[0m", 30)
			time.Sleep(400 * time.Millisecond)
			victims := 100 + rand.Intn(900)
			h.typeText(fmt.Sprintf("\033[36m[*] %d users redirected to fake site!\033[0m", victims), 30)
			h.typeText("\033[33m    (They'll never know they're on a fake page)\033[0m", 30)
		},
		func() {
			h.showPrompt()
			h.randomPause()
			h.typeText("reverse_shell --port=4444 --callback=attacker-server.com --obfuscate", 50)
			time.Sleep(500 * time.Millisecond)
			h.typeText("\033[36m[*] Establishing reverse shell connection...\033[0m", 30)
			time.Sleep(400 * time.Millisecond)
			h.spinner("Bypassing firewall", 2200*time.Millisecond)
			time.Sleep(300 * time.Millisecond)
			h.typeText("\033[32m[+] Shell connected!\033[0m", 30)
			time.Sleep(300 * time.Millisecond)
			h.typeText("\033[33m    bash-5.1$ whoami\033[0m", 30)
			time.Sleep(200 * time.Millisecond)
			h.typeText("\033[33m    root\033[0m", 30)
			time.Sleep(300 * time.Millisecond)
			h.typeText("\033[33m    (Sweet, sweet shell access)\033[0m", 30)
		},
		func() {
			h.showPrompt()
			h.randomPause()
			h.typeText("forensics_wipe --anti-forensic --secure-delete --cover-tracks", 50)
			time.Sleep(500 * time.Millisecond)
			h.typeText("\033[36m[*] Initiating anti-forensic procedures...\033[0m", 30)
			time.Sleep(300 * time.Millisecond)
			tasks := []string{"Clearing logs", "Wiping bash history", "Removing timestamps", "Overwriting slack space"}
			for _, task := range tasks {
				h.spinner(task, 1000*time.Millisecond)
			}
			time.Sleep(300 * time.Millisecond)
			h.typeText("\033[32m[+] All traces eliminated!\033[0m", 30)
			h.typeText("\033[33m    (It's like we were never here)\033[0m", 30)
		},
	}

	// Run a random sequence
	sequence := sequences[rand.Intn(len(sequences))]
	sequence()

	time.Sleep(1 * time.Second)
	printSeparator()
	time.Sleep(500 * time.Millisecond)
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
	bannerWidth := 65 // Width of the banner box

	lines := strings.SplitSeq(banner, "\n")
	for line := range lines {
		if line == "" {
			fmt.Println()
			continue
		}
		padding := max((termWidth-bannerWidth)/2, 0)
		fmt.Print(strings.Repeat(" ", padding))
		fmt.Println(line)
	}

	time.Sleep(1 * time.Second)
}

func visibleLength(s string) int {
	clean := ansi.ReplaceAllString(s, "")
	return utf8.RuneCountInString(clean)
}
