package builder

type Director struct {
	builder Builder
}

func (this Director) setBuilder(b Builder) {
	this.builder = b
}

func (this Director) buildPart(wa, do, wi int) {
	this.builder.setWalls(wa)
	this.builder.setDoors(do)
	this.builder.setWindows(wi)
}

func main() {
	dir := Director{}

	dir.setBuilder(villa{})
	dir.buildPart(1, 2, 3)

	dir.setBuilder(apartment{})
	dir.buildPart(2, 4, 8)
}
