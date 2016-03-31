package main

import (
	"fmt"

	"listing64/counters"
)

func main() {
	// 访问其它包中未暴露的类型
	counter := counters.alertCounter(10)

	// ./listing64.go:15: cannot refer to unexported name
	//                                         counters.alertCounter
	// ./listing64.go:15: undefined: counters.alertCounter

	fmt.Printf("Counter: %d\n", counter)
}
