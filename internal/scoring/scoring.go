package scoring

import (
	"sort"
	"sync"

	"cardinal-re/internal/competition"
)

// Struktur Entry Skor
type ScoreEntry struct {
	Rank     int
	TeamName string
	Score    int
}

var scoreMutex sync.Mutex

// Mendapatkan Leaderboard
func GetLeaderboard() []ScoreEntry {
	scoreMutex.Lock()
	defer scoreMutex.Unlock()

	teams := competition.GetTeams()
	var leaderboard []ScoreEntry

	for _, t := range teams {
		leaderboard = append(leaderboard, ScoreEntry{
			TeamName: t.Name,
			Score:    t.Score,
		})
	}

	// Urutkan berdasarkan skor tertinggi
	sort.Slice(leaderboard, func(i, j int) bool {
		return leaderboard[i].Score > leaderboard[j].Score
	})

	// Beri ranking
	for i := range leaderboard {
		leaderboard[i].Rank = i + 1
	}

	return leaderboard
}

// Update skor ketika tim mendapatkan Flag
func UpdateScore(teamToken string) {
	scoreMutex.Lock()
	defer scoreMutex.Unlock()

	for i, t := range competition.GetTeams() {
		if t.Token == teamToken {
			competition.GetTeams()[i].Score += 10 // Default skor tambahan
		}
	}
}
