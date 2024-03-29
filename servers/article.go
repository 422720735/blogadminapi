package servers

import (
	"blogadminapi/dbops"
	"blogadminapi/model"
	"database/sql"
	"math"
	"time"

	"github.com/astaxie/beego/logs"
)

/** 查询分页数据*/
func GetArticleLimitList(id, pageSize, current int, keyword string) (int, int, []*model.PostListRes, error) {
	var sq, idSql1, keySql string
	sq = "select count(`id`) from tb_post where status = 0"
	if id > 0 {
		idSql1 = " and category_id = ?"
	}
	if keyword != "" && id > 0 {
		keySql = " and title LIKE ? "
	} else if keyword != "" && id <= 0 {
		keySql = " and title LIKE ?"
	}
	sq = sq + idSql1 + keySql
	stmtOutCount, err := dbops.DbConn.Prepare(sq)
	var res []*model.PostListRes
	var total int
	if err != nil {
		logs.Error("article count", err.Error())
		return 0, 0, nil, err
	}
	if id > 0 && keyword != "" {
		stmtOutCount.QueryRow(id, "%"+keyword+"%").Scan(&total)
	} else if id > 0 && keyword == "" {
		stmtOutCount.QueryRow(id).Scan(&total)
	} else if id <= 0 && keyword != "" {
		stmtOutCount.QueryRow("%" + keyword + "%").Scan(&total)
	} else {
		stmtOutCount.QueryRow().Scan(&total)
	}
	stmtOutCount.Close()

	// 查询分页数据
	// stmtLimt, err := dbops.DbConn.Query("select id, title, tags, is_top, created, updated from tb_post order by id desc limit ?, ?", (current-1)*pageSize, pageSize)

	var baseSql, limitSql, idSql, likeSql string
	baseSql = "SELECT id, title, tags, is_top, created, updated, views, category_id from tb_post where is_top = 1 and status = 0 UNION all SELECT id, title, tags, is_top, created, updated, views, category_id from tb_post where is_top = 0 and status = 0"
	limitSql = " order by id desc limit ?, ?"
	if id > 0 {
		idSql = " and category_id = ?"
	}
	if keyword != "" && id > 0 {
		likeSql = " and title like ?"
	} else if keyword != "" && id <= 0 {
		likeSql = " title like ?"
	}

	sq = baseSql + idSql + likeSql + limitSql

	stmtLimt, err := dbops.DbConn.Prepare(sq)
	var row *sql.Rows
	var e error
	if id > 0 && keyword != "" {
		row, e = stmtLimt.Query(id, "%"+keyword+"%", (current-1)*pageSize, pageSize)
	} else if id > 0 && keyword == "" {
		row, e = stmtLimt.Query(id, (current-1)*pageSize, pageSize)
	} else if id <= 0 && keyword != "" {
		row, e = stmtLimt.Query("%"+keyword+"%", (current-1)*pageSize, pageSize)
	} else {
		row, e = stmtLimt.Query((current-1)*pageSize, pageSize)
	}
	if e != nil {
		logs.Error("article limit sq error", err.Error())
		return 0, 0, res, err
	}

	for row.Next() {
		ar := new(model.PostList)
		if err := row.Scan(
			&ar.Id,
			&ar.Title,
			&ar.Tags,
			&ar.IsTop,
			&ar.Created,
			&ar.Updated,
			&ar.Views,
			&ar.CategoryId,
		); err != nil {
			logs.Error("装填数据居然失败")
			return 0, 0, res, err
		}
		var top bool
		if ar.IsTop == 1 {
			top = true
		} else {
			top = false
		}
		cat := &model.PostListRes{
			Id:         ar.Id,
			Title:      ar.Title,
			Tags:       ar.Tags,
			IsTop:      top,
			Created:    ar.Created.Unix(),
			Updated:    ar.Updated.Unix(),
			Views:      ar.Views,
			CategoryId: ar.CategoryId,
		}
		res = append(res, cat)
	}
	stmtLimt.Close()
	count := int(math.Ceil(float64(total) / float64(pageSize)))
	return total, count, res, err
}

func UpdateArticleIstop(id int, isTop bool) error {
	stmtUpdate, err := dbops.DbConn.Prepare("UPDATE tb_post SET `is_top` = ? WHERE id = ?")
	if err != nil {
		return err
	}
	var top int
	if isTop {
		top = 1
	} else {
		top = 0
	}
	_, err = stmtUpdate.Exec(&top, &id)
	if err != nil {
		return err
	}
	defer stmtUpdate.Close()
	return nil
}

/**
获取文章详情
SELECT first_name, COALESCE(age, 0) FROM person;
使用IFNULL或者COALESCE
*/
func GetArticleInfo(id int) (*model.PostInfo, error) {
	stmtOut, err := dbops.DbConn.Prepare(`
		SELECT 
		id, user_id, title, url, content, tags, views, is_top, created, updated, category_id,
		COALESCE(types, -1), COALESCE(info, ''), COALESCE(image, '') 
		FROM tb_post WHERE id = ?
	`)

	if err != nil {
		logs.Warning("sql", err.Error())
		return nil, err
	}
	//ar.Created.Unix(),
	post := new(model.PostInfo)
	var c time.Time
	var u time.Time
	err = stmtOut.QueryRow(id).Scan(&post.Id, &post.UserId, &post.Title, &post.Url, &post.Content, &post.Tags, &post.Views, &post.IsTop, &c, &u, &post.CategoryId, &post.Types, &post.Info, &post.Image)
	post.Created = c.Unix()
	post.Updated = u.Unix()
	if err != nil {
		logs.Error("查询文章sql", err.Error())
		return nil, err
	}

	defer stmtOut.Close()
	return post, err
}

func DelArticleInfo(id int) error {
	stmtDel, err := dbops.DbConn.Prepare("update tb_post set `status` = 1 where id = ?")
	if err != nil {
		logs.Critical("删除文章sql", err.Error())
		return err
	}
	_, err = stmtDel.Exec(&id)
	if err != nil {
		logs.Critical("删除文章sql传递的id不正确", err.Error())
		return err
	}
	defer stmtDel.Close()
	return nil
}
