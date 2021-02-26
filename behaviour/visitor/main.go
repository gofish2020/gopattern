package main

import (
	"fmt"
	"reflect"
)

//----------------------访问者类-----------------

type Visitor interface {
	BuyApple(a Apple)
	BuyBanana(b Banana)
	BuyStawberry(c Strawberry)
}

//VipVisitor 会员，打折购买
type VipVisitor struct {
	discount float32
}

func (t VipVisitor) BuyApple(a Apple) {
	fmt.Println("苹果原价:", a.GetPrice(), "会员价格:", t.discount*float32(a.GetPrice()))
}

func (t VipVisitor) BuyBanana(b Banana) {
	fmt.Println("香蕉原价:", b.GetPrice(), "会员价格购:", t.discount*float32(b.GetPrice()))
}

func (t VipVisitor) BuyStawberry(c Strawberry) {
	fmt.Println("草莓原价:", c.GetPrice(), "会员价格", t.discount*float32(c.GetPrice()))
}

//NormalVisitor 普通用户，原价购买
type NormalVisitor struct {
}

func (t NormalVisitor) BuyApple(a Apple) {
	fmt.Println("苹果原价:", a.GetPrice(), "零售价格:", a.GetPrice())
}

func (t NormalVisitor) BuyBanana(b Banana) {
	fmt.Println("香蕉原价:", b.GetPrice(), "零售价格:", b.GetPrice())
}

func (t NormalVisitor) BuyStawberry(c Strawberry) {
	fmt.Println("草莓原价:", c.GetPrice(), "零售价格:", c.GetPrice())
}

//-------------------------水果接口-------------------
type Fruit interface {
	Accept(v Visitor)
}

//Apple 苹果类
type Apple struct {
	price int
}

func (t Apple) Accept(v Visitor) {
	v.BuyApple(t)
}
func (t Apple) GetPrice() int {
	return t.price
}

//Banana 香蕉类
type Banana struct {
	price int
}

func (t Banana) Accept(v Visitor) {
	v.BuyBanana(t)
}
func (t Banana) GetPrice() int {
	return t.price
}

//Strawberry 草莓类
type Strawberry struct {
	price int
}

func (t Strawberry) Accept(v Visitor) {
	v.BuyStawberry(t)
}
func (t Strawberry) GetPrice() int {
	return t.price
}
func main() {

	//VipVisitor 根据水果种类，调用不同的Buy方法购买商品
	fmt.Println("---------已知水果类型,直接购买----------")
	apple := Apple{price: 1}
	banana := Banana{price: 1}
	strawberry := Strawberry{price: 1}
	v := VipVisitor{discount: 0.5}
	v.BuyApple(apple)
	v.BuyBanana(banana)
	v.BuyStawberry(strawberry)

	fmt.Println("------反射方式，识别水果类型----------")
	//如果是下面这种如何实现呢？
	f := make([]Fruit, 3)
	f[0] = apple
	f[1] = banana
	f[2] = strawberry
	for _, fruit := range f {
		typ := reflect.TypeOf(fruit) //当水果类型越来越多，这里的逻辑需要不断的修改添加，通用性不好，有没有通用性更好的方法，一次完成，后续再也不用修改了。
		switch typ.Name() {
		case "Apple":
			v.BuyApple(fruit.(Apple))
		case "Banana":
			v.BuyBanana(fruit.(Banana))
		case "Strawberry":
			v.BuyStawberry(fruit.(Strawberry))
		}
	}

	fmt.Println("-----------访问者模式--------------")

	for _, fruit := range f {
		fruit.Accept(v)
	}
}
