package models

import "time"

type WkbImMergeMessage struct {
	Id string `json:"id"` 
	MergeId string `json:"merge_id"` // 合并id（雪花id）
	Content string `json:"content"` // 合并内容，格式与聊天相同
	UpdatedAt time.Time `json:"updated_at"` 
	CreatedAt time.Time `json:"created_at"` 
}