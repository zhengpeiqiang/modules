package models

import "time"

// 群组成员
type WkbImGroupMember struct {
	Id string `json:"id"` // id
	GroupId string `json:"group_id"` // 用户雪花ID，采用雪花算法
	UserId string `json:"user_id"` // 用户id
	MemberName string `json:"member_name"` // 群成员昵称
	MemberType string `json:"member_type"` // 成员类型：0 用户；1 机器人；2 超级管理员；3 管理员；4 客服；5 FAE；
	Avatar string `json:"avatar"` // 头像
	UpdatedAt time.Time `json:"updated_at"` // 更新时间
	CreatedAt time.Time `json:"created_at"` // 创建时间
}