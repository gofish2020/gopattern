package main

import (
	"fmt"

	"github.com/gofish2020/gopattern/creator/singleton"

	"github.com/gofish2020/gopattern/creator/prototype"

	bld "github.com/gofish2020/gopattern/creator/builder"

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
	b := bld.NewDirector(&bld.FamilyBuilder{Dinner: new(bld.BigDinner)})
	b.Construct().Print()
	b.SetBuilder(&bld.RestaurantBuilder{Dinner: new(bld.BigDinner)})
	b.Construct().Print()
	fmt.Println("建造者模式--另一种简单写法")
	bld.Construct(&bld.FamilyBuilder{Dinner: new(bld.BigDinner)}).Print()
	bld.Construct(&bld.RestaurantBuilder{Dinner: new(bld.BigDinner)}).Print()

	fmt.Println("原型模式")
	file := prototype.NewFile()
	clone := file.Clone() // 替换这里的函数,查看不同方法的结果
	fmt.Printf("file.Context: %v \nclone.Context: %v \n", file.Context, clone.Context)
	clone.Context = "克隆文件"
	fmt.Printf("file.Context: %v \nclone.Context: %v \n", file.Context, clone.Context)

	fmt.Println("单例模式")

	h1 := singleton.GetHungry()
	h2 := singleton.GetHungry()
	if h1 == h2 {
		fmt.Println("h1==h2")
	}

	l1 := singleton.GetLazy()
	l2 := singleton.GetLazy()
	if l1 == l2 {
		fmt.Println("l1==l2")
	}

	lo1 := singleton.GetOnce()
	lo2 := singleton.GetOnce()
	if lo1 == lo2 {
		fmt.Println("lo1==lo2")
	}
}
