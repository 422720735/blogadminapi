package common

import (
	"blogadminapi/json"
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
@params: num        返回的数量
@params: pageSize   请求每页显示的数量
@params: current    当前页
 */
func Page(allCount, num, pageSize, current int, inter interface{})([]byte, error) {
	data:=make(map[string]interface{})
	data["total"] = allCount
	data["num"] = num
	data["pageSize"] = pageSize
	data["current"] = current
	data["data"] = inter
	return json.Marshal(data)
}
