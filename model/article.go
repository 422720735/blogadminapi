/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-09-23 11:20:52
 * @LastEditTime: 2019-09-24 09:56:25
 * @LastEditors: Please set LastEditors
 */
package model

import "time"

type PostList struct {
	Id         int
	Title      string
	Tags       string
	IsTop      int8
	Views      int
	CategoryId int
	Created    time.Time
	Updated    time.Time
}

type PostListRes struct {
	Id         int    `json:"id"`
	Title      string `json:"title"`
	Tags       string `json:"tags"`
	IsTop      bool   `json:"isTop"`
	Created    int64  `json:"created"`
	Updated    int64  `json:"updated"`
	Views      int    `json:"views"`
	CategoryId int    `json:"categoryId"`
}

type PostInfo struct {
	Id         int    `json:"id"`
	UserId     int    `json:"userId"`
	Title      string `json:"title"`
	Url        string `json:"url"`
	Content    string `json:"content"`
	Tags       string `json:"tags"`
	Views      int    `json:"views`
	IsTop      int8   `json:"isTop"`
	Created    int64  `json:"created"`
	Updated    int64  `json:"updated"`
	CategoryId int    `json:"categoryId"`
	// Status     int8   `json:"status"` 软删除
	Types int8   `json:"types"`
	Info  string `json:"info"`
	Image string `json:"image"`
}
