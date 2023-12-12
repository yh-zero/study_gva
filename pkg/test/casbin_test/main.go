package main

import (
	"fmt"
	"github.com/casbin/casbin/v2"
)

func main() {
	e, err := casbin.NewEnforcer("./model.conf", "./policy.csv")

	sub := "zhangsan" // 想要访问资源的用户。
	obj := "data1"    // 将被访问的资源。
	act := "read"     // 用户对资源执行的操作。

	ok, err := e.Enforce(sub, obj, act)

	if err != nil {
		// 处理err
		fmt.Printf("------- err： %s", err)
	}

	if ok == true {
		// 允许alice读取data1
		fmt.Println("通过")

	} else {
		// 拒绝请求，抛出异常
		fmt.Println("未通过")
	}
}
