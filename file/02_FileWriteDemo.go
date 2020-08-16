package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.OpenFile("c:/myjack.abc", os.O_RDWR, 6)
	if err != nil {
		fmt.Println("OpenFile err:", err)
		return
	}
	defer f.Close()
	fmt.Println("success")
	n, err := f.WriteString("HelllloWord man\r\n") //注意回传换行，如果是windows是\r\n,如果linux是\n
	if err != nil {
		fmt.Println("writeString err:", err)
		return
	}
	fmt.Println("writeString n=", n)

	//seek获取文件的读写指针位置
	//参数1：表示文件的偏移量 正：向文件尾偏，负：向文件头偏
	//参数2：偏移的起始位置 ：io.SeekStart文件起始位置  io.SeekCurrent 文件当前位置 io.Seek
	//返回值：表示从文件起始位置到当前文件
	//off,_:=f.Seek(5,io.SeekStart)  //相当于光标到开头，然后向右偏移五个字符
	off, _ := f.Seek(-5, io.SeekEnd) //相当于光标到结尾，然后向左偏移五个字符
	fmt.Println("Seekoff:", off)

	n, _ = f.WriteAt([]byte("1111"), off) //文件指定偏移位置，写入 []byte,通常搭配Seek，参数1 待写入数据 参2：偏移位置，返回实际写出的字节数
	fmt.Println("Seekoff:", n)
}
