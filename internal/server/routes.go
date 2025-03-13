package server

import (
	"encoding/json"
	"html/template"
	"net/http"
	"sync"

	"cardinal-re/internal/competition"
	"cardinal-re/internal/monitor"
	"cardinal-re/internal/scoring"
)

// Struktur data untuk UI
type PageData struct {
	Title         string
	Message       string
	SystemInfo    monitor.SystemStats
	Network       monitor.NetworkStats
	Processes     []monitor.ProcessInfo
	Teams         []competition.Team
	Challenges    []competition.Challenge
	Announcements []competition.Announcement
	Leaderboard   []scoring.ScoreEntry
}

var mutex sync.Mutex

// Inisialisasi routing
func initRoutes() {
	http.HandleFunc("/", dashboardHandler)
	http.HandleFunc("/admin", adminDashboardHandler)
	http.HandleFunc("/api/stats", apiStatsHandler)
	http.HandleFunc("/api/flag", flagSubmissionHandler)
	http.HandleFunc("/api/scores", getScoresHandler)
}

// Handler dashboard utama untuk tim
func dashboardHandler(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		Title:       "Cyber Security CTF",
		Message:     "Selamat datang di kompetisi CTF!",
		SystemInfo:  monitor.GetSystemStats(),
		Network:     monitor.GetNetworkStats(),
		Processes:   monitor.GetRunningProcesses(),
		Leaderboard: scoring.GetLeaderboard(),
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, data)
}

// Handler dashboard admin
func adminDashboardHandler(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		Title:       "Admin Dashboard",
		Message:     "Kelola kompetisi dari sini",
		Teams:       competition.GetTeams(),
		Challenges:  competition.GetChallenges(),
		Leaderboard: scoring.GetLeaderboard(),
	}

	tmpl, err := template.ParseFiles("templates/admin.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, data)
}

// API Endpoint untuk mendapatkan data statistik sistem
func apiStatsHandler(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		SystemInfo: monitor.GetSystemStats(),
		Network:    monitor.GetNetworkStats(),
		Processes:  monitor.GetRunningProcesses(),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

// API Endpoint untuk menerima Flag submission
func flagSubmissionHandler(w http.ResponseWriter, r *http.Request) {
	var submission struct {
		Flag      string `json:"flag"`
		TeamToken string `json:"token"`
	}

	err := json.NewDecoder(r.Body).Decode(&submission)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	mutex.Lock()
	defer mutex.Unlock()

	// Pastikan ValidateFlag menerima dua parameter
	// valid, teamName := competition.ValidateFlag(submission.Flag, submission.TeamToken)
	valid := competition.ValidateFlag(submission.Flag, submission.TeamToken)
	if valid {
		scoring.UpdateScore(submission.TeamToken)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status": "success", "message": "Flag correct!"}`))
	} else {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte(`{"status": "fail", "message": "Invalid Flag!"}`))
	}
}

// API Endpoint untuk mendapatkan leaderboard skor
func getScoresHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(scoring.GetLeaderboard())
}
