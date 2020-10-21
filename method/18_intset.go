package main

import (
	"bytes"
	"fmt"
)

type Inset struct {
	words []uint
}

const (
	bitNum = (32 << (^uint(0) >> 63)) //根据平台自动判断决定是32还是64
)

func main() {

	var x, y Inset
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String())

	y.Add(1)
	y.Add(91)
	fmt.Println(y.String())

	x.UnionWith(&y)
	fmt.Println(x.String())
	fmt.Println(x.Has(9), x.Has(123)) //123找不到

	fmt.Println(&x)         //打印了*IntSet的指针，这个类型指针确实有自己定义的方法
	fmt.Println(x.String()) //直接调用了变量的String()方法
	fmt.Println(x)          //直接调用INSET类型，由于没有定义String 的方法，所以Println方法会直接以最原始的方式理解并打印。所以这种情况下&符号不能忘记，所以这里我们String方法绑定到IntSet对象上面更加合适一些

}

func (s *Inset) Has(x int) bool {
	word, bit := x/bitNum, uint(x%bitNum)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0

}

func (s *Inset) Add(x int) {
	word, bit := x/bitNum, uint(x%bitNum)

	for word >= len(s.words) { //第一次运行时s.words里面为Nil,所以此时的len里面是0

		s.words = append(s.words, 0) //第一次将0压入words数组
	}
	s.words[word] |= 1 << bit //将1移动bit
}

func (s *Inset) UnionWith(t *Inset) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// String returns the set as a string of the form "{1 2 3}".
func (s *Inset) String() string {
	var buf bytes.Buffer //Buffer 就像一个集装箱容器，可以存东西，取东西
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 { //先word和(1<<uint(j))与 512&1是1
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", bitNum*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}
