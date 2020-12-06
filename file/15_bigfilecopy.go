package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func cp(srcfile string, dstfile string) {
	opensrc, err := os.Open(srcfile)
	if err != nil {
		fmt.Println("opensrc err")
		return
	}
	reader := bufio.NewReader(opensrc)
	defer opensrc.Close()

	writedst, err := os.OpenFile(dstfile, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("write err")
		return
	}
	writer := bufio.NewWriter(writedst)

	defer writedst.Close()
	i, err := io.Copy(writer, reader)
	if err != nil {
		fmt.Println("io err")
		return
	} else {
		fmt.Printf("trans ok ......,file size%v", i)

	}
	writer.Flush()
}

func main() {
	srcfile := "c:/14.avi"
	dstfile := "d:/14.avi"
	cp(srcfile, dstfile)
}
