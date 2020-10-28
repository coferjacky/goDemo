package main

/*
一个类实现多个接口方法
*/
import (
	"fmt"
	"io"
)

//定义socket结构体
type Socket struct {
}

/*

io包内对Write interface进行了定义
******
type Writer interface {
	Writer1111(p []byte) (n int,err error)
}
type Closer interface {
	Close() error
}
*******

*/

//实现io.Writerj接口
func (s *Socket) Write(p []byte) (n int, err error) {
	return 0, nil
}

//实现io.Coloe接口
func (s *Socket) Close() error {
	return nil
}
func using_Writer(writer io.Writer) {
	writer.Write(nil)
}
func using_Closer(closer io.Closer) {
	closer.Close()
}
func main() {
	//实例化Socket结构体传递的指针变量给s
	s := new(Socket)
	fmt.Println(s)
	//相当于将s的传递值给接口变量赋值
	//接口变量write相当于调用了实现的Write方法
	using_Writer(s)
	//接口变量Coloser相当于调用了实现的close方法
	using_Closer(s)

}
