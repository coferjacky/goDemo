/*
统计指定目录内love的个数
*/

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	fmt.Println("请输入要找寻的目录：")
	var path string
	fmt.Scan(&path) //获取用户指定的目录名

	dir, _ := os.OpenFile(path, os.O_RDONLY, os.ModeDir) //只读打开该目录
	defer dir.Close()

	names, _ := dir.Readdir(-1) //读取当前目录下所有的文件名和目录名，存入names切片

	var AllLover int = 0
	for _, name := range names {
		if !name.IsDir() {
			s := name.Name() //文件名不包含路径
			if strings.HasSuffix(s, ".txt") {
				AllLover += readFile(path + s) //这里的readFile函数是来统计本次循环的读写关键字的次数
			}

		}
	}
	fmt.Printf("目录所有文件中共有%d个love \n", AllLover)
}

//从一个文件中逐行读取内容，统计该文件有多少个love
func readFile(s string) int {
	fp, err := os.Open(s)
	if err != nil {
		fmt.Println("open err", err)
		return -1
	}
	defer fp.Close()

	row := bufio.NewReader(fp) //创建一个reader
	var total int = 0          //统计Love总数的变量

	//循环按行读取文件内容存入buf中
	for {
		buf, err := row.ReadBytes('\n')
		if err != nil && err == io.EOF {
			break
		}
		//交给wordCount统计每行中love出现次数。累计各行love数
		total += wordCount(string(buf[:]))
	}
	return total

}

//从一行字符串中获取love出现的次数，该函数返回次数
func wordCount(s string) int {
	arr := strings.Fields(s)  //分割字符串，存入字符数组
	m := make(map[string]int) //创建map
	//对arr中每个单词进行循环，存入map中进行统计
	for i := 0; i < len(arr); i++ {
		m[arr[i]]++
	}
	//统计map中一共有多少个"love"
	for k, v := range m {
		if k == "love" {
			fmt.Printf("%s  :   %d \n", k, v)
			return m[k]
		}

	}
	return 0
}
