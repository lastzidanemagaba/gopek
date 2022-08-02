package web

import (
	authcontroller "github.com/antare74/gopek/controller/auth"
	homecontroller "github.com/antare74/gopek/controller/home"
	usercontroller "github.com/antare74/gopek/controller/user"
	"github.com/gin-gonic/gin"
)

func Routes(route *gin.Engine) {
	home := route.Group("/home")
	{
		home.GET("/", homecontroller.Index)
	}
	user := route.Group("/user")
	{
		user.GET("/", usercontroller.Index)
		user.GET("/croot", usercontroller.Croot)
	}
	auth := route.Group("/auth")
	{
		auth.GET("/", authcontroller.Index)
	}
}
