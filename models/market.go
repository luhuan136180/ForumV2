package models

import (
	"time"
)

type Skin struct {
	SkinID     int       `json:"skin_id"db:"skin_id"`
	SkinUrl    string    `json:"skin_Url"db:"skin_url"`
	CreateTime time.Time `json:"createTime"db:"create_time"`
	Status     int       `json:"status"db:"status"`
	Price      string    `json:"price"db:"price"`
	SkinAdd    string    `json:"skin_address"db:"skin_address"`
}

type Shop struct {
	SkinID      int    `json:"skin_id"db:"skin_id"`
	Status      int    `json:"status"db:"status"`
	UserAddress string `json:"user_address"db:"user_address"`
	Price       int    `json:"price"db:"price"`
}
type ShoppingInfo struct {
	SkinID      int    `json:"skin_id"db:"skin_id"`
	UserAddress string `json:"user_address"db:"user_address"`
	Status      int    `json:"status"db:"status"`
	SkinURL     string `json:"skin_Url"db:"skin_url"`
	Price       string `json:"price"`
	SkinAdd     string `json:"skin_address"db:"skin_address"`
}

//type AddSkin struct {
//	SkinID     int       `json:"skin_id"db:"skin_id"`
//	SkinUrl    string    `json:"skin_Url"db:"skin_url"`
//	Price      int       `json:"price"db:"price"`
//	SkinAdd    string    `json:"skin_address"db:"skin_address"`
//	Status     int       `json:"status"db:"status"`
//}
