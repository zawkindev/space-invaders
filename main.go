package main

import (
	s "space-invaders/screen"
)

func main() {
	helloText := s.Model{Cordinate: s.Position.Center, Content: "hello\tworld\nbye world!"}

	// for {
	// 	time.Sleep(16 * time.Millisecond) // 60 fps = 1:60
	// }
	s.Render(helloText)
}
