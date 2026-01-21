package main

import (
	_ "embed"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

//go:embed assets/beep.wav
var embeddedBeepWav []byte

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

func getAudioPlayer() string {
	players := []string{
		"paplay",  // PulseAudio
		"aplay",   // ALSA
		"ffplay",  // FFmpeg (with -nodisp -autoexit)
		"mpv",     // mpv
		"mplayer", // MPlayer
		"cvlc",    // VLC (command line, no interface)
		"afplay",  // macOS
	}

	for _, player := range players {
		if _, err := exec.LookPath(player); err == nil {
			return player
		}
	}
	return ""
}

func playAudioFile(audioPath string) error {
	player := getAudioPlayer()
	if player == "" {
		// No audio player found, fall back to terminal bell
		fmt.Print("\a")
		return nil
	}

	var cmd *exec.Cmd
	switch player {
	case "paplay", "aplay":
		cmd = exec.Command(player, audioPath)
	case "ffplay":
		cmd = exec.Command(player, "-nodisp", "-autoexit", "-loglevel", "quiet", audioPath)
	case "mpv":
		cmd = exec.Command(player, "--no-video", "--really-quiet", audioPath)
	case "mplayer":
		cmd = exec.Command(player, "-really-quiet", "-novideo", audioPath)
	case "cvlc":
		cmd = exec.Command(player, "--play-and-exit", "--quiet", audioPath)
	case "afplay":
		cmd = exec.Command(player, audioPath)
	default:
		// Fallback to terminal bell
		fmt.Print("\a")
		return nil
	}

	// Run in background, don't wait for completion
	if err := cmd.Start(); err != nil {
		// If audio playback fails, fall back to terminal bell
		fmt.Print("\a")
	}

	return nil
}

func beep() {
	// Try external file first (for development/customization)
	execPath, err := os.Executable()
	if err == nil {
		execDir := filepath.Dir(execPath)
		audioPath := filepath.Join(execDir, "assets", "beep.wav")

		if _, err := os.Stat(audioPath); err == nil {
			playAudioFile(audioPath)
			return
		}

		// Try relative path
		audioPath = "assets/beep.wav"
		if _, err := os.Stat(audioPath); err == nil {
			playAudioFile(audioPath)
			return
		}
	}

	// Use embedded audio file
	tmpFile, err := os.CreateTemp("", "hackerminal-beep-*.wav")
	if err != nil {
		// Fallback to terminal bell
		fmt.Print("\a")
		return
	}
	defer os.Remove(tmpFile.Name())
	defer tmpFile.Close()

	if _, err := tmpFile.Write(embeddedBeepWav); err != nil {
		fmt.Print("\a")
		return
	}

	tmpFile.Close()
	playAudioFile(tmpFile.Name())
}

func beepTimes(count int) {
	for i := 0; i < count; i++ {
		beep()
		if i < count-1 {
			time.Sleep(150 * time.Millisecond)
		}
	}
}

func dramaticBeep() {
	beepTimes(3)
}
