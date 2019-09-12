package common

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// 返回的数据
func Echo(c *gin.Context, code int, body interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  body,
		"time": time.Now().Unix(),
	})
}
