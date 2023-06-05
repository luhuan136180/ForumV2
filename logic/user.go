package logic

import (
	"database/sql"
	"errors"
	"fmt"
	"furumvv2/dao/mysql"
	"furumvv2/models"
	"strconv"
)

var CodeURLNOTEXIT = errors.New("没有购买")

func Login(user *models.User) (data *models.ResponseLogin, err error) {
	//对数据进行检验
	//1.合法性(user_address,)
	//2.该用户是否已经注册,从数据库中找
	exit, err := mysql.CheckUserExist(user.UserAddress)
	if err != nil {
		return nil, err
	}
	//3.没注册，就注册。注册了就登录流程
	//没注册
	if !exit {
		//数据库中没有注册，进入注册逻辑

		if data, err = SignUp(user); err != nil {
			return nil, err
		}

		//注册成功，返回token
		return data, nil
	}

	//数据库中已经注册，登录
	if data, err = mysql.Login(user); err != nil {
		return nil, err
	}
	if data.HeadPicture == "" {
		data.HeadPicture = "https://img1.baidu.com/it/u=1888856496,845797841&fm=253&fmt=auto&app=138&f=PNG?w=500&h=500"
	}
	return data, nil
}

func SignUp(user *models.User) (data *models.ResponseLogin, err error) {
	//附加默认值
	user.UserName = "默认用户" + user.UserAddress[:6]
	user.Balance = 0
	user.Picture = "https://img1.baidu.com/it/u=1888856496,845797841&fm=253&fmt=auto&app=138&f=PNG?w=500&h=500"
	data, err = mysql.InsertUser(user)
	data.HeadPicture = "https://img1.baidu.com/it/u=1888856496,845797841&fm=253&fmt=auto&app=138&f=PNG?w=500&h=500"
	return
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
	userInfoInside, err := mysql.GetUserInformationInside(user_address)
	if err != nil {
		return nil, err
	}
	//
	data = new(models.UserInformation)

	if userInfoInside.UserAddress.Valid {
		data.UserAddress = userInfoInside.UserAddress.String
	}
	if userInfoInside.UserName.Valid {
		data.UserName = userInfoInside.UserName.String
	}
	if userInfoInside.Age.Valid {
		data.Age = int(userInfoInside.Age.Int64)
	}
	if userInfoInside.Gender.Valid {
		data.Gender = userInfoInside.Gender.String
	}
	if userInfoInside.HeadPicture.Valid {
		data.HeadPicture = userInfoInside.HeadPicture.String
	}
	if userInfoInside.Signature.Valid {
		data.Signature = userInfoInside.Signature.String
	}
	if userInfoInside.Email.Valid {
		data.Email = userInfoInside.Email.String
	}
	if userInfoInside.Level.Valid {
		data.Level = int(userInfoInside.Level.Int64)
	}
	data.CreateTime = userInfoInside.CreateTime

	//
	if userInfoInside.BCGUrl.Valid {
		data.BackGroundPic = userInfoInside.BCGUrl.String
	}

	//

	return data, nil
}

func GetAllSkinByUser(user_address string) (data []*models.SkinListByUser, err error) {
	return mysql.GetAllSkinByUser(user_address)
}

func ChangeUserInformation(userprofile *models.UserProfile) (err error) {
	//从数据库中获取已存储的user——address的用户信息
	userInformation, err := mysql.GetUserInformationInside(userprofile.UserAddress)
	fmt.Println(*userInformation)
	updateProfile := new(models.UpdateProfile)

	//判断是否有修改
	if userInformation.UserName.String != userprofile.UserName {
		updateProfile.Username = sql.NullString{String: userprofile.UserName, Valid: true}
	} else {
		updateProfile.Username = sql.NullString{String: userprofile.UserName, Valid: true}
	}
	if userInformation.Email.String != userprofile.Email {
		updateProfile.Email = sql.NullString{String: userprofile.Email, Valid: true}
	} else {
		updateProfile.Email = sql.NullString{String: userprofile.Email, Valid: true}
	}
	ageprofile, err := strconv.Atoi(userprofile.Age)
	ageInfo := int(userInformation.Age.Int64)
	if ageInfo != ageprofile {
		updateProfile.Age = sql.NullInt64{Int64: int64(ageprofile), Valid: true}
	} else {
		updateProfile.Age = sql.NullInt64{Int64: int64(ageprofile), Valid: true}
	}
	if userInformation.Signature.String != userprofile.Signature {
		updateProfile.Signature = sql.NullString{String: userprofile.Signature, Valid: true}
	} else {
		updateProfile.Signature = sql.NullString{String: userprofile.Signature, Valid: true}
	}
	updateProfile.Gender = sql.NullString{String: userprofile.Gender, Valid: true}
	updateProfile.HeadPicture = sql.NullString{String: userprofile.HeadPicture, Valid: true}
	updateProfile.UserAddress = userprofile.UserAddress

	updateProfile.BackGroundPicture = sql.NullString{String: userprofile.BackGroundPciture, Valid: true}
	//开始修改
	fmt.Println("updateProfile:", updateProfile)

	if err = mysql.ChangeUserInformation(updateProfile); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func GetPostFromUserAdd(user_address string) (data []*models.GetPostByList, err error) {
	data, err = mysql.GetPostFromUserAdd(user_address)
	//fmt.Println(data)
	if err != nil {
		return nil, err
	}
	return
}

func changeStruct(userInfoInside *models.UserInformationInside) (data *models.UserInformation) {
	data = new(models.UserInformation)

	if userInfoInside.UserAddress.Valid {
		data.UserAddress = userInfoInside.UserAddress.String
	}
	if userInfoInside.UserName.Valid {
		data.UserName = userInfoInside.UserName.String
	}
	if userInfoInside.Age.Valid {
		data.Age = int(userInfoInside.Age.Int64)
	}
	if userInfoInside.Gender.Valid {
		data.Gender = userInfoInside.Gender.String
	}
	if userInfoInside.HeadPicture.Valid {
		data.HeadPicture = userInfoInside.HeadPicture.String
	}
	if userInfoInside.Signature.Valid {
		data.Signature = userInfoInside.Signature.String
	}
	if userInfoInside.Email.Valid {
		data.Email = userInfoInside.Email.String
	}
	if userInfoInside.Level.Valid {
		data.Level = int(userInfoInside.Level.Int64)
	}
	data.CreateTime = userInfoInside.CreateTime

	//
	if userInfoInside.BCGUrl.Valid {
		data.BackGroundPic = userInfoInside.BCGUrl.String
	}
	return
}

func ChangeUserBackGround(userprofile *models.ChangeBCGByUser) (responsedata *models.UserInformation, err error) {
	flag, err := mysql.EXitISBuyingBCG(userprofile)
	if err != nil {
		return nil, err
	}
	if flag == false {

		return nil, CodeURLNOTEXIT
	}
	//数据库中能查到数据，已经购买---可以替换
	data, err := mysql.ChangeUserBackGround(userprofile)
	if err != nil {
		return nil, err
	}

	responsedata = changeStruct(data)
	return
}

func ChangeUserPH(userprofile *models.ChangeHPByUser) (responsedata *models.UserInformation, err error) {
	flag, err := mysql.EXitISBuyingPH(userprofile)
	if err != nil {
		return nil, err
	}
	if flag == false {

		return nil, CodeURLNOTEXIT
	}

	//
	data, err := mysql.ChangeUserPH(userprofile)
	if err != nil {
		return nil, err
	}

	responsedata = changeStruct(data)
	return
}
