package main

import (
	"html/template"
	"net/http"
	"strconv"
	"sync"

	"power4/models"
)

var (
	game    *models.Game
	history []models.GameHistory
	mu      sync.Mutex
)

func main() {
	http.HandleFunc("/", startPage)
	http.HandleFunc("/start", startGame)
	http.HandleFunc("/play", playMove)
	http.HandleFunc("/win", winPage)
	http.HandleFunc("/draw", drawPage)
	http.HandleFunc("/rematch", rematch)
	http.HandleFunc("/history", showHistory)
	http.Handle("/static", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.ListenAndServe(":8080", nil)
}

func startPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/start.html"))
	tmpl.Execute(w, nil)
}
func startGame(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	p1 := r.FormValue("player1")
	p2 := r.FormValue("player2")
	diff := r.FormValue("difficulty")

	var rows, cols int
	switch diff {
	case "easy":
		rows, cols = 6, 7
	case "normal":
		rows, cols = 6, 9
	case "hard":
		rows, cols = 7, 8
	default:
		rows, cols = 6, 7
	}
	mu.Lock()
	game = models.NewGame(rows, cols, p1, p2)
	mu.Unlock()
	http.Redirect(w, r, "/play", http.StatusSeeOther)
}
func playMove(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	colStr := r.FormValue("column")

	mu.Lock()
	defer mu.Unlock()

	if colStr == "" {
		renderGame(w)
		return
	}
	col, err := strconv.Atoi(colStr)
	if err != nil {
		renderGame(w)
		return
	}
	game.Play(col)

	if game.Winner != 0 {
		recordWin()
		http.Redirect(w, r, "/win", http.StatusSeeOther)
		return
	}
	if game.Draw {
		recordDraw()
		http.Redirect(w, r, "/draw", http.StatusSeeOther)
		return
	}
	renderGame(w)
}
func renderGame(w http.ResponseWriter) {
	current := game.Player1
	if game.Turn == 2 {
		current = game.Player2
	}
	tmpl := template.Must(template.ParseFiles("templates/game.html"))
	tmpl.Execute(w, map[string]interface{}{
		"Grid":          game.Grid,
		"CureentPlayer": current,
		"Columns":       make([]int, game.Columns),
		"GravityUp":     game.GravityUp,
		"MoveCount":     game.MoveCount,
		"Player1":       game.Player1,
		"Player2":       game.Player2,
		"ScoreP1":       game.ScoreP1,
		"ScoreP2":       game.ScoreP2,
	})
}
func recordWin() {
	winner := game.Player1
	if game.Winner == 2 {
		winner = game.Player2
	}
	history = append(history, models.GameHistory{
		Player1:   game.Player1,
		Player2:   game.Player2,
		Winner:    winner,
		MoveCount: game.MoveCount,
		Gravity:   gravityLabel(game.GravityUp),
	})
}
func recordDraw() {
	history = append(history, models.GameHistory{
		Player1:   game.Player1,
		Player2:   game.Player2,
		Winner:    "Match Nul",
		MoveCount: game.MoveCount,
		Gravity:   gravityLabel(game.GravityUp),
	})
}
func gravityLabel(up bool) string {
	if up {
		return "Inversée"
	}
	return "Normale"
}
func winPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/win.html"))
	name := game.Player1
	if game.Winner == 2 {
		name = game.Player2
	}
	tmpl.Execute(w, map[string]interface{}{
		"Grid":       game.Grid,
		"WinnerName": name,
	})
}
func drawPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/draw.html"))
	tmpl.Execute(w, map[string]interface{}{
		"Grid": game.Grid,
	})
}
func rematch(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	old := game
	game = models.NewGame(old.Rows, old.Columns, old.Player1, old.Player2)
	game.ScoreP1 = old.ScoreP1
	game.ScoreP2 = old.ScoreP2
	mu.Unlock()
	http.Redirect(w, r, "/play", http.StatusSeeOther)
}
func showHistory(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/history.html"))
	tmpl.Execute(w, history)
}
