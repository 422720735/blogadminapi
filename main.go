package main

import (
	"blogadminapi/routers"
	"github.com/gin-gonic/gin"
)

func init()  {
	gin.SetMode(gin.DebugMode)
}
func main() {
	r := routers.InitRouter()
	r.Run(":4000")
}
