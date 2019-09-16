package servers

import (
	"blogadminapi/dbops"
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
		return false, err
	}
	_, err = stmtUpdate.Exec(n, id)
	if err != nil {
		return false, err
	}
	defer stmtUpdate.Close()
	return true, nil
}