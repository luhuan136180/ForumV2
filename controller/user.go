package controller

import (
	"errors"
	"fmt"
	"furumvv2/dao/mysql"
	"furumvv2/logic"
	"furumvv2/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"strconv"
)

//登录
//获取参数为一个签名和一个公钥地址
func LoginHandler(c *gin.Context) {

	//前端登录传递给后端的数据，获取参数
	userLogin := new(models.Login)
	//
	//

	if err := c.ShouldBindJSON(userLogin); err != nil {
		zap.L().Error("Login with invalid param", zap.Error(err))

		tanser, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseErrorWithMsg(c, CodeInvalidParam, err.Error())
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(tanser.Translate(trans)))
		return
	}
	//
	//fmt.Println("userLogin", userLogin)
	//后端存储到数据库中的一行数据对应的结构体
	user := new(models.User)
	user.UserAddress = userLogin.UserAddress
	//fmt.Println("user", user)
	//业务处理
	data, err := logic.Login(user)
	//fmt.Println(token)
	if err != nil {
		zap.L().Error("Logic.Login failed", zap.String("username", user.UserAddress), zap.Error(err))
		if errors.Is(err, mysql.ErrorUserNotExist) {
			ResponseError(c, CodeUserNotExist)
			return
		} else if errors.Is(err, mysql.ErrorInvalidPassword) {
			ResponseError(c, CodeInvalidPassword)
			return
		}
		ResponseErrorWithMsg(c, CodeMysql, err.Error())
		return
	}
	response := new(models.ResponseLoginData)
	if data.BackGroundPic.Valid {
		response.BackGroundPic = data.BackGroundPic.String
		response.UserAddress = data.UserAddress
		response.HeadPicture = data.HeadPicture
		response.UserName = data.UserName
	} else {
		response.UserAddress = data.UserAddress
		response.HeadPicture = data.HeadPicture
		response.UserName = data.UserName
		response.BackGroundPic = ""
	}
	fmt.Println("response", response)
	//返回请求
	ResponseSuccess(c, response)
}

func GetUserBalanceHandler(c *gin.Context) {
	user_address := c.Param("user_address")
	data, err := logic.GetUserBalance(user_address)
	if err != nil {
		zap.L().Error("Logic.GetUserBalance failed", zap.String("user_address", user_address), zap.Error(err))
		if errors.Is(err, mysql.ErrorUserNotExist) {
			ResponseError(c, CodeUserNotExist)
			return
		} else if errors.Is(err, mysql.ErrorInvalidPassword) {
			ResponseError(c, CodeInvalidPassword)
			return
		}
		ResponseErrorWithMsg(c, CodeMysql, err.Error())
		return
	}
	ResponseSuccess(c, data)
}

func AddBalanceHandler(c *gin.Context) {
	user_address := c.Param("user_address")
	amountS := c.Query("amount")
	amount, _ := strconv.Atoi(amountS)
	//
	data, err := logic.AddUserBalance(user_address, amount)
	if err != nil {
		zap.L().Error("logic.AddUserBalance(p) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}

func SubBalanceHandler(c *gin.Context) {
	user_address := c.Param("user_address")
	amountS := c.Query("amount")
	amount, _ := strconv.Atoi(amountS)

	data, err := logic.SubUserBalance(user_address, amount)
	if err != nil {
		zap.L().Error("logic.AddUserBalance(p) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}

//获取用户信息
func GetUserInformation(c *gin.Context) {
	user_address := c.Param("user_address")
	data, err := logic.GetUserInformation(user_address)
	if err != nil {
		zap.L().Error("logic.AddUserBalance(p) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}

//获取用户的所有拥有的皮肤
func GetAllSkinByUserHandler(c *gin.Context) {
	user_address := c.Param("user_address")
	data, err := logic.GetAllSkinByUser(user_address)
	if err != nil {
		zap.L().Error("logic.GetAllSkinByUser(p) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}

func GetPostFromUserAddHandler(c *gin.Context) {
	user_address := c.Param("user_address")
	//
	//data := &models.PostFromUser{
	//	UserAddress: "0xA7c2711DFE3B09Da2Ffce80E86ec0f18958AB151",
	//	UserName:    "老庄",
	//	Title:       "人工智能的概论讲解",
	//	PostID:      1573247996923904,
	//	Content:     "主贴内容，人工智能是一个好东西",
	//}
	//fmt.Println(*data)
	data, err := logic.GetPostFromUserAdd(user_address)
	num := len(data)
	if err != nil {
		zap.L().Error("logic.GetPostFromUserAdd(p) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	//
	ResponseSuccessGetPost(c, data, num)
}

//修改用户信息
func ChangeUserInformationHandler(c *gin.Context) {
	userprofile := new(models.UserProfile)
	if err := c.ShouldBindJSON(userprofile); err != nil {
		zap.L().Error("Change	UserInformation is failed", zap.Error(err))
		tanser, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseErrorWithMsg(c, CodeInvalidParam, err.Error())
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(tanser.Translate(trans)))
		return
	}
	//
	fmt.Println("userprofile:", &userprofile)
	//
	if err := logic.ChangeUserInformation(userprofile); err != nil {
		zap.L().Error("logic.ChangeUserInformation(p) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, "success change")
}

//修改用户的背景---用户和ksin——id的对应是在购买函数
func ChangeUserBackGroundHandler(c *gin.Context) {
	useraddress := c.Param("user_address")

	changeinfo := new(models.ChangeBCGByUser)
	changeinfo.UserAddress = useraddress
	if err := c.ShouldBindJSON(changeinfo); err != nil {
		zap.L().Error("Change	UserBackGround is failed", zap.Error(err))
		tanser, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseErrorWithMsg(c, CodeInvalidParam, err.Error())
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(tanser.Translate(trans)))
		return
	}
	fmt.Println("changeinfo:", changeinfo)
	data, err := logic.ChangeUserBackGround(changeinfo)
	if err != nil {
		zap.L().Error("logic.ChangeUserBackGround is failed", zap.Error(err))
		if err == logic.CodeURLNOTEXIT {
			fmt.Println("change..err:没有购买，拿不到")
			ResponseErrorWithMsg(c, CodeInvalidParam, err.Error())
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, err.Error())
		return
	}
	fmt.Println("data:", data)
	ResponseSuccess(c, data)

}

func ChangeUserPHHandler(c *gin.Context) {
	useraddress := c.Param("user_address")

	changeinfo := new(models.ChangeHPByUser)

	changeinfo.UserAddress = useraddress
	if err := c.ShouldBindJSON(changeinfo); err != nil {
		zap.L().Error("Change	UserBackGround is failed", zap.Error(err))
		tanser, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseErrorWithMsg(c, CodeInvalidParam, err.Error())
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(tanser.Translate(trans)))
		return
	}
	fmt.Println("changeinfo", changeinfo)
	//
	data, err := logic.ChangeUserPH(changeinfo)
	if err != nil {
		zap.L().Error("logic.ChangeUserBackGround is failed", zap.Error(err))
		if err == logic.CodeURLNOTEXIT {
			fmt.Println("change..err:没有购买，拿不到")
			ResponseErrorWithMsg(c, CodeInvalidParam, err.Error())
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, err.Error())
		return
	}
	fmt.Println("data:", data)
	ResponseSuccess(c, data)
}
