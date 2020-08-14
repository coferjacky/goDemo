/*
从用户给定目录下面将所有jpg后缀的文件拷贝到指定目录
*/
package main

import (
	"fmt"
	"io"
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
				cpMP32Dir(path+"/"+fileInfo.Name(), "./"+fileInfo.Name())
			}
		}
	}

}

//拷贝jgp文件到指定目录的操作
func cpMP32Dir(src, dst string) {
	//1打开读文件
	f_r, err := os.Open(src)
	if err != nil {
		fmt.Println("Open err", err)
		return
	}
	defer f_r.Close()
	//2 创建写文件
	f_w, err := os.Create(dst)
	if err != nil {
		fmt.Println("close err", err)
		return
	}
	defer f_w.Close()

	//3 从读文件中获取数据 放在缓存区
	//3.1 创建缓存区
	buf := make([]byte, 4096)
	//3.2循环从读文件中获取数据，”原封不动“写到写文件中
	for {
		n, err := f_r.Read(buf)
		if err != nil && err == io.EOF {
			fmt.Printf("读完毕 n= %d \n", n)
			return
		}
		f_w.Write(buf[:n]) //读多少写多少
	}
}
