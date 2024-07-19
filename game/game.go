package game

type Player struct {
	X, Y int
}

type Enemy struct {
	X, Y    int
	IsAlive bool
	Moved bool
}

type Bullet struct {
	X, Y     int
	IsActive bool
}
