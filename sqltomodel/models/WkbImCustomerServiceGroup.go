package models

import "time"

// 客服所在组关系表
type WkbImCustomerServiceGroup struct {
	Id int `json:"id"` // 自增id
	Uid int `json:"uid"` // wkb_realname_auth的uid
	CustomerServiceId int `json:"customer_service_id"` // 对于客服表主键
	RealName string `json:"real_name"` // 客服姓名
	GroupId int `json:"group_id"` // 所在技能组id
	Weight int `json:"weight"` // 优先级，数字越大越优先
	ReceptionNum string `json:"reception_num"` // 接待数量，默认为0
	Deleted int `json:"deleted"` // 是否删除:1=是，2=否
	UpdatedAt time.Time `json:"updated_at"` // 更新时间
	CreatedAt time.Time `json:"created_at"` // 创建时间
}