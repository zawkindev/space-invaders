package main

import (
	"fmt"
	s "space-invaders/screen"
)

var Position s.Position

func init() {
	Position = s.Position{
		Center: s.Cordinate{X: 0, Y: 0},
	}
}

func main() {
	fmt.Println("Ss.een size: ", s.ScreenSize())
}
