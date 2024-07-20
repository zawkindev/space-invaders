package main

import (
	"log"
	"math/rand"
	g "space-invaders/game"
	r "space-invaders/render"
	"time"
)

func main() {

	// define Size
	Height, Width, err := r.GetTerminalSize()
	if err != nil {
		log.Fatal(err)
	}
	// set up randomizer
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	// create player
	player := g.Player{X: Width / 2, Y: Height - 1}

	// create enemies
	enemies := make([]g.Enemy, 0)

	// create bullets
	bullets := make([]g.Bullet, 0)

	// game loop
	for {
		Height, Width, err = r.GetTerminalSize()
		if err != nil {
			log.Panic(err)
		}

		// shoot bullets
		bullets = append(bullets, g.Bullet{X: player.X, Y: player.Y - 1, IsActive: true})

		// add random enemies at random locations
		enemyNum := random.Intn(5)
		for i := 0; i < enemyNum; i++ {
			enemies = append(enemies, g.Enemy{X: random.Intn(Width - 1), Y: 0, IsAlive: true, Moved: false})
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

				case e.Y > Height-2:
					e.IsAlive = false

				case b.Y < 1:
					b.IsActive = false
				}
			}
		}

		activeEnemies := enemies[:0]
		for i := range enemies {
			e := &enemies[i]
			if e.IsAlive {
				e.Moved = false
				activeEnemies = append(activeEnemies, *e)
			}
		}

		enemies = activeEnemies

		r.Render(player, enemies, bullets, Width, Height)
		time.Sleep(60 * time.Millisecond)
		r.ClearScreen()
	}
}
