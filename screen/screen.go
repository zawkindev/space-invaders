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

type Position struct {
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

func ScreenSize() Cordinate {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	var cor Cordinate
	var temp string = ""

	for _, b := range out {
		switch {
		case b == ' ':
			cor.X, err = strconv.Atoi(temp)
			if err != nil {
				log.Fatal(err)
			}
			temp = ""
		case b == '\n':
			cor.Y, err = strconv.Atoi(temp)
			if err != nil {
				log.Fatal(err)
			}
		default:
			temp += string(b)
		}
	}

	return cor
}
