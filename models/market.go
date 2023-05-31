package models

import (
	"time"
)

type Skin struct {
	SkinID     int       `json:"skin_id"db:"skin_id"`
	SkinUrl    string    `json:"skin_Url"db:"skin_url"`
	CreateTime time.Time `json:"createTime"db:"create_time"`
	Status     int       `json:"status"db:"status"`
	Price      int       `json:"price"db:"price"`
}

type Shop struct {
	SkinID      int    `json:"skin_id"db:"skin_id"`
	Status      int    `json:"status"db:"status"`
	UserAddress string `json:"user_address"db:"user_address"`
	Price       int    `json:"price"db:"price"`
}
