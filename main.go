package main

import (
	s "space-invaders/screen"
	"time"
)

func main() {
	helloText := s.Model{Cordinate: s.Position.Center, View: "hello"}

	for {
		time.Sleep(16 * time.Millisecond) // 60 fps = 1:60
		s.Render(helloText)
	}

}
