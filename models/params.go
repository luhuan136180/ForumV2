package models

//登录参数
type ParamLogin struct {
	Key         string `json:"key"binding:"required"`
	UserAddress string `json:"user_address"binding:"required"`
}
