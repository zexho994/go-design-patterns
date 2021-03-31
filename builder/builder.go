package builder

type Builder interface {
	setWalls(n int)
	setDoors(n int)
	setWindows(n int)
}
