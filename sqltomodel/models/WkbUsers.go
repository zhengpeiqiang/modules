package models

import "time"

type WkbUsers struct {
	Id string `json:"id"` // 用户编号
	UserIdDiy string `json:"user_id_diy"` // 用户自定义id
	SysType int `json:"sys_type"` // 默认0为本站添加,其它为外部添加
	Name string `json:"name"` // 用户名
	Mobile string `json:"mobile"` // 新加手机号
	Email string `json:"email"` // 用户邮箱
	EmailStatus int `json:"email_status"` // 用户邮箱认证状态0：未认证1：待验证2：已经认证
	Password string `json:"password"` // 用户密码
	AlternatePassword string `json:"alternate_password"` // 支付密码
	Salt string `json:"salt"` // 随机码
	Status int `json:"status"` // 账户状态 0：未激活 1：已激活 2：禁用
	UsersType int `json:"users_type"` // 用户类型:1雇主 2威客
	AuthType string `json:"auth_type"` // 认证类型：1个人 2公司
	Profession string `json:"profession"` // 专业
	GoodAt string `json:"good_at"` // 擅长
	Abstract string `json:"abstract"` // 个人简介
	OverdueDate time.Time `json:"overdue_date"` // 找回密码邮件过期时间
	ValidationCode string `json:"validation_code"` // 找回密码随机码
	ExpireDate time.Time `json:"expire_date"` // 重置密码邮件过期时间
	ResetPasswordCode string `json:"reset_password_code"` // 重置密码验证随机码
	RememberToken string `json:"remember_token"` 
	LastLoginTime time.Time `json:"last_login_time"` // 最后一次登录时间
	Source int `json:"source"` // 注册来源 1-来自pc 2-来自手机
	IsEmployee int `json:"is_employee"` // 0 普通用户（自由注册）， 1 员工， 2 老板 3 普通用户（通过分享注册）通过分享注册的用户状态为1
	ShareNum string `json:"share_num"` // //用户分享码
	ShareUid int `json:"share_uid"` // 邀请当前用户id
	IsShareMoney int `json:"is_share_money"` // 是否已经给分享人打款
	ShareMoney float32 `json:"share_money"` // 分享获得的金额
	WechatOpenid string `json:"wechat_openid"` // 第三方登陆 微信授权openid
	UnionId string `json:"union_id"` // 开放平台与公众平台 UnionID
	IsSubscribe int `json:"is_subscribe"` // 是否关注 1是 0否
	IsMovie int `json:"is_movie"` // 是否在直播1是,0否
	SugarGift int `json:"sugar_gift"` // 是否发放礼品 1是 0否
	CreatedAt time.Time `json:"created_at"` // 创建时间
	UpdatedAt time.Time `json:"updated_at"` // 更新时间
	ImUserType int `json:"im_user_type"` // 用户类型:1=普通用户,2=客服
}