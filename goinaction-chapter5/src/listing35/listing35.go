// 本代码演示如何将字节缓冲区内容输出.
package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	// 定义字节缓冲区.
	var b bytes.Buffer

	// 向缓冲区中填充字符串, 字节数组;
	// Write 为缓冲区的方法, 由于需要修改内容, 要求的对象是指针;
	// 结构体中的自身的方法自动转换值类型与引用类型;
	// 对象为值类型还是引用类型在定义时已确定.
	b.Write([]byte("Hello"))

	// Fprintf 用于向一个输出通道中追加字符串.
	fmt.Fprintf(&b, " World!")

	// 将后者的内容拷贝到前者.
	io.Copy(os.Stdout, &b)

	// 新增的演示示例
	useWithPointer()
}

// 使用指针引用类型的演示
func useWithPointer()  {
	b := &bytes.Buffer{}
	b.Write([]byte("Hello"))

	fmt.Fprintf(b, " World!")
	io.Copy(os.Stdout, b)
}
