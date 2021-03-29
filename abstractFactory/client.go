package abstractFactory

func shopping() {
	var wf FurnitureFactory = WoodFactory{}
	wf.getFurniture("table")
	wf.getFurniture("chair")
	wf.getFurniture("sofa")

	var lf FurnitureFactory = LeatherFactory{}
	lf.getFurniture("table")
	lf.getFurniture("chair")
	lf.getFurniture("sofa")
}
