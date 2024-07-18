package main

import (
	"math/rand"
	g "space-invaders/game"
	r "space-invaders/render"
	"time"
)

func main() {
	// set up randomizer
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	// create enemies
	enemies := make([]g.Enemy, random.Intn(5))
	for i := 0; i < 5; i++ {
		enemies[i] = g.Enemy{X: random.Intn(20), Y: 0}
	}

	// create player
	player := g.Player{X: 20, Y: 19}

	// game loop
	for {
		r.Render(enemies, player)
		time.Sleep(17 * time.Millisecond)
		r.ClearScreen()
	}
}
