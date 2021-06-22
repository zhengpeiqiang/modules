package models

import "time"

// 系统回复设置表
type WkbImReplySetting struct {
	Id int `json:"id"` // 自增主键
	CustomerServiceFlag int `json:"customer_service_flag"` // 客服无应答状态:1=开启,2=关闭
	CustomerServiceTimeout int `json:"customer_service_timeout"` // 客服无应答超时(分钟)
	CustomerServiceText string `json:"customer_service_text"` // 客服无应答自动回复
	CustomerFlag int `json:"customer_flag"` // 客户无应答状态:1=开启,2=关闭
	CustomerTimeout int `json:"customer_timeout"` // 客户无应答超时(分钟)
	CustomerText string `json:"customer_text"` // 客户无应答自动回复
	Deleted int `json:"deleted"` // 是否删除:1=删除,2=否
	UpdatedAt time.Time `json:"updated_at"` // 更新时间
	CreatedAt time.Time `json:"created_at"` // 创建时间
	NotWorkingFlag int `json:"not_working_flag"` // 非工作时间开启状态:1=开启,2=关闭
	NotWorkingText string `json:"not_working_text"` // 非工作时间回复内容
}