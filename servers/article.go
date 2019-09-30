/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-09-23 09:31:19
 * @LastEditTime: 2019-09-23 09:31:19
 * @LastEditors: your name
 */
package servers

import (
	"blogadminapi/dbops"
	"blogadminapi/model"
	"database/sql"
	"math"
	"strings"

	"github.com/astaxie/beego"
)

// import "blogadminapi/dbops"

/** 查询分页数据*/
func GetArticleLimitList(id, pageSize, current int, keyword string) (int, int, []*model.PostListRes, error) {
	// id == 0 就表示查询全部
	// SELECT * FROM tb_post WHERE category_id = 2 AND title LIKE '%d%' or url LIKE '%baidu%' ORDER BY id DESC LIMIT 0, 2;
	// select count(`id`) from tb_post
	// select count(`id`) from tb_post where like %%
	// select count(`id`) from tb_post where and
	var sq, idSql1, keySql string
	sq = "select count(`id`) from tb_post where"
	if id > 0 {
		idSql1 = " category_id = ?"
	}
	if keyword != "" && id > 0 {
		keySql = " and titlt LIKE ?"
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
	beego.Info("======", sq)
	stmtOutCount, err := dbops.DbConn.Prepare(sq)
	var res []*model.PostListRes
	var total int
	if err != nil {

		return 0, 0, nil, err
	}

	if id > 0 && keyword != "" {
		beego.Info("123355", id, "%"+keyword+"%")
		stmtOutCount.QueryRow(id, "%"+keyword+"%").Scan(&total)
	} else if id > 0 && keyword == "" {
		stmtOutCount.QueryRow(id).Scan(&total)
	} else if id <= 0 && keyword != "" {
		stmtOutCount.QueryRow("%" + keyword + "%").Scan(&total)
	} else {
		stmtOutCount.QueryRow().Scan(&total)
	}
	stmtOutCount.Close()
	// stmtLimt, err := dbops.DbConn.Query("select id, title, tags, is_top, created, updated from tb_post order by id desc limit ?, ?", (current-1)*pageSize, pageSize)

	if id > 0 {
		sq = "select id, title, tags, is_top, created, updated, views, category_id from tb_post where category_id = ? order by id desc limit ?, ?"
	} else {
		sq = "select id, title, tags, is_top, created, updated, views, category_id from tb_post order by id desc limit ?, ?"
	}

	var baseSql, limitSql, idSql, likeSql string
	baseSql = "select id, title, tags, is_top, created, updated, views, category_id from tb_post where"
	limitSql = " and order by id desc limit ?, ?"
	idSql = ""
	likeSql = ""
	if id > 0 {
		idSql = " category_id = ?"
	}
	if keyword != "" {
		likeSql = " and title like ?"
	}
	sq = baseSql + idSql + likeSql + limitSql
	beego.Info(sq, "sq")
	if id <= 0 && keyword == "" {
		sq = strings.Replace(sq, "where", "", -1)
		sq = strings.TrimSpace(sq)
	}

	stmtLimt, err := dbops.DbConn.Prepare(sq)
	beego.Info("++++++++123", sq)

	var row *sql.Rows
	beego.Info("rowrowrowrowrow")

	var e error
	if id > 0 && keyword != "" {
		beego.Info("11111111")

		row, e = stmtLimt.Query(id, "%"+keyword+"%", (current-1)*pageSize, pageSize)
	} else if id > 0 && keyword == "" {
		beego.Info("22222")

		row, e = stmtLimt.Query(id, (current-1)*pageSize, pageSize)
	} else if id <= 0 && keyword != "" {
		beego.Info("3333")

		row, e = stmtLimt.Query("%"+keyword+"%", (current-1)*pageSize, pageSize)
	} else {
		beego.Info("4444")

		row, e = stmtLimt.Query((current-1)*pageSize, pageSize)
	}
	beego.Info("5555")
	if e != nil {
		beego.Info("666666666")
		return 0, 0, res, err
	}
	beego.Info("456")
	for row.Next() {
		ar := new(model.PostList)
		/*
			Id      int
			Title   string
			Tags    string
			IsTop   int8
			Created time.Time
			Updated time.Time
		*/
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
			beego.Info("10 12 ")
			return 0, 0, res, err
		}
		cat := &model.PostListRes{
			Id:         ar.Id,
			Title:      ar.Title,
			Tags:       ar.Tags,
			IsTop:      ar.IsTop,
			Created:    ar.Created.Unix(),
			Updated:    ar.Updated.Unix(),
			Views:      ar.Views,
			CategoryId: ar.CategoryId,
		}
		res = append(res, cat)
	}
	stmtLimt.Close()
	beego.Info("789")
	count := int(math.Ceil(float64(total) / float64(pageSize)))

	return total, count, res, err
}
