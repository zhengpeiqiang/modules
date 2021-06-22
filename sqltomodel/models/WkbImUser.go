package models

import "time"

type WkbImUser struct {
	Id int `json:"id"` 
	MiniprogramAppid string `json:"miniprogram_appid"` 
	UserId int `json:"user_id"` 
	UnionId string `json:"union_id"` 
	Openid string `json:"openid"` 
	UpdatedAt time.Time `json:"updated_at"` 
	CreatedAt time.Time `json:"created_at"` 
}