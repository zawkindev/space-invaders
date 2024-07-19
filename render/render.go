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

func Render(player g.Player, enemies []g.Enemy, bullets []g.Bullet) {

	// create and fill the matrix to represent empty screen
	var matrix [Height][Width]byte
	for i := 0; i < Height; i++ {
		for j := 0; j < Width; j++ {
			matrix[i][j] = '.'
		}
	}

	// place enemies to matrix
	for _, e := range enemies {
		if e.IsAlive {
			matrix[e.Y][e.X] = 'E'
		}
	}

	// place bullets to matrix
	for _, b := range bullets {
		if b.IsActive {
			matrix[b.Y][b.X] = '|'
		}
	}

	// place player
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
