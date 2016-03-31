package main

import (
    "fmt"

    "listing68/counters"
)

func main() {
    // 对于在其它包中类型，通常需要为它定义可以获取实例的方法
    counter := counters.New(10)

    fmt.Printf("Counter: %d\n", counter)
}
