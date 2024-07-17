package main

import (
	"math/rand"
	g "space-invaders/game"
	r "space-invaders/render"
	"time"
)

func main() {
	enemies := make([]g.Enemey, 5)
	for i := 0; i < 5; i++ {
		source := rand.NewSource(time.Now().UnixNano())
		r := rand.New(source)
		enemies[i] = g.Enemey{X: r.Intn(20), Y: 0}
	}

	r.Render(enemies)

}
