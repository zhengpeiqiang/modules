package models

import "time"

// 优先接入设置信息表
type WkbImReceptionGroup struct {
	Id int `json:"id"` // 自增id
	Channel string `json:"channel"` // 渠道
	GroupId int `json:"group_id"` // 技能组id
	Weight int `json:"weight"` // 权重
	Deleted int `json:"deleted"` // 是否删除:1=删除,2=否
	UpdatedAt time.Time `json:"updated_at"` // 更新时间
	CreatedAt time.Time `json:"created_at"` // 创建时间
}