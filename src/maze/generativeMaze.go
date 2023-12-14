package maze

import (
	"math/rand"
	"openpair/cheese-maze/src/mouse"
)

//! Directions for moving in the maze
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

//! Depth First Search - to create navigable spaces
func (m *Maze) dfs(x, y int) {
    // Mark the current cell as a path
        // The origin position, where the mouse is, will always be the start point for DFS.
    m.Grid[y][x] = "⬜️" 

    //! Randomize the order of directions to ensure different maze patterns
    /*
        This code snippet is responsible for randomizing the order of the directions in which the maze paths will be carved. The randomization ensures that each generated maze has a unique pattern, as the direction chosen to extend the path from any given cell will vary each time the maze is generated.

        Variables:
        - directions: This is an array or slice containing the four cardinal directions (up, down, left, right), each represented as a struct with `dx` and `dy` values. These values are used to move in the maze grid.
        - i, j: These are indices for elements in the `directions` slice. They are used to swap two elements in the `directions` array.

        Process:
        1. `rand.Shuffle` is a function provided by Go's `math/rand` package, used to randomize the order of elements in a slice.
        2. `len(directions)` determines the length of the `directions` slice, which is the total number of directions.
        3. The lambda function `func(i, j int)` is a callback provided to `rand.Shuffle`. It defines how the elements in the slice should be swapped during the shuffling process.
        4. Inside the lambda function, `directions[i], directions[j] = directions[j], directions[i]` swaps the elements at indices `i` and `j` in the `directions` slice. 
        - This swapping occurs multiple times in different combinations as `rand.Shuffle` iterates over the slice, effectively randomizing the order of the directions.
        5. As a result, each call to generate a maze will likely start carving paths in different directions, leading to a variety of maze patterns.
    */
    rand.Shuffle(len(directions), func(i, j int) {
        directions[i], directions[j] = directions[j], directions[i]
    })

    //! Iterate through each direction
    /*
        This loop iterates through each of the four cardinal directions (up, down, left, right) to carve out paths in the maze. The directions are represented by `dir`, which has two properties: `dx` and `dy`. These properties are used to determine the next cell's coordinates in the given direction.

        Variables:
        - dir: Represents the current direction from the set of directions. 
            `dir.dx` and `dir.dy` are the differences in the x and y coordinates for this direction, respectively.
        - nx, ny: These are the coordinates of the next cell in the maze. They are calculated based on the current cell's coordinates (`x`, `y`) and the direction (`dir`). 
                The expression `x + 2*dir.dx` and `y + 2*dir.dy` moves two steps from the current cell in the direction of `dir`. This ensures that we leave one cell as a wall between two path cells, maintaining the maze structure.

        Loop Process:
        1. For each direction, calculate the next cell's coordinates (`nx`, `ny`).
        2. Check if the next cell (`nx`, `ny`) is within the bounds of the maze. This ensures we don't access cells outside the maze's grid.
        3. If the next cell is within bounds, check if it is currently a wall (`"🟫"`). A wall cell indicates that it has not been visited yet.
        4. If the next cell is a wall, convert the wall cell between the current and the next cell into a path (`"⬜️"`). This 'carves' a path in the maze.
        - The wall cell to be converted is the immediate neighbor of the current cell in the direction of `dir`, calculated as `(y+dir.dy, x+dir.dx)`.
        5. After carving the path, continue the depth-first search from the next cell (`nx`, `ny`) by recursively calling `m.dfs(nx, ny)`. This step explores further paths from the next cell, continuing the maze generation process.
    */

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


