/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-09-16 10:50:30
 * @LastEditTime: 2019-09-23 11:16:48
 * @LastEditors: Please set LastEditors
 */
package servers

import (
	"blogadminapi/dbops"
	"blogadminapi/model"
	"log"
	"time"
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

func SetSystemConfig(title, url, keywords, description, email, qq, start string) error {
	/*
		UPDATE tb_config
		SET
		`value` = '556'
		WHERE
		`name` = 'title';
	*/
	if title != "" {
		stmt, err := dbops.DbConn.Prepare("update tb_config set `value` = ? where `name` = 'title' ")
		defer stmt.Close()
		if err != nil {
			log.Print(err)
			return err
		}
		_, err = stmt.Exec(title)
		if err != nil {
			return err
		}
	}

	if url != "" {
		stmt, err := dbops.DbConn.Prepare("update tb_config set `value` = ? where `name` = 'url' ")
		defer stmt.Close()
		if err != nil {
			log.Print(err)
			return err
		}
		_, err = stmt.Exec(url)
		if err != nil {
			return err
		}
	}

	if keywords != "" {
		stmt, err := dbops.DbConn.Prepare("update tb_config set `value` = ? where `name` = 'keywords' ")
		defer stmt.Close()
		if err != nil {
			log.Print(err)
			return err
		}
		_, err = stmt.Exec(keywords)
		if err != nil {
			return err
		}
	}

	if description != "" {
		stmt, err := dbops.DbConn.Prepare("update tb_config set `value` = ? where `name` = 'description' ")
		defer stmt.Close()
		if err != nil {
			log.Print(err)
			return err
		}
		_, err = stmt.Exec(description)
		if err != nil {
			return err
		}
	}

	if email != "" {
		stmt, err := dbops.DbConn.Prepare("update tb_config set `value` = ? where `name` = 'email' ")
		defer stmt.Close()
		if err != nil {
			log.Print(err)
			return err
		}
		_, err = stmt.Exec(email)
		if err != nil {
			return err
		}
	}

	stmt, err := dbops.DbConn.Prepare("update tb_config set `value` = ? where `name` = 'start' ")
	defer stmt.Close()
	if err != nil {
		log.Print(err)
		return err
	}
	_, err = stmt.Exec(start)
	if err != nil {
		return err
	}

	return nil
}

func GetCategory() ([]*model.Category, error) {
	rows, err := dbops.DbConn.Query("SELECT id, `name`, created, updated FROM tb_category where status = ?", 0)
	if err != nil {
		return nil, err
	}
	var res []*model.Category
	for rows.Next() {
		var name string
		var id int
		var created, updated time.Time
		if err := rows.Scan(&id, &name, &created, &updated); err != nil {
			return res, err
		}
		cat := &model.Category{Id: id, Name: name, Created: created.Unix(), Updated: updated.Unix()}
		res = append(res, cat)
	}
	defer rows.Close()
	return res, nil
}
