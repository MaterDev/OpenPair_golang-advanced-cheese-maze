package main

import (
	"openpair/cheese-maze/src/maze"
	"openpair/cheese-maze/src/mouse"
	"openpair/cheese-maze/src/play"
)

// ! Main gameplay loop
func main() {
	// intialize a new mouse
	mouse := mouse.Mouse{X: 1, Y: 1}
	// intialize a new maze
	width, height := 10, 10
	maze := maze.CreateMaze(width, height)

	play.Start(mouse, maze)
}
