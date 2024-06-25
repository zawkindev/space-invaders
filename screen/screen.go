package console

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strconv"
)

type Cordinate struct {
	X, Y int
}

type position struct {
	Top         Cordinate
	Bottom      Cordinate
	Right       Cordinate
	Left        Cordinate
	Center      Cordinate
	TopRight    Cordinate
	TopLeft     Cordinate
	BottomRight Cordinate
	BottomLeft  Cordinate
}

type Model struct {
	Cordinate Cordinate
	View      string
}

func (m *Model) constructor() {
	viewWidth := Width - len(m.View)/2
	viewHeight := func() int {
		var n int
		for _, b := range m.View {
			if b == '\n' {
				n++
			}
		}
		if n == 0 {
			return 1
		}
		return n
	}()

	m.Cordinate.X = viewWidth - (MaxX - m.Cordinate.X)
	m.Cordinate.Y = viewHeight - (MaxY + m.Cordinate.Y)
}

type Line struct {
	Contents  []byte
	hasObject bool
}

var Position position
var Width, Height, MaxX, MaxY int

func init() {
	Width, Height = ScreenSize()
	MaxX, MaxY = Width/2, Height/2

	Position = position{
		Top:         Cordinate{X: 0, Y: MaxY},
		Bottom:      Cordinate{X: 0, Y: -MaxY},
		Right:       Cordinate{X: MaxX, Y: 0},
		Left:        Cordinate{X: -MaxX, Y: 0},
		Center:      Cordinate{X: 0, Y: 0},
		TopRight:    Cordinate{X: MaxX, Y: MaxY},
		TopLeft:     Cordinate{X: -MaxX, Y: MaxY},
		BottomRight: Cordinate{X: MaxX, Y: -MaxY},
		BottomLeft:  Cordinate{X: -MaxX, Y: -MaxY},
	}
}

func ScreenSize() (int, int) {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	var temp string = ""
	var width, height int

	for _, b := range out {
		switch {
		case b == ' ':
			height, err = strconv.Atoi(temp)
			if err != nil {
				log.Fatal(err)
			}
			temp = ""
		case b == '\n':
			width, err = strconv.Atoi(temp)
			if err != nil {
				log.Fatal(err)
			}
		default:
			temp += string(b)
		}
	}

	return width, height
}

func clearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func printScreen(matrix *[]Line) {
	m := *matrix
	var screen string

	for i := 0; i < Height; i++ {
		for j := 0; j < Width; j++ {
			if j == Width-1 || !m[i].hasObject {
				screen += string('\n')
				continue
			}
			screen += string(m[i].Contents[j])
		}
	}

	fmt.Print(screen)

}

func Render(models ...Model) {
	clearScreen()

	matrix := make([]Line, Height)
	for i := range matrix {
		matrix[i].Contents = make([]byte, Width)
	}

	for _, el := range models {
		el.constructor()
	}

	printScreen(&matrix)
}
