package main
import (
	"fmt"
)
/*
由于此处涉及到多行所以使用反引号（数字键的最右侧行）
 */
func main(){
	var temp string
	temp=`x:=10
	y:=20
	z:=30`
	fmt.Println(temp)

	//通用打印格式  %T表示值类型，%v表示值默认格式
	str :="steven"
	fmt.Printf("%T ,%v \n",str,str)
	//rune是字节数4,表示单个unicode字符，int32别名
	var a rune='一'
	fmt.Printf("%T,%v \n",a,a)

	var b byte='b'
	fmt.Printf("%T,%v \n",b,b)

	var c int32=98
	fmt.Printf("%T,%v \n",c,c)

	//布尔打印格式  %t布尔的值

	var flag bool
	fmt.Printf("%T,%t \n",flag,flag)
	flag=true
	fmt.Printf("%T,%t \n",flag,flag)

	/*整形打印格式 %d表示十进制，%nd 表示该长度为n，不足n补0，超过以实际为准 %b表示二进制
      %o 8进制数
	%X 十六进制大写
	%x 十六进制小写
	%c 该值对应的unicode码值
	%q 该值对应单引号扩起来的go语法字符字面值，必要时采用安全转义表示
	%U 表示为unicode格式：U+1234 等价 U+%O4X


	 */
	c1:=123
	c2:=97
	fmt.Printf("%T,%d \n",c1,c1)
	fmt.Printf("%T,%4d \n",c1,c1)
	fmt.Printf("%T,%04d \n",c1,c1)
	fmt.Printf("%T,%b \n",c1,c1)
	fmt.Printf("%T,%o \n",c1,c1)
	fmt.Printf("%T,%c \n",c2,c2)
	fmt.Printf("%T,%q \n",c2,c2)
	fmt.Printf("%T,%x \n",c1,c1)
	fmt.Printf("%T,%X \n",c1,c1)
	fmt.Printf("%T,%U \n",a,a)
}
