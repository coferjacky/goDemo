package main

import (
	"fmt"
)

func main() {
	i := 0               //i放在这里就能再for循环外面使用了
	for ; i <= 10; i++ { //此处省略初始语句
		fmt.Printf("%d\n", i)
	}
	fmt.Printf("%d", i)

	j := 0
	for ; ; j++ { //条件语句被省略以后 就是无限循环了
		if j > 20 {
			break
		}
		fmt.Printf("新值为%d", j)
	}

	var k int
	for { //全部都省了 无限循环了，里面一定要有 结束语句 否则就是死循环了
		if k > 10 {
			break
		}
		fmt.Printf("第三种情况%d", k)

		k++
	}
	fmt.Print("\n")

	var s int
	for s <= 10 { //省略初始化语句和省略了条件语句
		fmt.Print(s)
		s++
	}
	fmt.Print("\nfor循环的range格式\n")
	str := "121212Abc曹菲   "
	for i, value := range str { //for循环的range格式对string slice array map channel进行迭代循环。array slice string 返回索引,map 返回键值，channel返回通道值
		fmt.Printf("第%d位的ASCII值为%d,字符是%c \n", i, value, value)
	}

	fmt.Print("\n来，我们整个例子  求0...10的和\n")
	sum := 0
	for i := 0; i <= 100; i++ {
		sum += i
	}
	fmt.Printf("0..10的和为%d", sum)

	fmt.Print("\n求1...40的3的倍数和\n")
	q := 1
	suma := 0
	fmt.Print("0")
	for q <= 40 {
		if q%3 == 0 {
			suma += q
			if q < 39 {
				fmt.Printf("+%d", q)
			} else {
				fmt.Printf("+%d=%d\n", q, suma)
			}
		}
		q++
	}

	//截取竹竿:每次砍掉1.5米，32米的竹竿，砍几次后剩余不足4米
	fmt.Print("\n截取竹竿，每次砍掉1.5米，3米的竹竿，砍几次后剩余不足4米\n")
	count := 0
	for i := 32.0; i > 4; i -= 1.5 {
		count++
	}
	fmt.Print(count)

	//打印直角三角形
	fmt.Print("\n打印直角三角形\n")
	lines := 8
	for i := 0; i < lines; i++ {
		for j := 0; j < i+1; j++ {
			fmt.Print("* ")
		}
		fmt.Print("\n")
	}

	//打印九九除法表
	fmt.Print("\n打印九九除法表\n")
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d ", i, j, i*j)
		}
		fmt.Print("\n")
	}

	//嵌套循环来输出 2-50的素数（不借用goto语句来输出）
	fmt.Print("\n嵌套循环来输出 2-50的素数\n")
	var i1, i2 int
	for i1 = 2; i1 <= 50; i1++ {
		for i2 = 2; i2 <= (i1 / i2); i2++ {
			if i1%i2 == 0 {
				break
			}
		}
		if i2 > (i1 / i2) {
			fmt.Printf("%d\t", i1)
		}
	}

	//嵌套循环来输出 2-50的素数（借用goto语句来输出）
	fmt.Print("\n嵌套循环来输出 2-50的素数(GOTO语句)\n")
	var C, c int
	C = 1
LOOP: //跳转到LOOP时不时重新一轮循环，而是继续当前循环，所以这里的c不能写到循环体的初始化文件里面
	for C < 50 {
		C++
		for c = 2; c < C; c++ {
			if C%c == 0 {
				goto LOOP
			}
		}
		fmt.Printf("%d \t", C)
	}

}
