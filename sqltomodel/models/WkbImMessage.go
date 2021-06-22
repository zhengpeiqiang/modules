package models

import "time"

type WkbImMessage struct {
	Id string `json:"id"` // IM消息编号
	FromUid int `json:"from_uid"` // 消息发送人uid
	ToUid int `json:"to_uid"` // 消息接收人uid
	Content string `json:"content"` // 消息内容
	Status int `json:"status"` // 消息状态    1-未读    2-已读
	CreatedAt time.Time `json:"created_at"` // IM消息创建时间
}