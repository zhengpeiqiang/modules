package models

import "time"

// 技能组信息表
type WkbImSkillGroup struct {
	Id int `json:"id"` // 自增主键
	GroupCode string `json:"group_code"` // 技能组编码
	GroupName string `json:"group_name"` // 技能组名称
	GroupStrategy string `json:"group_strategy"` // 策略:1=随机,2=按优先级,3=轮流,4=进行会话最少客服优先
	IsShow string `json:"is_show"` // 是否显示（0 否；1 是；）
	Deleted string `json:"deleted"` // 是否删除:1=删除,2=否
	UpdatedAt time.Time `json:"updated_at"` // 更新时间
	CreatedAt time.Time `json:"created_at"` // 创建时间
}