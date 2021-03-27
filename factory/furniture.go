package factory

import "fmt"

type Furniture interface {
	name() string
	get() Furniture
}

type chair struct {
}

type sofa struct {
}

type table struct {
}

func (chair) get() Furniture {
	return chair{}
}

func (chair) name() string {
	return "chair"
}

func (sofa) get() Furniture {
	return sofa{}
}

func (sofa) name() string {
	return "sofa"
}
func (table) get() Furniture {
	return table{}
}

func (table) name() string {
	return "table"
}
func Shopping() {
	// 如果客户想要椅子
	// chair := chair{}
	// chair.get()
	// 如果客户想要沙发
	// sofa := sofa{}
	// sofa.get()
	// 如果客户想要桌子
	// table := table{}
	// table.get()

	factor := FurnitureFactor{}
	chair := factor.getFurniture("chair")
	fmt.Println("get a " + chair.name())
	table := factor.getFurniture("table")
	fmt.Println("get a " + table.name())
	sofa := factor.getFurniture("sofa")
	fmt.Println("get a " + sofa.name())
}
