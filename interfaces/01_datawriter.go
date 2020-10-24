package main

import "fmt"

/*
方法名(参数列表1) 返回值列表1
实现接口的条件：
1 接口方法必须与实现接口的类型方法格式一致
2 接口所有方法均被实现
*/

//定义一个Data_Writer接口
type Data_Writer interface {
	Write_Data(data interface{}) error
	//Can_Write() bool
}
type file struct {
}

func (d *file) Write_Data(data interface{}) error {
	fmt.Println("wirte data:", data)
	return nil
}
func main() {
	//实例化结构体再赋值给f，这时f类型就是*file
	f := new(file)
	fmt.Printf("%T\n", f)
	//声明Data_write类型接口变量writer
	var writer Data_Writer
	//将*file类型变量赋值给DataWriter接口的Writer,虽然变量类型不一致，但是Writer是一个接口，f已经完全实现了DataWrite的所有方法，因此赋值成功
	writer = f

	writer.Write_Data("data")
}
