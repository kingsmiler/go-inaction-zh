package main

import (
    "fmt"

    "listing74/entities"
)

func main() {
    a := entities.Admin{
        Rights: 10,
    }

    // 大写字母开头的方法或属性为公有，可直接访问。
    a.Name = "Bill"
    a.Email = "bill@email.com"

    fmt.Printf("User: %v\n", a)
}
