package main

import (
	"fmt"
)

type Boss struct {
}

func (t Boss) Buy() {
	fmt.Println("买菜")
}

type Waiter struct {
}

func (t Waiter) Wash() {
	fmt.Println("洗菜")
}

type Chef struct {
}

func (t Chef) Cook() {
	fmt.Println("烹饪")
}

type Cleaner struct {
}

func (t Cleaner) Clean() {
	fmt.Println("清理")
}

type Restaurant struct {
	b  Boss
	w  Waiter
	c  Chef
	cl Cleaner
}

func (t Restaurant) EnjoyService() {
	t.b.Buy()
	t.w.Wash()
	t.c.Cook()
	t.cl.Clean()
}

func main() {

	customer := Restaurant{b: Boss{}, w: Waiter{}, c: Chef{}, cl: Cleaner{}}
	customer.EnjoyService() //客人享受服务，具体里面是怎样的一种关系，客人是不需要关心
	return
}
