package main

import (
	s "space-invaders/screen"
	"time"
)

func main() {
	models := make([]s.Model, 0, 10)
	helloText := s.Model{Cordinate: s.Position.Center, View: "hello"}
	models = append(models, helloText)
	s.Render(helloText)

	for {
		time.Sleep(0 * time.Millisecond)
		s.Render(helloText)
	}
}
