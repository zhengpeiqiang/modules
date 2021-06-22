package models

import "time"

// 会话报表
type WkbImSessionReport struct {
	Id int `json:"id"` // 自增id
	CustomerUid int `json:"customer_uid"` // 用户id
	CustomerNickname string `json:"customer_nickname"` // 用户昵称
	CustomerServiceName string `json:"customer_service_name"` // 接待客服名称
	CustomerServiceUid int `json:"customer_service_uid"` // 接待客服id
	Channel string `json:"channel"` // 渠道
	VisitStart string `json:"visit_start"` // 访问开始时间
	VisitEnd string `json:"visit_end"` // 访问截至时间
	Seconds int `json:"seconds"` // 秒数
	UpdatedAt time.Time `json:"updated_at"` // 更新时间
	CreatedAt time.Time `json:"created_at"` // 创建时间
}