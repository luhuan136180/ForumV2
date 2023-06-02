package models

import (
	"database/sql"
	"time"
)

type Post struct {
	Title         string `json:"title"db:"title"`
	Content       string `json:"content"db:"content"binding:"required"`
	AuthorAddress string `json:"author_address"db:"author_address"binding:"required"`
	//postkey是我后端根据  作者，内容，时间错，生成的
	PostKey    string    `json:"post_key"db:"post_key"`
	CreateTime time.Time `json:"create_time" db:"create_time"`
	//topic 先默认为1
	TopicID int64 `json:"topic_id"db:"topic_id"`
	//post_id 用于标记一个帖子的所有消息体
	PostID int64 `json:"post_id"db:"post_id"`
}

//创建帖子的结构体
type CtreatePost struct {
	Title         string `json:"title"db:"title"`
	Content       string `json:"content"db:"content"binding:"required"`
	AuthorAddress string `json:"author_address"db:"author_address"binding:"required"`
	PictureURL    string `json:"picture_url" db:"url"`
	PostKey       string `json:"post_key"db:"post_key"`
	PostID        int64  `json:"post_id"db:"post_id"`
	//topic 先默认为1
	TopicID int64 `json:"topic_id"db:"topic_id"`
	Status  int   `json:"status"db:"status"`
}

type CreateResponse struct {
	Title   string `json:"title"`
	PostKey string `json:"post_key"`
}

type PostPicture struct {
	PostID int64  `json:"post_id"db:"post_id"`
	URl    string `json:"url"db:"url"`
}

//获取主贴信息的结构体
type GetPost struct {
	Title         sql.NullString `json:"title"db:"title"`
	Content       string         `json:"content"db:"content"binding:"required"`
	AuthorName    string         `json:"author_name"db:"user_name"binding:"required"`
	PostID        int64          `json:"post_id"db:"post_id"`
	PictureURL    string         `json:"picture_url"db:"url"`
	AuthorAddress string         `json:"author_address"db:"user_address"`
}

//获取帖子的每一条内容的结构体
type GetMessage struct {
	Title         string `json:"title"db:"title"`
	Content       string `json:"content"db:"content"binding:"required"`
	AuthorAddress string `json:"author_address"db:"author_address"binding:"required"`
	TopicID       int64  `json:"topic_id"db:"topic_id"`
}

//获取
type UserAddressList []struct {
	PostKey string `json:"post_key"`
}
