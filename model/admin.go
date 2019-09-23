/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-09-12 16:35:32
 * @LastEditTime: 2019-09-23 11:13:31
 * @LastEditors: Please set LastEditors
 */
package model

type SystemConfig struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

type SystemConfigFilters struct {
	Title       string `json:"title"`
	Url         string `json:"url"`
	Keywords    string `json:"keywords"`
	Description string `json:"description"`
	Email       string `json:"email"`
	Start       string `json:"start"`
	Qq          string `json:"qq"`
}

type Category struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Created int64  `json:"created"`
	Updated int64  `json:"updated"`
}
