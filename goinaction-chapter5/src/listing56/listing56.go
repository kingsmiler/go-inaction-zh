// 子类使用父类实现的接口示例。
package main

import (
    "fmt"
)

// 接口
type notifier interface {
    notify()
}

// 父类
type user struct {
    name, email string
}

// 父类实现接口
func (u *user) notify() {
    fmt.Printf("Sending user email to %s<%s>\n", u.name, u.email)
}

// 子类
type admin struct {
    user
    level string
}

func main() {
    // 创建子类实例。
    ad := admin{
        user: user{
            name:  "john smith",
            email: "john@yahoo.com",
        },
        level: "super",
    }

    // 父类的接口实现被子类继承，方法能被直接使用。
    sendNotification(&ad)
}

// 统一的接口调用方法。
func sendNotification(n notifier) {
    n.notify()
}
