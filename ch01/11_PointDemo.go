package main
import "fmt"

type Student struct {
	name string
	age int
	married bool
	sex int8
}
func main(){
	var a int =120
	var ip *int
	ip=&a
	fmt.Printf("a的类型%T,值是%v \n",a,a)
	fmt.Printf("&a的类型是%T,值是%v \n ",&a,&a)
	fmt.Printf("ip的类型是%T,值是%v \n",ip,ip)
	fmt.Printf("*ip的类型是%T,值是%v \n",*ip,*ip)
	fmt.Printf("&a变量的类型是%T,值是%v \n",*&a,*&a)
	fmt.Println(a,&a,*&a)
	fmt.Println(ip,&ip,*ip,*(&ip),&(*ip))

	//-------------结构体案例----------------------
	var s1=Student{
		name:    "steven",
		age:     0,
		married: false,
		sex:     0,
	}
	var s2=Student{
		name:    "sunny",
		age:     0,
		married: false,
		sex:     0,
	}
	var j *Student=&s1 //将s1的地址 赋给j指针变量
	var k *Student=&s2




	fmt.Printf("s1类型为%T，值为%v \n",s1,s1)
	fmt.Printf("s2类型为%T，值为%v \n",s2,s2)
	fmt.Printf("j类型为%T，值为%v \n",j,j)
	fmt.Printf("k类型为%T，值为%v \n",k,k)
	fmt.Printf("*j类型为%T，值为%v \n",*j,*j)
	fmt.Printf("*k类型为%T，值为%v \n",*k,*k)
	fmt.Println(s1.name,s1.age,s1.married,s1.sex)
	fmt.Println(j.name,j.age,j.married,j.sex)
	fmt.Println((*j).sex,(*j).married)
	fmt.Printf("&j类型为%T,值为%v \n",&j,&j)
	fmt.Printf("&k类型为%T,值为%v \n",&k,&k)
	fmt.Println(&j.name)
	fmt.Println(&k.name)

}
