package models

import "time"

// 客服成就列表
type WkbImAchievement struct {
	Id string `json:"id"` 
	Uid string `json:"uid"` // 客服id
	CustomerUid string `json:"customer_uid"` // 客户id
	SkillGroupId string `json:"skill_group_id"` // 技能组id
	Integral string `json:"integral"` // 可以获得积分分值
	UseTime string `json:"use_time"` // 用时，秒
	Status string `json:"status"` // 接单状态（0 无状态；1 处理中；2 完成；3 失败；）
	UpdatedAt time.Time `json:"updated_at"` 
	CreatedAt time.Time `json:"created_at"` 
}