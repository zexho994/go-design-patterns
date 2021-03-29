package abstractFactory

type FurnitureFactory interface {
	getFurniture(f string) Furniture
}

type WoodFactory struct {
}

func (WoodFactory) getFurniture(f string) Furniture {
	if f == "chair" {
		return woodChair{}
	}
	if f == "sofa" {
		return woodSofa{}
	}
	if f == "table" {
		return woodTable{}
	}
	return nil
}

type LeatherFactory struct {
}

func (LeatherFactory) getFurniture(f string) Furniture {
	if f == "chair" {
		return leatherChair{}
	}
	if f == "sofa" {
		return leatherSofa{}
	}
	if f == "table" {
		return leatherTable{}
	}
	return nil
}
