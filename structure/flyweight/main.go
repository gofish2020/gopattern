package main

import (
	"fmt"
)

//FlyWeight 接口
type FlyWeight interface {
	Operation()
}

//ConcreteFlyWeight FlyWeight接口的实现
type ConcreteFlyWeight struct {
	Key string
}

func (t ConcreteFlyWeight) Operation() {
	fmt.Println(t.Key)
}

//FlyWeightFactory 共享资源，核心在于map这个数据结构
type FlyWeightFactory struct {
	flys map[string]FlyWeight
}

func (t FlyWeightFactory) GetFlyWeight(key string) FlyWeight {
	if f, ok := t.flys[key]; !ok {
		f = ConcreteFlyWeight{Key: key}
		t.flys[key] = f
		return f
	} else {
		return f
	}
}

//Len 打印对象个数
func (t FlyWeightFactory) Len() int {
	return len(t.flys)
}

func main() {

	factory := FlyWeightFactory{flys: make(map[string]FlyWeight)}
	factory.GetFlyWeight("A").Operation()
	factory.GetFlyWeight("A").Operation()
	factory.GetFlyWeight("B").Operation()
	factory.GetFlyWeight("C").Operation()
	fmt.Println(factory.Len())
	return
}
