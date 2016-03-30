// 接口的使用示例之一.
package main

import (
	"fmt"
)

// 定义一个通知的接口.
type notifier interface {
	notify()
}

// 定义用户的结构体
type user struct {
	name  string
	email string
}

// 用户的结构体实现通知的接口.
func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n",
		u.name,
		u.email)
}

func main() {
	// 创建一个值类型的用户对象u
	u := user{"Bill", "bill@email.com"}

	// 值类型的对象u并没有实现通知接口,
	// 但如果直接调用的话也是没有问题的, Go会自动帮忙转换
	u.notify()

	// 但如果非直接调用, 参数传递导致类型改变,那方法调用就可能报错
	sendNotification(u)
}

// notifier 类型的话, 就必须使用指针; 可修改为User类型.
func sendNotification(n notifier) {
	n.notify()
}
