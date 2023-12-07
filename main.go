package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

// * Represents the game maze.
type Maze struct {
	// The actual Grid itself
	Grid   [][]string
	Width  int
	Height int
}

// * Mouse represents the player character's coordinates
type Mouse struct {
	X, Y int
}

// ! Main gameplay loop
func main() {
	// intialize a new mouse
	mouse := Mouse{X: 1, Y: 1}
	// intialize a new maze
	maze := createMaze()

	// bufio.NewReader will allow us a buffer to hold the userinput
	reader := bufio.NewReader(os.Stdin)

	// an infinite loop (rounds of gameplay, per user input)
	for {
		// clear the terminal
		clearScreen()

		// print the maze
		printMaze(maze, mouse)
		fmt.Printf("Current position: x:%d, y:%d \n", mouse.X, mouse.Y)

		// ask player for input
		fmt.Print("Move (WASD): ")
		// receive input - will pause loop until player provides unput
		input, _ := reader.ReadString('\n')
		// handle user input
		handleInput(strings.TrimSpace(input), &mouse, maze)

		// evaluate if the cheese has been acquired by mouse
		if mouse.X == 1 && mouse.Y == 8 {
			fmt.Println("The cheese has been found!!!!")
			break
		}
	}
}

// ! Function to handle screen clearing/reloading
func clearScreen() {
	var cmd *exec.Cmd
	// runtime.GOOS: programmatically dermine what the current operating system is.
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	// Before command is ran, need to store the OS standard output inside of the cmd object.
	cmd.Stdout = os.Stdout
	// ! Run Command
	cmd.Run()
}

// ! Function to print maze
// * The mouse will be added to the Grid at print-time
func printMaze(inc_maze Maze, inc_mouse Mouse) {
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

// ! Function that will return the maze structure.
func createMaze() Maze {
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
		Width:  10,
		Height: 10,
	}
	return layout
}

// ! Function to handle changing the mouse position in the grind.
// Will take in an input of `W/A/S/D`, which will correlate to moving in one of 4 directions.
func handleInput(input string, inc_mouse *Mouse, maze Maze) {
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
