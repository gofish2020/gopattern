package main

import (
	"fmt"
)

//Expression 表达式
type Expression interface {
	Interpreter(a, b int) int
}

//加法
type AddExpression struct {
}

func (t AddExpression) Interpreter(a, b int) int {
	return a + b
}

//减法
type SubExpression struct {
}

func (t SubExpression) Interpreter(a, b int) int {
	return a - b
}

//乘法
type MulExpression struct {
}

func (t MulExpression) Interpreter(a, b int) int {
	return a * b
}

//幂
type PowerExpression struct {
}

func (t PowerExpression) Interpreter(a, b int) int {
	result := 1
	for i := 0; i < b; i++ {
		result *= a
	}
	return result
}

type CalcParser struct {
	expr1 Expression
	expr2 Expression
}

// 1 + 2 * 3
func (t CalcParser) interperter(a, b, c int) int { //强调行为模式　－－　解释的过程(解释器模式)　行为已经组织好了，只需要按步解释
	return t.expr1.Interpreter(a, t.expr2.Interpreter(b, c))
}

func main() {
	add := AddExpression{}
	power := PowerExpression{}
	p := CalcParser{expr1: add, expr2: power}
	fmt.Println(p.interperter(1, 2, 3))
	p.expr2 = MulExpression{}
	fmt.Println(p.interperter(1, 2, 3))
}
