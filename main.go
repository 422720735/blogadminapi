package main

import (
	"blogadminapi/lib"
	"blogadminapi/logger"
	"blogadminapi/routers"

	"github.com/gin-gonic/gin"
)

func init() {
	gin.SetMode(gin.DebugMode)
	logger.Init()
}

func main() {
	r := routers.InitRouter()
	port := lib.Conf.Read("site", "httpport")
	r.Run(":" + port)
}
