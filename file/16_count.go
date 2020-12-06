package main

import (
	"bufio"

	"fmt"
	"io"
	"os"
)

type CharCount struct {
	ChCount    int
	NumCount   int
	SpaceCount int
	OtherCount int
	ChinaCount int
}

func main() {
	fileName := "c:/myjack.txt"
	open, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("open err%v", err)
		return
	}
	defer open.Close()
	var count CharCount
	reader := bufio.NewReader(open)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		for _, v := range str {
			switch {
			case v >= 'a' && v <= 'z':
				fallthrough //穿透处理
			case v >= 'A' && v <= 'Z':
				count.ChCount++

			case v >= '0' && v <= '9':
				count.NumCount++
			case v == ' ' || v == '\t':
				count.SpaceCount++

			case v >= 0x4e00 && v <= 0x9fbb:
				count.ChinaCount++
			default:
				count.OtherCount++
			}

		}

	}
	fmt.Printf("英文个数%d,空格个数%d,数字个数%d,其他字符个数%d,中文字符个数%d", count.ChCount, count.SpaceCount,
		count.NumCount, count.OtherCount, count.ChinaCount)
}
