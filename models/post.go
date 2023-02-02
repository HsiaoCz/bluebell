package models

import "time"

// 内存对其  相同类型的尽量放在一起
type Post struct {
	Status      int32     `json:"status" db:"status"`
	ID          int64     `json:"id" db:"post_id"`
	AuthorID    int64     `json:"author_id" db:"author_id"`
	CommunityID int64     `json:"community_id" db:"community_id" binding:"required"`
	Title       string    `json:"title" db:"title" binding:"required"`
	Content     string    `json:"content" db:"content" binding:"required"`
	CreateTime  time.Time `json:"create_time" db:"create_time"`
}

type ApiPostDetail struct {
	AuthorName      string             `json:"author_name" `
	VoteNumber      int64              `json:"vote_number"`
	Post            `json:"post"`      //嵌入帖子的结构体
	CommunityDetail `json:"community"` //嵌入社区信息
}
