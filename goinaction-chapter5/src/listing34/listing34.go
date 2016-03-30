// 示例： curl 命令简化版本，使用 io.Reader 和 io.Writer 。
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// 在 init 方法在 main 方法之前被调用，可在该方法中作初始化与校验等。
func init() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ./example2 <url>")
		os.Exit(-1)
	}
}

// 程序主入口。
func main() {
	transport := &http.Transport{Proxy: nil}
	client := &http.Client{Transport: transport}

	request, err := http.NewRequest("GET", os.Args[1], nil)
	response, err := client.Do(request)

	checkError(err)
	fmt.Println("Read ok")

	if response.Status != "200 OK" {
		fmt.Println(response.Status)
		os.Exit(2)
	}
	fmt.Println("Reponse ok")

	var buf [512]byte
	reader := response.Body
	for {
		n, err := reader.Read(buf[0:])
		if err != nil {
			os.Exit(0)
		}
		fmt.Print(string(buf[0:n]))
	}

	os.Exit(0)

}

func checkError(err error) {
	if err != nil {
		if err == io.EOF {
			return
		}
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}

func simpleGet() {
	// 创建一个 GET 请求
	r, err := http.Get(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	// 将响应体拷贝到标准输出
	io.Copy(os.Stdout, r.Body)
	if err := r.Body.Close(); err != nil {
		fmt.Println(err)
	}
}
