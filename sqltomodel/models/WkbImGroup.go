package models

import "time"

// 群聊
type WkbImGroup struct {
	Id string `json:"id"` // id
	GroupId string `json:"group_id"` // 用户雪花ID，采用雪花算法
	GroupType string `json:"group_type"` // 1 群聊；2 私聊
	Name string `json:"name"` // 群名称
	State string `json:"state"` // 0 初始状态，1 进行中，2 完成，3 未解决问题，4 删除
	Notice string `json:"notice"` // 群公告
	UserId string `json:"user_id"` // 如果是私聊，则记录群主的用户id
	OtherUserId string `json:"other_user_id"` // 如果是私聊，则记录对方的用户id
	LastChatId string `json:"last_chat_id"` // 关联本群聊最后一句聊天的id
	LastFromUserId string `json:"last_from_user_id"` 
	LastObjectName string `json:"last_object_name"` 
	LastContentText string `json:"last_content_text"` 
	LastChatTime time.Time `json:"last_chat_time"` 
	UpdatedAt time.Time `json:"updated_at"` // 更新时间
	CreatedAt time.Time `json:"created_at"` // 创建时间
}