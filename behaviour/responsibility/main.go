package main

import (
	"fmt"
)

type ChainResponsibility interface {
	Approve(money int)
}

//[0,100)
type Staff struct {
	next ChainResponsibility
}

func (t Staff) Approve(money int) {
	if money >= 0 && money < 100 {
		fmt.Println("Staff审批完成")
	} else {
		fmt.Println("Staff无权限审批，传递给上一级")
		t.next.Approve(money)
	}
	return
}

//[100,1000)
type Manager struct {
	next ChainResponsibility
}

func (t Manager) Approve(money int) {

	if money >= 100 && money < 1000 {
		fmt.Println("Manager审批完成")
	} else {
		fmt.Println("Manager无权限审批，传递给上一级")
		t.next.Approve(money)
	}
	return
}

// >=1000
type CEO struct {
}

func (t CEO) Approve(money int) {
	if money >= 1000 {
		fmt.Println("CEO审批完成")
	} else {
		fmt.Println("CEO驳回审批")
	}
	return
}

func main() {
	//定义Staff对象，并持有Manager,定义Manager对象,并持有CEO,形成链表
	s := &Staff{next: &Manager{next: &CEO{}}}
	s.Approve(100)
	fmt.Println("--------------")
	s.Approve(1000)
	return
}
