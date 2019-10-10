/*
 https://www.jianshu.com/p/1f9915818992
*/
package v2

import (
	"blogadminapi/common"
	myjwt "blogadminapi/middleware/jkt"
	"blogadminapi/servers"
	"blogadminapi/transform"
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/gin-gonic/gin"
)

const (
	ExpireTime = 3600
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
		// 生产token
		token := newToken(c, user.Username, user.Password)
		if token == "" {
			common.Echo(c, common.G_ParamErr, "登录失败")
			return
		}
		common.Echo(c, common.G_Success, token)
		return
	}
	common.Echo(c, common.G_ParamErr, "登录失败")
}

// 生成令牌
func newToken(c *gin.Context, user, pwd string) string {
	//CreateToken
	j := &myjwt.JWT{
		[]byte("newtrekWang"),
	}
	claims := myjwt.CustomClaims{
		ID:       1,
		Username: user,
		Password: pwd,
	}
	claims.IssuedAt = time.Now().Unix()
	claims.ExpiresAt = time.Now().Add(time.Second * time.Duration(ExpireTime)).Unix()
	token, err := j.CreateToken(claims)

	if err != nil {
		return ""
	}

	return token
}
