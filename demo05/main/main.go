package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	//定义源文件
	filePath1 := "C:/Users/86158/Desktop/书籍包/test02.txt"
	//定义目标文件
	filePath2 := "C:/Users/86158/Desktop/书籍包/test03.txt"

	//对文件进行读取
	content, err := ioutil.ReadFile(filePath1)
	if err != nil {
		fmt.Println("读取出错，错误为", err)
		return
	}

	//对文件进行写入
	err1 := ioutil.WriteFile(filePath2, content, 0666)
	if err1 != nil {
		fmt.Println("写入出错，错误为", err)
		return
	}
}