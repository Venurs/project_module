package routers

import (
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	//"github.com/justinas/nosurf"

	"kjlive-service/controllers"
	"kjlive-service/middlewares"
)

var DB = make(map[string]string)

// 初始化路由
func setupRouter() *gin.Engine {
	r := gin.New()
	// 使用用户信息中间件
	r.Use(middlewares.UserInfoMiddleware())
	// 使用Logger
	r.Use(middlewares.Logger())
	r.Use(middlewares.ApiInterfaceMiddleware())
	//r.Use(middlewares.CsrfMiddleware())
	liveApi := r.Group("/liveapi")
	{
		liveApi.GET("/live_info", controllers.LiveInfoController())
		liveApi.GET("/live_constant_info", controllers.LiveConstantInfoController())
		liveApi.GET("/user_info", controllers.UserInfoController())
		liveApi.GET("/new_user_info", controllers.NewUserInfoController())
	}

	return r
}

func Run() {
	r := setupRouter()
	//nosurf.CookieName := "ucservicecsrftoken"
	//endless.ListenAndServe(":8080", nosurf.New(r))
	endless.ListenAndServe(":3888", r)
}

func GetRouter() *gin.Engine {
	r := setupRouter()
	return r
}