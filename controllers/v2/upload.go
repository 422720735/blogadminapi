/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-10-11 14:14:04
 * @LastEditTime: 2019-10-12 10:31:35
 * @LastEditors: Please set LastEditors
 */
package v2

import (
	"blogadminapi/common"
	"io"
	"log"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/astaxie/beego/logs"

	"github.com/gin-gonic/gin"
)

func Upload(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		logs.Error("formdata", err.Error())
		common.Echo(c, common.G_ParamErr, err.Error())
		return
	}

	ext := path.Ext(header.Filename)
	name := strconv.FormatInt(time.Now().Unix(), 10)

	filename := "./assets/" + name + ext

	out, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
		common.Echo(c, common.G_ParamErr, "上传失败")
		return
	}

	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
		common.Echo(c, common.G_ParamErr, "上传失败")
		return
	}

	common.Echo(c, common.G_Success, "/assets/"+name+ext)
}
