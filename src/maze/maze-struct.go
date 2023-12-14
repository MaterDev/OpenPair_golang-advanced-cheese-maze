package maze

// * Represents the game maze.
type Maze struct {
	// The actual Grid itself
	Grid   [][]string
	Width  int
	Height int
}

// ! Function that will return the static maze structure.
// ! Deprecated
func CreateMaze(width int, height int) Maze {
	// Simple static maze layout
	layout := Maze{
		Grid: [][]string{
			{"🟫", "🟫", "🟫", "🟫", "🟫", "🟫", "🟫", "🟫", "🟫", "🟫"},
			{"🟫", "⬜️", "⬜️", "⬜️", "⬜️", "⬜️", "⬜️", "⬜️", "⬜️", "🟫"},
			{"🟫", "⬜️", "🟫", "🟫", "🟫", "🟫", "🟫", "🟫", "⬜️", "🟫"},
			{"🟫", "⬜️", "🟫", "🟫", "🟫", "🟫", "🟫", "🟫", "⬜️", "🟫"},
			{"🟫", "⬜️", "🟫", "🟫", "🟫", "🟫", "🟫", "🟫", "🟫", "🟫"},
			{"🟫", "⬜️", "⬜️", "⬜️", "⬜️", "⬜️", "🟫", "🟫", "🟫", "🟫"},
			{"🟫", "🟫", "🟫", "🟫", "🟫", "⬜️", "🟫", "🟫", "🟫", "🟫"},
			{"🟫", "🟫", "🟫", "🟫", "🟫", "⬜️", "🟫", "🟫", "🟫", "🟫"},
			{"🟫", "🧀", "⬜️", "⬜️", "⬜️", "⬜️", "🟫", "🟫", "🟫", "🟫"},
			{"🟫", "🟫", "🟫", "🟫", "🟫", "🟫", "🟫", "🟫", "🟫", "🟫"},
		},
		Width:  width,
		Height: height,
	}
	return layout
}