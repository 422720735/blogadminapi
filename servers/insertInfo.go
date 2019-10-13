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

/**
SELECT * from tb_post where is_top = 1 UNION all SELECT * from tb_post where is_top = 0 ORDER BY id desc LIMIT 1, 8;
*/
// 普通当前数据不置顶
//  置顶为 ture == 1，否则为false == 0。
func InsertAritcle(title, tag, url, image, content string, cid int, isTop bool) error {
	var top int
	if isTop {
		top = 1
	} else {
		top = 0
	}
	stmtIns, err := dbops.DbConn.Prepare(`
		INSERT INTO tb_post(
			user_id, 
			title, 
			url,
			content, 
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
		return err
	}
	_, err = stmtIns.Exec(1, &title, &url, &content, &tag, 0, 0, &top, &cid, nil, nil, &image)
	if err != nil {
		return err
	}
	defer stmtIns.Close()
	return nil
}
