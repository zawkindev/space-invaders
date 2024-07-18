package render

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	g "space-invaders/game"
)

const Width = 40  // screen width
const Height = 20 // screen height

func Render(enemies []g.Enemy, player g.Player) {

	// make and fill the matrix to represent empty screen
	var matrix [Height][Width]byte
	for i := 0; i < Height; i++ {
		for j := 0; j < Width; j++ {
			matrix[i][j] = '.'
		}
	}

	// place enemies and player to matrix
	for _, e := range enemies {
		matrix[e.Y][e.X] = 'E'
	}

	matrix[player.Y][player.X] = 'P'

	// print the matrix
	for i := 0; i < Height; i++ {
		line := ""
		for j := 0; j < Width; j++ {
			line += string(matrix[i][j])
		}
		fmt.Println(line)
	}
}

func ClearScreen() {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	default:
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}
