package abstract

import (
	"fmt"
)

//Food -- 抽象的食物接口
type Food interface {
	Eat()
}

//Drug -- 抽象的药物接口
type Drug interface {
	Take()
}

//Factory 负责生产药物类和食物类
type Factory interface {
	NewFood() Food
	NewDrug() Drug
}

//-----Food 实现结构体----
type meat struct { // 肉
}

func (t meat) Eat() {
	fmt.Println("Eat meat")
}

type fruit struct { //水果
}

func (t fruit) Eat() {
	fmt.Println("Eat fruit")
}

// -----Drug 实现结构体-----

type feverdrug struct { //发烧药

}

func (t feverdrug) Take() {
	fmt.Println("Take fever drug")
}

type colddrug struct { //感冒药
}

func (t colddrug) Take() {
	fmt.Println("Take cold drug")
}

//------ Factory 实现结构体 ------

type FirstFactory struct { // 第一个工厂

}

func (t *FirstFactory) NewFood() Food {
	return meat{}
}

func (t *FirstFactory) NewDrug() Drug {
	return feverdrug{}
}

type SecondFactory struct { // 第二个工厂

}

func (t *SecondFactory) NewFood() Food {
	return fruit{}
}

func (t *SecondFactory) NewDrug() Drug {
	return colddrug{}
}
