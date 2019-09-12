package json

//json相关的函数
import (
	"encoding/json"
)

//解析json
func Unmarshal(data []byte) (map[string]interface{}, error) {
	msg := make(map[string]interface{})
	err := json.Unmarshal(data, &msg)
	if err != nil {
		return nil, err
	}
	return msg, nil
}

//编码json
func Marshal(data interface{}) ([]byte, error) {
	msg, err := json.Marshal(&data)
	if err != nil {
		return nil, err
	}
	return msg, nil
}
