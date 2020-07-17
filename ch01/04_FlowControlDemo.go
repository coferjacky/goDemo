package main

import (
	"fmt"
)

func main() {
	/*
		case的DEMO:学生评级
	*/
	grade := ""
	score := 78.5
	switch {
	case score >= 90:
		grade = "A"
	case score >= 80:
		grade = "B"
	case score >= 70:
		grade = "C"
	default:
		grade = "E"
	}
	fmt.Printf("等级为：%s\n", grade)
	fmt.Println("最终评价")
	switch grade {
	case "A":
		fmt.Println("优秀")
	case "B":
		fmt.Println("一般")
	default:
		fmt.Println("超差")
	}
	/*
		判断某年某月的天数
	*/
	year := 2008
	mouth := 4
	days := 0
	switch mouth {
	case 1, 3, 5, 7, 8, 10, 12:
		days = 31
	case 4, 6, 9, 11:
		days = 30
	case 2:
		if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
			days = 29
		} else {
			days = 28
		}
	default:
		days = -1
	}
	fmt.Printf("%d年%d月的天数为：%d\n", year, mouth, days)
}
