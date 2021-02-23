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

type Parser struct {
	add   Expression
	sub   Expression
	mul   Expression
	power Expression
}

// 1 + 2 * 3
func (t Parser) interperter(a, b, c int) int { //强调行为模式　－－　解释的过程(解释器模式)　行为已经组织好了，只需要按步解释
	return t.add.Interpreter(a, t.mul.Interpreter(b, c))
}

func main() {
	add := AddExpression{}
	sub := SubExpression{}
	mul := MulExpression{}
	power := PowerExpression{}
	p := Parser{add: add, sub: sub, mul: mul, power: power}
	fmt.Println(p.interperter(1, 2, 3))

	//强调结构模式　－－　构造的过程(组合模式) 自己去组织行为
	//3^4 +1 -1
	fmt.Println(sub.Interpreter(add.Interpreter(power.Interpreter(3, 4), 1), 1))

}
