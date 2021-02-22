package main

import "fmt"

//这里以计算器为例，包含两种功能 （加法、减法）
/*
//这种代码优点，如果只有这两种计算方法，编写简单容易理解。如果后续计算器包含很多计算方法(乘法、除法) Cal 结构体就需要不同的修改。违背了代码放开闭合原则。好的代码，应该是支持扩展，减少对原有代码修改，可以快速满足不同用户的需求。
type Cal struct{

}

func (t Cal)Add(a,b int) int{
	return a+b
}

func (t Cal)Sub(a,b int)int{
	return a-b
}
*/

//Strategy 定义策略接口
type Strategy interface {
	Calc(a, b int) int
}

//AddStrategy 加法策略
type AddStrategy struct {
}

func (t AddStrategy) Calc(a, b int) int {
	return a + b
}

//SubStrategy 减法策略
type SubStrategy struct {
}

func (t SubStrategy) Calc(a, b int) int {
	return a - b
}

//Calculator 计算器
type Calculator struct {
	s Strategy
}

func (t *Calculator) setStrategy(s Strategy) {
	t.s = s
}
func (t Calculator) GetResult(a, b int) int {
	return t.s.Calc(a, b)
}
func main() {
	add := AddStrategy{}
	sub := SubStrategy{}
	//计算器通过setStrategy设置不同策略，解耦了计算器和算法实现类
	cal := &Calculator{}
	cal.setStrategy(add)
	fmt.Println("加法策略结果:", cal.GetResult(1, 1))
	cal.setStrategy(sub)
	fmt.Println("减法策略结果:", cal.GetResult(1, 1))
}
