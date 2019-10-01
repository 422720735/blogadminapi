package main

import (
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
	r.Run(":4000")
}
