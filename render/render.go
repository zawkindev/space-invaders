package render

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"

	g "github.com/zawkindev/space-invaders/game"
)

// Declare both functions, implementations are in platform-specific files

// func GetTerminalSize() (int, int, error) {
// 	if runtime.GOOS == "windows" {
// 		return getTerminalSizeWindows()
// 	}
// 	return getTerminalSizeUnix()
// }

func ClearScreen() {
	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		if err := cmd.Run(); err != nil {
			fmt.Print("\033[H\033[2J")
		}
	default:
		fmt.Print("\033[H\033[2J")
	}
}

func Render(player g.Player, enemies []g.Enemy, bullets []g.Bullet, Width, Height int) {
	var buffer strings.Builder
	buffer.Grow(Height * (Width + 1))

	matrix := make([][]byte, Height)
	for i := range matrix {
		matrix[i] = make([]byte, Width)
		for j := range matrix[i] {
			matrix[i][j] = ' '
		}
	}

	// place game objects
	for _, e := range enemies {
		if e.IsAlive && e.Y < Height {
			matrix[e.Y][e.X] = '#' // Block character for enemies
		}
	}

	for _, b := range bullets {
		if b.IsActive && b.Y > 0 {
			matrix[b.Y][b.X] = '^' // Arrow for bullets
		}
	}

	matrix[player.Y][player.X] = 'A' // Triangle for player

	// Build frame
	buffer.WriteString("\033[H")
	for i := 0; i < Height; i++ {
		buffer.Write(matrix[i])
		buffer.WriteByte('\n')
	}

	fmt.Print(buffer.String())
}
