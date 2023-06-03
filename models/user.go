package models

import (
	"database/sql"
	"time"
)

//已经不到这个结构体时用于干啥的了
type User struct {
	UserAddress string `db:"user_address"`
	UserName    string `db:"user_name"`
	Balance     int    `db:"balance"`
	Picture     string `db:"picture_url"`
}

//获取余额的结构体
type GetBalance struct {
	UserName    string `json:"user_name"db:"user_name"`
	UserAddress string `json:"user_address"db:"user_address"`
	Balance     int    `json:"balance"db:"balance""`
}

//登录的结构体
type Login struct {
	UserAddress string `json:"user_address"` //地址
	Key         string `json:"hash"`         //签名
	Time        string `json:"time"`         //时间戳
}

//用户信息显示结构体
type UserInformation struct {
	UserAddress string    `json:"user_address"db:"user_address"`
	UserName    string    `json:"user_name"db:"user_name"`
	CreateTime  time.Time `json:"create_time"db:"create_time"`
	Age         int       `json:"age"db:"age"`
	Email       string    `json:"eamil"db:"email"`
	Gender      string    `json:"gender"db:"gender"` //0:男 ；1：女
	Signature   string    `json:"signature"db:"signature"`
	HeadPicture string    `json:"head_picture"db:"picture_url"` //头像
	Level       int       `json:"level"db:"level"`
	Experience  int       `json:"experience"db:"experience"`
}

type UserInformationInside struct {
	UserAddress sql.NullString `json:"user_address"db:"user_address"`
	UserName    sql.NullString `json:"user_name"db:"user_name"`
	Balance     sql.NullInt64  `json:"balance"db:"balance"`
	CreateTime  time.Time      `json:"create_time"db:"create_time"`
	Age         sql.NullInt64  `json:"age"db:"age"`
	Email       sql.NullString `json:"eamil"db:"email"`
	Gender      sql.NullString `json:"gender"db:"gender"` //0:男 ；1：女
	Signature   sql.NullString `json:"signature"db:"signature"`
	HeadPicture sql.NullString `json:"head_picture"db:"picture_url"` //头像
	Level       sql.NullInt64  `json:"level"db:"level"`
	Experience  sql.NullInt64  `json:"experience"db:"experience"`
}

//展示用户拥有的皮肤结构体
type SkinListByUser struct {
	UserAddress string `json:"user_address"db:"user_address"`
	SkinUrl     string `json:"skin_Url"db:"skin_url"`
	Status      int    `json:"status"db:"status"`
	SkinID      int    `json:"skin_id"db:"skin_id"`
}

//显示用户发表的帖子集合
type PostFromUser struct {
	UserAddress string `json:"user_address"db:"author_address"`
	UserName    string `json:"user_name"db:"user_name"`
	PostID      int64  `json:"post_id"db:"post_id"`
	Title       string `json:"title"db:"title"`
	Content     string `json:"content"db:"content"`
}

//用户信息修改的josn收集结构体
type UserProfile struct {
	UserAddress string `json:"user_address"db:"user_address"binding:"required"`
	//昵称修改
	UserName    string `json:"user_name"db:"user_name"binding:"required"`
	Email       string `json:"email"db:"email"binding:"required"`
	Age         string `json:"age"db:"age"binding:"required"`
	Gender      string `json:"gender"db:"gender"binding:"required"`
	Signature   string `json:"signature"db:"signature"binding:"required"`
	HeadPicture string `json:"head_picture"db:"picture_url"binding:"required"`
}

// 定义更新数据的结构体类型
type UpdateProfile struct {
	UserAddress string         `db:"user_address"`
	Username    sql.NullString `db:"user_name"`
	Email       sql.NullString `db:"email"`
	Gender      sql.NullString `db:"gender"`
	Signature   sql.NullString `db:"signature"`
	HeadPicture sql.NullString `json:"head_picture"db:"picture_url"`
	Age         sql.NullInt64  `db:"age"`
}
