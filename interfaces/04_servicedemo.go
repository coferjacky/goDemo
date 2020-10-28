package main

//定义日志器的结构
type Logger struct {
}

type Service interface {
	Start()
	Log(string)
}

type Game_Service struct {
	Logger
}

func (g *Logger) Log(s string) {

}
func (g *Game_Service) Start() {

}
func main() {

	g := new(Game_Service)
	var s Service
	s = g
	s.Start()
	s.Log("fs")

}
