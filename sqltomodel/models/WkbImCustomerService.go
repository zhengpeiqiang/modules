package models

import "time"

// 客服表
type WkbImCustomerService struct {
	Id int `json:"id"` // 自增id
	Uid string `json:"uid"` // wkb_realname_auth的uid
	RealName string `json:"real_name"` // 真实姓名(冗余)
	Nickname string `json:"nickname"` // 昵称
	IsOnline string `json:"is_online"` // 是否在线（0 不在线；1 在线；）
	IsManage string `json:"is_manage"` // 是否管理员:1=是,2=否
	SeeCustomerService string `json:"see_customer_service"` // 查看其他客服权限:1=有,2=否
	SeeCustomer string `json:"see_customer"` // 是否查看用户权限:1 可查看；2 不可见；
	ReceptionNum string `json:"reception_num"` // 接待数量，默认为0
	Deleted string `json:"deleted"` // 是否删除:1=删除,2=否
	UpdatedAt time.Time `json:"updated_at"` // 更新时间
	CreatedAt time.Time `json:"created_at"` // 创建时间
}