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
	enemies := make([]g.Enemy, 0)

	// create bullets
	bullets := make([]g.Bullet, 0)

	// game loop
	for {

		// shoot bullets
		bullets = append(bullets, g.Bullet{X: player.X, Y: player.Y - 1, IsActive: true})

		// add random enemies at random locations
		enemyNum := random.Intn(5)
		for i := 0; i < enemyNum; i++ {
			enemies = append(enemies, g.Enemy{X: random.Intn(r.Width-1), Y: 0, IsAlive: true, Moved: false})
		}

		// collision detection
		for i := range bullets {
			b := &bullets[i]
			b.Y -= 1
			for j := range enemies {
				e := &enemies[j]
				if !e.Moved {
					e.Y += 1
					e.Moved = true
				}
				switch {
				case e.X == b.X && e.Y == b.Y:
					e.IsAlive = false
					b.IsActive = false

				case e.Y > r.Height-2:
					e.IsAlive = false

				case b.Y < 1:
					b.IsActive = false
				}
			}
		}

		for i := range enemies {
			e := &enemies[i]
			e.Moved = false
		}

		r.Render(player, enemies, bullets)
		time.Sleep(1000 * time.Millisecond)
		r.ClearScreen()
	}
}
