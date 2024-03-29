package web

import "github.com/gin-gonic/gin"

// UserHandler 定义所有与User相关的路由
type UserHandler struct {
}

// RegisterUserRoutes 注册用户相关路由
func (u *UserHandler) RegisterUserRoutes(server *gin.Engine) {
	// 路由分组简化路由
	ug := server.Group("/users")
	ug.POST("/signup", u.SignUp)
	ug.POST("/login", u.Login)
	ug.POST("/edit", u.Edit)
	ug.GET("/profile", u.Profile)
}

func (u *UserHandler) SignUp(c *gin.Context) {
}

func (u *UserHandler) Login(c *gin.Context) {

}

func (u *UserHandler) Edit(c *gin.Context) {

}

func (u *UserHandler) Profile(c *gin.Context) {

}
