package models

type Game struct {
	Grid      [][]int
	Rows      int
	Columns   int
	Turn      int
	Player1   string
	Player2   string
	Winner    int
	Draw      bool
	GravityUp bool
	MoveCount int
	ScoreP1   int
	ScoreP2   int
}

type GameHistory struct {
	Player1   string
	Player2   string
	Winner    string
	MoveCount int
	Gravity   string
}

func NewGame(rows, cols int, p1, p2 string) *Game {
	grid := make([][]int, rows)
	for i := range grid {
		grid[i] = make([]int, cols)
	}
	return &Game{
		Grid:    grid,
		Rows:    rows,
		Columns: cols,
		Player1: p1,
		Player2: p2,
		Turn:    1,
	}
}

func (g *Game) Play(col int) bool {
	if col < 0 || col >= g.Columns || g.Winner != 0 || g.Draw {
		return false
	}
	var rowRange []int
	if g.GravityUp {
		for i := 0; i < g.Rows; i++ {
			rowRange = append(rowRange, i)
		}
	} else {
		for i := g.Rows - 1; i >= 0; i-- {
			rowRange = append(rowRange, i)
		}
	}
	for _, row := range rowRange {
		if g.Grid[row][col] == 0 {
			g.Grid[row][col] = g.Turn
			g.MoveCount++
			if g.MoveCount%5 == 0 {
				g.GravityUp = !g.GravityUp
			}
			if g.CheckWin(row, col) {
				g.Winner = g.Turn
				if g.Turn == 1 {
					g.ScoreP1++
				} else {
					g.ScoreP2++
				}
			} else if g.CheckDraw() {
				g.Draw = true
			} else {
				g.Turn = 3 - g.Turn
			}
			return true
		}
	}
	return false
}

func (g *Game) CheckWin(r, c int) bool {
	player := g.Grid[r][c]
	directions := [][]int{{0, 1}, {1, 0}, {1, 1}, {1, -1}}
	for _, d := range directions {
		count := 1
		for i := 1; i < 4; i++ {
			nr, nc := r+i*d[0], c+i*d[1]
			if nr >= 0 && nr < g.Rows && nc >= 0 && nc < g.Columns && g.Grid[nr][nc] == player {
				count++
			} else {
				break
			}
		}
		for i := 1; i < 4; i++ {
			nr, nc := r-i*d[0], c-i*d[1]
			if nr >= 0 && nr < g.Rows && nc >= 0 && nc < g.Columns && g.Grid[nr][nc] == player {
				count++
			} else {
				break
			}
		}
		if count >= 4 {
			return true
		}
	}
	return false
}

func (g *Game) CheckDraw() bool {
	for _, row := range g.Grid {
		for _, cell := range row {
			if cell == 0 {
				return false
			}
		}
	}
	return true
}
