package abstractFactory

type leatherChair struct {
}

type leatherSofa struct {
}

type leatherTable struct {
}

func (leatherChair) get() Furniture {
	return woodChair{}
}

func (leatherSofa) get() Furniture {
	return woodSofa{}
}

func (leatherTable) get() Furniture {
	return woodTable{}
}
