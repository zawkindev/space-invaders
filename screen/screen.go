package console

import (
	"log"
	"os"
	"os/exec"
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

var Position position
var Width, Height, MaxX, MaxY int

func init() {
	Width, Height = ScreenSize()
  MaxX, MaxY = Width/2, Height/2

	Position = position{
		Top:    Cordinate{X: 0, Y: MaxY},
		Bottom:    Cordinate{X: 0, Y: -MaxY},
		Right:    Cordinate{X: MaxX, Y: 0},
    Left: Cordinate{X: -MaxX, Y: 0},
		Center: Cordinate{X: 0, Y: 0},
    TopRight: Cordinate{X: MaxX, Y: MaxY},
    TopLeft: Cordinate{X: -MaxX, Y: MaxY},
    BottomRight: Cordinate{X: MaxX, Y: -MaxY},
    BottomLeft: Cordinate{X: -MaxX, Y: -MaxY},
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
			width, err = strconv.Atoi(temp)
			if err != nil {
				log.Fatal(err)
			}
			temp = ""
		case b == '\n':
			height, err = strconv.Atoi(temp)
			if err != nil {
				log.Fatal(err)
			}
		default:
			temp += string(b)
		}
	}

	return width, height
}
