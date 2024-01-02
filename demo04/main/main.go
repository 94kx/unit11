package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//写入文件
	//打开文件
	file, err := os.OpenFile("C:/Users/86158/Desktop/书籍包/test02.txt", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil { //打开文件失败
		fmt.Println("文件打开失败，失败原因为", err)
		return
	}

	//及时关闭文件
	defer file.Close()

	//创建写入流
	writer := bufio.NewWriter(file)
	//写入字符串
	for i := 0; i < 10; i++ {
		writer.WriteString("你好！中国")
	}
	//流带有缓存区，需刷新缓存区
	writer.Flush()
}
