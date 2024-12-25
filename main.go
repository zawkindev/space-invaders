package main

import (
	"log"
	"math"
	"math/rand"
	"os"
	"time"

	g "github.com/zawkindev/space-invaders/game"
	"github.com/zawkindev/space-invaders/keyboard"
	r "github.com/zawkindev/space-invaders/render"
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

	// create tickers with faster timings
	enemyMoveTicker := time.NewTicker(500 * time.Millisecond)  // Faster enemy movement
	bulletMoveTicker := time.NewTicker(30 * time.Millisecond)  // Faster bullet movement
	gameLoopTicker := time.NewTicker(16666 * time.Microsecond) // ~60 FPS
	cleanupTicker := time.NewTicker(500 * time.Millisecond)    // More frequent cleanup
	defer enemyMoveTicker.Stop()
	defer bulletMoveTicker.Stop()
	defer gameLoopTicker.Stop()
	defer cleanupTicker.Stop()

	// Create bullet pool to reduce allocations
	bulletPool := make([]g.Bullet, 0, 100)

	// initialize keyboard
	keyboard.SetRawMode()
	defer keyboard.ResetRawMode()

	// Buffer size chosen to prevent blocking during rapid key presses
	PressedKey := make(chan string, 10)
	var currentKey string
	moveTicker := time.NewTicker(16 * time.Millisecond) // ~60 FPS for smooth movement
	defer moveTicker.Stop()

	// Goroutine for non-blocking keyboard input handling
	go func() {
		var buf [3]byte
		lastKey := ""
		for {
			n, err := os.Stdin.Read(buf[:])
			if err != nil {
				log.Fatal(err)
				return
			}

			// Handle special keys (arrow keys, etc.) which send 3 bytes
			if n == 3 {
				if buf[0] == keyboard.KeyEsc[0] && buf[1] == '[' {
					key := string([]byte{buf[0], buf[1], buf[2]})
					// Debounce repeated key events
					if key == lastKey {
						continue
					}
					lastKey = key
					PressedKey <- key
				}
			} else if n == 1 {
				// Single byte input handling (ESC key or key release)
				if buf[0] == keyboard.KeyEsc[0] {
					PressedKey <- keyboard.KeyEsc
				} else {
					PressedKey <- "RELEASE"
					lastKey = ""
				}
			} else {
				PressedKey <- "RELEASE"
				lastKey = ""
			}
		}
	}()

	// Main game loop with event-driven architecture
	for {
		// Collision detection between bullets and enemies
		for i := range bullets {
			b := &bullets[i]
			for j := range enemies {
				e := &enemies[j]
				switch {
				// Hit detection with 1-unit proximity threshold
				case e.X == b.X && e.Y == b.Y || math.Abs(float64(e.X-b.X)) == 1 && math.Abs(float64(e.Y-b.Y)) == 1:
					e.IsAlive = false
					b.IsActive = false
				// Boundary checks
				case e.Y > Height-2:
					e.IsAlive = false
				case b.Y < 1:
					b.IsActive = false
				}
			}
		}

		// Event multiplexing using select
		select {
		case <-gameLoopTicker.C:
			r.ClearScreen()
			Height, Width, _ = r.GetTerminalSize()
			r.Render(player, enemies, bullets, Width, Height)

		case <-bulletMoveTicker.C:
			// In-place bullet movement and cleanup using slice tricks
			activeBullets := bullets[:0]
			for i := range bullets {
				b := &bullets[i]
				if b.IsActive {
					b.Y -= 1
					if b.Y > 0 {
						activeBullets = append(activeBullets, *b)
					}
				}
			}
			bullets = activeBullets

			// Automatic fire mechanism with pool-based bullet management
			if len(bullets) < cap(bulletPool) {
				bullets = append(bullets, g.Bullet{X: player.X, Y: player.Y - 1, IsActive: true})
			}

		case <-enemyMoveTicker.C:
			// More efficient enemy management
			activeEnemies := enemies[:0]
			for i := range enemies {
				e := &enemies[i]
				if e.IsAlive {
					e.Y += 1
					if e.Y < Height-1 {
						activeEnemies = append(activeEnemies, *e)
					}
				}
			}
			enemies = activeEnemies

			// Spawn new enemies
			if len(enemies) < 40 { // Increased from 20 to 40 maximum enemies
				enemyNum := random.Intn(3) + 2 // Changed from (2)+1 to (3)+2, so 2-4 enemies at a time
				for i := 0; i < enemyNum; i++ {
					enemies = append(enemies, g.Enemy{
						X:       random.Intn(Width - 1),
						Y:       0,
						IsAlive: true,
					})
				}
			}

		case <-cleanupTicker.C:
			// Periodic cleanup of inactive objects
			bullets = cleanupBullets(bullets)
			enemies = cleanupEnemies(enemies)

		case <-moveTicker.C:
			// Player movement with boundary checks
			if currentKey != "" {
				switch currentKey {
				case keyboard.KeyArrowLeft:
					if player.X > 0 {
						player.X -= 1
					}
				case keyboard.KeyArrowRight:
					if player.X < Width-1 {
						player.X += 1
					}
				case keyboard.KeyArrowUp:
					if player.Y > 0 {
						player.Y -= 1
					}
				case keyboard.KeyArrowDown:
					if player.Y < Height-1 {
						player.Y += 1
					}
				}
			}

		case key := <-PressedKey:
			// Input state management
			if key == "RELEASE" {
				currentKey = ""
			} else if key == keyboard.KeyEsc {
				return
			} else {
				currentKey = key
			}
		}
	}
}

func cleanupBullets(bullets []g.Bullet) []g.Bullet {
	active := bullets[:0]
	for _, b := range bullets {
		if b.IsActive {
			active = append(active, b)
		}
	}
	return active
}

func cleanupEnemies(enemies []g.Enemy) []g.Enemy {
	active := enemies[:0]
	for _, e := range enemies {
		if e.IsAlive {
			active = append(active, e)
		}
	}
	return active
}
