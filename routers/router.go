package routers

import (
	cv1 "blogadminapi/controllers/v2"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(Cors())
	v1 := r.Group("/api/admin/v2")
	{
		v1.GET("/system/get", cv1.GetSystemOrg)
		v1.POST("/system/post", cv1.SetSystemOrg)
		v1.GET("/tag/get", cv1.GetTag)
		v1.GET("/tag/del", cv1.DelTag)
		v1.POST("/tag/post", cv1.SetTag)
		v1.GET("/articleList/get", cv1.GetArticleList)
	}
	return r
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}
