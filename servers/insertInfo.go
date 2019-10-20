package servers

import (
	"blogadminapi/dbops"
	"github.com/astaxie/beego/logs"
)

/**
SELECT * from tb_post where is_top = 1 UNION all SELECT * from tb_post where is_top = 0 ORDER BY id desc LIMIT 1, 8;
*/
// 普通当前数据不置顶
//  置顶为 ture == 1，否则为false == 0。
func InsertArticle(title, tag, url, image, content string, cid int, isTop bool) error {
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

func UpdateArticle(id int, title, tag, url, image, content string, cid int, isTop bool) error {
	stmtUp, err := dbops.DbConn.Prepare(`
		UPDATE
		 tb_post
		SET
		  title = ?, 
			url = ?,
			content = ?, 
			tags = ?, 
			is_top = ?,
			updated = now(),
			category_id = ?,
			image = ?
		  WHERE id = ?
	`)
	//stmtUp, err:= dbops.DbConn.Prepare("UPDATE tb_post SET `name` = ?, updated = now() WHERE `id` = ?")
	if err != nil {
		logs.Critical("sql拼写错误", err.Error())
		return err
	}
	_, err = stmtUp.Exec(&title, &url, &content, &tag, &isTop, &cid, &image, &id)
	if err != nil {
		logs.Warning("编辑文章错误", err.Error())
		return err
	}
	defer stmtUp.Close()
	return nil
}
