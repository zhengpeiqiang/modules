package models

import "time"

// 聊天内容
type WkbImGroupChatMessage struct {
	Id string `json:"id"` // id
	GroupId string `json:"group_id"` 
	FromUserId string `json:"from_user_id"` // 发送用户 Id
	ToUserId string `json:"to_user_id"` // 目标 Id
	ObjectName string `json:"object_name"` // 消息类型 文本消息 RC:TxtMsg 、 图片消息 RC:ImgMsg 、语音消息 RC:VcMsg 、图文消息 RC:ImgTextMsg 、位置消息 RC:LBSMsg 、添加联系人消息 RC:ContactNtf 、提示条通知消息 RC:InfoNtf 、资料通知消息 RC:ProfileNtf 、通用命令通知消息 RC:CmdNtf
	ContentText string `json:"content_text"` 
	HiddenUserId string `json:"hidden_user_id"` // 不可见的人user_id
	Content string `json:"content"` // 消息类型
	ImageUri string `json:"imageUri"` 
	W int `json:"w"` 
	H int `json:"h"` 
	Name string `json:"name"` 
	Size int `json:"size"` 
	FileUrl string `json:"fileUrl"` 
	Types string `json:"types"` 
	Duration int `json:"duration"` 
	RemoteUrl string `json:"remoteUrl"` 
	ExtraMsg string `json:"extra_msg"` // 额外的信息
	ChannelType string `json:"channel_type"` // 会话类型
	MsgTimestamp string `json:"msg_timestamp"` // 服务端收到客户端发送消息时的服务器时间（1970年到现在的毫秒数）
	MsgUid string `json:"msg_uid"` // 可通过 msgUID 确定消息唯一
	SensitiveType int `json:"sensitive_type"` // 消息中是否含有敏感信息，0 为不包含，1 为含有屏蔽敏感词，2 为含有替换敏感词
	Source string `json:"source"` // 标识消息的发送源头
	CurrentStatus int `json:"current_status"` // 是否当前消息（0 历史消息；1 当前消息；）
	Status string `json:"status"` // 消息状态（1 已发送；2 已完成；3 已取消，非消息发送失败；4 已失效，非消息发送失败；5 删除）
	UpdatedAt time.Time `json:"updated_at"` // 更新时间
	CreatedAt time.Time `json:"created_at"` // 创建时间
}