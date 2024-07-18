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

	// game loop
	for {
		bullets = append(bullets, g.Bullet{X: player.X, Y: player.Y - 1})

		activeBullets := bullets[:0] // using the same underlying array to avoid allocations
		for i := range bullets {
			b := &bullets[i]
			if b.Y > 0 {
				b.Y -= 1
				activeBullets = append(activeBullets, *b)
			}
		}
		bullets = activeBullets

		r.Render(player, enemies, bullets)
		time.Sleep(300 * time.Millisecond)
		r.ClearScreen()
	}
}
