/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-09-23 09:31:19
 * @LastEditTime: 2019-10-08 17:09:53
 * @LastEditors: Please set LastEditors
 */
package servers

import (
	"blogadminapi/dbops"
	"blogadminapi/model"
	"database/sql"
	"fmt"
	"math"
	"strings"

	"github.com/astaxie/beego/logs"
)

// import "blogadminapi/dbops"

/** 查询分页数据*/
func GetArticleLimitList(id, pageSize, current int, keyword string) (int, int, []*model.PostListRes, error) {
	var sq, idSql1, keySql string
	sq = "select count(`id`) from tb_post where"
	if id > 0 {
		idSql1 = " category_id = ?"
	}
	if keyword != "" && id > 0 {
		keySql = " and title LIKE ? "
	} else if keyword != "" && id <= 0 {
		keySql = " title LIKE ?"
	}
	sq = sq + idSql1 + keySql
	if id <= 0 && keyword == "" {
		// 替换where
		sq = strings.Replace(sq, "where", "", -1)
		// 去除收尾空格
		sq = strings.TrimSpace(sq)
	}
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
	baseSql = "SELECT id, title, tags, is_top, created, updated, views, category_id from tb_post where is_top = 1 UNION all SELECT id, title, tags, is_top, created, updated, views, category_id from tb_post where is_top = 0"
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
	// if id <= 0 && keyword == "" {
	// 	sq = strings.Replace(sq, "where", "", -1)
	// 	sq = strings.TrimSpace(sq)
	// }

	fmt.Printf(sq)
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
