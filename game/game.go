package game

type Player struct {
	X, Y int
}

type Enemy struct {
	X, Y  int
	alive bool
}

type Bullet struct {
	X, Y   int
	Active bool
}
