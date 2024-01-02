package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	//定义源文件
	filePath1 := "C:/Users/86158/Desktop/书籍包/test03.txt"
	//定义目标文件
	filePath2 := "C:/Users/86158/Desktop/书籍包/test04.txt"

	//打开源文件
	file, err := os.Open(filePath1)
	if err != nil {
		fmt.Println("源文件打开出错")
		return
	}
	//及时关闭文件
	defer file.Close()

	//打开目标文件
	file1, err1 := os.OpenFile(filePath2, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err1 != nil {
		fmt.Println("目标文件打开出错")
		return
	}
	//及时关闭文件
	defer file1.Close()

	//创建读入流
	reader := bufio.NewReader(file)
	//创建写出流
	writer := bufio.NewWriter(file1)
	for {
		content, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		writer.WriteString(content)
	}
	//刷新写入流的缓存区
	writer.Flush()
}
