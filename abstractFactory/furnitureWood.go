package abstractFactory

type woodChair struct {
}

type woodSofa struct {
}

type woodTable struct {
}

func (woodChair) get() Furniture {
	return woodChair{}
}

func (woodSofa) get() Furniture {
	return woodSofa{}
}

func (woodTable) get() Furniture {
	return woodTable{}
}
