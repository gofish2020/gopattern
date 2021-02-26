package main

import "fmt"

/*
	这里以文档编辑器为例
*/

//History 历史类
type History struct {
	data string
}

func (t History) GetHistory() string {
	return t.data
}

//Doc 文档类
type Doc struct {
	body string
}

func (t *Doc) SetBody(body string) {
	t.body = body
}
func (t *Doc) GetBody() string {
	return t.body
}

func (t *Doc) CreateHistory() History {
	return History{data: t.body}
}

func (t *Doc) RecoverHistory(h History) {
	t.body = h.data
}

//Editor 编辑器类
type Editor struct {
	d     Doc       //持有文档对象
	his   []History //持有历史列表
	index int
}

//Write 写文字
func (t *Editor) Write(txt string) {
	fmt.Print("写操作>>")
	t.d.SetBody(t.d.GetBody() + txt)
	t.BackUp() //备份当前记录
	t.Show()
}

//Delete 清空内容
func (t *Editor) Delete() {
	fmt.Print("清空操作>>")
	t.d.SetBody("")
	t.BackUp()
	t.Show()
}

//BackUp 创建备份
func (t *Editor) BackUp() {
	t.his = append(t.his, t.d.CreateHistory()) //文档对象备份记录，保存到his中
	t.index++
}

//Ctrl+z 恢复上次记录
func (t *Editor) Undo() {
	if t.index <= 0 {
		return
	}
	fmt.Print("回撤操作>>")
	t.index--
	t.d.RecoverHistory(t.his[t.index]) //　his中记录，恢复到文档对象中
	t.Show()
}

func (t *Editor) Show() { //　打印内容
	fmt.Println(t.d.body)
}

func main() {
	fmt.Println("<<开始编写文章>>")
	editor := Editor{index: -1, d: Doc{body: ""}}
	editor.Write("我有一个梦想,")
	editor.Write("改变世界!!!")
	editor.Delete()
	editor.Undo()
	editor.Undo()
	return
}
