package handlers

import (
	"openpair/cheese-maze/src/maze"
	"openpair/cheese-maze/src/mouse"
)

// ! Function to handle changing the mouse position in the grind.
// Will take in an input of `W/A/S/D`, which will correlate to moving in one of 4 directions.
func HandleInput(input string, inc_mouse *mouse.Mouse, maze maze.Maze) {
	switch input {
	// Only move mouse if target space is "⬜️" or "🧀".
	case "w":
		// Move mouse up. x = x, y = y - 1
		if maze.Grid[inc_mouse.Y-1][inc_mouse.X] == "⬜️" || maze.Grid[inc_mouse.Y-1][inc_mouse.X] == "🧀" {
			inc_mouse.Y--
		}
	case "a":
		// Move mouse left. x = x - 1, y = y
		if maze.Grid[inc_mouse.Y][inc_mouse.X-1] == "⬜️" || maze.Grid[inc_mouse.Y][inc_mouse.X-1] == "🧀" {
			inc_mouse.X--
		}
	case "s":
		// Move mouse down. x = x, y = y + 1
		if maze.Grid[inc_mouse.Y+1][inc_mouse.X] == "⬜️" || maze.Grid[inc_mouse.Y+1][inc_mouse.X] == "🧀" {
			inc_mouse.Y++
		}
	case "d":
		// Move mouse right. x = x + 1, y = y
		if maze.Grid[inc_mouse.Y][inc_mouse.X+1] == "⬜️" || maze.Grid[inc_mouse.Y][inc_mouse.X+1] == "🧀" {
			inc_mouse.X++
		}
	}
}
