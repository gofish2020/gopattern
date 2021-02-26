package main

import "fmt"

type Switcher struct { //将状态托管给State进行处理
	s State
}

func (t *Switcher) GetState() State {
	return t.s
}

func (t *Switcher) SetState(s State) {
	t.s = s
}

func (t *Switcher) TurnOn() {
	t.s.SwitchOn(t)
}

func (t *Switcher) TurnOff() {
	t.s.SwitchOff(t)
}

//State 不再是单纯的一个状态标识符，抽象为一个接口
type State interface { // State 负责控制Switcher开关状态
	SwitchOn(c *Switcher)
	SwitchOff(c *Switcher)
}

type On struct {
}

func (t On) SwitchOn(c *Switcher) {
	fmt.Println("WARNNING，无需再开启")
}

func (t On) SwitchOff(c *Switcher) {
	fmt.Println("OK...关闭成功")
	c.SetState(Off{})
}

type Off struct {
}

func (t Off) SwitchOn(c *Switcher) {
	fmt.Println("OK...开启成功")
	c.SetState(On{})

}

func (t Off) SwitchOff(c *Switcher) {
	fmt.Println("WARNNING，无需再关闭")
}

func main() {
	s := Switcher{s: Off{}} //默认关闭状态
	s.TurnOn()
	s.TurnOn()
	s.TurnOff()
	s.TurnOff()

}
