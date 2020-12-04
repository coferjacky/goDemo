package main

import "fmt"

type student struct {
}

func TypeJudge(item ...interface{}) {
	for index, v := range item {
		switch v.(type) {
		case bool:
			fmt.Printf("第%v个参数是bool类型，值是%v", index, v)
		case float32:
			fmt.Printf("第%v个参数是float32类型，值是%v", index, v)
		case float64:
			fmt.Printf("第%v个参数是float64类型，值是%v", index, v)
		case int, int32, int64:
			fmt.Printf("第%v个参数是int类型，值是%v", index, v)
		case string:
			fmt.Printf("第%v个参数是string类型，值是%v", index, v)
		case student:
			fmt.Println("第%v个参数是student类型，值是%v", v)
		case *student:
			fmt.Println("第%v个参数是*student类型，值是%v", v)
		default:
			fmt.Printf("第%v个参数值确定", index)
		}
	}
}
func main() {
	var n1 float32 = 1.1
	var n2 float64 = 12.3
	var n4 int32 = 30
	var name string = "tom"
	address := "北京"
	n5 := 300
	stu1 := student{}
	stu2 := &student{}
	TypeJudge(n1, n2, n4, n5, name, address, stu1, stu2)
}
