package models



type WkbImAttention struct {
	Id string `json:"id"` // 联系人关系编号
	Uid int `json:"uid"` // 用户编号
	FriendUid int `json:"friend_uid"` // 好友uid编号
}