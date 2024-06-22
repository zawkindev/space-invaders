package main

import (
	"fmt"
	s "space-invaders/screen"
)

func main() {
	helloText := s.Model{Cordinate: s.Position.Center, View: "hello"}
	fmt.Printf("width: %d, height: %d\n", s.Width, s.Height)
  s.Render(helloText)
}
