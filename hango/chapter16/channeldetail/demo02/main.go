package  main
/*
select 可以解决从管道取数据的阻塞问题
 */
import "fmt"

func main(){
	intchan:=make(chan int,10)
	for i:=0;i<10;i++{
		intchan<-i
	}
	stringchan:=make(chan string,5)
	for i:=0;i<5;i++{
		stringchan<-"hello"+fmt.Sprintf("%d",i)
	}
	//传统方法在遍历管道时，如果不关闭会阻塞而导致deadlock
	//有的场景无法确定关闭管道的
	for {
		select {
		case v := <-intchan:
			fmt.Printf("从intchan读取数据%d\n", v)
		case v := <-stringchan:
			fmt.Printf("从stringchan读取数据%s\n", v)
		default:
			fmt.Println("取不到了")
		}
	}