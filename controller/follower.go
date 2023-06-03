package controller

import (
	"furumvv2/logic"
	"furumvv2/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

func GetFollowerListHandler(c *gin.Context) {
	//获取信息
	user_address := c.Param("user_address")
	//业务流程
	data, err := logic.GetFollowerList(user_address)
	if err != nil {
		zap.L().Error("logic.GetFollowerList(p) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, data)
}

func AddFollowerByUserIDHandler(c *gin.Context) {
	//读取被关注者和关注者
	followerInfo := new(models.Follower)
	if err := c.ShouldBindJSON(followerInfo); err != nil {
		zap.L().Error("AddFollower is falied", zap.Error(err))
		tanser, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseErrorWithMsg(c, CodeInvalidParam, err.Error())
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(tanser.Translate(trans)))
		return
	}
	//
	data, err := logic.AddFollowerByUserID(followerInfo)
	if err != nil {
		zap.L().Error("logic.AddFollowerByUserID(p) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	if data == "exit" {
		ResponseError(c, CodeFollowerExit)
		return
	}
	ResponseSuccess(c, data)
}
