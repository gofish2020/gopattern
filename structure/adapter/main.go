package main

import (
	"fmt"
)

type USB interface {
	Work()
}

//对外接口
func Charge(usb USB) {
	usb.Work()
}

//TypeC 没有实现USB接口,不可以充电
type TypeC struct {
}

func (t TypeC) IPhone() {
	fmt.Println("Charge IPhone")
}

//AdapterTypeC 实现USB接口,可以充电
type AdapterTypeC struct { //两个结构体组合一起,形成新的结构体
	TypeC
}

func (t AdapterTypeC) Work() {
	t.IPhone()
}

//上面的另一种写法
type AdapterTypeC1 struct { //TypeC是新接口的一个成员变量
	tc TypeC
}

func (t AdapterTypeC1) Work() {
	t.tc.IPhone()
}

//Micro 实现USB接口,可以充电
type Micro struct {
}

func (t Micro) Work() {
	fmt.Println("Charge Android")
}

func main() {
	Charge(Micro{})
	Charge(AdapterTypeC{})
	Charge(AdapterTypeC1{tc: TypeC{}})
	//	Charge(TypeC{}) //只能给实现USB接口的供电
	return
}
