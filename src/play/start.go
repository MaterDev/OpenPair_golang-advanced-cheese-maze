package play

import (
	"bufio"
	"fmt"
	"openpair/cheese-maze/src/handlers"
	"openpair/cheese-maze/src/maze"
	"openpair/cheese-maze/src/mouse"
	"os"
	"strings"
)

func Start (mouse mouse.Mouse, maze maze.Maze) {
	// bufio.NewReader will allow us a buffer to hold the userinput
	reader := bufio.NewReader(os.Stdin)
	
	// an infinite loop (rounds of gameplay, per user input)
	for {
		// clear the terminal
		handlers.ClearScreen()

		// print the maze
		handlers.PrintMaze(maze, mouse)
		fmt.Printf("Current position: x:%d, y:%d \n", mouse.X, mouse.Y)

		// ask player for input
		fmt.Print("Move (WASD): ")
		// receive input - will pause loop until player provides unput
		input, _ := reader.ReadString('\n')
		// handle user input
		handlers.HandleInput(strings.TrimSpace(input), &mouse, maze)

		// evaluate if the cheese has been acquired by mouse
		if mouse.X == 1 && mouse.Y == 8 {
			fmt.Println("The cheese has been found!!!!")
			break
		}
	}
}