package main

import (
	"fmt"
)

//Modem 猫，负责访问网页
type Modem struct {
}

func (t Modem) Access(url string) {
	fmt.Println("正常访问：" + url)
}

//Proxy 代理了Modem的网络访问功能
type Proxy struct {
	blacklist []string
	m         Modem
}

//FilterAccess 可以看到这里增加了控制，限制某些网站的访问
func (t Proxy) FilterAccess(url string) {
	for _, v := range t.blacklist {
		if v == url {
			fmt.Println("禁止访问：" + url)
			return
		}
	}
	t.m.Access(url) //最终访问网络还是使用的Modem
}

func main() {

	p := Proxy{blacklist: []string{"www.bilibili.com", "www.youku.com"}, m: Modem{}}
	p.FilterAccess("www.baidu.com")
	p.FilterAccess("www.bilibili.com")
	return
}
