package servers

import (
	"blogadminapi/dbops"
	"blogadminapi/model"
)

func GetSystem() ([]*model.SystemConfig, error) {
	rows, err := dbops.DbConn.Query("SELECT id, `name`, `value` FROM tb_config")
	if err != nil {
		return nil, err
	}
	var res []*model.SystemConfig
	for rows.Next() {
		var name, value string
		var id int
		if err := rows.Scan(&id, &name, &value); err != nil {
			return res, err
		}
		c := &model.SystemConfig{Id: id, Name: name, Value: value}
		res = append(res, c)
	}
	defer rows.Close()
	return res, nil

}
