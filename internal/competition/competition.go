package competition

import (
	"sync"
)

// Struktur Tim
type Team struct {
	ID       int
	Name     string
	Token    string
	Score    int
	Attempts int
}

// Struktur Soal
type Challenge struct {
	ID     int
	Title  string
	Flag   string
	Points int
}

// Struktur Pengumuman
type Announcement struct {
	ID      int
	Message string
}

// Database kompetisi (sementara, bisa pakai DB)
var (
	teams         = []Team{}
	challenges    = []Challenge{}
	announcements = []Announcement{}
	teamMutex     sync.Mutex
)

// Tambah Tim
func AddTeam(name, token string) {
	teamMutex.Lock()
	defer teamMutex.Unlock()
	teams = append(teams, Team{ID: len(teams) + 1, Name: name, Token: token, Score: 0, Attempts: 0})
}

// Tambah Soal
func AddChallenge(title, flag string, points int) {
	challenges = append(challenges, Challenge{ID: len(challenges) + 1, Title: title, Flag: flag, Points: points})
}

// Validasi Flag
func ValidateFlag(flag, teamToken string) bool {
	teamMutex.Lock()
	defer teamMutex.Unlock()

	for _, ch := range challenges {
		if ch.Flag == flag {
			// Beri tim yang benar skor tambahan
			for i, t := range teams {
				if t.Token == teamToken {
					teams[i].Score += ch.Points
					teams[i].Attempts++
					return true
				}
			}
		}
	}
	return false
}

// Dapatkan Daftar Tim
func GetTeams() []Team {
	return teams
}

// Dapatkan Daftar Soal
func GetChallenges() []Challenge {
	return challenges
}

// Tambah Pengumuman
func AddAnnouncement(message string) {
	announcements = append(announcements, Announcement{ID: len(announcements) + 1, Message: message})
}

// Dapatkan Daftar Pengumuman
func GetAnnouncements() []Announcement {
	return announcements
}
