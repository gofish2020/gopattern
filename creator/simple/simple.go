package simple

import (
	"fmt"
)

type FoodKind int

const (
	MeatKind FoodKind = iota
	FruitKind
	VegetableKind
	NutKind
)

//Food 食物接口
type Food interface {
	Eat()
}

//肉-食物接口的实现
type meat struct {
}

func (t *meat) Eat() {
	fmt.Println("Eat meat")
}

//水果-食物接口的实现
type fruit struct {
}

func (t *fruit) Eat() {
	fmt.Println("Eat fruit")
}

//蔬菜-食物接口的实现
type vegetable struct {
}

func (t *vegetable) Eat() {
	fmt.Println("Eat vegetable")
}

//坚果-食物接口的实现
type nut struct {
}

func (t nut) Eat() {
	fmt.Println("Eat nut")
}

//NewFood 根据需要的类型,返回实现对象
func NewFood(k FoodKind) Food {
	switch k {
	case MeatKind:
		return &meat{}
	case FruitKind:
		return &fruit{}
	case VegetableKind:
		return &vegetable{}
	case NutKind:
		return nut{} // &nut{}  两种写法都可以
	}
	return nil
}
