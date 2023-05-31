package models

import "time"

type User struct {
	UserAddress string `db:"user_address"`
	UserName    string `db:"user_name"`
	Balance     int    `db:"balance"`
}

type GetBalance struct {
	UserName    string `json:"user_name"db:"user_name"`
	UserAddress string `json:"user_address"db:"user_address"`
	Balance     int    `json:"balance"db:"balance""`
}

type Login struct {
	UserAddress string `json:"user_address"` //地址
	Key         string `json:"hash"`         //签名
	Time        string `json:"time"`         //时间戳
}

type UserInformation struct {
	UserAddress string    `json:"user_address"db:"user_address"`
	UserName    string    `json:"user_name"db:"user_name"`
	Balance     int       `json:"balance"db:"balance"`
	CreateTime  time.Time `json:"create_time"db:"create_time"`
}

type SkinListByUser struct {
	UserAddress string `json:"user_address"db:"user_address"`
	SkinUrl     string `json:"skin_Url"db:"skin_url"`
	Status      int    `json:"status"db:"status"`
	SkinID      int    `json:"skin_id"db:"skin_id"`
}
