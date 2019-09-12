package routers

import (
	v1 "blogadminapi/controllers/v1"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api/admin/v1", v1.GetSystemOrg)
	{
		v1.GET("/system")
	}

	return r
}
