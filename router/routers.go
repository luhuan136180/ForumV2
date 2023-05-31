package router

import (
	"furumvv2/controller"
	"furumvv2/logger"
	middleWares "furumvv2/middlerWares"
	"github.com/gin-gonic/gin"
)

func SetUpRouter(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) //设置成发布模式
	}

	r := gin.New()
	//使用自己编写的logger
	//r.Use(logger.GinLogger(), logger.GinRecovery(true), middleWares.RateLimitMiddleware(2*time.Second, 1))
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	//分组
	v1 := r.Group("/api/v1") //组
	v1.Use(middleWares.Proxy())
	v1.GET("/helloworld", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello world!",
		})
	})

	//注册
	//v1.POST("/signup", controller.SignUpHandler)
	//登录
	v1.POST("/login", controller.LoginHandler)

	//继续访问该项目的内容需要先登录，所以需要认证处理，添加中间件
	//v1.Use(middleWares.JWTAuthMiddleware()) //应用JWT认证中间件
	{ //通过认证后可以使用的功能
		//创建帖子------注意：：：：需要修改成表单传入
		v1.POST("/post/create", controller.CreatePostHandler)
		//查询帖子  /post/list?page=x&size=y
		v1.GET("/post/list", controller.GetPostsListHandler)
		////查询一些帖子---模糊查询--对内容版本
		v1.GET("post/like-content/:word", controller.GetPostByContentLIKEHandler)
		//查询一些帖子---模糊查询--对标题
		v1.GET("post/like-title/:word", controller.GetPostByTitleLIKEHandler)
		//查询整个帖子（主贴+回复）
		v1.GET("/post/getpost/:postid", controller.GetPostByPostID)
		//对帖子发表评论
		v1.POST("post/:postID/response/", controller.CreateResponseHandler)
		//查询用户信息
		v1.GET("/user/:user_address/getuserInformation", controller.GetUserInformation)
		//修改昵称

		//
		//查询一个用户的余额
		v1.GET("/user/:user_address/balance", controller.GetUserBalanceHandler)

		//进行交易，改变用户余额--添加
		v1.POST("/user/:user_address/add_balance", controller.AddBalanceHandler)
		////进行交易，改变用户余额--支出
		v1.PUT("/user/:user_address/sub_amount", controller.SubBalanceHandler)
		//
		//上传文件？？---纯上传API
		v1.POST("/file", controller.UploadFile)
		//上传文件，同时记录上传用户等信息
		//v1.POST("/file/user", controller.UploadFileWithAuthor)

		//市场显示
		v1.GET("/market/skins/:status", controller.GetAllSkinListHanlder)
		//展示用户个人的所拥有皮肤
		v1.GET("/user/:user_address/skinList", controller.GetAllSkinByUserHandler)
		//买皮肤---此API废弃
		v1.GET("market/skins/shop", controller.ShopSkinByUserHandler)
	}

	return r
}
