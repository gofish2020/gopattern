package main

import "fmt"

//Recorder 记录类
type Recorder struct {
	data []string
}

func (t *Recorder) Add(str string) {
	t.data = append(t.data, str)
}

func (t *Recorder) Remove(index int) {
	if t.data == nil || index < 0 || index >= len(t.data) {
		return
	}
	t.data = append(t.data[:index], t.data[index+1:]...)
}

//Iterator 迭代类，通过迭代可以取出Recorder中的所有的记录，但是又不会影响原有类结构代码
type Iterator struct {
	*Recorder
	index int
}

func (t *Iterator) Reset() {
	t.index = 0
}

func (t *Iterator) Next() string {
	str := t.data[t.index]
	t.index++
	return str
}

func (t *Iterator) HasNext() bool {
	if t.index < 0 || t.index >= len(t.data) {
		return false
	}
	return true
}

func main() {
	rec := Recorder{}
	rec.Add("first")
	rec.Add("second")
	rec.Add("third")
	rec.Add("fourth")
	rec.Add("fifth")

	iter := Iterator{index: 0, Recorder: &rec}
	for iter.HasNext() {
		fmt.Println(iter.Next())
	}
	iter.Reset()
	fmt.Println("----------")
	rec.Remove(1) //iter.Remove(1) 效果等价。如果这里的Recorder不是指针，两个的作用不一样的，因为操作的对象是两个不同的对象
	for iter.HasNext() {
		fmt.Println(iter.Next())
	}
}
