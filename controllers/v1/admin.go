package v1

import (
	"blogadminapi/common"
	"blogadminapi/servers"
	"log"

	"github.com/gin-gonic/gin"
)

func GetSystemOrg(c *gin.Context) {
	res, err := servers.GetSystem()
	if err != nil {
		log.Printf("Error in AddNewVideo: %s", err)
		// 打印日志
	}
	common.Echo(c, 200, res)
}

func SetSystemOrg(c *gin.Context) {
	msg, err := common.Unmarshal(c)
	if err != nil {
		common.Echo(c, common.G_ParamErr, "参数不合法")
		return
	}
	// 获取post body参数
}
