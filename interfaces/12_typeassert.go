package main

import "fmt"

//接口断言将接口转换成另外的接口，也可以将接口转换为另外的类型。接口转换再开发中特别常见，使用非常频繁

//t:=i.(T)  i代表接口变量，T代表转换的目标类型 t:表示转换后变量。如果i没有完全实现T接口的方法，这个语言会触发宕机，而且不太友好。
//另外还有一种写法t,ok:=i.(T)，这里如果i没有完全实现T接口方法

//将接口转换为其他接口
//飞行动物接口
type Flyer interface {
	Fly()
}

//行走动物接口
type Walker interface {
	Walker()
}

//定义鸟类
type bird struct {
}

func (b *bird) Walker() {
	fmt.Println("bird:walk")
}
func (b *bird) Fly() {
	fmt.Println("bird:fly")
}

//定义猪类
type pig struct {
}

func (p *pig) Walker() {
	fmt.Println("pig:walk")
}

func main() {
	animals := map[string]interface{}{
		"bird": new(bird),
		"pig":  new(pig),
	}
	for name, obj := range animals {
		f, is_Flayer := obj.(Flyer)
		w, is_Walker := obj.(Walker)
		fmt.Printf("name:%s is flayer:%v, is Walker:%v\n", name, is_Flayer, is_Walker)

		if is_Flayer {
			f.Fly()

		}
		if is_Walker {
			w.Walker()
		}
	}
}
