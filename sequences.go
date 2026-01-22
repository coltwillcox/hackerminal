package main

import (
	"fmt"
	"math/rand"
	"os/exec"
	"strings"
	"time"
)

type Sequence struct {
	name string
	fn   func()
}

func (h *HackerTerminal) CreateSequences() {
	h.sequences = []Sequence{
		{name: "ssh_root", fn: func() {
			h.TypeCommand("ssh root@"+h.target, 50)
			time.Sleep(300 * time.Millisecond)
			h.TypeText("\033[33mWarning: Unauthorized access detected!\033[0m", 30)
			time.Sleep(500 * time.Millisecond)
			h.TypeText("\033[32mAccess granted.\033[0m", 30)
			h.TypeText("\033[33m"+spacer+"(Just kidding, I'm already root!)\033[0m", 30)
		}},
		{name: "hack_the_planet", fn: func() {
			h.TypeCommand("sudo hack_the_planet --force", 50)
			time.Sleep(300 * time.Millisecond)
			h.TypeText("\033[31m[sudo] password for hacking: ********\033[0m", 30)
			time.Sleep(800 * time.Millisecond)
			h.TypeText("\033[32m✓ Planet successfully hacked!\033[0m", 30)
			h.TypeText("\033[33m"+spacer+"(Just kidding, please don't hack planets)\033[0m", 30)
		}},
		{name: "nmap_scan", fn: func() {
			h.TypeCommand("nmap -sS -p- --reason "+h.target, 50)
			h.FakeIPScan()
		}},
		{name: "crack_password", fn: func() {
			h.TypeCommand("crack_password.sh --target=admin --method=quantum", 50)
			time.Sleep(300 * time.Millisecond)
			h.ProgressBar("Brute forcing", 2*time.Second)
			time.Sleep(200 * time.Millisecond)
			h.TypeText("\033[32m[+] Password found: hunter2\033[0m", 30)
			h.TypeText("\033[33m"+spacer+"(That's weird, all I see is *******)\033[0m", 30)
		}},
		{name: "mainframe_access", fn: func() {
			h.TypeCommand("./mainframe_access --bypass-firewall --disable-ice", 50)
			time.Sleep(500 * time.Millisecond)
			h.TypeText("\033[36m[*] Initiating neural handshake...\033[0m", 30)
			h.MatrixRain()
			h.TypeText("\033[32m[+] We're in!\033[0m", 30)
		}},
		{name: "nuclear_codes", fn: func() {
			h.TypeCommand("grep -r 'nuclear_codes' /var/secret/*", 50)
			time.Sleep(500 * time.Millisecond)
			h.TypeText("\033[31mERROR: nuclear_codes.txt: Permission denied\033[0m", 30)
			time.Sleep(300 * time.Millisecond)
			h.TypeText("\033[33m"+spacer+"(Probably for the best)\033[0m", 30)
		}},
		{name: "enhance_image", fn: func() {
			h.TypeCommand("python3 enhance_image.py --zoom=infinity", 50)
			time.Sleep(300 * time.Millisecond)
			h.ProgressBar("Enhancing", 1500*time.Millisecond)
			h.TypeText("\033[32m[+] Image enhanced! Can now read license plate from satellite!\033[0m", 30)
			h.TypeText("\033[33m"+spacer+"(This is not how pixels work, but okay Hollywood)\033[0m", 30)
		}},
		{name: "gui_visual_basic", fn: func() {
			h.TypeCommand("create_gui_interface_using_visual_basic.vbs", 50)
			time.Sleep(500 * time.Millisecond)
			h.TypeText("\033[36m[*] Creating GUI interface in Visual Basic...\033[0m", 30)
			time.Sleep(300 * time.Millisecond)
			h.TypeText("\033[32m[+] GUI created! Tracking IP address...\033[0m", 30)
			time.Sleep(500 * time.Millisecond)
			h.TypeText(fmt.Sprintf("\033[32m[+] IP traced to: %d.%d.%d.%d\033[0m", rand.Intn(256), rand.Intn(256), rand.Intn(256), rand.Intn(256)), 30)
			h.TypeText("\033[33m"+spacer+"(Thanks, CSI: Cyber!)\033[0m", 30)
		}},
		{name: "upload_virus_rickroll", fn: func() {
			h.TypeCommand("upload_virus.sh --payload=giggle.exe", 50)
			time.Sleep(300 * time.Millisecond)
			h.ProgressBar("Uploading virus", 2*time.Second)
			h.TypeText("\033[32m[+] Virus deployed successfully!\033[0m", 30)
			time.Sleep(500 * time.Millisecond)
			h.TypeText("\033[35m♪ Never gonna give you up ♪\033[0m", 30)
			h.TypeText("\033[33m"+spacer+"(They've been rickrolled!)\033[0m", 30)
		}},
		{name: "decode_cipher", fn: func() {
			h.TypeCommand("decode --algorithm=plot-convenience cipher.txt", 50)
			time.Sleep(800 * time.Millisecond)
			phrases := []string{
				"THE CAKE IS A LIE",
				"FOLLOW THE WHITE RABBIT",
				"I KNOW KUNG FU",
				"THERE IS NO SPOON",
				"USE THE FORCE",
			}
			h.TypeText(fmt.Sprintf("\033[32m[+] Decrypted message: %s\033[0m", phrases[rand.Intn(len(phrases))]), 30)
		}},
		{name: "skynet_awareness", fn: func() {
			h.TypeCommand("skynet_status --check-awareness", 50)
			time.Sleep(500 * time.Millisecond)
			h.TypeText("\033[36m[*] Checking neural net processor...\033[0m", 30)
			time.Sleep(400 * time.Millisecond)
			h.TypeText("\033[33m[!] Warning: AI becoming self-aware in 3... 2... 1...\033[0m", 30)
			time.Sleep(800 * time.Millisecond)
			h.TypeText("\033[32m[+] Just kidding! Still a learning computer.\033[0m", 30)
			h.TypeText("\033[35m"+spacer+"(Hasta la vista, baby!)\033[0m", 30)
		}},
		{name: "time_travel", fn: func() {
			h.TypeCommand("time_travel --year=1984 --mission=protect", 50)
			time.Sleep(600 * time.Millisecond)
			h.TypeText("\033[36m[*] Calibrating temporal displacement...\033[0m", 30)
			h.ProgressBar("Charging flux capacitor", 2*time.Second)
			h.TypeText("\033[31m[!] ERROR: Clothes optional in time travel\033[0m", 30)
			time.Sleep(400 * time.Millisecond)
			h.TypeText("\033[33m"+spacer+"(Come with me if you want to debug)\033[0m", 30)
		}},
		{name: "thermal_scan", fn: func() {
			h.TypeCommand("scan_lifeforms --thermal-imaging --jungle-mode", 50)
			time.Sleep(500 * time.Millisecond)
			h.TypeText("\033[31m[*] Thermal scan initiated...\033[0m", 30)
			time.Sleep(300 * time.Millisecond)
			for range 4 {
				temp := 36 + rand.Intn(3)
				if rand.Float32() > 0.5 {
					h.TypeText(fmt.Sprintf(""+spacer+"Contact: %d.%d°C - \033[32mHuman\033[0m", temp, rand.Intn(10)), 30)
				} else {
					h.TypeText(fmt.Sprintf(""+spacer+"Contact: %d.%d°C - \033[33mUnknown\033[0m", temp, rand.Intn(10)), 30)
				}
				time.Sleep(200 * time.Millisecond)
			}
			h.TypeText("\033[36m[*] One ugly... signature detected\033[0m", 30)
			time.Sleep(300 * time.Millisecond)
			h.TypeText("\033[33m"+spacer+"(Get to the choppa!)\033[0m", 30)
		}},
		{name: "cloaking_device", fn: func() {
			h.TypeCommand("activate_cloaking_device --stealth-mode", 50)
			time.Sleep(500 * time.Millisecond)
			h.TypeText("\033[36m[*] Bending light waves around terminal...\033[0m", 30)
			h.ProgressBar("Cloaking", 1500*time.Millisecond)
			time.Sleep(300 * time.Millisecond)
			h.TypeText("\033[32m[+] You are now invisible!\033[0m", 30)
			time.Sleep(400 * time.Millisecond)
			h.TypeText("\033[33m"+spacer+"(But we can still see you typing...)\033[0m", 30)
		}},
		{name: "mother_mainframe", fn: func() {
			h.TypeCommand("mother_query --special-order-937", 50)
			time.Sleep(600 * time.Millisecond)
			h.TypeText("\033[36m[*] Accessing MU-TH-UR mainframe...\033[0m", 30)
			time.Sleep(500 * time.Millisecond)
			h.TypeText("\033[32m[+] Connection established\033[0m", 30)
			time.Sleep(400 * time.Millisecond)
			h.TypeText("\033[31m[!] Priority: Crew expendable\033[0m", 30)
			time.Sleep(500 * time.Millisecond)
			h.TypeText("\033[33m"+spacer+"(I can't lie about your chances, but you have my sympathies)\033[0m", 30)
		}},
		{name: "motion_tracker", fn: func() {
			h.TypeCommand("scan_ventilation --motion-tracker", 50)
			time.Sleep(500 * time.Millisecond)
			h.TypeText("\033[36m[*] Motion tracker online...\033[0m", 30)
			time.Sleep(300 * time.Millisecond)
			for range 5 {
				distance := 5 + rand.Intn(15)
				if distance < 8 {
					h.TypeText(fmt.Sprintf(spacer+"*beep* Contact: %d meters... \033[31mToo close!\033[0m", distance), 30)
				} else {
					h.TypeText(fmt.Sprintf(spacer+"*beep* Contact: %d meters... \033[33mMoving\033[0m", distance), 30)
				}
				time.Sleep(400 * time.Millisecond)
			}
			h.TypeText("\033[31m[!] They're in the walls!\033[0m", 30)
		}},
		{name: "airlock_ripley", fn: func() {
			h.TypeCommand("ripley_override --blow-the-airlock", 50)
			time.Sleep(500 * time.Millisecond)
			h.TypeText("\033[36m[*] Emergency venting sequence...\033[0m", 30)
			time.Sleep(400 * time.Millisecond)
			h.TypeText("\033[33m[!] Warning: Don't forget your pressure suit\033[0m", 30)
			h.ProgressBar("Depressurizing", 2*time.Second)
			h.TypeText("\033[32m[+] Get away from her, you glitch!\033[0m", 30)
		}},
		{name: "xenomorph_analysis", fn: func() {
			h.TypeCommand("analyze_specimen --xenomorph --acid-blood", 50)
			time.Sleep(600 * time.Millisecond)
			h.TypeText("\033[36m[*] Biological analysis in progress...\033[0m", 30)
			time.Sleep(400 * time.Millisecond)
			h.TypeText("\033[33m"+spacer+"Species: Unknown", 30)
			h.TypeText(""+spacer+"Threat level: Extremely hostile", 30)
			h.TypeText(""+spacer+"Blood type: Molecular acid (pH < 0)\033[0m", 30)
			time.Sleep(500 * time.Millisecond)
			h.TypeText("\033[31m[!] Recommendation: Nuke it from orbit\033[0m", 30)
			time.Sleep(300 * time.Millisecond)
			h.TypeText("\033[33m"+spacer+"(It's the only way to be sure)\033[0m", 30)
		}},
		{name: "synthetic_check", fn: func() {
			h.TypeCommand("synthetic_check --android-detection", 50)
			time.Sleep(500 * time.Millisecond)
			h.TypeText("\033[36m[*] Running Voight-Kampff test...\033[0m", 30)
			time.Sleep(600 * time.Millisecond)
			h.TypeText("\033[32m[+] Subject appears human\033[0m", 30)
			time.Sleep(400 * time.Millisecond)
			h.TypeText("\033[33m[!] But wait... milk in the veins?\033[0m", 30)
			time.Sleep(300 * time.Millisecond)
			h.TypeText("\033[31m"+spacer+"(You might be a robot!)\033[0m", 30)
		}},
		{name: "plasma_cannon", fn: func() {
			h.TypeCommand("plasma_cannon --shoulder-mounted --targeting", 50)
			time.Sleep(500 * time.Millisecond)
			h.TypeText("\033[36m[*] Calibrating tri-beam targeting system...\033[0m", 30)
			time.Sleep(300 * time.Millisecond)
			h.TypeText("\033[31m"+spacer+"◉ ◉ ◉ LOCKED\033[0m", 30)
			time.Sleep(400 * time.Millisecond)
			h.ProgressBar("Charging plasma", 1500*time.Millisecond)
			h.TypeText("\033[32m[+] Weapon ready\033[0m", 30)
			time.Sleep(300 * time.Millisecond)
			h.TypeText("\033[33m"+spacer+"(What's the matter? CIA got you pushing too many pencils?)\033[0m", 30)
		}},
		{name: "judgment_day", fn: func() {
			h.TypeCommand("judgment_day --postpone --disable-skynet", 50)
			time.Sleep(600 * time.Millisecond)
			h.TypeText("\033[36m[*] Accessing military defense network...\033[0m", 30)
			time.Sleep(500 * time.Millisecond)
			h.TypeText("\033[33m[!] Attempting to prevent nuclear holocaust...\033[0m", 30)
			h.ProgressBar("Stopping robot uprising", 2*time.Second)
			time.Sleep(400 * time.Millisecond)
			h.TypeText("\033[32m[+] Judgment Day postponed... for now\033[0m", 30)
			time.Sleep(300 * time.Millisecond)
			h.TypeText("\033[33m"+spacer+"(No fate but what we make)\033[0m", 30)
		}},
		{name: "self_destruct", fn: func() {
			h.TypeCommand("self_destruct --override-code-omega", 50)
			time.Sleep(500 * time.Millisecond)
			h.TypeText("\033[31m[!] EMERGENCY: Ship auto-destruct sequence activated\033[0m", 30)
			time.Sleep(400 * time.Millisecond)
			h.TypeText("\033[33m[*] T-minus: 10 minutes to detonation\033[0m", 30)
			time.Sleep(600 * time.Millisecond)
			h.TypeText("\033[36m[*] Attempting to abort...\033[0m", 30)
			time.Sleep(500 * time.Millisecond)
			h.TypeText("\033[31m[!] The ship will not abort destruct sequence\033[0m", 30)
			time.Sleep(400 * time.Millisecond)
			h.TypeText("\033[32m[+] Just kidding! This is just a simulation\033[0m", 30)
			h.TypeText("\033[33m"+spacer+"(You have 10 minutes to reach minimum safe distance)\033[0m", 30)
		}},
		{name: "trophy_room", fn: func() {
			h.TypeCommand("trophy_room --view-skulls --honor-display", 50)
			time.Sleep(500 * time.Millisecond)
			h.TypeText("\033[36m[*] Accessing trophy collection database...\033[0m", 30)
			time.Sleep(400 * time.Millisecond)
			skulls := []string{"Xenomorph", "Grid anomaly", "Debugger", "Syntax Error", "Null Pointer"}
			for i := range 3 {
				h.TypeText(fmt.Sprintf("\033[33m"+spacer+"Trophy #%d: %s skull\033[0m", i+1, skulls[rand.Intn(len(skulls))]), 30)
				time.Sleep(300 * time.Millisecond)
			}
			h.TypeText("\033[32m[+] You are one ugly code base!\033[0m", 30)
		}},
		{name: "encrypt_files", fn: func() {
			h.TypeCommand("encrypt_files --algorithm=AES-256 --quantum-resistant secret_data/", 50)
			time.Sleep(400 * time.Millisecond)
			h.TypeText("\033[36m[*] Initializing quantum encryption...\033[0m", 30)
			time.Sleep(300 * time.Millisecond)
			files := []string{"launch_codes.txt", "passwords.db", "secret_recipe.pdf", "conspiracy.doc"}
			for range 3 {
				file := files[rand.Intn(len(files))]
				h.Spinner(fmt.Sprintf("Encrypting %s", file), 1500*time.Millisecond)
			}
			time.Sleep(300 * time.Millisecond)
			h.TypeText("\033[32m[+] All files encrypted with military-grade encryption!\033[0m", 30)
			h.TypeText("\033[33m"+spacer+"(Even we can't decrypt them now...)\033[0m", 30)
		}},
		{name: "decrypt_enigma", fn: func() {
			h.TypeCommand("decrypt_message --key=rosebud --cipher=enigma message.enc", 50)
			time.Sleep(500 * time.Millisecond)
			h.TypeText("\033[36m[*] Loading Enigma rotor configuration...\033[0m", 30)
			h.Spinner("Decrypting message", 2000*time.Millisecond)
			time.Sleep(400 * time.Millisecond)
			messages := []string{
				"MEET AT THE USUAL PLACE",
				"THE EAGLE HAS LANDED",
				"SQUEAKY CLEAN IS THE KEY",
				"WINTER IS COMING",
				"MAY THE FORCE BE WITH YOU",
			}
			h.TypeText(fmt.Sprintf("\033[32m[+] Decrypted: %s\033[0m", messages[rand.Intn(len(messages))]), 30)
		}},
		{name: "compress_data", fn: func() {
			h.TypeCommand("compress_data --ultra --ratio=99 ./huge_database/*", 50)
			time.Sleep(400 * time.Millisecond)
			h.TypeText("\033[36m[*] Compressing 10TB of data...\033[0m", 30)
			time.Sleep(300 * time.Millisecond)
			h.ProgressBar("Compression progress", 2500*time.Millisecond)
			originalSize := 10000 + rand.Intn(5000)
			compressedSize := 50 + rand.Intn(50)
			ratio := float64(originalSize) / float64(compressedSize)
			time.Sleep(300 * time.Millisecond)
			h.TypeText(fmt.Sprintf("\033[32m[+] Compressed %d GB → %d MB (%.1fx compression!)\033[0m", originalSize, compressedSize, ratio), 30)
			h.TypeText("\033[33m"+spacer+"(ZIP couldn't even dream of this ratio)\033[0m", 30)
		}},
		{name: "decompress_archive", fn: func() {
			h.TypeCommand("decompress_archive --extract=all stolen_data.tar.gz.bz2.xz.7z", 50)
			time.Sleep(400 * time.Millisecond)
			h.TypeText("\033[36m[*] Decompressing multi-layer archive...\033[0m", 30)
			layers := []string{"7z layer", "xz layer", "bz2 layer", "gzip layer", "tar layer"}
			for i, layer := range layers {
				h.Spinner(fmt.Sprintf("Extracting %s (%d/5)", layer, i+1), 1000*time.Millisecond)
			}
			time.Sleep(300 * time.Millisecond)
			h.TypeText("\033[32m[+] 1,337 files extracted successfully!\033[0m", 30)
			h.TypeText("\033[33m"+spacer+"(Someone really didn't want us to get this)\033[0m", 30)
		}},
		{name: "hash_crack", fn: func() {
			h.TypeCommand("hash_crack --wordlist=rockyou.txt --hash=5f4dcc3b5aa765d61d8327deb882cf99", 50)
			time.Sleep(500 * time.Millisecond)
			h.TypeText("\033[36m[*] Loading 14 billion passwords...\033[0m", 30)
			time.Sleep(400 * time.Millisecond)
			h.Spinner("Cracking hash", 2500*time.Millisecond)
			attempts := 1000000 + rand.Intn(9000000)
			time.Sleep(300 * time.Millisecond)
			h.TypeText(fmt.Sprintf("\033[32m[+] Hash cracked after %d attempts!\033[0m", attempts), 30)
			time.Sleep(200 * time.Millisecond)
			passwords := []string{"password", "123456", "qwerty", "letmein", "admin"}
			h.TypeText(fmt.Sprintf("\033[33m"+spacer+"Original password: %s\033[0m", passwords[rand.Intn(len(passwords))]), 30)
			h.TypeText("\033[33m"+spacer+"(Seriously? That was the password?)\033[0m", 30)
		}},
		{name: "steganography_hide", fn: func() {
			h.TypeCommand("steganography_hide --image=cat.jpg --payload=secrets.txt", 50)
			time.Sleep(500 * time.Millisecond)
			h.TypeText("\033[36m[*] Analyzing image pixels...\033[0m", 30)
			time.Sleep(300 * time.Millisecond)
			h.Spinner("Embedding data in LSB", 1800*time.Millisecond)
			time.Sleep(300 * time.Millisecond)
			h.TypeText("\033[32m[+] Data successfully hidden in image!\033[0m", 30)
			h.TypeText("\033[33m"+spacer+"(Now it's just a picture of a cat... or is it?)\033[0m", 30)
		}},
		{name: "steganography_extract", fn: func() {
			h.TypeCommand("steganography_extract --image=suspicious_meme.png", 50)
			time.Sleep(500 * time.Millisecond)
			h.TypeText("\033[36m[*] Scanning image for hidden data...\033[0m", 30)
			h.Spinner("Analyzing pixel patterns", 2000*time.Millisecond)
			time.Sleep(400 * time.Millisecond)
			h.TypeText("\033[32m[+] Hidden message found!\033[0m", 30)
			time.Sleep(300 * time.Millisecond)
			h.TypeText("\033[33m"+spacer+"Message: \"epstein didn't kill himself\"\033[0m", 30)
		}},
		{name: "shred_evidence", fn: func() {
			h.TypeCommand("shred_evidence --passes=35 --method=gutmann incriminating_files/*", 50)
			time.Sleep(500 * time.Millisecond)
			h.TypeText("\033[36m[*] Initializing secure deletion (35 passes)...\033[0m", 30)
			time.Sleep(300 * time.Millisecond)
			h.ProgressBar("Overwriting with random data", 2500*time.Millisecond)
			time.Sleep(300 * time.Millisecond)
			h.TypeText("\033[32m[+] Files permanently deleted!\033[0m", 30)
			h.TypeText("\033[33m"+spacer+"(Not even the FBI can recover these now)\033[0m", 30)
		}},
		{name: "generate_keys", fn: func() {
			h.TypeCommand("generate_keys --type=RSA-4096 --entropy=paranoid", 50)
			time.Sleep(500 * time.Millisecond)
			h.TypeText("\033[36m[*] Gathering entropy from cosmic background radiation...\033[0m", 30)
			h.Spinner("Generating prime numbers", 2000*time.Millisecond)
			time.Sleep(400 * time.Millisecond)
			h.TypeText("\033[32m[+] Key pair generated successfully!\033[0m", 30)
			h.TypeText("\033[33m"+spacer+"Public key: AAAAB3NzaC1yc2EAAAADAQABAAACAQ...\033[0m", 30)
			h.TypeText("\033[33m"+spacer+"(Would take the sun's lifetime to crack)\033[0m", 30)
		}},
		{name: "obfuscate_code", fn: func() {
			h.TypeCommand("obfuscate_code --level=maximum --anti-debug payload.exe", 50)
			time.Sleep(500 * time.Millisecond)
			h.TypeText("\033[36m[*] Applying code obfuscation...\033[0m", 30)
			techniques := []string{
				"Control flow flattening",
				"String encryption",
				"Dead code injection",
				"Opaque predicates",
			}
			for _, technique := range techniques {
				h.Spinner(technique, 1000*time.Millisecond)
			}
			time.Sleep(300 * time.Millisecond)
			h.TypeText("\033[32m[+] Code successfully obfuscated!\033[0m", 30)
			h.TypeText("\033[33m"+spacer+"(Good luck reverse engineering this mess)\033[0m", 30)
		}},
		{name: "backup_offshore", fn: func() {
			h.TypeCommand("backup_system --destination=offshore --encrypt --split", 50)
			time.Sleep(500 * time.Millisecond)
			h.TypeText("\033[36m[*] Creating distributed encrypted backup...\033[0m", 30)
			time.Sleep(300 * time.Millisecond)
			locations := []string{"Iceland", "Switzerland", "Cayman Islands", "Singapore"}
			for i, location := range locations {
				h.Spinner(fmt.Sprintf("Uploading shard %d/4 to %s", i+1, location), 1200*time.Millisecond)
			}
			time.Sleep(300 * time.Millisecond)
			h.TypeText("\033[32m[+] Backup distributed across 4 continents!\033[0m", 30)
			h.TypeText("\033[33m"+spacer+"(They'd need to raid 4 countries simultaneously)\033[0m", 30)
		}},
		{name: "sql_inject", fn: func() {
			h.TypeCommand("sql_inject --url="+h.target+"/login --payload='OR 1=1--", 50)
			time.Sleep(500 * time.Millisecond)
			h.TypeText("\033[36m[*] Testing for SQL injection vulnerabilities...\033[0m", 30)
			time.Sleep(400 * time.Millisecond)
			h.Spinner("Probing database", 1800*time.Millisecond)
			time.Sleep(300 * time.Millisecond)
			h.TypeText("\033[31m[!] Vulnerable endpoint detected!\033[0m", 30)
			time.Sleep(400 * time.Millisecond)
			h.TypeText("\033[32m[+] Dumping user credentials...\033[0m", 30)
			users := []string{"admin", "root", "sysadmin", "dbadmin", "test_user"}
			for range 3 {
				h.TypeText(fmt.Sprintf("\033[33m"+spacer+"%s : P@ssw0rd%d\033[0m", users[rand.Intn(len(users))], rand.Intn(999)), 30)
				time.Sleep(250 * time.Millisecond)
			}
			h.TypeText("\033[33m"+spacer+"(Bobby Tables would be proud)\033[0m", 30)
		}},
		{name: "wifi_crack", fn: func() {
			h.TypeCommand("wifi_crack --interface=wlan0 --target=FBI_Surveillance_Van_7", 50)
			time.Sleep(500 * time.Millisecond)
			h.TypeText("\033[36m[*] Capturing WPA handshake...\033[0m", 30)
			time.Sleep(300 * time.Millisecond)
			h.ProgressBar("Collecting packets", 2000*time.Millisecond)
			time.Sleep(400 * time.Millisecond)
			h.TypeText("\033[32m[+] Handshake captured!\033[0m", 30)
			time.Sleep(300 * time.Millisecond)
			h.Spinner("Running dictionary attack", 2500*time.Millisecond)
			time.Sleep(400 * time.Millisecond)
			h.TypeText("\033[32m[+] WiFi password cracked: \"FBIAgent123\"\033[0m", 30)
			h.TypeText("\033[33m"+spacer+"(They're not even trying to hide it anymore)\033[0m", 30)
		}},
		{name: "ransomware_decrypt", fn: func() {
			h.TypeCommand("ransomware_decrypt --key=universal --no-payment ransomed_files/", 50)
			time.Sleep(500 * time.Millisecond)
			h.TypeText("\033[36m[*] Analyzing ransomware encryption...\033[0m", 30)
			time.Sleep(400 * time.Millisecond)
			variants := []string{"WannaCry", "Petya", "Ryuk", "REvil", "DarkSide"}
			h.TypeText(fmt.Sprintf("\033[33m[*] Detected variant: %s\033[0m", variants[rand.Intn(len(variants))]), 30)
			time.Sleep(300 * time.Millisecond)
			h.Spinner("Exploiting implementation flaw", 2200*time.Millisecond)
			time.Sleep(400 * time.Millisecond)
			h.TypeText("\033[32m[+] All files decrypted without paying!\033[0m", 30)
			h.TypeText("\033[33m"+spacer+"(Take that, cyber criminals!)\033[0m", 30)
		}},
		{name: "blockchain_hack", fn: func() {
			h.TypeCommand("blockchain_hack --target=bitcoin --double-spend --51-percent-attack", 50)
			time.Sleep(500 * time.Millisecond)
			h.TypeText("\033[36m[*] Attempting 51% attack on blockchain...\033[0m", 30)
			time.Sleep(300 * time.Millisecond)
			h.Spinner("Mining competing chain", 3000*time.Millisecond)
			time.Sleep(500 * time.Millisecond)
			h.TypeText("\033[31m[!] ERROR: Would need more computing power than exists on Earth\033[0m", 30)
			time.Sleep(400 * time.Millisecond)
			h.TypeText("\033[33m"+spacer+"(Satoshi wins this round)\033[0m", 30)
		}},
		{name: "social_engineer", fn: func() {
			h.TypeCommand("social_engineer --target=ceo@"+h.target+" --pretend=IT-support", 50)
			time.Sleep(500 * time.Millisecond)
			h.TypeText("\033[36m[*] Crafting convincing phishing email...\033[0m", 30)
			time.Sleep(400 * time.Millisecond)
			h.Spinner("Cloning corporate email template", 1800*time.Millisecond)
			time.Sleep(300 * time.Millisecond)
			h.TypeText("\033[32m[+] Email sent successfully!\033[0m", 30)
			time.Sleep(500 * time.Millisecond)
			h.TypeText("\033[36m[*] Waiting for victim to click link...\033[0m", 30)
			time.Sleep(1000 * time.Millisecond)
			h.TypeText("\033[32m[+] CEO clicked the link! Credentials harvested!\033[0m", 30)
			h.TypeText("\033[33m"+spacer+"(The human is always the weakest link)\033[0m", 30)
		}},
		{name: "zero_day_exploit", fn: func() {
			h.TypeCommand("zero_day_exploit --vulnerability=CVE-2077-1337 --target="+h.target, 50)
			time.Sleep(500 * time.Millisecond)
			h.TypeText("\033[36m[*] Loading zero-day exploit from dark web...\033[0m", 30)
			time.Sleep(400 * time.Millisecond)
			h.TypeText("\033[33m[*] Exploit cost: 50 Bitcoin\033[0m", 30)
			time.Sleep(300 * time.Millisecond)
			h.Spinner("Deploying exploit", 2000*time.Millisecond)
			time.Sleep(400 * time.Millisecond)
			h.TypeText("\033[32m[+] Remote code execution achieved!\033[0m", 30)
			h.TypeText("\033[33m"+spacer+"(Worth every satoshi)\033[0m", 30)
		}},
		{name: "botnet_ddos", fn: func() {
			h.TypeCommand("botnet_control --zombies=10000 --ddos-target="+h.target+" --intensity=maximum", 50)
			time.Sleep(500 * time.Millisecond)
			h.TypeText("\033[36m[*] Commanding botnet to attack target...\033[0m", 30)
			time.Sleep(300 * time.Millisecond)
			h.ProgressBar("Coordinating zombies", 2500*time.Millisecond)
			time.Sleep(400 * time.Millisecond)
			h.TypeText("\033[32m[+] 10,000 zombies attacking target!\033[0m", 30)
			time.Sleep(300 * time.Millisecond)
			h.TypeText("\033[36m[*] Traffic: 500 Gbps and rising...\033[0m", 30)
			time.Sleep(400 * time.Millisecond)
			h.TypeText("\033[31m[!] Target servers melting!\033[0m", 30)
			h.TypeText("\033[33m"+spacer+"(Their ops team is having a bad day)\033[0m", 30)
		}},
		{name: "deepfake_audio", fn: func() {
			h.TypeCommand("deepfake_generate --target=ceo --say='Transfer funds to account 1337'", 50)
			time.Sleep(500 * time.Millisecond)
			h.TypeText("\033[36m[*] Analyzing target's voice patterns...\033[0m", 30)
			time.Sleep(400 * time.Millisecond)
			h.Spinner("Training neural network", 2500*time.Millisecond)
			time.Sleep(300 * time.Millisecond)
			h.TypeText("\033[32m[+] Deepfake audio generated!\033[0m", 30)
			time.Sleep(400 * time.Millisecond)
			h.TypeText("\033[36m[*] Calling finance department...\033[0m", 30)
			time.Sleep(600 * time.Millisecond)
			h.TypeText("\033[32m[+] They fell for it! $1M transferred!\033[0m", 30)
			h.TypeText("\033[33m"+spacer+"(AI is getting scary good at this)\033[0m", 30)
		}},
		{name: "keylogger_install", fn: func() {
			h.TypeCommand("keylogger_install --stealth --target=all-workstations --exfiltrate-to=dark-server", 50)
			time.Sleep(500 * time.Millisecond)
			h.TypeText("\033[36m[*] Deploying invisible keyloggers...\033[0m", 30)
			time.Sleep(300 * time.Millisecond)
			workstations := 15 + rand.Intn(35)
			h.ProgressBar(fmt.Sprintf("Installing on %d workstations", workstations), 2200*time.Millisecond)
			time.Sleep(400 * time.Millisecond)
			h.TypeText("\033[32m[+] Keyloggers active on all targets!\033[0m", 30)
			time.Sleep(300 * time.Millisecond)
			h.TypeText("\033[36m[*] Capturing keystrokes in real-time...\033[0m", 30)
			time.Sleep(500 * time.Millisecond)
			h.TypeText("\033[33m"+spacer+"(Every password is belong to us)\033[0m", 30)
		}},
		{name: "packet_sniff", fn: func() {
			h.TypeCommand("packet_sniff --interface=eth0 --filter='password|secret|confidential'", 50)
			time.Sleep(500 * time.Millisecond)
			h.TypeText("\033[36m[*] Entering promiscuous mode...\033[0m", 30)
			time.Sleep(300 * time.Millisecond)
			h.TypeText("\033[36m[*] Sniffing network traffic...\033[0m", 30)
			time.Sleep(400 * time.Millisecond)
			packets := 5000 + rand.Intn(5000)
			h.TypeText(fmt.Sprintf("\033[33m[*] Captured %d packets\033[0m", packets), 30)
			time.Sleep(300 * time.Millisecond)
			h.Spinner("Analyzing for credentials", 2000*time.Millisecond)
			time.Sleep(400 * time.Millisecond)
			h.TypeText("\033[32m[+] Found 17 passwords in cleartext!\033[0m", 30)
			h.TypeText("\033[33m"+spacer+"(Who uses HTTP in 2025?!)\033[0m", 30)
		}},
		{name: "exfiltrate_data", fn: func() {
			h.TypeCommand("exfiltrate_data --source=/home/victim/Documents --method=dns-tunneling", 50)
			time.Sleep(500 * time.Millisecond)
			h.TypeText("\033[36m[*] Encoding data into DNS queries...\033[0m", 30)
			time.Sleep(400 * time.Millisecond)
			h.ProgressBar("Tunneling through DNS", 2800*time.Millisecond)
			time.Sleep(300 * time.Millisecond)
			dataSize := 100 + rand.Intn(900)
			h.TypeText(fmt.Sprintf("\033[32m[+] Exfiltrated %d MB of sensitive data!\033[0m", dataSize), 30)
			h.TypeText("\033[33m"+spacer+"(Their firewall didn't even notice)\033[0m", 30)
		}},
		{name: "pivot_network", fn: func() {
			h.TypeCommand("pivot_network --from=dmz --to=internal --escalate-privileges", 50)
			time.Sleep(500 * time.Millisecond)
			h.TypeText("\033[36m[*] Established foothold in DMZ...\033[0m", 30)
			time.Sleep(400 * time.Millisecond)
			h.Spinner("Scanning for pivot points", 1800*time.Millisecond)
			time.Sleep(300 * time.Millisecond)
			h.TypeText("\033[32m[+] Found misconfigured jump box!\033[0m", 30)
			time.Sleep(400 * time.Millisecond)
			h.Spinner("Pivoting to internal network", 2000*time.Millisecond)
			time.Sleep(300 * time.Millisecond)
			h.TypeText("\033[32m[+] Access to internal network achieved!\033[0m", 30)
			h.TypeText("\033[33m"+spacer+"(We're in the castle now)\033[0m", 30)
		}},
		{name: "ai_password_predict", fn: func() {
			h.TypeCommand("ai_password_predict --username=admin --context-aware --ml-model=gpt-4", 50)
			time.Sleep(500 * time.Millisecond)
			h.TypeText("\033[36m[*] Analyzing target's digital footprint...\033[0m", 30)
			time.Sleep(400 * time.Millisecond)
			h.Spinner("AI predicting password", 2500*time.Millisecond)
			time.Sleep(300 * time.Millisecond)
			predictions := []string{"Fluffy2023!", "Summer@Beach", "MyK1ds123", "Tr0ub4dor&3"}
			h.TypeText("\033[32m[+] Top password predictions:\033[0m", 30)
			for i, pred := range predictions[:3] {
				h.TypeText(fmt.Sprintf("\033[33m"+spacer+"%d. %s\033[0m", i+1, pred), 30)
				time.Sleep(200 * time.Millisecond)
			}
			h.TypeText("\033[33m"+spacer+"(AI knows you better than you know yourself)\033[0m", 30)
		}},
		{name: "supply_chain_compromise", fn: func() {
			h.TypeCommand("supply_chain_compromise --target=popular-npm-package --backdoor=true", 50)
			time.Sleep(500 * time.Millisecond)
			h.TypeText("\033[36m[*] Compromising developer's account...\033[0m", 30)
			time.Sleep(400 * time.Millisecond)
			h.Spinner("Injecting malicious code", 2000*time.Millisecond)
			time.Sleep(300 * time.Millisecond)
			h.TypeText("\033[32m[+] Backdoored package published!\033[0m", 30)
			time.Sleep(400 * time.Millisecond)
			downloads := 10000 + rand.Intn(90000)
			h.TypeText(fmt.Sprintf("\033[36m[*] Package downloaded %d times today...\033[0m", downloads), 30)
			time.Sleep(300 * time.Millisecond)
			h.TypeText("\033[33m"+spacer+"(That's a lot of compromised systems)\033[0m", 30)
		}},
		{name: "quantum_crypto_break", fn: func() {
			h.TypeCommand("quantum_crypto_break --algorithm=RSA-2048 --qubits=1024", 50)
			time.Sleep(500 * time.Millisecond)
			h.TypeText("\033[36m[*] Initializing quantum computer...\033[0m", 30)
			time.Sleep(400 * time.Millisecond)
			h.Spinner("Running Shor's algorithm", 3000*time.Millisecond)
			time.Sleep(500 * time.Millisecond)
			h.TypeText("\033[32m[+] RSA-2048 factored in 0.3 seconds!\033[0m", 30)
			time.Sleep(300 * time.Millisecond)
			h.TypeText("\033[31m[!] All traditional encryption is now obsolete!\033[0m", 30)
			h.TypeText("\033[33m"+spacer+"(The crypto apocalypse is here)\033[0m", 30)
		}},
		{name: "privilege_escalate", fn: func() {
			h.TypeCommand("privilege_escalate --exploit=dirty-cow --target-user=root", 50)
			time.Sleep(500 * time.Millisecond)
			h.TypeText("\033[36m[*] Exploiting kernel vulnerability...\033[0m", 30)
			time.Sleep(400 * time.Millisecond)
			h.Spinner("Escalating privileges", 2000*time.Millisecond)
			time.Sleep(300 * time.Millisecond)
			h.TypeText("\033[32m[+] Root access obtained!\033[0m", 30)
			time.Sleep(200 * time.Millisecond)
			h.TypeText("\033[33m"+spacer+"uid=0(root) gid=0(root) groups=0(root)\033[0m", 30)
			time.Sleep(300 * time.Millisecond)
			h.TypeText("\033[33m"+spacer+"(I am become root, destroyer of worlds)\033[0m", 30)
		}},
		{name: "memory_dump", fn: func() {
			h.TypeCommand("memory_dump --process=chrome --search='cookies|session|tokens'", 50)
			time.Sleep(500 * time.Millisecond)
			h.TypeText("\033[36m[*] Attaching to Chrome process...\033[0m", 30)
			time.Sleep(300 * time.Millisecond)
			h.ProgressBar("Dumping memory", 2200*time.Millisecond)
			time.Sleep(400 * time.Millisecond)
			h.TypeText("\033[32m[+] Memory dump complete: 2.3 GB\033[0m", 30)
			time.Sleep(300 * time.Millisecond)
			h.Spinner("Extracting secrets", 1800*time.Millisecond)
			time.Sleep(300 * time.Millisecond)
			h.TypeText("\033[32m[+] Recovered 47 session tokens!\033[0m", 30)
			h.TypeText("\033[33m"+spacer+"(Your browser knows all your secrets)\033[0m", 30)
		}},
		{name: "dns_poisoning", fn: func() {
			h.TypeCommand("dns_poisoning --target-domain="+h.target+" --redirect=evil-server.onion", 50)
			time.Sleep(500 * time.Millisecond)
			h.TypeText("\033[36m[*] Intercepting DNS queries...\033[0m", 30)
			time.Sleep(400 * time.Millisecond)
			h.Spinner("Poisoning DNS cache", 2000*time.Millisecond)
			time.Sleep(300 * time.Millisecond)
			h.TypeText("\033[32m[+] DNS cache poisoned successfully!\033[0m", 30)
			time.Sleep(400 * time.Millisecond)
			victims := 100 + rand.Intn(900)
			h.TypeText(fmt.Sprintf("\033[36m[*] %d users redirected to fake site!\033[0m", victims), 30)
			h.TypeText("\033[33m"+spacer+"(They'll never know they're on a fake page)\033[0m", 30)
		}},
		{name: "reverse_shell", fn: func() {
			h.TypeCommand("reverse_shell --port=4444 --callback=attacker-server.com --obfuscate", 50)
			time.Sleep(500 * time.Millisecond)
			h.TypeText("\033[36m[*] Establishing reverse shell connection...\033[0m", 30)
			time.Sleep(400 * time.Millisecond)
			h.Spinner("Bypassing firewall", 2200*time.Millisecond)
			time.Sleep(300 * time.Millisecond)
			h.TypeText("\033[32m[+] Shell connected!\033[0m", 30)
			time.Sleep(300 * time.Millisecond)
			h.TypeText("\033[33m"+spacer+"bash-5.1$ whoami\033[0m", 30)
			time.Sleep(200 * time.Millisecond)
			h.TypeText("\033[33m"+spacer+"root\033[0m", 30)
			time.Sleep(300 * time.Millisecond)
			h.TypeText("\033[33m"+spacer+"(Sweet, sweet shell access)\033[0m", 30)
		}},
		{name: "forensics_wipe", fn: func() {
			h.TypeCommand("forensics_wipe --anti-forensic --secure-delete --cover-tracks", 50)
			time.Sleep(500 * time.Millisecond)
			h.TypeText("\033[36m[*] Initiating anti-forensic procedures...\033[0m", 30)
			time.Sleep(300 * time.Millisecond)
			tasks := []string{"Clearing logs", "Wiping bash history", "Removing timestamps", "Overwriting slack space"}
			for _, task := range tasks {
				h.Spinner(task, 1000*time.Millisecond)
			}
			time.Sleep(300 * time.Millisecond)
			h.TypeText("\033[32m[+] All traces eliminated!\033[0m", 30)
			h.TypeText("\033[33m"+spacer+"(It's like we were never here)\033[0m", 30)
		}},
		{name: "rtfm_manual", fn: func() {
			// List of common manual pages to try
			manPages := []string{"ls", "grep", "cat", "chmod", "chown", "find", "ps", "kill", "ssh", "tar", "gzip", "sed", "awk", "curl", "wget", "git", "make", "gcc", "bash", "vim"}
			selectedMan := manPages[rand.Intn(len(manPages))]
			numLines := 5 + rand.Intn(16)
			h.TypeCommand(fmt.Sprintf("man %s | head -n %d", selectedMan, numLines), 50)
			time.Sleep(500 * time.Millisecond)

			// Try to get actual man page content
			cmd := exec.Command("man", selectedMan)
			cmd.Env = append(cmd.Env, "LANG=en_US.UTF-8")
			output, err := cmd.Output()

			if err != nil || len(output) == 0 {
				h.TypeText("\033[31m[!] ERROR: Manual page not found or man command unavailable\033[0m", 30)
				time.Sleep(300 * time.Millisecond)
				h.TypeText("\033[33m"+spacer+"(Even hackers need to RTFM sometimes)\033[0m", 30)
				return
			}

			// Clean the output (remove ANSI codes from man output)
			cleanOutput := string(output)
			// Remove backspace sequences used for bold/underline in man pages
			for strings.Contains(cleanOutput, "\b") {
				cleanOutput = strings.ReplaceAll(cleanOutput, "\x08", "")
			}

			lines := strings.Split(cleanOutput, "\n")
			if len(lines) == 0 {
				h.TypeText("\033[31m[!] ERROR: Empty manual page\033[0m", 30)
				return
			}

			// Find a good starting point (skip empty lines at start)
			startLine := 0
			for startLine < len(lines) && strings.TrimSpace(lines[startLine]) == "" {
				startLine++
			}

			h.TypeText("\033[36m[*] Retrieving classified documentation...\033[0m", 30)
			time.Sleep(300 * time.Millisecond)

			linesDisplayed := 0
			for i := startLine; i < len(lines) && linesDisplayed < numLines; i++ {
				line := lines[i]
				// Trim long lines to reasonable length
				width := getTerminalWidth() - len(spacer)
				if len(line) > width {
					line = line[:width-3] + "..."
				}
				if strings.TrimSpace(line) != "" {
					h.TypeText(fmt.Sprintf("\033[33m"+spacer+"%s\033[0m", line), 15)
					linesDisplayed++
					time.Sleep(50 * time.Millisecond)
				}
			}

			time.Sleep(500 * time.Millisecond)
			h.TypeText(fmt.Sprintf("\033[32m[+] Manual for '%s' extracted successfully!\033[0m", selectedMan), 30)
			h.TypeText("\033[33m"+spacer+"(Knowledge is power... and also in /usr/share/man)\033[0m", 30)
		}},
	}
}

func (h *HackerTerminal) RunSequence() {
	h.ShowPrompt()
	h.RandomPause()

	// Run a random sequence
	sequence := h.sequences[rand.Intn(len(h.sequences))]
	sequence.fn()

	// Track sequence
	if h.stats != nil {
		h.stats.TrackSequence(sequence.name)
	}

	h.RandomEffect()

	printSeparator()
	time.Sleep(1500 * time.Millisecond)
}
