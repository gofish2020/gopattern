package main

import (
	"fmt"

	"github.com/gofish2020/gopattern/creator/method"

	"github.com/gofish2020/gopattern/creator/simple"
)

func main() {
	fmt.Println("简单工厂模式")
	simple.NewFood(simple.MeatKind).Eat()
	simple.NewFood(simple.FruitKind).Eat()
	simple.NewFood(simple.VegetableKind).Eat()
	simple.NewFood(simple.NutKind).Eat()

	fmt.Println("工厂方法模式")
	method.MeatFactory{}.NewFood(method.MeatKind).Eat()
	method.FruitFactory{}.NewFood(method.FruitKind).Eat()
	method.VegetableFactory{}.NewFood(method.VegetableKind).Eat()
	method.NutFactory{}.NewFood(method.NutKind).Eat()
}
