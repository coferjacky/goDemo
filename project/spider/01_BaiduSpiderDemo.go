package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	var start, end int
	fmt.Print("请输入爬取的起始页")
	fmt.Scan(&start)
	fmt.Print("请输入爬取终止页")
	fmt.Scan(&end)
	working(start, end)
}

func working(start int, end int) {
	fmt.Printf("正在爬取第%d页到%d页...", start, end)

	//循环爬取每页数据
	for i := start; i <= end; i++ {
		url := "https://tieba.baidu.com/f?kw=%E4%B9%A6&ie=utf-8&pn=" + strconv.Itoa((i-1)*50)
		result, err := httpGet(url)

		if err != nil {
			fmt.Printf("http get err:", err)
			continue
		}
		//将读到的整个网页保存成一个文件，读到一页保存一页
		f, err := os.Create("第" + strconv.Itoa(i) + "页.html")
		f.WriteString(result)
		f.Close() //保存好一个文件马上关闭一个文件
	}
}

func httpGet(url string) (result string, err error) {
	resp, err1 := http.Get(url)
	if err1 != nil {
		err = err1
		return
	}
	defer resp.Body.Close()
	time.Sleep(time.Second)
	buf := make([]byte, 4096)
	//循环读取网页信息传给调用者
	for {
		n, err2 := resp.Body.Read(buf)
		if n == 0 {
			fmt.Printf("读取网页完成")
			break
		}
		if err2 != nil && err2 != io.EOF {
			err = err2
			return
		}
		//累加每一次循环读到buf数据，存入result 一次性返回
		result += string(buf[:n])
	}
	return
}
