package middleware

import (
	"Go_Practice/app/tool"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PermissionCheck(permisssion string) gin.HandlerFunc {
	return func(c *gin.Context) {

		var company = c.GetHeader("company")
		if company == "" {
			c.JSON(http.StatusOK, "header miss company")
			c.Abort()
			return
		}
		org := "it"
		depart := "rd1"
		method := c.Request.Method

		// 从数据库中读取&判断
		// 加载策略规则
		err := tool.Enforcer.LoadPolicy()
		if err != nil {
			c.JSON(http.StatusOK, "loadPolicy error")
			panic(err)
		}

		// 验证策略规则
		result, err := tool.Enforcer.Enforce(org, depart, permisssion, method)
		if err != nil {
			c.JSON(http.StatusOK, "No permission found")
			c.Abort()
			return
		}

		fmt.Println(":===>", result)

		if !result {
			// TODO 添加到缓存中
			c.JSON(http.StatusOK, "access denied")
			c.Abort()
			return
		}

		c.Next()
	}
}
