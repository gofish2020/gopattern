package builder

import (
	"fmt"
)

//定义造饭的行为
type Builder interface {
	Hungry()
	Buy()
	Cook()
	Eat()
	Wash()
}

//由director决定整个造饭的过程
type director struct {
	builder Builder
}

func (t *director) SetBuilder(b Builder) {
	t.builder = b
}

func (t *director) Dinner() {
	t.builder.Hungry() //1.饿了
	t.builder.Buy()    //2.买菜
	t.builder.Cook()   //3.烧菜
	t.builder.Eat()    //4.开吃
	t.builder.Wash()   //5.刷碗
}

func NewDirector(b Builder) *director { //这里director 我故意小写,这样必须调用该函数才能生成director对象
	return &director{builder: b}
}

///-----------buidler的实现结构体↓↓↓↓↓↓↓↓---------------
//家庭造饭行为
type Family struct {
}

func (t Family) Hungry() {
	fmt.Println("今天在家吃~~")
}

func (t Family) Buy() {
	fmt.Println("[家]超市买菜")
}

func (t Family) Cook() {
	fmt.Println("[家]小锅炒菜")
}

func (t Family) Eat() {
	fmt.Println("[家]一个人吃")
}

func (t Family) Wash() {
	fmt.Println("[家]一个人洗")
}

//餐馆造饭行为
type Restaurant struct {
}

func (t Restaurant) Hungry() {
	fmt.Println("今天下馆子~~")
}

func (t Restaurant) Buy() {
	fmt.Println("[餐馆]批发市场买菜")
}

func (t Restaurant) Cook() {
	fmt.Println("[餐馆]大锅爆炒")
}

func (t Restaurant) Eat() {
	fmt.Println("[餐馆]客人用餐")
}

func (t Restaurant) Wash() {
	fmt.Println("[餐馆]后厨刷碗")
}
