package main

import "fmt"

type ChatRoom struct {
	users []*User
}

func (t *ChatRoom) Connect(u *User) {
	fmt.Println(u.name, "-加入聊天室")
	t.users = append(t.users, u)
}

func (t *ChatRoom) SendMsg(fromWho *User, msg string) {
	for _, v := range t.users {
		if v != fromWho {
			v.Lister(fromWho, msg)
		}
	}
}

type User struct {
	name     string
	chatRoom *ChatRoom
}

func (t *User) Login() {
	t.chatRoom.Connect(t)
}

func (t *User) Talk(msg string) {
	t.chatRoom.SendMsg(t, msg)
}

func (t *User) Lister(fromWho *User, msg string) {
	fmt.Print("[", t.name, "]对话框:")
	fmt.Println(fromWho.name, "说:", msg)
}

func main() {
	chat := &ChatRoom{}
	userA := User{name: "A", chatRoom: chat}
	userB := User{name: "B", chatRoom: chat}

	userA.Login()
	userB.Login()

	userA.Talk("这里真好看")
	userB.Talk("是的")
}
