package render

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	g "space-invaders/game"
	"syscall"
	"unsafe"
)

type winsize struct {
	Row    uint16
	Col    uint16
	Xpixel uint16
	Ypixel uint16
}

func createMatrix(height, width int) [][]byte {
	matrix := make([][]byte, height)
	for i := range matrix {
		matrix[i] = make([]byte, width)
	}
	return matrix
}

func Render(player g.Player, enemies []g.Enemy, bullets []g.Bullet, Width, Height int) {
	matrix := createMatrix(Height, Width)
	for i := 0; i < Height; i++ {
		for j := 0; j < Width; j++ {
			matrix[i][j] = ' '
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
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func GetTerminalSize() (int, int, error) {
	ws := &winsize{}
	retCode, _, errno := syscall.Syscall(
		syscall.SYS_IOCTL,
		os.Stdout.Fd(),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(ws)),
	)
	if int(retCode) == -1 {
		return 0, 0, errno
	}
	return int(ws.Row), int(ws.Col), nil
}
