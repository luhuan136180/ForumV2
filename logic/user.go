package logic

import (
	"furumvv2/dao/mysql"
	"furumvv2/models"
	"furumvv2/pkg/jwt"
)

func Login(user *models.User) (token string, err error) {
	//对数据进行检验
	//1.合法性(user_address,)
	//2.该用户是否已经注册,从数据库中找
	exit, err := mysql.CheckUserExist(user.UserAddress)
	if err != nil {
		return "", err
	}
	//3.没注册，就注册。注册了就登录流程
	//没注册
	if !exit {
		//数据库中没有注册，进入注册逻辑
		if err = SignUp(user); err != nil {
			return "", err
		}
		//注册成功，返回token
		return jwt.GenToken(user.UserAddress, user.UserName)
	}
	//数据库中已经注册，登录
	if err := mysql.Login(user); err != nil {
		return "", nil
	}
	return jwt.GenToken(user.UserAddress, user.UserName)
}

func SignUp(user *models.User) (err error) {
	//附加默认值
	user.UserName = "默认用户" + user.UserAddress[:6]
	user.Balance = 0
	return mysql.InsertUser(user)
}

func GetUserBalance(user_address string) (data *models.GetBalance, err error) {
	return mysql.GetUserBalance(user_address)
}

func AddUserBalance(user_address string, amount int) (data *models.GetBalance, err error) {
	return mysql.AddUserBalance(user_address, amount)
}

func SubUserBalance(user_address string, amount int) (data *models.GetBalance, err error) {
	return mysql.SubUserBalance(user_address, amount)
}

func GetUserInformation(user_address string) (data *models.UserInformation, err error) {
	return mysql.GetUserInformation(user_address)
}

func GetAllSkinByUser(user_address string) (data []*models.SkinListByUser, err error) {
	return mysql.GetAllSkinByUser(user_address)
}
