/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-09-16 10:50:30
 * @LastEditTime: 2019-09-16 10:50:30
 * @LastEditors: your name
 */
package v2

import (
	"blogadminapi/agin"
	"blogadminapi/common"
	"blogadminapi/servers"
	"blogadminapi/transform"
	"log"
	"strconv"

	"github.com/astaxie/beego/logs"

	"github.com/gin-gonic/gin"
)

func GetSystemOrg(c *gin.Context) {
	res, err := servers.GetSystem()
	if err != nil {
		log.Printf("Error in AddNewVideo: %s", err)
		// 打印日志
		common.Echo(c, common.G_ParamErr, "查詢失敗")
		return
	}
	// 數據再次加工，當前是個切片，轉化成map
	data := agin.Process(res)
	common.Echo(c, common.G_Success, data)
}

func SetSystemOrg(c *gin.Context) {
	msg, err := common.Unmarshal(c)
	if err != nil {
		common.Echo(c, common.G_ParamErr, "参数不合法")
		return
	}
	// 获取post body参数
	title, err := transform.InterToString(msg["title"])
	if err != nil {
		common.Echo(c, common.G_ParamErr, "标题不合法")
		return
	}
	url, err := transform.InterToString(msg["url"])
	if err != nil {
		common.Echo(c, common.G_ParamErr, "网址不合法")
		return
	}
	keywords, err := transform.InterToString(msg["keywords"])
	if err != nil {
		common.Echo(c, common.G_ParamErr, "关键字不合法")
		return
	}
	description, err := transform.InterToString(msg["description"])
	if err != nil {
		common.Echo(c, common.G_ParamErr, "描述不合法")
		return
	}
	email, err := transform.InterToString(msg["email"])
	if err != nil {
		common.Echo(c, common.G_ParamErr, "邮箱不合法")
		return
	}
	qq, err := transform.InterToString(msg["qq"])
	if err != nil {
		common.Echo(c, common.G_ParamErr, "qq不合法")
		return
	}
	start, err := transform.InterToString(msg["start"])
	if err != nil {
		common.Echo(c, common.G_ParamErr, "开关状态不合法")
		return
	}
	data := servers.SetSystemConfig(title, url, keywords, description, email, qq, start)
	if data != nil {
		common.Echo(c, common.G_ParamErr, "更新失败~！")
		return
	}
	common.Echo(c, common.G_Success, "修改成功~！")
}

func GetTag(c *gin.Context) {
	res, err := servers.GetCategory()
	if err != nil {
		log.Printf("Error in AddNewVideo: %s", err)
		// 打印日志
		common.Echo(c, common.G_ParamErr, "查詢失敗")
		return
	}
	common.Echo(c, common.G_Success, res)
}

func DelTag(c *gin.Context) {
	idStr := c.Query("id")
	if idStr == "" {
		logs.Warning("id is undefined")
		common.Echo(c, common.G_ParamErr, "参数不合法")
		return
	}
	id, _ := strconv.Atoi(idStr)
	err := servers.DelteleTag(id)
	if err != nil {
		common.Echo(c, common.G_ParamErr, "删除失败")
		return
	}
	common.Echo(c, common.G_Success, "删除成功")
}

func SetTag(c *gin.Context) {
	msg, err := common.Unmarshal(c)
	if err != nil {
		log.Printf("Error in AddNewVideo: %s", err)
		// 打印日志
		common.Echo(c, common.G_ParamErr, "查詢失敗")
		return
	}

	name, err := transform.InterToString(msg["name"])
	if err != nil {
		common.Echo(c, common.G_ParamErr, "标题不合法")
		return
	}
	count, err := servers.GetCountTag(name)
	if err != nil {
		common.Echo(c, common.G_ParamErr, "参数不正确")
		return
	}
	if count > 0 {
		common.Echo(c, common.G_ParamErr, "该名称的数据已存在，不能重复添加！~")
		return
	}
	id, err := transform.InterToInt(msg["id"])
	/**
	* 传递了id就是修改没有传递就是新增
	 */
	if id == -1 {
		result, err := servers.Inset(name)
		if result && err == nil {
			common.Echo(c, common.G_Success, "新增成功")
			return
		} else {
			common.Echo(c, common.G_ParamErr, "参数不合法")
			return
		}
	} else {
		result, err := servers.Update(name, id)
		if result && err == nil {
			common.Echo(c, common.G_Success, "编辑成功")
			return
		} else {
			common.Echo(c, common.G_ParamErr, "参数不合法")
			return
		}
	}
}
