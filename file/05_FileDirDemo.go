package main

/*
打开目录OpenFile，把目录看成一个文件
参数1：打开目录的路径
参数2：打开权限 O_RDONLY O_WRONLY  O_RDWR
参数3：os.ModeDir  打开模式 os.ModeDir是文件夹打开模式
返回一个可以读写目录的文件指针
*常用的的方法：Readdir

下例列明一个目录
*/
import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("请输入待查询的目录路径:")
	var path string
	fmt.Scan(&path) //获取用户输入的的目录路径
	//打开目录
	file, err := os.OpenFile(path, os.O_RDONLY, os.ModeDir)
	if err != nil {
		fmt.Println("opener:", err)
		return
	}
	defer file.Close()

	//读取目录项目
	info, err := file.Readdir(-1)
	if err != nil {
		fmt.Println("opendrer:", err)
		return
	}
	//遍历readdir切片
	for _, fileInfo := range info {
		if fileInfo.IsDir() {
			fmt.Println(fileInfo.Name(), "是一个目录：")
		} else {
			fmt.Println(fileInfo.Name(), "是一个文件")
		}
	}

	//找目录中的.jpg文件
	for _, fileInfo := range info {
		if !fileInfo.IsDir() {
			if strings.HasSuffix(fileInfo.Name(), ".jpg") {
				fmt.Println("jpg文件有", fileInfo.Name())
			}
		}
	}
}
