package main

import (
	"fmt"
)

type Showable interface {
	Show()
}

//普通的女生
type Girl struct {
}

func (t *Girl) Show() {
	fmt.Print("素颜女生")
}

//MakeUpDecorator 化妆
type MakeUpDecorator struct {
	Showable
}

func (t *MakeUpDecorator) Show() {
	fmt.Print("打粉底(")
	t.Showable.Show()
	fmt.Print(")")
}

//LispDecorator 口红
type LispDecorator struct {
	Showable
}

func (t *LispDecorator) Show() {
	fmt.Print("涂口红(")
	t.Showable.Show()
	fmt.Print(")")
}

func main() {

	girl := Girl{}
	girl.Show()
	fmt.Println("")

	l := LispDecorator{Showable: &MakeUpDecorator{Showable: &Girl{}}}
	l.Show()
	fmt.Println("")
	return
}
