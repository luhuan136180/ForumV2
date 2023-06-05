package models

import "database/sql"

//登录参数
type ParamLogin struct {
	Key         string `json:"key"binding:"required"`
	UserAddress string `json:"user_address"binding:"required"`
}

//ResponseLogin
type ResponseLogin struct {
	HeadPicture   string         `json:"head_picture"db:"picture_url"`
	UserName      string         `json:"user_name"db:"user_name"`
	UserAddress   string         `db:"user_address"`
	BackGroundPic sql.NullString `json:"bcg_url"db:"background_url"`
}

type ResponseLoginData struct {
	HeadPicture   string `json:"head_picture"db:"picture_url"`
	UserName      string `json:"user_name"db:"user_name"`
	UserAddress   string `db:"user_address"`
	BackGroundPic string `json:"bcg_url"db:"background_url"`
}
