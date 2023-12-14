package maze

import (
	"math/rand"
	"openpair/cheese-maze/src/mouse"
)

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
    // Mark the current cell as a path
    m.Grid[y][x] = "⬜️" 

    // Randomize the order of directions to ensure different maze patterns
    rand.Shuffle(len(directions), func(i, j int) {
        directions[i], directions[j] = directions[j], directions[i]
    })

    // Iterate through each direction
    for _, dir := range directions {
        // Calculate the coordinates of the next cell in this direction
        nx, ny := x + 2*dir.dx, y + 2*dir.dy

        // Check if the next cell is within the maze bounds
        if nx >= 0 && ny >= 0 && nx < m.Width && ny < m.Height {
            // Check if the next cell is a wall (has not been visited)
            if m.Grid[ny][nx] == "🟫" {
                // Convert the wall between the current cell and the next cell into a path
                // This "carves" the path in the maze
                m.Grid[y+dir.dy][x+dir.dx] = "⬜️"

                // Continue the depth-first search from the next cell
                m.dfs(nx, ny)
            }
        }
    }
}


