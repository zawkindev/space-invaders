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

	// create player
	player := g.Player{X: 20, Y: 19}

	// create enemies
	enemyNum := random.Intn(5)
	enemies := make([]g.Enemy, enemyNum)
	for i := 0; i < enemyNum; i++ {
		enemies[i] = g.Enemy{X: random.Intn(20), Y: 0}
	}

	// create bullets
	bullets := make([]g.Bullet, 0)
	bullets = append(bullets, g.Bullet{X: rand.Intn(10), Y: 15, Active: true})

	// game loop
	for {
		r.Render(player, enemies, bullets)
		time.Sleep(17 * time.Millisecond)
		r.ClearScreen()
	}
}
