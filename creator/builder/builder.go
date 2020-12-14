package builder

import (
	"fmt"
)

//BigDinner 大餐对象
type BigDinner struct {
	taste string
	price string
}

func (t *BigDinner) Cook(tips string) { //同上
	t.taste = tips
}

func (t *BigDinner) Pay(price string) {
	t.price = price
}

func (t BigDinner) Print() {
	fmt.Println("口味:", t.taste, "价格:", t.price)
}

//Builder 建造者的目的创建BigDinner对象
type Builder interface {
	Pay()
	Cook()
	CreateDinner() *BigDinner
}

//director 定义Builder抽象如何组合函数,创建BigDinner对象
type director struct {
	builder Builder
}

func (t *director) SetBuilder(b Builder) {
	t.builder = b
}

func (t director) Construct() *BigDinner {
	t.builder.Pay()                 //1.支付
	t.builder.Cook()                //2.烹饪
	return t.builder.CreateDinner() //产生一顿丰盛的晚宴
}

func NewDirector(b Builder) *director { //这里director 我故意小写,这样必须调用该函数才能生成director对象
	return &director{builder: b}
}

//简单版的构造函数

func Construct(b Builder) *BigDinner {
	b.Pay()
	b.Cook()
	return b.CreateDinner()
}

///-----------buidler的实现结构体↓↓↓↓↓↓↓↓---------------
//FamilyBuilder 接口的家庭实现
type FamilyBuilder struct {
	Dinner *BigDinner
}

func (t *FamilyBuilder) Pay() {
	t.Dinner.Pay("Free")
}

func (t *FamilyBuilder) Cook() {
	t.Dinner.Cook("清淡")
}

func (t *FamilyBuilder) CreateDinner() *BigDinner {
	return t.Dinner
}

//RestaurantBuilder 接口的餐馆实现
type RestaurantBuilder struct {
	Dinner *BigDinner
}

func (t *RestaurantBuilder) Pay() {
	t.Dinner.Pay("100cny")
}

func (t *RestaurantBuilder) Cook() {
	t.Dinner.Cook("酸辣")
}

func (t *RestaurantBuilder) CreateDinner() *BigDinner {
	return t.Dinner
}
