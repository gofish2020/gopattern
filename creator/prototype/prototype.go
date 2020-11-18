package prototype

type File struct {
	Context string
}

func (t *File) Clone() *File { // 返回 fake 不同步
	fake := *t
	return &fake
}
func (t *File) Clone1() *File { // 返回新File 不同步
	return &File{Context: t.Context}
}

func (t *File) Clone2() *File { //返回 t 同步改
	return t
}

func (t *File) Clone3() File { // 返回 *t 不同步
	return *t
}

func (t File) Clone4() *File { // 返回 &t 不同步
	return &t
}

func (t File) Clone5() File { // 返回 t 不同步
	return t
}

func NewFile() *File {
	return &File{Context: "源文件"}
}
