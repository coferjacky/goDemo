package main

import (
	"../packageDemo"
	"fmt"
	"os"
	"strconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64) //将string转换为float64
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf:%v\n", err) //将错误默认输出到屏幕
			os.Exit(1)
		}
		f := packageDemo.Fahrenheit(t)
		c := packageDemo.Celsius(t)
		fmt.Printf("%s=%s,%s=%s\n", f, packageDemo.FToC(f), c, packageDemo.CToF(c))

	}

}
