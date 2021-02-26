package main

import "fmt"

/*
	观察者模式：当[被观察者]有什么情况，通知[给观察者]。很类似于系统回调。该模式和中介模式代码实现上也有类似之处
*/

//Subject　被观察者抽象
type Subject interface {
	Add(o Observer)
	Remove(o Observer)
	Notify()
}

type ConcreteSubject struct {
	obs []Observer
}

func (t *ConcreteSubject) Add(o Observer) { //注册观察者
	t.obs = append(t.obs, o)
}
func (t *ConcreteSubject) Remove(o Observer) { //删除观察者
	for k, v := range t.obs {
		if v == o {
			t.obs = append(t.obs[0:k], t.obs[k+1:]...)

			return
		}
	}
	return
}

func (t *ConcreteSubject) Notify() { //通知观察者
	for _, v := range t.obs {
		v.Response("你有一条消息")
	}
}

//Observer　观察者
type Observer interface {
	Response(msg string)
}

//ConcreteObserver　实例化的观察者
type ConcreteObserver struct {
	username string
}

func (t *ConcreteObserver) Response(msg string) {
	fmt.Println(t.username, ":", msg)
}

func main() {
	ob1 := &ConcreteObserver{username: "小明"}
	ob2 := &ConcreteObserver{username: "小汪"}
	sub := &ConcreteSubject{}
	sub.Notify()
	sub.Add(ob1)
	sub.Add(ob2)
	sub.Notify()
	fmt.Println("移除观察者:", ob1.username)
	sub.Remove(ob1)
	sub.Notify()
	return
}
