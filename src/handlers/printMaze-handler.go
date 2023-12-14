package handlers

import (
	"fmt"
	"openpair/cheese-maze/src/maze"
	"openpair/cheese-maze/src/mouse"
)

// ! Function to print maze
// * The mouse will be added to the Grid at print-time
func PrintMaze(inc_maze maze.Maze, inc_mouse mouse.Mouse) {
	// TODO: create struct for cheese and add cheese at print-time, using coordinates.
	for y, line := range inc_maze.Grid {
		for x, char := range line {
			// will print mouse `M` to replace character at given coordinate
			if inc_mouse.X == x && inc_mouse.Y == y {
				fmt.Print("🐭")
			} else {
				// when looping over a string, the characters are converted into unicode decimals that represent the character.
				// ? Unicode: https://en.wikipedia.org/wiki/List_of_Unicode_characters
				// string() will turn unicode numbers back into the appropriate glyph
				fmt.Print(string(char))
			}
		}
		fmt.Println()
	}
}