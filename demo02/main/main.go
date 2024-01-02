package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	//读取文件
	content, err := ioutil.ReadFile("C:/Users/86158/Desktop/书籍包/test01.txt")
	//返回的内容为[]byte,err

	if err != nil {
		fmt.Println("读取出错，错误为：", err)
	} else {
		fmt.Printf("%v\n", content)
		fmt.Printf("%v", string(content))
	}
}
