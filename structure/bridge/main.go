package main

import (
	"fmt"
)

//Shape 形状接口
type Shape interface {
	Show()
}

//Shape的实现类型
type Circle struct {
}

func (t Circle) Show() {
	fmt.Println("圆形")
}

//Shape的实现类型
type Square struct {
}

func (t Square) Show() {
	fmt.Println("正方形")
}

//Pen 笔接口
type Pen interface {
	Draw()
}

//Pen的实现类型
type BlackPen struct {
}

func (t BlackPen) Draw() {

	fmt.Println("黑色")
}

//Pen的实现类型
type YellowPen struct {
}

func (t YellowPen) Draw() {
	fmt.Println("黄色")
}

//Bridge 桥接Pen画笔和Shape形状
type Bridge struct {
	p Pen
	s Shape
}

func (t *Bridge) SetPen(p Pen) {
	t.p = p
}

func (t *Bridge) SetShape(s Shape) {
	t.s = s
}

func (t Bridge) Image() {
	t.p.Draw()
	t.s.Show()
	fmt.Println("------------")
}

func main() {
	b := Bridge{}
	b.SetPen(BlackPen{})
	b.SetShape(Circle{})
	b.Image()

	b.SetPen(BlackPen{})
	b.SetShape(Square{})
	b.Image()

	b.SetPen(YellowPen{})
	b.SetShape(Circle{})
	b.Image()

	b.SetPen(BlackPen{})
	b.SetShape(Circle{})
	b.Image()
	return
}
