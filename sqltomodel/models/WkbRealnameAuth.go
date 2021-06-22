package models

import "time"

type WkbRealnameAuth struct {
	Id string `json:"id"` // 自增id
	Uid int `json:"uid"` // 用户id
	Admin int `json:"admin"` // 管理员id
	InvitationUserId int `json:"invitation_user_id"` // 邀请人用户id
	Username string `json:"username"` // 用户名
	Realname string `json:"realname"` // 用户真实姓名
	Email string `json:"email"` // 邮箱
	Tel string `json:"tel"` // 电话
	CardNumber string `json:"card_number"` // 用户证件号
	CardFrontSide string `json:"card_front_side"` // 身份证正面
	CardBackDside string `json:"card_back_dside"` // 身份证背面
	ValidationImg string `json:"validation_img"` // 持证验证图片
	BusinessCard string `json:"business_card"` // 名片
	District string `json:"district"` // 地理区域
	Status int `json:"status"` // 认证状态 0：待验证 1：成功 2：失败
	CardType int `json:"card_type"` // 证件类型  1-身份证  2-护照
	Type int `json:"type"` // 认证类型  1-身份认证  2-企业认证
	FaeCategoryId int `json:"fae_category_id"` // fae申请版主应用分类
	Usertype int `json:"usertype"` // 用户类型
	Reason string `json:"reason"` // 审核失败理由
	Resume string `json:"resume"` // 简历
	Province int `json:"province"` // 省
	City int `json:"city"` // 市
	CompanyAddress string `json:"company_address"` // 所在公司
	WorkingYear string `json:"working_year"` // 工作年限
	AuthTime time.Time `json:"auth_time"` // 认证通过时间
	FailTime time.Time `json:"fail_time"` // 审核失败时间
	IntegralOne int `json:"integral_one"` // FAE认证积分
	IntegralTwo int `json:"integral_two"` // 邀请人积分
	RecordText string `json:"record_text"` // 记录文本
	NoviceVideo int `json:"novice_video"` // 新手视频
	Class int `json:"class"` // 分类
	CreatedAt time.Time `json:"created_at"` // 创建时间
	UpdatedAt time.Time `json:"updated_at"` // 更新时间
}