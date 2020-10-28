package main

import (
	"errors"
	"fmt"
	"os"
)

/*
便于扩展输出方式的日志系统

*/

//1、日志对外接口部分
//声明日志写入器接口，可以对外使用，日志的输出可以有多个设备，这个写入器就是用来实现一个日志输出设备
type Log_Writer interface {
	Write(data interface{}) error
}

//日志器结构  日志器使用write_List记录输出到哪些设备上
type Loggger struct {
	//这个日志器用到的日志写入器切片
	write_List []Log_Writer
}

//注册一个日志写入器，使用日志器方法Register_Writer()将一个日志写入器注册到日志器中Logger
func (l *Loggger) Register_Writer(writer Log_Writer) {
	l.write_List = append(l.write_List, writer)
}

//将一个data类型的数据写入日志
func (l *Loggger) Log(data interface{}) {

	//遍历所有注册的写入器
	for _, writer := range l.write_List {
		//将日志输出到每个写入器里面
		writer.Write(data)
	}
}

//创建日志器实例
func New_Loggger() *Loggger {
	return &Loggger{}
}

//2、文件写入器部分

//申明文件写入器，在结构体总保存一个文件句柄，以方便每次写入时进行操作
type file_Writer struct {
	file *os.File
}

//设置文件写入器写入的文件名
func (f *file_Writer) Set_File(filename string) (err error) {
	//如果文件已经打开，关闭前一个文件，如果重复调用这个函数，需要将之前的文件关闭掉再开新文件
	if f.file != nil {
		f.file.Close()
	}
	//创建一个文件并保存句柄
	f.file, err = os.Create(filename)
	return err
}

//实现Log_Writer的write()方法
func (f *file_Writer) Write(data interface{}) error {
	//日志文件没有创建成功
	if f.file == nil {
		return errors.New("FILE NOT CREATED")
	}
	//将数据序列化为字符串
	str := fmt.Sprintf("%v \n", data)
	//将数组以字节数组写入文件中
	_, err := f.file.Write([]byte(str))
	return err
}

//创建文件写入器实例
func new_File_Write() *file_Writer {
	return &file_Writer{}
}
