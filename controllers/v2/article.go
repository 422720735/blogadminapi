package v2

import (
	"blogadminapi/common"
	"blogadminapi/servers"
	"github.com/gin-gonic/gin"
	"strconv"
)

/**
分页查询文章
id传递的是0 或者没有传递 查询全部，分页查询
返回查询到的总条数
*/
func GetArticleList(c *gin.Context) {
	idStr := c.Query("id")
	pageSizeStr := c.Query("pageSize")
	currentStr := c.Query("current")
	if idStr == "" || pageSizeStr == "" || currentStr == "" {
		common.Echo(c, common.G_ParamErr, "参数不合法")
		return
	}
	id, _ := strconv.Atoi(idStr)
	pageSize, _ := strconv.Atoi(pageSizeStr)
	current, _ := strconv.Atoi(currentStr)
	servers.GetArticleLimitList(id, pageSize, current)
	common.Echo(c, common.G_Success, gin.H{
		"id": id,
		"current": current,
		"pageSize": pageSize,
	})

}
