/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-09-23 09:31:19
 * @LastEditTime: 2019-09-24 14:24:19
 * @LastEditors: Please set LastEditors
 */
package v2

import (
	"blogadminapi/common"
	"blogadminapi/servers"
	"strconv"

	"github.com/astaxie/beego/logs"

	"github.com/gin-gonic/gin"
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
	keyword := c.Query("keyword")
	if idStr == "" || pageSizeStr == "" || currentStr == "" {
		logs.Warning("分页参数不合法")
		common.Echo(c, common.G_ParamErr, "参数不合法")
		return
	}
	id, _ := strconv.Atoi(idStr)
	pageSize, _ := strconv.Atoi(pageSizeStr)
	current, _ := strconv.Atoi(currentStr)
	total, count, res, err := servers.GetArticleLimitList(id, pageSize, current, keyword)
	if err != nil {
		logs.Error(err.Error())
		common.Echo(c, common.G_ParamErr, "查询数据失败")
		return
	}
	// 组装分页数据
	data := common.Page(total, count, pageSize, current, res)
	common.Echo(c, common.G_Success, data)
}
