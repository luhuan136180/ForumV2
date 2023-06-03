package models

import "database/sql"

//关注的列表
type FollowerList struct {
	UserAddress      string `json:"user_address"db:"user_address"`
	FollowerList     []string
	FollowerInfoList []*FollowerInfo `json:"follower_list"`
}

//关注者-被关注者
type Follower struct {
	UserAddress     string `json:"user_address"db:"user_address"`
	FollowerAddress string `json:"follower_address"db:"follower_address"`
}

type FollowerInfo struct {
	FollowerAddress string         `json:"follower_address"db:"user_address"`
	FollowerName    string         `json:"follower_name"db:"user_name"`
	Signature       sql.NullString `json:"signature"db:"signature"`
	HeadPicture     sql.NullString `json:"head_picture"db:"picture_url"`
}
