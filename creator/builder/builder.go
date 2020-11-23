package builder

import (
	"fmt"
)

//BigDinner 大餐对象
type BigDinner struct {
	where string
	taste string
}

func (t *BigDinner) Buy(tips string) { //因为要修改变量where值,所以t的类型是指针
	t.where = tips
}

func (t *BigDinner) Cook(tips string) { //同上
	t.taste = tips
}

func (t BigDinner) Print() { // 因为只是获取where和who的值,所以t是值类型
	fmt.Printf("购物:%s , 味道:%s \n", t.where, t.taste)
}

//Builder 建造者的目的创建BigDinner对象
type Builder interface {
	Buy()
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
	t.builder.Buy()                 //1.买菜
	t.builder.Cook()                //2.烹饪
	return t.builder.CreateDinner() //产生一顿丰盛的晚宴
}

func NewDirector(b Builder) *director { //这里director 我故意小写,这样必须调用该函数才能生成director对象
	return &director{builder: b}
}

///-----------buidler的实现结构体↓↓↓↓↓↓↓↓---------------
//FamilyBuilder 接口的家庭实现
type FamilyBuilder struct {
	Dinner *BigDinner
}

func (t *FamilyBuilder) Buy() {
	t.Dinner.Buy("超市")
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

func (t *RestaurantBuilder) Buy() {
	t.Dinner.Buy("批发")
}

func (t *RestaurantBuilder) Cook() {
	t.Dinner.Cook("酸辣")
}

func (t *RestaurantBuilder) CreateDinner() *BigDinner {
	return t.Dinner
}
