package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	//打开文件
	file, err := os.Open("C:/Users/86158/Desktop/书籍包/test01.txt")

	if err != nil { //打开失败
		fmt.Println("文件打开失败，失败的原因为", err)
	}
	//函数退出后，防止文件泄漏
	defer file.Close()

	//创建流
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n') //从流中读取字符串，读取到一个换行就结束
		if err == io.EOF {                  //读取到文件结尾
			break
		}
		//当没有读到文件末尾时，正常输出文件内容
		fmt.Println(str)
	}
	fmt.Println("文件读取成功，已全部读取完毕")
}
