package main

import (
	"fmt"
	"time"
)

func main() {
	type Employee struct {
		ID        int
		Name      string
		Address   string
		DoB       time.Time
		Position  string
		Salary    int
		ManagerID int
	}
	var dilbert Employee
	dilbert.Salary -= 5000
	position := &dilbert.Position
	*position = "senior " + *position
	//点操作符和指向结构体的指针一起工作
	var empoyeeOfTheMonth *Employee = &dilbert
	empoyeeOfTheMonth.Position += "(proactive team player)"

	fmt.Println("直接给成员赋值：", dilbert.Salary)
	fmt.Println("通过指针访问成员", *position)

}
