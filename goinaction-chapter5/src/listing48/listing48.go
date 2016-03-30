// 通过接口实现多态行为的示例.
package main

import (
	"fmt"
)

// 接口
type notifier interface {
	notify()
}

// 结构体
type user struct {
	name  string
	email string
}

// 结构体 user 实现 notifier 接口.
func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n",
		u.name,
		u.email)
}

// 管理信息结构体
type admin struct {
	name  string
	email string
}

// 实现接口 notifier
func (a *admin) notify() {
	fmt.Printf("Sending admin email to %s<%s>\n",
		a.name,
		a.email)
}

func main() {
	bill := user{"Bill", "bill@email.com"}
	sendNotification(&bill)

	lisa := admin{"Lisa", "lisa@email.com"}
	sendNotification(&lisa)
}

// 为接口定义统一的调用方法.
func sendNotification(n notifier) {
	n.notify()
}
