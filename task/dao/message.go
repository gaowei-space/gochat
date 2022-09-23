/**
 * Created by lock
 * Date: 2019-09-22
 * Time: 22:53
 */
package dao

import (
	"gochat/db"
)

var dbIns = db.GetDb("gochat")

type Message struct {
	Id           int    `gorm:"primary_key" json:"id,omitempty"`
	Seq          string `json:"seq,omitempty"`
	RoomId       int    `json:"room_id,omitempty"`
	ReceiverType string `json:"receiver_type,omitempty"`
	FromUserName string `json:"from_user_name,omitempty"`
	FromUserId   int    `json:"from_user_id,omitempty"`
	ToUserName   string `json:"to_user_name,omitempty"`
	ToUserId     int    `json:"to_user_id,omitempty"`
	MsgType      int    `json:"msg_type,omitempty"`
	Msg          string `json:"msg,omitempty"`
	Status       int    `json:"status,omitempty"`
	SendTime     string `json:"send_time,omitempty"`
	db.DbGoChat
}

func (u *Message) TableName() string {
	return "messages"
}

func (u *Message) Add() (msgId int, err error) {
	if err = dbIns.Table(u.TableName()).Create(&u).Error; err != nil {
		return 0, err
	}
	return u.Id, nil
}
