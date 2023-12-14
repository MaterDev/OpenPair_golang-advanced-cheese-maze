package maze

import (
	"math/rand"
	"openpair/cheese-maze/src/mouse"
)

type Cell struct {
    X, Y     int
    Visited  bool
}

// Directions for moving in the maze
var directions = []struct {
    dx int
    dy int
}{
    {0, -1}, // up
    {1, 0},  // right
    {0, 1},  // down
    {-1, 0}, // left
}

func NewMaze(width, height int) *Maze {
    grid := make([][]string, height)
    for i := range grid {
        grid[i] = make([]string, width)
        for j := range grid[i] {
            grid[i][j] = "🟫" // Initialize all cells as walls
        }
    }
    return &Maze{Grid: grid, Width: width, Height: height}
}

func (m *Maze) GenerateMaze(mouse mouse.Mouse) {
    // Starting from mouse position
    m.dfs(mouse.X, mouse.Y)
}

func (m *Maze) dfs(x, y int) {
    // Mark the current cell as part of the path and as visited
    m.Grid[y][x] = "⬜️" 

    // Shuffle directions to ensure randomness
    rand.Shuffle(len(directions), func(i, j int) {
        directions[i], directions[j] = directions[j], directions[i]
    })

    for _, dir := range directions {
        nx, ny := x + 2*dir.dx, y + 2*dir.dy

        // Check if the new cell is within bounds and has not been visited
        if nx >= 0 && ny >= 0 && nx < m.Width && ny < m.Height && m.Grid[ny][nx] == "🟫" {
            // Remove the wall between the current cell and the new cell
            m.Grid[y+dir.dy][x+dir.dx] = "⬜️"
            m.dfs(nx, ny)
        }
    }
}

