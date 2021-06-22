package models

import "time"

// 关键词回复设置
type WkbImKeyword struct {
	Id int `json:"id"` // 自增主键
	Keyword string `json:"keyword"` // 关键词,多个以;分割
	GroupId int `json:"group_id"` // data_flag=1,跳转技能组id
	Condition int `json:"condition"` // 满足条件:1=包含
	DataFlag int `json:"data_flag"` // 数据标识:1=跳转,2=回复
	AnswerText string `json:"answer_text"` // data_flag=2,回复内容
	RobotCode string `json:"robot_code"` // 机器人编号
	Deleted int `json:"deleted"` // 是否删除:1=删除,2=否
	UpdatedAt time.Time `json:"updated_at"` // 更新时间
	CreatedAt time.Time `json:"created_at"` // 创建时间
}