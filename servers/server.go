package servers

import (
	"blogadminapi/model"
	"fmt"
)

func GetSystem() (*model.SystemConfig, error) {

	var res *model.SystemConfig
	if err != nil {
		fmt.Println(err)
		return res, err
	}

	for stmtOut.Next() {
		var id int
		var name, value string
		if err := stmtOut.Scan(&id, &name, &value); err != nil {
			return res, err
		}

		c := &model.SystemConfig{Id: id, Name: name, Value: value}
		res = append(res, c)
	}

	fmt.Println(res)

	defer stmtOut.Close()
	return res, nil

}
