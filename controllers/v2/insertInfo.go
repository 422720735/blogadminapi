/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-10-12 10:17:37
 * @LastEditTime: 2019-10-12 18:00:57
 * @LastEditors: Please set LastEditors
 */
package v2

import (
	"blogadminapi/common"
	"blogadminapi/servers"
	"blogadminapi/transform"

	"github.com/astaxie/beego/logs"
	"github.com/gin-gonic/gin"
)

func AddOfUpdateArticleInfo(c *gin.Context) {
	msg, err := common.Unmarshal(c)
	if err != nil {
		common.Echo(c, common.G_ParamErr, "参数不合法")
		return
	}
	id, _ := transform.InterToInt(msg["id"])
	title, err := transform.InterToString(msg["title"])
	if err != nil || title == "" {
		logs.Warning("标题不正确", err.Error())
		common.Echo(c, common.G_ParamErr, "标题不正确")
		return
	}
	categoryId, err := transform.InterToInt(msg["categoryId"])
	if err != nil {
		logs.Warning("获取categoryId失败", err.Error())
		common.Echo(c, common.G_ParamErr, "获取categoryId失败")
		return
	}

	url, err := transform.InterToString(msg["url"])
	if err != nil {
		logs.Warning("获取url失败", err.Error())
		common.Echo(c, common.G_ParamErr, "获取url失败")
		return
	}

	isTop, err := transform.InterToBool(msg["isTop"])
	if err != nil {
		logs.Warning("获取isTop失败", err.Error())
		common.Echo(c, common.G_ParamErr, "获取isTop失败")
		return
	}

	tags, err := transform.InterToString(msg["tags"])
	if err != nil {
		logs.Warning("获取tags失败", err.Error())
		common.Echo(c, common.G_ParamErr, "获取tags失败")
		return
	}

	image, err := transform.InterToString(msg["image"])
	if err != nil {
		logs.Warning("获取image失败", err.Error())
		common.Echo(c, common.G_ParamErr, "获取image失败")
		return
	}

	content, err := transform.InterToString(msg["content"])
	if err != nil {
		logs.Warning("获取content失败", err.Error())
		common.Echo(c, common.G_ParamErr, "获取content失败")
		return
	}

	if id > 0 {
		err := servers.UpdateArticle(id, title, tags, url, image, content, categoryId, isTop)
		if err != nil {
			logs.Error("update article err", err.Error())
			common.Echo(c, common.G_ParamErr, "修改数据失败")
			return
		}
		common.Echo(c, common.G_Success, "修改数据成功")
	} else {
		err = servers.InsertArticle(title, tags, url, image, content, categoryId, isTop)
		if err != nil {
			logs.Error("add article err", err.Error())
			common.Echo(c, common.G_ParamErr, "新增数据失败")
			return
		}
		common.Echo(c, common.G_Success, "新增数据成功")
	}

}
