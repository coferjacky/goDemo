package main

//接口嵌套组合-将多个接口放一个接口里面
import "io"

//声明一个设备结构,用来模拟一个虚拟设备
type device struct {
}

//实现io.Writer的write()方法
func (d *device) Write(p []byte) (n int, err error) {
	return 0, nil
}

//实现io.Closer的close()方法
func (d *device) Close() error {
	return nil
}

func main() {
	//声明写入关闭器 并赋予device，由于device实现了io.WriteClose,则device指针就可以隐形的转为io.WriteCloser接口
	var wc io.WriteCloser = new(device)
	wc.Write(nil)
	wc.Close()
	var write_Only io.Writer = new(device)
	write_Only.Write(nil)
}
