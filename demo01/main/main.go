package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("C:/Users/86158/Desktop/书籍包/test01.txt")

	if err != nil {
		fmt.Println("文件打开出错，对应错误为", err)
	} else {
		fmt.Printf("文件 = %v", file)
	}

	err2 := file.Close()
	if err2 != nil {
		fmt.Println("文件关闭出错，错误为：", err2)
	}
}
