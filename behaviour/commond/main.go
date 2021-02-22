package main

import "fmt"

//------------命令接口-------------
type Command interface {
	Do()
	UnDo()
}

//SwitchCommand　开关控制
type SwitchCommand struct {
	d Device
}

func (t SwitchCommand) Do() {
	t.d.Open()
}

func (t SwitchCommand) UnDo() {
	t.d.Close()
}

//VolumnCommand 音量控制
type VolumnCommand struct {
	d Device
}

func (t VolumnCommand) Do() {
	t.d.VolumnUp()
}

func (t VolumnCommand) UnDo() {
	t.d.VolumnDown()
}

//ChannelCommand 频道控制
type ChannelCommand struct {
	d Device
}

func (t ChannelCommand) Do() {
	t.d.ChannelUp()
}

func (t ChannelCommand) UnDo() {
	t.d.ChannelDown()
}

//----------------设备接口-------------

type Device interface {
	Open()
	Close()
	VolumnUp()
	VolumnDown()
	ChannelUp()
	ChannelDown()
}

//-------电视机-------
type TV struct {
}

func (t TV) Open() {
	fmt.Println("打开TV")
}

func (t TV) Close() {
	fmt.Println("关闭TV")
}

func (t TV) VolumnUp() {
	fmt.Println("TV音量+")
}

func (t TV) VolumnDown() {
	fmt.Println("TV音量-")
}

func (t TV) ChannelUp() {
	fmt.Println("TV频道+")
}

func (t TV) ChannelDown() {
	fmt.Println("TV频道-")
}

//-------收音机------
type Radio struct {
}

func (t Radio) Open() {
	fmt.Println("打开Radio")
}

func (t Radio) Close() {
	fmt.Println("关闭Radio")
}

func (t Radio) VolumnUp() {
	fmt.Println("Radio音量+")
}

func (t Radio) VolumnDown() {
	fmt.Println("Radio音量-")
}

func (t Radio) ChannelUp() {
	fmt.Println("Radio频道+")
}

func (t Radio) ChannelDown() {
	fmt.Println("Radio频道-")
}

//-----控制类------
type Controller struct {
	switchC  Command
	volumnC  Command
	channelC Command
}

func (t Controller) buttonOkHold() {
	fmt.Print("长按Ok键...")
	t.switchC.Do()
}
func (t Controller) buttonOkClick() {
	fmt.Print("单击Ok键...")
	t.switchC.UnDo()
}

func (t Controller) buttonUpClick() {
	fmt.Print("单击↑按键...")
	t.volumnC.Do()
}
func (t Controller) buttonDownClick() {
	fmt.Print("单击↓按键...")
	t.volumnC.UnDo()
}

func (t Controller) buttonRightClick() {
	fmt.Print("单击→按键...")

	t.channelC.Do()
}
func (t Controller) buttonLeftClick() {
	fmt.Print("单击←按键...")
	t.channelC.UnDo()
}

func main() {

	tv := TV{}

	c := Controller{
		switchC:  SwitchCommand{d: tv},
		volumnC:  VolumnCommand{d: tv},
		channelC: ChannelCommand{d: tv},
	}
	c.buttonOkHold()
	c.buttonUpClick()
	c.buttonRightClick()
	c.buttonOkClick()

	radio := Radio{}
	fmt.Println("-----------")
	c = Controller{
		switchC:  SwitchCommand{d: radio},
		channelC: VolumnCommand{d: radio},  //调整按键功能【左右】键负责音量
		volumnC:  ChannelCommand{d: radio}, //调整按键功能【上下】键负责频道
	}

	c.buttonOkHold()
	c.buttonUpClick()
	c.buttonRightClick()
	c.buttonOkClick()

	return
}
