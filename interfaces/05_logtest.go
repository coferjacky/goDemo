package main

import "io"

/*
便于扩展输出方式的日志系统

*/
//声明日志写入器接口
type Log_Writer interface {
	Write(data interface{}) error
}

//日志器
type Logger struct {
	//这个日志器用到的日志写入器
	write_List []Log_Writer
}

//注册一个日志写入器
func (l *Logger) Register_Writer(writer Log_Writer) {
	l.write_List = append(l.write_List, writer)
}

//将一个data类型的数据写入日志
func (l *Logger) Log(data interface{}) {
	for _, writer := range l.write_List {
		writer.Write(data)
	}
}

func New_Logger() *Logger {
	return &Logger{}
}
