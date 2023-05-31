package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func UploadFile(c *gin.Context) {
	// 单个文件
	file, err := c.FormFile("file")
	if err != nil {
		//日志记录
		zap.L().Debug("**c.FormFile(\"file\") error", zap.Any("err", err))
		zap.L().Error("***获取上传的文件", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	//
	// 保存文件到本地
	dst := fmt.Sprintf("./tmp/%s", file.Filename)
	err = c.SaveUploadedFile(file, dst)
	if err != nil {
		zap.L().Error("c.SaveUploadedFile(file, dst) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	data := "上传文件成功，文件名：" + file.Filename
	//响应
	ResponseSuccess(c, data)

}
