/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-10-08 16:07:12
 * @LastEditTime: 2019-10-08 18:00:17
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

func Login(c *gin.Context) {
	msg, err := common.Unmarshal(c)
	if err != nil {
		logs.Alert("login not params unknown err", err.Error())
		common.Echo(c, common.G_ParamErr, "查詢失敗")
		return
	}

	username, err := transform.InterToString(msg["username"])
	if err != nil {
		common.Echo(c, common.G_ParamErr, "参数不合法")
		return
	}

	password, err := transform.InterToString(msg["password"])
	if err != nil {
		common.Echo(c, common.G_ParamErr, "参数不合法")
		return
	}

	if username == "" || password == "" {
		common.Echo(c, common.G_ParamErr, "用户名或密码不能为空")
		return
	}
	user, err := servers.SeleltUsers(username, " e10adc3949ba59abbe56e057f20f883e")
	if err != nil {
		logs.Alert("登录查询失败", err.Error())
		common.Echo(c, common.G_ParamErr, "登录失败")
		return
	}
	if user.Password == " e10adc3949ba59abbe56e057f20f883e" && user.Username == username {
		// 保存到sessign token
		common.Echo(c, common.G_Success, "登录成功")
		return
	}
	common.Echo(c, common.G_ParamErr, "登录失败")
}
