package factory

type FurnitureFactor struct {
}

func (FurnitureFactor) getFurniture(f string) Furniture {
	if f == "chair" {
		return chair{}
	}
	if f == "sofa" {
		return sofa{}
	}
	if f == "table" {
		return table{}
	}
	return nil
}
