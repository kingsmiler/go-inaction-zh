// 子类对父类接口实现的覆盖。
package main

import (
    "fmt"
)

type notifier interface {
    notify()
}

// 父类
type user struct {
    name, email string
}

// 父类实现
func (u *user) notify() {
    fmt.Printf("Sending user email to %s<%s>\n", u.name, u.email)
}

// 子类
type admin struct {
    user
    level string
}

// 子类实现
func (a *admin) notify() {
    fmt.Printf("Sending admin email to %s<%s>\n", a.name, a.email)
}

func main() {
    ad := admin{
        user: user{
            name:  "john smith",
            email: "john@yahoo.com",
        },
        level: "super",
    }

    // 直接调用时，使用的是子类自身的实现
    sendNotification(&ad)

    // 通过使用父类名称来调用父类的实现
    ad.user.notify()

    // 直接调用时，使用的是子类自身的实现
    ad.notify()
}

func sendNotification(n notifier) {
    n.notify()
}
