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

var CHAIR chair
var TABLE table
var SOFA sofa

// 带缓存的工厂
func (FurnitureFactor) getFurnitureCache(f string) Furniture {
	if f == "chair" {
		return CHAIR
	}
	if f == "sofa" {
		return TABLE
	}
	if f == "table" {
		return SOFA
	}
	return nil
}

// 静态工厂
func GetFurnitureStatic(f string) Furniture {
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
