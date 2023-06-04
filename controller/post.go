package controller

import (
	"fmt"
	"furumvv2/logic"
	"furumvv2/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"strconv"
)

func CreatePostHandler(c *gin.Context) {
	//获取参数和参数的校验
	post := new(models.CtreatePost)                //创建
	if err := c.ShouldBindJSON(post); err != nil { //获取 标题，内容，作者address，图片url，
		zap.L().Error("Login with invalid param", zap.Error(err))
		tanser, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseErrorWithMsg(c, CodeInvalidParam, err.Error())
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(tanser.Translate(trans)))
		return
	}

	//创建帖子的业务流程
	data, err := logic.CreatePost(post)
	if err != nil {
		zap.L().Error("logic.CreatPost(p) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	//响应
	ResponseSuccess(c, data)
}

//
func GetPostsListHandler(c *gin.Context) {
	fmt.Println("hello")
	page, size := GetPageInfo(c)
	//获取页数和每页显示量
	//处理业务逻辑，查询
	//fmt.Println(page, size)
	data, err := logic.GetPostsList(page, size)
	if err != nil {
		zap.L().Error("logic.GetPostsList(p) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}

//对主贴的内容 进行模糊查询
func GetPostByContentLIKEHandler(c *gin.Context) {
	//模糊查询内容，对主贴
	word := c.Param("word")
	data, err := logic.GetPostByContentLIKE(word)
	if err != nil {
		zap.L().Error("logic.GetPostByContentLIKE(p) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	fmt.Println("data:===", word)
	ResponseSuccess(c, data)
}

func GetPostByTitleLIKEHandler(c *gin.Context) {
	word := c.Param("word")
	data, err := logic.GetPostByTitleLIKE(word)
	if err != nil {
		zap.L().Error("logic.GetPostByTitleLIKE(p) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}

func CreateResponseHandler(c *gin.Context) {
	postid := c.Param("postID")
	post := new(models.CtreatePost)                //创建
	if err := c.ShouldBindJSON(post); err != nil { //获取  内容，作者address，
		zap.L().Error("Login with invalid param", zap.Error(err))
		tanser, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseErrorWithMsg(c, CodeInvalidParam, err.Error())
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(tanser.Translate(trans)))
		return
	}
	postID, _ := strconv.ParseInt(postid, 10, 64)
	post.PostID = postID

	fmt.Println("postResponse creeate:", post)
	//
	data, err := logic.CreateResponseByPostID(post)
	if err != nil {
		zap.L().Error("logic.CreateResponseByPostID(p) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}

//查询单个帖子的所有内容
func GetPostByPostID(c *gin.Context) {
	postid := c.Param("postid")

	data, err := logic.GetPostByPostID(postid)
	if err != nil {
		zap.L().Error("logic.GetPostByPostID(p) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	//for _, val := range data {
	//	fmt.Println(val)
	//}

	ResponseSuccess(c, data)
}
