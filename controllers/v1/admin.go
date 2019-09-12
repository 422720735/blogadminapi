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
