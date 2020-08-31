package main

import (
	"fmt"
	"image"
	"image/color"
)

/*
整形
	有符号
		int8
		int16
		int32
		int64
	无符号
		uint8   就是我们熟悉的byte
		uint16
		uint32
		uint64



浮点型
	float32
	float64

布尔型(禁止整数强转为bool类型)
字符串
	1、字符转义
		\r	回车符
    	\n	换行符
		\t	制表符
		\'	单引号
		\" 双引号
		\\反斜杠
	2、字符串实现基于utf-8编码
	3、定义多行字符串

字符类型
	uint8 或byte 这代表ASCII码的一个字符

	rune	代表一个UTF-8字符 实际就是一个int32

切片
结构体
函数
map
channel

*/
/*
输出正弦图形
*/
const size = 300

func main() {
	pic := image.NewGray(image.Rect(0, 0, size, size))
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			pic.SetGray(x, y, color.Gray{255})
		}
	}
	//如下表示的的转义符号
	fmt.Println("str:=\"c:\\go\\bin\\go.ext\"")

}
