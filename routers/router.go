package routers

import (
	cv1 "blogadminapi/controllers/v1"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/api/admin/v1")
	{
		v1.GET("/system", cv1.GetSystemOrg)
		v1.POST("/system", cv1.SetSystemOrg)
	}

	return r
}
