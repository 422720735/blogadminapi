/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-10-08 16:45:55
 * @LastEditTime: 2019-10-08 18:00:52
 * @LastEditors: Please set LastEditors
 */
package servers

import (
	"blogadminapi/dbops"
	"blogadminapi/model"

	"github.com/astaxie/beego/logs"
)

func SeleltUsers(user, pwd string) (*model.User, error) {
	// 查询用户，密码
	// SELECT username, `password` from tb_user WHERE username = 'admin' and  `password` = 'e10adc3949ba59abbe56e057f20f883e';
	stmtOut, err := dbops.DbConn.Prepare("select username, `password` from tb_user where username = ? and `password` = ?")
	res := new(model.User)
	if err != nil {
		logs.Error("login is user err", err.Error())
		return res, err
	}
	err = stmtOut.QueryRow(&user, &pwd).Scan(&res.Username, &res.Password)
	if err != nil {
		logs.Error("select user err", err.Error())
		return res, err
	}
	return res, nil
}
