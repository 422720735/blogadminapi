package qny

import (
	"blogadminapi/lib"
	"github.com/astaxie/beego/logs"
	"github.com/gin-gonic/gin"
	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
	"io"
)

func LoadQiNiu(c *gin.Context, key string, data io.Reader, size int64) (string, error) {
	AccessKey := lib.Conf.Read("qiniu", "AccessKey")
	SecretKey := lib.Conf.Read("qiniu", "SecretKey")
	bucket := lib.Conf.Read("qiniu", "Scope")
	//上传凭证,关于凭证这块大家可以去看看官方文档
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := qbox.NewMac(AccessKey, SecretKey)
	upToken := putPolicy.UploadToken(mac)
	cfg := storage.Config{}
	//七牛云存储空间设置首页有存储区域
	cfg.Zone = &storage.ZoneHuanan
	//不启用HTTPS域名
	cfg.UseHTTPS = false
	//不使用CND加速
	cfg.UseCdnDomains = false
	//构建上传表单对象
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	// 可选
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": "github logo",
		},
	}
	err := formUploader.Put(c, &ret, upToken, key, data, size, &putExtra)
	if err != nil {
		logs.Critical("图片上传七牛云失败", err.Error())
		return "", err
	}
	return "/" + ret.Key, nil
}
