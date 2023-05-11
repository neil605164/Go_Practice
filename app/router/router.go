package router

import (
	"Go_Practice/app/handler"
	"Go_Practice/app/middleware"
	"Go_Practice/app/tool"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ProviderRouter(r *gin.Engine) {
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"code": 400, "message": "Bad Request"})
	})
	r.GET("/", middleware.PermissionCheck, handler.GetAdmin)

	auth := r.Group("/api")
	{
		// 模拟添加一条Policy策略
		auth.POST("acs", func(c *gin.Context) {
			subject := "tom"
			object := "/api/routers"
			action := "POST"
			result, err := tool.Enforcer.AddPolicy(subject, object, action)
			if err != nil {
				c.JSON(http.StatusOK, err)
				return
			}
			if !result {
				c.JSON(http.StatusOK, "fail")
				return
			}
			c.JSON(http.StatusOK, "success")
		})
		// 模拟删除一条Policy策略
		auth.DELETE("acs", func(c *gin.Context) {
			result, err := tool.Enforcer.RemovePolicy("tom", "/api/routers", "POST")
			if err != nil {
				c.JSON(http.StatusOK, err)
				return
			}
			if !result {
				c.JSON(http.StatusOK, "fail")
				return
			}
			c.JSON(http.StatusOK, "success")
		})
		// 获取路由列表
		auth.POST("/routers", middleware.PermissionCheck, func(c *gin.Context) {
			type data struct {
				Method string `json:"method"`
				Path   string `json:"path"`
			}
			var datas []data
			routers := r.Routes()
			for _, v := range routers {
				var temp data
				temp.Method = v.Method
				temp.Path = v.Path
				datas = append(datas, temp)
			}

			c.JSON(http.StatusOK, datas)
			return
		})
	}

	// 定义路由组
	user := r.Group("/api/v1")
	// 使用访问控制中间件
	user.Use(middleware.PermissionCheck)
	{
		user.POST("user", func(c *gin.Context) {
			c.JSON(200, gin.H{"code": 200, "message": "user add success"})
		})
		user.DELETE("user/:id", func(c *gin.Context) {
			id := c.Param("id")
			c.JSON(200, gin.H{"code": 200, "message": "user delete success " + id})
		})
		user.PUT("user/:id", func(c *gin.Context) {
			id := c.Param("id")
			c.JSON(200, gin.H{"code": 200, "message": "user update success " + id})
		})
		user.GET("user/:id", func(c *gin.Context) {
			id := c.Param("id")
			c.JSON(200, gin.H{"code": 200, "message": "user Get success " + id})
		})
	}
}
