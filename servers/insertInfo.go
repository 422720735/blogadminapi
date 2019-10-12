/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-10-12 17:02:49
 * @LastEditTime: 2019-10-12 17:55:55
 * @LastEditors: Please set LastEditors
 */
package servers

import (
	"blogadminapi/dbops"
)

// 普通当前数据不置顶
func OrdinaryInsertAritcle(title, tag, image, content string, cid, isTop int) {
	stmtIns, err := dbops.DbConn.Prepare(`
		INSERT INTO tb_post(
			user_id, 
			'title', 
			'url',
			'content', 
			tags, 
			views,
			status,
			is_top,
			created,
			updated,
			category_id,
			types,
			info,
			image) value (
				?,
				?,
				?,
				?,
				?,
				?,
				?,
				?,
				now(),
				now(),
				?,
				?,
				?,
				?)
	`)
	if err != nil {
		return
	}
	stmtIns.Exec(1, &title, &url, &content, &tag, 0, 0, &isTop, &cid, nil, nil, &image)
	defer stmtIns.Close()
}

// 当前数据置顶的
func IsTopInsertAritcle(title, tag, image, content string, cid, isTop int) {

}
