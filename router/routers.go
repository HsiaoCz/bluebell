package router

import (
	controllors "bluebell/controllors"
	"bluebell/logger"
	"bluebell/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Setup() *gin.Engine {
	//当app中的mode 和gin框架中的release 模式相同的时候
	//将gin 框架设置成release模式
	if viper.GetString("app.mode") == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	v1 := r.Group("/api/v1")
	//注册
	v1.POST("/signup", controllors.SingUpHandler)
	//登录
	v1.POST("/login", controllors.LoginHandler)

	//认证中间件
	v1.Use(middlewares.JWTAuthMiddleware())
	{
		v1.GET("/community", controllors.CommunityHandler)
		v1.GET("/community/:id", controllors.CommunityDetailHandler)
		v1.POST("/post", controllors.CreatePostHandler)
		v1.GET("/post/:id", controllors.GetPostDetailHandler)
		v1.GET("/posts/", controllors.GetPostListHandler)
        //根据时间或分数获取帖子列表
        v1.GET("/posts1",controllors.GetPostListHandler1)
		v1.POST("/vote", controllors.PostVoteHandler)

	}

	r.GET("/", middlewares.JWTAuthMiddleware(), func(c *gin.Context) {
		//如果是登录的用户，判断请求头中是否有有效的JWT
		c.String(http.StatusOK, "pong")
	})
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": 404,
		})
	})
	return r
}
