package mysql

import "errors"

var (
	ErrorUserExist       = errors.New("用户已存在")
	ErrorUserNotExist    = errors.New("用户不存在")
	ErrorInvalidPassword = errors.New("用户名或地址错误")
	ErrorInvalidID       = errors.New("无效的ID")
	ErrorNotEnoughAmount = errors.New("没有足够的余额，无法购买，请充值")
	ErrorSkinISExit      = errors.New("皮肤已经购买")
)
