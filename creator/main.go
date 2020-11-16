package main

import (
	"fmt"

	"github.com/gofish2020/gopattern/creator/abstract"

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

	fmt.Println("抽象工厂模式")
	first := &abstract.FirstFactory{} //这里用的&符号,因为实现的指针类型
	first.NewFood().Eat()
	first.NewDrug().Take()
	second := &abstract.SecondFactory{}
	second.NewFood().Eat()
	second.NewDrug().Take()

}
