package common

import (
	"github.com/gin-gonic/gin"
	js "blogadminapi/json"
	"io/ioutil"
)

//获取body
func Unmarshal(c *gin.Context) (map[string]interface{}, error) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return nil, err
	}
	msg := make(map[string]interface{})
	// err = json.Unmarshal(body, &msg)
	msg, err = js.Unmarshal(body)
	if err != nil {
		return nil, err
	}
	return msg, nil
}
