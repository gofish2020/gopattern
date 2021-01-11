package main

import "fmt"

//Node 节点
type Node interface {
	Add(child Node)
	Print(space int)
}

type File struct {
	name string
}

//Add
func (t *File) Add(child Node) {
	//do nothing
}

//Print
func (t *File) Print(space int) {
	for i := 0; i < space; i++ {
		fmt.Print(" ")
	}
	fmt.Println(t.name)
}

type Folder struct {
	name  string
	child []Node
}

//Add add file or add folder
func (t *Folder) Add(child Node) {
	t.child = append(t.child, child)
}

//Print
func (t *Folder) Print(space int) {
	for i := 0; i < space; i++ {
		fmt.Print(" ")
	}
	fmt.Println(t.name)
	space++
	for _, v := range t.child {
		v.Print(space)
	}
}

func main() {

	diver := Folder{name: "D:"}
	//两个文件
	f1 := File{name: "1.txt"}
	f2 := File{name: "test.txt"}
	//文件夹
	fd1 := Folder{name: "文档"}
	ff1 := File{name: "2.txt"}
	ff2 := File{name: "3.txt"}
	ff3 := File{name: "4.txt"}
	fd1.Add(&ff1)
	fd1.Add(&ff2)
	fd1.Add(&ff3)
	//追加到父文件夹中
	diver.Add(&f1)
	diver.Add(&f2)
	diver.Add(&fd1)
	//打印目录层次
	diver.Print(0)
	return
}
