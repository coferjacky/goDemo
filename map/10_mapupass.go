package main

import (
	"fmt"
)

/*
key 表示用户名 唯一的 不可以重复
如果一个用户名存在，则将其密码修改为888888，如果不存在就增加一个用户信息。包括nickname 和 密码pwd
*/
func modifyUser(users map[string]map[string]string, name string) {
	if users[name] != nil {
		//找到这个用户了
		users[name]["pwd"] = "888888"
	} else {
		users[name] = make(map[string]string, 2)
		users[name]["pwd"] = "888888"
		users[name]["nickname"] = "昵称-" + name
	}
}
func main() {
	users := make(map[string]map[string]string, 10)
	users["smith"] = make(map[string]string, 2)
	users["smith"]["pwd"] = "99999"
	users["smith"]["nickname"] = "boysmith"

	modifyUser(users, "tom")
	modifyUser(users, "jack")
	modifyUser(users, "jackbaul")
	modifyUser(users, "smith")
	fmt.Println(users)

}
