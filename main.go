package main

import (
	"fmt"
	s "space-invaders/screen"
)



func main() {
	fmt.Printf("width: %d, height: %d\n", s.Width, s.Height)
  fmt.Println(s.Position.Bottom)
}
