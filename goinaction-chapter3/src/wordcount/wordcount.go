// 输入输出工具包的使用示例。
package main

import (
    "fmt"
    "io/ioutil"
    "os"

    "words"
)

// main 函数是整个程序的入口。
func main() {
    filename := os.Args[1]

    contents, err := ioutil.ReadFile(filename)
    if err != nil {
        fmt.Println("打开文件时发生错误:", err)
        return
    }

    text := string(contents)

    count := words.CountWords(text)
    fmt.Printf("在你的文件中有 %d 个单词。\n", count)
}
