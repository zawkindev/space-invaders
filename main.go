package main

import (
	"fmt"
	c "space-invaders/console"
)

var Position c.Position

func init() {
	Position = c.Position{
		Center: c.Cordinate{X: 0, Y: 0},
	}
}

func main() {
	fmt.Println("Screen size: ", c.ScreenSize())
}
