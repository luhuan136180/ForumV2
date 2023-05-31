package controller

import (
	"errors"
	"furumvv2/dao/mysql"
	"furumvv2/logic"
	"furumvv2/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"strconv"
)

func GetAllSkinListHanlder(c *gin.Context) {
	status := c.Param("status")
	Status, _ := strconv.Atoi(status)
	data, err := logic.GetAllSkinList(Status)
	if err != nil {
		zap.L().Error("logic.GetAllSkinList(p) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, data)
}

func ShopSkinByUserHandler(c *gin.Context) {
	shopInformation := new(models.Shop)
	if err := c.ShouldBindJSON(shopInformation); err != nil {
		zap.L().Error("Shoping Skin by User_address is failed", zap.Error(err))
		tanser, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseErrorWithMsg(c, CodeInvalidParam, err.Error())
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(tanser.Translate(trans)))
		return
	}
	//fmt.Println(shopInformation)  --购买的业务逻辑
	err := logic.ShopSkinByUser(shopInformation)
	if err != nil {
		zap.L().Error("logic.GetAllSkinList(p) failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorSkinISExit) {
			ResponseError(c, CodeISShopping)
			return
		}
		if errors.Is(err, mysql.ErrorNotEnoughAmount) {
			ResponseError(c, CodeNOAmount)
		}
		ResponseErrorWithMsg(c, CodeMysql, err.Error())
		return
	}

	//
	ResponseSuccess(c, shopInformation)
}
