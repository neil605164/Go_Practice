package main

import (
	"log"

	"github.com/casbin/casbin/v2"
)

func main() {
	// /depts
	sub := "neil"         // 想要访问资源的用户。
	company := "companyc" // 将被访问的资源。
	org := "orgc"
	depart := "departc"
	act := "GetUserList" // 用户对资源执行的操作。
	e, _ := casbin.NewEnforcer("resources/model.conf", "resources/p.csv")

	ok, _ := e.Enforce(sub, company, org, depart, act)
	if ok {
		log.Println("运行通过")
	} else {
		log.Println("运行不通过")
	}

}
