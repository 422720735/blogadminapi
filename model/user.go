/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-10-08 17:05:48
 * @LastEditTime: 2019-10-08 17:05:48
 * @LastEditors: your name
 */
package model

import "time"

type User struct {
	Id         int
	Username   string
	Password   string
	Email      string
	LoginCount int
	LastTime   time.Time
	LastIp     string
	State      int8
	Created    time.Time
	Updated    time.Time
}
