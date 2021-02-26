package main

import "fmt"

type Product interface {
	Design()
	Develop()
	Test()
	Release()
}

type ConcreteProductA struct {
}

func (t ConcreteProductA) Design() {
	fmt.Println("设计A")
}
func (t ConcreteProductA) Develop() {
	fmt.Println("开发A")
}
func (t ConcreteProductA) Test() {
	fmt.Println("测试A")
}
func (t ConcreteProductA) Release() {
	fmt.Println("发布A")
}

type ConcreteProductB struct {
}

func (t ConcreteProductB) Design() {
	fmt.Println("设计B")
}
func (t ConcreteProductB) Develop() {
	fmt.Println("开发B")
}
func (t ConcreteProductB) Test() {
	fmt.Println("测试B")
}
func (t ConcreteProductB) Release() {
	fmt.Println("发布B")
}

//模板类　-- 定义一个通用的模板方法
type Template struct {
	Product
}

func (t *Template) SetProduct(p Product) {
	t.Product = p

}
func (t Template) TemplateMethod() { //产品的发布流程都是一样的，每个产品的内部细节并不相同，留给各个产品自己去实现。。这个模式类似于门面模式代码
	t.Product.Design()
	t.Product.Develop()
	t.Product.Test()
	t.Product.Release()
}
func main() {
	a := ConcreteProductA{}
	b := ConcreteProductB{}
	t := Template{a}
	t.TemplateMethod()
	t.SetProduct(b)
	t.TemplateMethod()
}
