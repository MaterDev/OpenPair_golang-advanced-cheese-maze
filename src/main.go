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
	width, height := 30, 30

	// Static maze
	// maze := maze.CreateMaze(width, height)
	// play.Start(mouse, maze)

	// Generative maze
	maze := maze.NewMaze(width, height)

	maze.GenerateMaze(mouse)

	play.Start(mouse, *maze)
}
