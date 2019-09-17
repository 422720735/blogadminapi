package middleware

import (
	"encoding/json"
	"fmt"
	"log"
	"runtime"
	"strings"
)

//日志等级
const (
	LEVEL_CRITICAL = iota
	LEVEL_ERROR
	LEVEL_WARNING
	LEVEL_INFO
	LEVEL_DEBUG
)

//日志模式
const (
	MODEL_PRO = iota
	MODEL_INFO
	MODEL_DEV
)

var (
	g_level = map[int]string{
		LEVEL_CRITICAL: "CRITICAL",
		LEVEL_ERROR:    "ERROR",
		LEVEL_WARNING:  "WARNING",
		LEVEL_INFO:     "INFO",
		LEVEL_DEBUG:    "DEBUG",
	}
	level = LEVEL_DEBUG
	model = MODEL_DEV
)

//调用log的服务器名字
var serverName string

/*
*初始化日志logServerAddr日志服务器地址
*参数说明:
*@param:nsqdAddr		nsqd地址
*@param:name			调用log包的服务器地址
 */
func InitLog(nsqdAddr string, name string) error {
	serverName = name
	return InitProducer(nsqdAddr)
}

//设置日志等级
func SetLogLevel(lev int) {
	level = lev
}

/*
*设置日志模式
*参数说明:
*@param:mod		模式 MODEL_PRO:只向日志服务器发送日志  MODEL_INFO:向日志服务器发送日志切输出到控制台 MODEL_DEV:只输出到控制台
 */
func SetLogModel(mod int) {
	model = mod
}

//危险的
func Critical(format string, args ...interface{}) {
	if level >= LEVEL_CRITICAL {
		logStr := logFormat(LEVEL_CRITICAL, fmt.Sprintf(format, args...))
		if model == MODEL_DEV || model == MODEL_INFO {
			log.Println(logStr)
		}
		if model == MODEL_INFO || model == MODEL_PRO {
			//组装一个json格式
			logMsg := make(map[string]string)
			logMsg["name"] = serverName
			logMsg["body"] = logStr
			logMsg["level"] = g_level[LEVEL_CRITICAL]
			send, err := json.Marshal(logMsg)
			if err != nil {
				log.Fatal("marshal err ", err)
			}
			go Publish("log", send)
		}
	}
}

//错误
func Error(format string, args ...interface{}) {
	if level >= LEVEL_ERROR {
		logStr := logFormat(LEVEL_ERROR, fmt.Sprintf(format, args...))
		if model == MODEL_DEV || model == MODEL_INFO {
			log.Println(logStr)
		}
		if model == MODEL_INFO || model == MODEL_PRO {
			//组装一个json格式
			logMsg := make(map[string]string)
			logMsg["name"] = serverName
			logMsg["body"] = logStr
			logMsg["level"] = g_level[LEVEL_ERROR]
			send, err := json.Marshal(logMsg)
			if err != nil {
				log.Fatal(err)
			}
			go Publish("log", send)
		}
	}
}

//警告
func Warn(format string, args ...interface{}) {
	if level >= LEVEL_WARNING {
		logStr := logFormat(LEVEL_WARNING, fmt.Sprintf(format, args...))
		if model == MODEL_DEV || model == MODEL_INFO {
			log.Println(logStr)
		}
		if model == MODEL_INFO || model == MODEL_PRO {
			//组装一个json格式
			logMsg := make(map[string]string)
			logMsg["name"] = serverName
			logMsg["body"] = logStr
			logMsg["level"] = g_level[LEVEL_WARNING]
			send, err := json.Marshal(logMsg)
			if err != nil {
				log.Fatal(err)
			}
			go Publish("log", send)
		}
	}
}

//提示
func Info(format string, args ...interface{}) {
	if level >= LEVEL_INFO {
		logStr := logFormat(LEVEL_INFO, fmt.Sprintf(format, args...))
		if model == MODEL_DEV || model == MODEL_INFO {
			log.Println(logStr)
		}
		if model == MODEL_INFO || model == MODEL_PRO {
			//组装一个json格式
			logMsg := make(map[string]string)
			logMsg["name"] = serverName
			logMsg["body"] = logStr
			logMsg["level"] = g_level[LEVEL_INFO]
			send, err := json.Marshal(logMsg)
			if err != nil {
				log.Fatal(err)
			}
			go Publish("log", send)
		}
	}
}

//调试
func Debug(format string, args ...interface{}) {
	if level >= LEVEL_DEBUG {
		logStr := logFormat(LEVEL_DEBUG, fmt.Sprintf(format, args...))
		if model == MODEL_DEV || model == MODEL_INFO {
			log.Println(logStr)
		}
		if model == MODEL_INFO || model == MODEL_PRO {
			//组装一个json格式
			logMsg := make(map[string]string)
			logMsg["name"] = serverName
			logMsg["body"] = logStr
			send, err := json.Marshal(logMsg)
			if err != nil {
				log.Fatal(err)
			}
			go Publish("log", send)
		}
	}
}

func logFormat(level int, msg string) string {
	fname := ""
	pc, path, line, ok := runtime.Caller(2) // 去掉两层，当前函数和日志的接口函数
	if ok {
		if f := runtime.FuncForPC(pc); f != nil {
			fname = f.Name()
		}
	}
	funcName := lastFname(fname)
	t := time.CurTimeStr()
	path = getFilePath(path)
	format := fmt.Sprintf("%s [%s] %s %s %d ▶ %s", t, g_level[level], path, funcName, line, msg)
	return format
}

func lastFname(fname string) string {
	flen := len(fname)
	n := strings.LastIndex(fname, ".")
	if n+1 < flen {
		return fname[n+1:]
	}
	return fname
}

func getFilePath(path string) string {
	s := strings.Split(path, "src")
	return s[1]
}
