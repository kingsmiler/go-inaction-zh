// Golang 中的继承实现，将一个类型嵌入到另一类型中。
package main

import (
    "fmt"
)

// 定义一个父类
type user struct {
    name  string
    email string
}

// 父类的方法
func (u *user) notify() {
    fmt.Printf("Sending user email to %s<%s>\n", u.name, u.email)
}

// 定义子类
type admin struct {
    user // Embedded Type
    level string
}

func main() {
    // 获取一个子类实例。
    ad := admin{
        // 作为成员字段，嵌入的类型不能使用指针，这是加上&将报错
        user: user{
            name:  "john smith",
            email: "john@yahoo.com",
        },
        level: "super",
    }

    // 以属性的方式来访问父类的属性与方法
    fmt.Println(ad.user.name)
    ad.user.notify()

    // 子类也可以直接访问
    fmt.Println(ad.name)
    ad.notify()

    // 除了对象初始化时，父类不能使用指针外，
    // 子类可以随意操作父类的属性与方法。
    changeName1(ad)
    fmt.Println("changeName1:", ad.name)
    changeName2(&ad)
    fmt.Println("changeName2:", ad.name)
}

func changeName1(user admin)  {
    user.name = "robot wang"
}

func changeName2(user *admin)  {
    user.name = "robot wang"
}