package main

import (
	libs "gin_crud/lib"
	"gin_crud/routers"
	"net/http"
	"os"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func init() {
	//设置日志格式为json格式　自带的只有两种样式logrus.JSONFormatter{}和logrus.TextFormatter{}
	logrus.SetFormatter(&logrus.JSONFormatter{})
	//指定日志输出
	file, err := os.OpenFile("./xx.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		//fmt.Println("open log file failed, err:", err)
		return
	}
	// 输出日志
	log.Out = file
	gin.SetMode(gin.ReleaseMode)
	// gin的日志也输入
	gin.DefaultWriter = log.Out
	// 设置日志级别
	log.Level = logrus.DebugLevel

}

func main() {
	r := routers.InitRouter()
	r.Use(CorsMiddleware())
	r.Run(":" + libs.Conf.Read("site", "httpport"))
}
func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		var filterHost = [...]string{"http://localhost.*", "http://*.hfjy.com"}
		// filterHost 做过滤器，防止不合法的域名访问
		var isAccess = false
		for _, v := range filterHost {
			match, _ := regexp.MatchString(v, origin)
			if match {
				isAccess = true
			}
		}
		if isAccess {
			// 核心处理方式
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
			c.Header("Access-Control-Allow-Methods", "GET, OPTIONS, POST, PUT, DELETE")
			c.Set("content-type", "application/json")
		}
		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}

		c.Next()
	}
}
