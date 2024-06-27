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
	Content   string
	View      []Row
	Width     int
	Height    int
}

func (m *Model) constructor() {
	m.Height, m.Width = func() (int, int) {
		var h, w, tempW int
		for _, b := range m.Content {
			tempW++
			if b == '\n' {
				if w < tempW-1 {
					w = tempW - 1
				}
				tempW = 0
				h++
			}
		}
		if h == 0 {
			return 1, len(m.Content)
		} else if h == 1 && m.Content[len(m.Content)-1] != '\n' {
			return h + 1, w
		}
		return h, w
	}()

	m.Cordinate.X = Width - (MaxX - m.Cordinate.X)
	m.Cordinate.Y = Height - (MaxY + m.Cordinate.Y)
}

type Row struct {
	Columns   []byte
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

func printScreen(matrix *[]Row) {
	m := *matrix
	for i := 0; i < Height; i++ {
		if m[i].hasObject {
			for j := 0; j < Width; j++ {
				current := m[i].Columns[j]
				if current == 0 {
					fmt.Print(" ")
				} else {
					fmt.Print(string(current))
				}
			}
		} else {
			fmt.Println("")
		}

	}
}

func Render(models ...Model) {
	clearScreen()

	matrix := make([]Row, Height)
	for i := range matrix {
		matrix[i].Columns = make([]byte, Width)
	}

	for _, m := range models {
		m.constructor()

		fmt.Println("width & height of model: ", m.Width, m.Height)

		for i := 0; i < m.Height; i++ {
			row := m.Cordinate.Y - m.Height/2 + i
			matrix[row].hasObject = true

			fmt.Println("current row: ", row)

			for j := 0; j < m.Width; j++ {
				column := m.Cordinate.X - m.Width/2 + j
				matrix[row].Columns[column] = m.Content[j]

				fmt.Println("current column: ", column)
			}
		}
	}

	printScreen(&matrix)
}
