package main

import (
	"fmt"
	s "space-invaders/screen"
	"time"
)

func main() {
	models := make([]s.Model, 0, 10)
	helloText := s.Model{Cordinate: s.Position.Center, View: "hello"}
	models = append(models, helloText)
	fmt.Printf("width: %d, height: %d\n", s.Width, s.Height)
	s.Render(helloText)

	for {
		time.Sleep(600 * time.Millisecond)
		s.Render(helloText)
	}
}
