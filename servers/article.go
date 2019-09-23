/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-09-23 09:31:19
 * @LastEditTime: 2019-09-23 18:24:50
 * @LastEditors: Please set LastEditors
 */
package servers

import (
	"blogadminapi/dbops"
	"blogadminapi/model"
	"database/sql"
	"math"
	"unsafe"
)

// import "blogadminapi/dbops"

/** 查询分页数据*/
func GetArticleLimitList(id, pageSize, current int) (int, int, []*model.PostListRes, error) {
	// id == 0 就表示查询全部
	sq := "select count(`id`) from tb_post"
	if id > 0 {
		sq = "select count(`id`) from tb_post where category_id = ?"
	}
	stmtOutCount, err := dbops.DbConn.Prepare(sq)
	var res []*model.PostListRes
	var count int64
	if err != nil {
		return 0, 0, nil, err
	}

	if id > 0 {
		stmtOutCount.QueryRow(id).Scan(&count)
	} else {
		stmtOutCount.QueryRow().Scan(&count)
	}
	stmtOutCount.Close()

	// stmtLimt, err := dbops.DbConn.Query("select id, title, tags, is_top, created, updated from tb_post order by id desc limit ?, ?", (current-1)*pageSize, pageSize)

	if id > 0 {
		sq = "select id, title, tags, is_top, created, updated from tb_post where category_id = ? order by id desc limit ?, ?"
	} else {
		sq = "select id, title, tags, is_top, created, updated from tb_post order by id desc limit ?, ?"
	}

	stmtLimt, err := dbops.DbConn.Prepare(sq)

	var row *sql.Rows
	var e error
	if id > 0 {
		row, e = stmtLimt.Query(id, (current-1)*pageSize, pageSize)
	} else {
		row, e = stmtLimt.Query((current-1)*pageSize, pageSize)
	}
	if e != nil {
		return 0, 0, res, err
	}

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
		); err != nil {
			return 0, 0, res, err
		}
		cat := &model.PostListRes{
			Id:      ar.Id,
			Title:   ar.Title,
			Tags:    ar.Tags,
			IsTop:   ar.IsTop,
			Created: ar.Created.Unix(),
			Updated: ar.Updated.Unix(),
		}
		res = append(res, cat)
	}
	stmtLimt.Close()
	pageSize = int(math.Ceil(float64(count) / float64(pageSize)))
	total := *(*int)(unsafe.Pointer(&count))
	return total, pageSize, res, err
}
