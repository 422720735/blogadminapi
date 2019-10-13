/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-09-12 15:20:07
 * @LastEditTime: 2019-10-12 16:59:33
 * @LastEditors: Please set LastEditors
 */
package routers

import (
	cv1 "blogadminapi/controllers/v2"
	"blogadminapi/middleware/jkt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(Cors())
	// 登陆
	r.POST("/api/admin/login", cv1.Login)
	v2 := r.Group("/api/admin/v2")
	v2.Use(jkt.JWTAuth())
	{
		v2.GET("/system/get", cv1.GetSystemOrg)
		v2.POST("/system/post", cv1.SetSystemOrg)
		v2.GET("/tag/get", cv1.GetTag)
		v2.GET("/tag/del", cv1.DelTag)
		v2.POST("/tag/post", cv1.SetTag)
		v2.GET("/articleList/get", cv1.GetArticleList)
		v2.POST("/article/info", cv1.Upload)
		v2.POST("/article/info/add", cv1.AddArticleInfo)
		v2.POST("/article/isTop/update", cv1.UpdateArticleIstop)
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
