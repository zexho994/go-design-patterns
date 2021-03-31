package builder

type House interface {
	setWalls(nums int)
	setWindows(nums int)
	setDoors(nums int)
}

type villa struct {
	walls   int
	windows int
	doors   int
}

func (v villa) setWalls(nums int) {
	v.walls = nums
	// 建造villa的wall
}

func (v villa) setWindows(nums int) {
	v.windows = nums
	// 建造villa的windows
}

func (v villa) setDoors(nums int) {
	v.doors = nums
	// 建造villa的doors
}

func (v villa) build() {
	// build the villa
}

type apartment struct {
	walls   int
	windows int
	doors   int
}

func (a apartment) setWalls(nums int) {
	a.walls = nums
	// 建造apartment的wall
}

func (a apartment) setWindows(nums int) {
	a.windows = nums
	// 建造apartment的windows
}

func (a apartment) setDoors(nums int) {
	a.doors = nums
	// 建造apartment的doors
}

func (a apartment) build() {
	// build the apartment
}
func client() {
	v := villa{}
	v.setDoors(2)
	v.setWindows(4)
	v.setWalls(8)
	v.build()

	a := apartment{}
	a.setDoors(1)
	a.setWindows(2)
	a.setWalls(4)
	a.build()
}
