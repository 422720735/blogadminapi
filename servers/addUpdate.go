/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-09-16 14:03:09
 * @LastEditTime: 2019-09-23 10:36:05
 * @LastEditors: Please set LastEditors
 */
package servers

import (
	"blogadminapi/dbops"

	"github.com/astaxie/beego/logs"
)

func GetCountTag(name string) (int, error) {
	stmtOut, err := dbops.DbConn.Prepare("SELECT COUNT(`name`) FROM tb_category WHERE `name` = ?")
	if err != nil {
		return 0, err
	}
	var num int
	err = stmtOut.QueryRow(name).Scan(&num)
	if err != nil {
		return num, err
	}
	defer stmtOut.Close()
	return num, nil
}

// 新增方法
func Inset(n string) (bool, error) {
	stmtIns, err := dbops.DbConn.Prepare("INSERT INTO tb_category(`name`, created, updated) value (?, NOW(), NOW())")
	if err != err {
		return false, err
	}
	_, err = stmtIns.Exec(n)

	if err != nil {
		return false, err
	}
	defer stmtIns.Close()
	return true, nil
}

// 编辑方法
func Update(n string, id int) (bool, error) {
	// UPDATE tb_category SET `name`  = 'aes', updated WHERE id = 4;
	stmtUpdate, err := dbops.DbConn.Prepare("UPDATE tb_category SET `name` = ?, updated = now() WHERE `id` = ?")
	if err != nil {
		logs.Warning("update tag not sql err", err.Error())
		return false, err
	}
	_, err = stmtUpdate.Exec(n, id)
	if err != nil {
		logs.Warning("update tag not exec err", err.Error())
		return false, err
	}
	defer stmtUpdate.Close()
	return true, nil
}

func DelteleTag(id int) error {
	stmtUpdate, err := dbops.DbConn.Prepare("update tb_category set status = ? where id = ?")
	if err != nil {
		logs.Warning("delete tag not sql err", err.Error())
		return err
	}
	_, err = stmtUpdate.Exec(1, id)
	if err != nil {
		logs.Alert("soft delete tag not err", err.Error())
		return err
	}
	defer stmtUpdate.Close()
	return nil

}
