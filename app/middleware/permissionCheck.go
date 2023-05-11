package middleware

import (
	"Go_Practice/app/tool"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PermissionCheck(c *gin.Context) {

	var userName = c.GetHeader("userName")
	if userName == "" {
		c.JSON(http.StatusOK, "header miss userName")
		c.Abort()
		return
	}
	path := c.Request.URL.Path
	method := c.Request.Method

	// 从数据库中读取&判断
	// 加载策略规则
	err := tool.Enforcer.LoadPolicy()
	if err != nil {
		c.JSON(http.StatusOK, "loadPolicy error")
		panic(err)
	}
	// 验证策略规则
	result, err := tool.Enforcer.Enforce(userName, path, method)
	if err != nil {
		c.JSON(http.StatusOK, "No permission found")
		c.Abort()
		return
	}
	if !result {
		// TODO 添加到缓存中
		c.JSON(http.StatusOK, "access denied")
		c.Abort()
		return
	}

	c.Next()
}
