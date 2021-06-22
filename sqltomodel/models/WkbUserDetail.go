package models

import "time"

type WkbUserDetail struct {
	Id string `json:"id"` // 用户详情编号
	Uid int `json:"uid"` // 用户编号
	Realname string `json:"realname"` // 真实姓名
	RealnameStatus int `json:"realname_status"` // 表示真实姓名是否公开 0表示不公开 1表示公开
	Sex int `json:"sex"` // 0表示女 1表示男
	Mobile string `json:"mobile"` // 手机号码
	Email string `json:"email"` // 邮箱
	Co string `json:"co"` // 公司
	Address string `json:"address"` // 收货地址
	WorkYears int `json:"work_years"` // 工作年限
	MobileStatus int `json:"mobile_status"` // 0表示不公开 1表示公开
	Nickname string `json:"nickname"` // app端用户昵称
	WxNickname string `json:"wx_nickname"` // 微信昵称第一时记录的
	Qq string `json:"qq"` // 用户qq
	QqStatus int `json:"qq_status"` // qq是否公开 0表示不公开 1表示公开
	Wechat string `json:"wechat"` // 用户微信号
	WechatStatus int `json:"wechat_status"` // qq是否公开 0表示不公开 1表示公开
	CardNumber string `json:"card_number"` // 身份证号码
	Province int `json:"province"` // 用户省份
	City int `json:"city"` // 用户城市
	Area int `json:"area"` // 用户地区
	Avatar string `json:"avatar"` // 用户头像
	Autograph string `json:"autograph"` // 个人签名-问答头衔
	Introduce string `json:"introduce"` // 个人简介-问答简介
	AnswerCount int `json:"answer_count"` // 用户回答总数
	FieldId string `json:"field_id"` // 擅长领域之前是id
	Sign string `json:"sign"` // 个人标签
	ToMeAskMoney float32 `json:"to_me_ask_money"` // 向我提提问需要支付金额
	IsToMeAsk int `json:"is_to_me_ask"` // 是否同意向我追问1是，2否
	IsMinuteListen int `json:"is_minute_listen"` // 是否规定分钟内可以免费听回答1是，2否
	IsOpenListen int `json:"is_open_listen"` // 开发时,回答超过30天可免费收听1是，2否
	IsModerator int `json:"is_moderator"` // 是否版主1是
	EmployeePraiseRate int `json:"employee_praise_rate"` // 做为服务商的好评数量
	EmployerPraiseRate int `json:"employer_praise_rate"` // 做为雇主的好评数量
	ReceiveTaskNum int `json:"receive_task_num"` // 承接任务数量
	PublishTaskNum int `json:"publish_task_num"` // 发布任务数量
	ShopStatus int `json:"shop_status"` // 店铺状态: -1.管理员禁用店铺 0.未开启店铺 1.开启店铺 2.关闭店铺
	Balance float32 `json:"balance"` // 用户资产余额
	BalanceStatus int `json:"balance_status"` // 资产冻结 0表示未冻结 1表示资金被冻结
	LastLoginTime time.Time `json:"last_login_time"` // 最后一次登录时间
	Backgroundurl string `json:"backgroundurl"` // 空间个人资料背景图片
	AlternateTips int `json:"alternate_tips"` // 支付提示 0：提示 1：不提示
	EmployeeNum int `json:"employee_num"` // 累计雇佣量
	Integral int `json:"integral"` // 积分
	TotalIntegral int `json:"total_integral"` // 总积分数
	ReviveValue int `json:"revive_value"` // 用户生命值
	AllMoney float32 `json:"all_money"` // 用户总收入
	AskMeMoney float32 `json:"ask_me_money"` // 对我提问需要支付金额
	CreatedAt time.Time `json:"created_at"` // 创建时间
	UpdatedAt time.Time `json:"updated_at"` // 更新时间
	RyToken string `json:"ry_token"` // 融云的token
	Author string `json:"author"` // 作者
	Signature string `json:"signature"` // 个人签名
	Remark string `json:"remark"` // 备注
	ImUserType int `json:"im_user_type"` // 用户类型:1=普通用户,2=客服,3=客服（管理员）
}