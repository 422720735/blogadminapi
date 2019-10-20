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

/*
* 分页返回的数据
@params: allCount   总数量
@params: count   	总页数
@params: pageSize   请求每页显示的数量

@params: current    当前页
*/
func Page(allCount, count, pageSize, current int, inter interface{}) map[string]interface{} {
	data := make(map[string]interface{})
	data["total"] = allCount
	data["pageSize"] = pageSize
	data["count"] = count
	data["current"] = current
	data["data"] = inter

	return data
}
