package main

import (
	"log"

	"github.com/casbin/casbin/v2"
)

func main() {
	// 想要访问资源的用户。
	role := "posa"        // 角色名称，根据角色策略文件选择合适的角色。
	company := "companya" // 将被访问的资源。
	org := "orga"
	act := "GetUserList" // 用户对资源执行的操作。

	e, err := casbin.NewEnforcer("resources/model.conf", "resources/p.csv")
	if err != nil {
		panic(err)
	}
	// Enable debugging
	e.EnableLog(true)

	ok, err := e.Enforce(role, org, company, act)
	if err != nil {
		panic(err)
	}

	if ok {
		log.Println("运行通过")
	} else {
		log.Println("运行不通过")
	}

}
