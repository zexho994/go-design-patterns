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

var c chair
var t table
var s sofa

// 带缓存的工厂
func (FurnitureFactor) getFurnitureCache(f string) Furniture {
	if f == "chair" {
		return c
	}
	if f == "sofa" {
		return s
	}
	if f == "table" {
		return t
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
