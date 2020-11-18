package main

import (
	"fmt"

	"github.com/gofish2020/gopattern/creator/prototype"

	"github.com/gofish2020/gopattern/creator/builder"

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

	fmt.Println("建造者模式")
	b := builder.NewDirector(builder.Family{})
	b.Dinner()
	b.SetBuilder(builder.Restaurant{})
	b.Dinner()

	fmt.Println("原型模式")
	file := prototype.NewFile()
	clone := file.Clone() // 替换这里的函数,查看不同方法的结果
	fmt.Printf("file.Context: %v \nclone.Context: %v \n", file.Context, clone.Context)
	clone.Context = "克隆文件"
	fmt.Printf("file.Context: %v \nclone.Context: %v \n", file.Context, clone.Context)

}
