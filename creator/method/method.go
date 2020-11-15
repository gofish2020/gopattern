package method

import "fmt"

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

//-------------干饭人的分割线-----------
//食物工厂
type Factory interface {
	NewFood(k FoodKind) Food //这个函数是不是和简单工厂模式一模一样,而且分割线上面的代码也是一抹一样
}

//肉厂
type MeatFactory struct {
}

func (t MeatFactory) NewFood(k FoodKind) Food {
	return &meat{}
}

//水果厂
type FruitFactory struct {
}

func (t FruitFactory) NewFood(k FoodKind) Food {
	return &fruit{}
}

//蔬菜厂
type VegetableFactory struct {
}

func (t VegetableFactory) NewFood(k FoodKind) Food {
	return &vegetable{}
}

//坚果厂
type NutFactory struct {
}

func (t NutFactory) NewFood(k FoodKind) Food {
	return nut{}
}
