package main

import (
	s "space-invaders/screen"
	"time"
)

func main() {
	models := make([]s.Model, 0)
	helloText := s.Model{Cordinate: s.Position.Center, View: "hello"}
	models = append(models, helloText)
	s.Render(helloText)

	for {
		time.Sleep(16 * time.Millisecond)
		s.Render(helloText)
	}
}
