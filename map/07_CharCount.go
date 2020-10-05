package main

//统计字符数和 每种字段长度进行对应的计数
import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	counts := make(map[rune]int)    //建立一个map key:rune类型，value:int类型
	var utflen [utf8.UTFMax + 1]int //utf8长度 这里声明了utflen的数组 5个元素
	invalid := 0                    //无效utf-8的数量
	in := bufio.NewReader(os.Stdin) //创建读取器，并将其与标准输入绑定
	for {
		r, n, err := in.ReadRune() //r:解码的rune值 n:字符utf-8编码后长度nbytes err:错误值、error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount:%v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++ //读取一个rune就将 解码的值放入键，然后值+1，相同编码就进行值累加
		utflen[n]++ //读取一个n则放进数组，读取相同的n，则对数组中的值进行累加
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-9 characters \n", invalid)
	}

}
