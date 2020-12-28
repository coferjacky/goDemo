package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

//定义全局的pool
var pool *redis.Pool

func init() {

	//当启动程序时，就初始化连接池
	pool = &redis.Pool{
		MaxIdle:     8,
		MaxActive:   0,
		IdleTimeout: 100,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "localhost:6379")
		},
	}

}
func main() {
	//先从pool取出一个连接
	conn := pool.Get()
	defer conn.Close()

	_, err := conn.Do("Set", "name", "tomcat")
	if err != nil {
		fmt.Println(err)
		return
	}
	getName, err := redis.String(conn.Do("Get", "name"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(getName)
}
