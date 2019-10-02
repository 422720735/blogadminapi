package v2

import (
	"blogadminapi/common"
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
	common.Echo(c, common.G_Success, "ok")
}
