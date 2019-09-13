package v2

import (
	"blogadminapi/agin"
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
		common.Echo(c, common.G_ParamErr, "查詢失敗")
		return
	}
	// 數據再次加工，當前是個切片，轉化成map
	data := agin.Process(res)
	common.Echo(c, common.G_Success, data)
}

func SetSystemOrg(c *gin.Context) {
	_, err := common.Unmarshal(c)
	if err != nil {
		common.Echo(c, common.G_ParamErr, "参数不合法")
		return
	}
	// 获取post body参数
	common.Echo(c, common.G_Success, "123")
}
