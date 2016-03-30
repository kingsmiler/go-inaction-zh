// 示例演示不能使用指针的场景.
package main

import "fmt"

// duration 是基本的整形类型.
type duration int

// 格式化显示 duration 类型的值.
func (d *duration) pretty() string {
	return fmt.Sprintf("Duration: %d", *d)
}

func main() {
	duration(42).pretty()
	// 换句话说, 对基本类型的扩展方法, 不要使用指针.
	// ./listing46.go:15: cannot call pointer method on duration(42)
	// ./listing46.go:15: cannot take the address of duration(42)
}
