package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"time"
)

// Stats tracks user statistics and achievements
type Stats struct {
	Version            string                      `json:"version"`
	TotalSessions      int                         `json:"total_sessions"`
	TotalCommands      int                         `json:"total_commands"`
	TotalUptimeSeconds int64                       `json:"total_uptime_seconds"`
	SequencesTriggered map[string]int              `json:"sequences_triggered"`
	EffectsTriggered   map[string]int              `json:"effects_triggered"`
	Achievements       []UnlockedAchievement       `json:"achievements"`
	LongestSessionSecs int64                       `json:"longest_session_seconds"`
	CurrentStreak      int                         `json:"current_streak"`
	LastSessionDate    string                      `json:"last_session_date"`
	SessionStartTime   time.Time                   `json:"-"`
	LastSaveTime       time.Time                   `json:"-"`
	CurrentCommands    int                         `json:"-"`
	CurrentSequences   map[string]int              `json:"-"`
	FilePath           string                      `json:"-"`
	OnNotification     func(string, string, int64) `json:"-"` // Callback for achievement notifications
}

// UnlockedAchievement represents an achievement that has been unlocked
type UnlockedAchievement struct {
	ID         string    `json:"id"`
	UnlockedAt time.Time `json:"unlocked_at"`
}

// Achievement defines an achievement
type Achievement struct {
	ID          string
	Name        string
	Description string
	Icon        string
	Check       func(*Stats) bool
}

var achievements = []Achievement{
	// Rookie
	{ID: "im_in", Name: "I'm In!", Description: "Complete your first hacking sequence", Icon: "🔓",
		Check: func(s *Stats) bool { return s.TotalCommands >= 1 }},
	{ID: "script_kiddie", Name: "Script Kiddie", Description: "Type 100 commands", Icon: "⌨️",
		Check: func(s *Stats) bool { return s.TotalCommands >= 100 }},
	{ID: "password_hunter", Name: "Password Hunter", Description: "Crack 10 passwords", Icon: "🔑",
		Check: func(s *Stats) bool { return s.getTotalSequenceCount("password") >= 10 }},
	{ID: "green_screen", Name: "Green Screen", Description: "Trigger 5 CRT effects", Icon: "💚",
		Check: func(s *Stats) bool { return s.EffectsTriggered["crt_scan"] >= 5 }},

	// Intermediate
	{ID: "glitch_master", Name: "Glitch Master", Description: "Experience 50 screen glitches", Icon: "🤖",
		Check: func(s *Stats) bool { return s.EffectsTriggered["glitch"] >= 50 }},
	{ID: "matrix_fan", Name: "Matrix Fan", Description: "See Matrix rain 25 times", Icon: "🌧️",
		Check: func(s *Stats) bool { return s.EffectsTriggered["matrix_rain"] >= 25 }},
	{ID: "virus_deployer", Name: "Virus Deployer", Description: "Deploy 50 fake viruses", Icon: "☠️",
		Check: func(s *Stats) bool { return s.getTotalSequenceCount("virus") >= 50 }},
	{ID: "splash_collector", Name: "Splash Collector", Description: "See 20 ASCII splash screens", Icon: "💀",
		Check: func(s *Stats) bool { return s.EffectsTriggered["splash"] >= 20 }},

	// Advanced
	{ID: "hack_planet", Name: "Hack the Planet!", Description: "Run for 1 hour total", Icon: "🏆",
		Check: func(s *Stats) bool { return s.TotalUptimeSeconds >= 3600 }},
	{ID: "pentagon_infiltrator", Name: "Pentagon Infiltrator", Description: "Hack pentagon 50 times", Icon: "🏛️",
		Check: func(s *Stats) bool { return s.SequencesTriggered["pentagon"] >= 50 }},
	{ID: "command_master", Name: "Command Master", Description: "Type 1000 commands", Icon: "⚡",
		Check: func(s *Stats) bool { return s.TotalCommands >= 1000 }},
	{ID: "dedicated_hacker", Name: "Dedicated Hacker", Description: "Complete 50 sessions", Icon: "🎭",
		Check: func(s *Stats) bool { return s.TotalSessions >= 50 }},

	// Rare
	{ID: "night_owl", Name: "Night Owl", Description: "Start a session between 2-4 AM", Icon: "🕐",
		Check: func(s *Stats) bool {
			hour := s.SessionStartTime.Hour()
			return hour >= 2 && hour < 4
		}},
	{ID: "marathon_runner", Name: "Marathon Runner", Description: "Run a single session for 30 minutes", Icon: "🏃",
		Check: func(s *Stats) bool {
			if s.SessionStartTime.IsZero() {
				return false
			}
			return time.Since(s.SessionStartTime).Seconds() >= 1800
		}},
	{ID: "lucky_seven", Name: "Lucky Seven", Description: "Complete exactly 7 sessions", Icon: "🎲",
		Check: func(s *Stats) bool { return s.TotalSessions == 7 }},
}

// Helper method to count sequences by partial name match
func (s *Stats) getTotalSequenceCount(keyword string) int {
	total := 0
	for name, count := range s.SequencesTriggered {
		if len(name) >= len(keyword) {
			// Simple substring match
			for i := 0; i <= len(name)-len(keyword); i++ {
				if name[i:i+len(keyword)] == keyword {
					total += count
					break
				}
			}
		}
	}
	return total
}

// getConfigDir returns the application config directory (platform-specific)
func getConfigDir() (string, error) {
	configBase, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	configDir := filepath.Join(configBase, "hackerminal")
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return "", err
	}

	return configDir, nil
}

// LoadStats loads statistics from disk
func LoadStats() (*Stats, error) {
	configDir, err := getConfigDir()
	if err != nil {
		return nil, err
	}

	filePath := filepath.Join(configDir, "stats.json")
	stats := &Stats{
		Version:            "1.2.3",
		SequencesTriggered: make(map[string]int),
		EffectsTriggered:   make(map[string]int),
		Achievements:       []UnlockedAchievement{},
		SessionStartTime:   time.Now(),
		CurrentSequences:   make(map[string]int),
		FilePath:           filePath,
	}

	// Try to load existing stats
	data, err := os.ReadFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			// First run, return fresh stats
			return stats, nil
		}
		return nil, err
	}

	if err := json.Unmarshal(data, stats); err != nil {
		return nil, err
	}

	// Update version to current version
	stats.Version = "1.2.3"
	stats.SessionStartTime = time.Now()
	stats.CurrentSequences = make(map[string]int)
	stats.FilePath = filePath

	return stats, nil
}

// Save writes statistics to disk
func (s *Stats) Save() error {
	// Update session duration
	now := time.Now()
	delta := int64(now.Sub(s.LastSaveTime).Seconds())
	s.TotalUptimeSeconds += delta
	s.LastSaveTime = now

	// Update longest session
	if !s.SessionStartTime.IsZero() {
		sessionDuration := int64(time.Since(s.SessionStartTime).Seconds())
		if sessionDuration > s.LongestSessionSecs {
			s.LongestSessionSecs = sessionDuration
		}
	}

	// Update streak
	today := time.Now().Format("2006-01-02")
	if s.LastSessionDate != today {
		yesterday := time.Now().AddDate(0, 0, -1).Format("2006-01-02")
		if s.LastSessionDate == yesterday {
			s.CurrentStreak++
		} else if s.LastSessionDate != "" {
			s.CurrentStreak = 1
		} else {
			s.CurrentStreak = 1
		}
		s.LastSessionDate = today
	}

	data, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(s.FilePath, data, 0644)
}

// TrackCommand increments command counter
func (s *Stats) TrackCommand() {
	s.TotalCommands++
	s.CurrentCommands++
}

// TrackSequence records a sequence trigger
func (s *Stats) TrackSequence(name string) {
	s.SequencesTriggered[name]++
	s.CurrentSequences[name]++
	s.checkAchievements()
}

// TrackEffect records an effect trigger
func (s *Stats) TrackEffect(effect string) {
	s.EffectsTriggered[effect]++
	s.checkAchievements()
}

// checkAchievements checks if any new achievements should be unlocked
func (s *Stats) checkAchievements() {
	for _, achievement := range achievements {
		// Check if already unlocked
		if s.hasAchievement(achievement.ID) {
			continue
		}

		// Check if conditions are met
		if achievement.Check(s) {
			s.unlockAchievement(achievement)
		}
	}
}

// hasAchievement checks if an achievement is already unlocked
func (s *Stats) hasAchievement(id string) bool {
	for _, ua := range s.Achievements {
		if ua.ID == id {
			return true
		}
	}
	return false
}

// unlockAchievement unlocks a new achievement and triggers notification
func (s *Stats) unlockAchievement(achievement Achievement) {
	s.Achievements = append(s.Achievements, UnlockedAchievement{
		ID:         achievement.ID,
		UnlockedAt: time.Now(),
	})

	// Format achievement notification
	notification := s.formatAchievement(achievement)

	// Call the callback if set
	if s.OnNotification != nil {
		s.OnNotification(notification, "\033[1;33m", 2000)
	}
}

// formatAchievement formats an achievement notification as a string
func (s *Stats) formatAchievement(achievement Achievement) (result string) {
	result += "\n\033[1;33m"
	result += "╔════════════════════════════════════════════════════╗\n"
	result += fmt.Sprintf("║  %s ACHIEVEMENT UNLOCKED!                          ║\n", achievement.Icon)
	result += "╠════════════════════════════════════════════════════╣\n"
	result += fmt.Sprintf("║  %-49s ║\n", achievement.Name)
	result += fmt.Sprintf("║  %-49s ║\n", achievement.Description)
	result += "╚════════════════════════════════════════════════════╝"
	result += "\033[0m\n"
	return result
}

// PrintStats displays current session and all-time statistics
func (s *Stats) PrintStats() {
	// Format notification
	notification := s.formatCurrentStats()

	// Call the callback if set
	if s.OnNotification != nil {
		s.OnNotification(notification, "\033[38;5;46m", 0)
	}
}

// formatCurrentStats returns current statistics formatted
func (s *Stats) formatCurrentStats() (result string) {
	sessionDuration := int64(0)
	if !s.SessionStartTime.IsZero() {
		sessionDuration = int64(time.Since(s.SessionStartTime).Seconds())
	}

	result += "\n\033[1;32m╔════════════════════════════════════════════════════╗\n"
	result += "║              HACKERMINAL STATISTICS                ║\n"
	result += "╠════════════════════════════════════════════════════╣\n"
	result += "║                                                    ║\n"
	result += "║  SESSION STATS                                     ║\n"
	result += fmt.Sprintf("║    Commands typed:     %-7d                     ║\n", s.CurrentCommands)
	result += fmt.Sprintf("║    Session time:       %-7s                     ║\n", formatDuration(sessionDuration))
	if len(s.CurrentSequences) > 0 {
		result += "║                                                    ║\n"
		result += "║  SEQUENCES THIS SESSION                            ║\n"
		// Show top 3 sequences
		type seqCount struct {
			name  string
			count int
		}
		var sequences []seqCount
		for name, count := range s.CurrentSequences {
			sequences = append(sequences, seqCount{name, count})
		}
		sort.Slice(sequences, func(i, j int) bool {
			return sequences[i].count > sequences[j].count
		})
		for i := 0; i < len(sequences) && i < 3; i++ {
			result += fmt.Sprintf("║    %-20s x%-4d                      ║\n", truncate(sequences[i].name, 20), sequences[i].count)
		}
	}
	result += "║                                                    ║\n"
	result += "║  ALL-TIME STATS                                    ║\n"
	result += fmt.Sprintf("║    Total sessions:     %-7d                     ║\n", s.TotalSessions)
	result += fmt.Sprintf("║    Total commands:     %-7d                     ║\n", s.TotalCommands)
	result += fmt.Sprintf("║    Total time:         %-7s                     ║\n", formatDuration(s.TotalUptimeSeconds))
	result += fmt.Sprintf("║    Longest session:    %-7s                     ║\n", formatDuration(s.LongestSessionSecs))
	result += fmt.Sprintf("║    Current streak:     %-7s                     ║\n", formatStreak(s.CurrentStreak))
	result += fmt.Sprintf("║    Achievements:       %-7s                     ║\n", formatAchievements(s.Achievements))
	result += "║                                                    ║\n"

	// Show recent achievements
	if len(s.Achievements) > 0 {
		result += "║  RECENT ACHIEVEMENTS                               ║\n"
		start := 0
		if len(s.Achievements) > 3 {
			start = len(s.Achievements) - 3
		}
		for i := start; i < len(s.Achievements); i++ {
			for _, achievement := range achievements {
				if achievement.ID == s.Achievements[i].ID {
					result += fmt.Sprintf("║    %s %-44s ║\n", achievement.Icon, truncate(achievement.Name, 42))
					break
				}
			}
		}
		result += "║                                                    ║\n"
	}

	result += "╚════════════════════════════════════════════════════╝\n\n\033[0m"

	return
}

// formatDuration formats seconds into a readable duration
func formatDuration(seconds int64) string {
	if seconds < 60 {
		return fmt.Sprintf("%ds", seconds)
	} else if seconds < 3600 {
		return fmt.Sprintf("%dm %ds", seconds/60, seconds%60)
	} else {
		hours := seconds / 3600
		minutes := (seconds % 3600) / 60
		return fmt.Sprintf("%dh %dm", hours, minutes)
	}
}

// truncate truncates a string to a maximum length
func truncate(s string, max int) string {
	if len(s) <= max {
		return s
	}
	return s[:max-3] + "..."
}

// StartNewSession initializes a new session and tracks session-related achievements
func (s *Stats) StartNewSession() {
	s.TotalSessions++
	s.SessionStartTime = time.Now()
	s.LastSaveTime = time.Now()
	s.CurrentCommands = 0
	s.CurrentSequences = make(map[string]int)
	s.checkAchievements()
}

// formatStreak formats a streak duration as a day count string
func formatStreak(duration int) (result string) {
	suffix := ""
	if duration > 1 {
		suffix = "s"
	}
	result = fmt.Sprintf("%d day%s", duration, suffix)
	return
}

// formatAchievements formats achievement progress as unlocked/total count
func formatAchievements(unlockedAchievements []UnlockedAchievement) string {
	return fmt.Sprintf("%d / %d", len(unlockedAchievements), len(achievements))
}
