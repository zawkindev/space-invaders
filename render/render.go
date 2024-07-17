package render

import (
	"fmt"
	g "space-invaders/game"
)

const Width = 40  // screen width
const Height = 20 // screen height

func Render(enemies []g.Enemey) {

	// make and fill the matrix to represent empty screen
	var matrix [Width][Height]byte
	for i := 0; i < Width; i++ {
		for j := 0; j < Height; j++ {
			matrix[i][j] = ' '
		}
	}

	// place enemies to matrix
	for _, e := range enemies {
		matrix[e.X][e.Y] = 'E'
	}

	// print the matrix
	for i := 0; i < Width; i++ {
		line := ""
		for j := 0; j < Height; j++ {
			line += string(matrix[i][j])
		}
		fmt.Println(line)
	}

}
