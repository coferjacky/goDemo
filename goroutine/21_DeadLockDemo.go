package main

/*
死锁引起的3个情况
1 单go程自己死锁
	channel应该至少2个以上的go程总通讯，否则死
2 go程间channal访问顺序导致的死锁
	使用channel一起读写，保证另外一端写（读）操作，同时有机会执行，否则死锁
3 多go程序总交叉死锁
     交叉持有对方执行权
4、尽量不要将互斥锁、读写锁与channel混用，会发生隐形死锁

*/

func main() {
	//演示第三种死锁
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		for {
			select {
			case num := <-ch1:
				ch2 <- num
			}
		}
	}()
	for {
		select {
		case num := <-ch2:
			ch1 <- num
		}
	}
}
