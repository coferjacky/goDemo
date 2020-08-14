/*
找到文件后缀为.jpg的文件并显示文件名
*/
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("请输入文件所在的路径：")
	var path string
	fmt.Scan(&path)

	//打开目录
	f, err := os.OpenFile(path, os.O_RDONLY, os.ModeDir)
	if err != nil {
		fmt.Println("openFile error:", err)
		return
	}
	defer f.Close()

	//读取目录项
	info, err := f.Readdir(-1)
	if err != nil {
		fmt.Println("Readdir err:", err)
		return
	}

	//变量返回切片
	for _, fileInfo := range info {
		if !fileInfo.IsDir() {
			if strings.HasSuffix(fileInfo.Name(), ".jpg") {
				fmt.Println("jpg文件有：", fileInfo.Name())
			}
		}
	}

}
