package main

import (
    "fmt"

    "listing71/entities"
)

func main() {
    u := entities.User{
        Name:  "Bill",
        // 其它包中的类型，小写的属性均为不可以访问的，email将报语法错误。
        email: "bill@email.com",
    }

    fmt.Printf("User: %v\n", u)
}
