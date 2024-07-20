package main

import (
	"log"
	"math"
	"math/rand"
	g "space-invaders/game"
	r "space-invaders/render"
	"time"

	"github.com/nsf/termbox-go"
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

	// create tickers
	enemyMoveTicker := time.NewTicker(500 * time.Millisecond)
	gameLoopTicker := time.NewTicker(60 * time.Millisecond)
	defer enemyMoveTicker.Stop()
	defer gameLoopTicker.Stop()

	// initialize keyboard
	if err := termbox.Init(); err != nil {
		panic(err)
	}
	defer termbox.Close()

	// create channel for key event
	PressedKey := make(chan termbox.Key, 10)
	go func() {
		for {
			if ev := termbox.PollEvent(); ev.Type == termbox.EventKey {
				PressedKey <- ev.Key
			}
		}
	}()

	// game loop
	for {
		for i := range bullets {
			b := &bullets[i]
			for j := range enemies {
				e := &enemies[j]
				switch {
				case e.X == b.X && e.Y == b.Y || math.Abs(float64(e.X-b.X)) == 1 && math.Abs(float64(e.Y-b.Y)) == 1:
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
				activeEnemies = append(activeEnemies, *e)
			}
		}
		enemies = activeEnemies

		select {
		case <-gameLoopTicker.C:
			// clear screen
			r.ClearScreen()

			// define screen size
			Height, Width, err = r.GetTerminalSize()
			if err != nil {
				log.Panic(err)
			}

			// shoot bullets
			bullets = append(bullets, g.Bullet{X: player.X, Y: player.Y - 1, IsActive: true})

			// collision detection && move bullets
			for i := range bullets {
				b := &bullets[i]
				b.Y -= 1
			}

			// render objects
			r.Render(player, enemies, bullets, Width, Height)

		case <-enemyMoveTicker.C:
			// add random enemies at random locations
			enemyNum := random.Intn(3)
			for i := 0; i < enemyNum; i++ {
				enemies = append(enemies, g.Enemy{X: random.Intn(Width - 1), Y: 0, IsAlive: true})
			}

			for i := range enemies {
				e := &enemies[i]
				e.Y += 1
			}

		case key := <-PressedKey:
			if key == termbox.KeyEsc {
				return
			}
			switch key {
			case termbox.KeyArrowLeft:
				if player.X > 0 {
					player.X -= 1
				}
			case termbox.KeyArrowRight:
				if player.X < Width-1 {
					player.X += 1
				}
			case termbox.KeyArrowUp:
				if player.Y > 0 {
					player.Y -= 1
				}
			case termbox.KeyArrowDown:
				if player.Y < Height-1 {
					player.Y += 1
				}
			}
		}
	}
}
