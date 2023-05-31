package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
)

const ContextUserIDKey = "userID"
const ContextUserAddressKey = "userAddress"

var ErrorUserNotLogin = errors.New("用户未登录")

func GetCurrentUser(c *gin.Context) (userAddress string, err error) {
	uaddress, ok := c.Get(ContextUserAddressKey)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	userAddress, ok = uaddress.(string) //强转
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	return
}

func GetPageInfo(c *gin.Context) (int64, int64) {
	pageStr := c.Query("page")
	sizeStr := c.Query("size")

	var (
		page int64
		size int64
		err  error
	)

	page, err = strconv.ParseInt(pageStr, 10, 64)
	if err != nil {
		page = 1
	}
	size, err = strconv.ParseInt(sizeStr, 10, 64)
	if err != nil {
		size = 10
	}
	return page, size
}
