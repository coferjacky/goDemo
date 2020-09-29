package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

/*
由于此处涉及到多行所以使用反引号（数字键的最右侧行）
*/
func main() {
	var temp string
	temp = `x:=10
	y:=20
	z:=30`
	fmt.Println(temp)

	//通用打印格式  %T表示值类型，%v表示值默认格式
	str := "steven"
	fmt.Printf("%T ,%v \n", str, str)
	//rune是字节数4,表示单个unicode字符，int32别名
	var a rune = '一'
	fmt.Printf("%T,%v \n", a, a)

	var b byte = 'b'
	fmt.Printf("%T,%v \n", b, b)

	var c int32 = 98
	fmt.Printf("%T,%v \n", c, c)

	//布尔打印格式  %t布尔的值

	var flag bool
	fmt.Printf("%T,%t \n", flag, flag)
	flag = true
	fmt.Printf("%T,%t \n", flag, flag)

	/*整形打印格式 %d表示十进制，%nd 表示该长度为n，不足n补0，超过以实际为准 %b表示二进制
	      %o 8进制数
		%X 十六进制大写
		%x 十六进制小写
		%c 该值对应的unicode码值
		%q 该值对应单引号扩起来的go语法字符字面值，必要时采用安全转义表示
		%U 表示为unicode格式：U+1234 等价 U+%O4X



	*/
	c1 := 123
	c2 := 97
	fmt.Printf("%T,%d \n", c1, c1)
	fmt.Printf("%T,%4d \n", c1, c1)
	fmt.Printf("%T,%04d \n", c1, c1)
	fmt.Printf("%T,%b \n", c1, c1)
	fmt.Printf("%T,%o \n", c1, c1)
	fmt.Printf("%T,%c \n", c2, c2)
	fmt.Printf("%T,%q \n", c2, c2)
	fmt.Printf("%T,%x \n", c1, c1)
	fmt.Printf("%T,%X \n", c1, c1)
	fmt.Printf("%T,%U \n", a, a)

	/*
	   %b 无小数部分，二进制指数的科学计数法
	   %E 科学计数法
	   %f 有六位小数的部分
	   %F 和%f一致
	   %.3e 有3位小数的的科学计数法
	   %g 根据实际情况采用%e或%f格式
	   %G 根据实际情况采用%E或%F格式

	*/
	c3 := 123.12345678
	fmt.Printf("%b \n", c3)
	fmt.Printf("%E \n", c3)
	fmt.Printf("%f \n", c3)
	fmt.Printf("%F \n", c3)
	fmt.Printf("%.3e \n", c3)
	fmt.Printf("%g \n", c3)
	fmt.Printf("%G \n", c3)
	/*
	   复数格式打印
	*/
	var value complex64 = 2.1 + 12i
	value2 := complex(2.1, 12)
	fmt.Println(real(value))
	fmt.Println(imag(value))
	fmt.Println(value2)
	/*
		字符串和字节数组的打印格式
	*/
	s0 := []byte{'x', 'y', 'z', 'Z'}
	s1 := "欢迎大家学些区块链"
	fmt.Printf("%s \n", s1)
	fmt.Printf("%q \n", s1)
	fmt.Printf("%x \n", s1) //两个字节用两个字符十六进制数表示 使用a-f
	fmt.Printf("%X \n", s1) //两个字节用两个字符十六进制数表示 使用A-F

	fmt.Printf("%T,%s \n", s0, s0)
	fmt.Printf("%T,%q \n", s0, s0)
	fmt.Printf("%T,%x \n", s0, s0)

	//类型转换
	/*
	   var  a int =100
	   b:=float64(a)  //int转换成float64型
	   c:=string(a)   //int型转换成string类型
	*/
	chinese := 90
	english := 80.9
	avg := (chinese + int(english)) / 2
	avg2 := (float64(chinese) + english) / 2
	fmt.Printf("%T,%d \n", avg, avg)
	fmt.Printf("%T,%f \n", avg2, avg2)

	/*
		int数值转换为string类型，注意go语言不允许字符串转int类型
	*/
	d1 := 97
	d2 := 19968
	result := string(d1)
	result2 := string(d2)
	fmt.Println(result)
	fmt.Println(result2)

	/*
		枚举
	*/
	const (
		Unknown = 0
		female  = 1
		male    = 2
	)
	const (
		a11 = 10
		b11
		c11
	)
	const (
		a22 = iota //第一次出现时0
		b22 = iota //第二次出现时1
		c22 = iota //第三次出现时2
	)
	fmt.Println(Unknown, female, male)
	fmt.Println(a11, b11, c11)
	fmt.Println(a22, b22, c22)

	//字符串处理函数-指定分隔符拆分
	str1 := "i love my work and i love my family too"
	ret := strings.Split(str1, " ")
	for _, s := range ret {
		fmt.Println(s)
	}

	//字符串处理函数-空格拆分:返回将字符串按照空白（unicode.IsSpace确定，可以是一到多个连续的空白字符）分割的多个字符串。如果字符串全部是空白或者是空字符串的话，会返回空切片。
	ret1 := strings.Fields(str)

	for _, s := range ret1 {
		fmt.Println(s)
	}

	//判断字符串结束标志
	flg := strings.HasSuffix("TEST.ABc", ".ABc") //分大小写的
	fmt.Println(flg)
	//判断字符串起始标志
	flg1 := strings.HasPrefix("test.AbC", "test")
	fmt.Println(flg1)

	//len函数返回字节数目
	s := "Hello, 世界"
	fmt.Println(s[0:])
	fmt.Printf("%c\n", s[11])

	//utf8解码器的使用
	for i := 0; i < len(s); {
		inString, size2 := utf8.DecodeRuneInString(s[i:]) //返回切片首字符，并返回该字符的所占字节数（UTF编码后长度）
		fmt.Printf("%d\t%c\t%d\t", i, inString, size2)

		i += size2
		fmt.Println(i)
	}

	//使用range来统计字符串字符个数
	n := 0
	for range s { //健值返回都省略
		n++

	}
	fmt.Println(n)

	//

}
