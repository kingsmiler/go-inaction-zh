// 示例：如何在自定义结构体中申明方法。
package main

import (
	"fmt"
)

// 定义 user 结构体。
type user struct {
	name, email string
}

// 声明结构体方法，在方法前用括号定义一个结构体变量。
func (u user) notify() {
	fmt.Printf("Sending User Email To %s<%s>\n",
		u.name,
		u.email,
	)
}

// 如果要修改结构体的属性，则结构体变量为指针类型。
func (u *user) changeEmail(email string) {
	u.email = email
}

func main() {
	// 值类型变量
	bill := user{"Bill", "bill@email.com"}
	bill.notify()

	// 修改变量
	bill.changeEmail("bill@newdomain.com")
	bill.notify()

	// 指针类型变量
	lisa := &user{"Lisa", "lisa@email.com"}
	lisa.notify()

	// 修改变量
	lisa.changeEmail("lisa@newdomain.com")
	lisa.notify()
}
