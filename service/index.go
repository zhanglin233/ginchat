package service

import (
	"chatgpt/models"
	"html/template"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetIndex
// @Tags 首页
// @Success 200 {string} welcome
// @Router /index [get]
func GetIndex(c *gin.Context) {
	ind, err := template.ParseFiles("index.html", "views/chat/head.html")
	if err != nil {
		panic(err)
	}
	ind.Execute(c.Writer, "index")
	// c.JSON(200, gin.H{
	// 	"message": "welcome",
	// })
}

// GetIndex
// @Tags 注册页
// @Success 200 {string} welcome
// @Router /toRegister [get]
func ToRegister(c *gin.Context) {
	ind, err := template.ParseFiles("register.html")
	if err != nil {
		panic(err)
	}
	ind.Execute(c.Writer, "toRegister")
}

// GetIndex
// @Tags 聊天页
// @Success 200 {string} welcome
// @Router /toChat [get]
func ToChat(c *gin.Context) {
	ind, err := template.ParseFiles("views/chat/index.html",
		"views/chat/head.html",
		"views/chat/tabmenu.html",
		"views/chat/concat.html",
		"views/chat/group.html",
		"views/chat/profile.html",
		"views/chat/main.html",
		"views/chat/createcom.html",
		"views/chat/userinfo.html",
		"views/chat/foot.html")
	if err != nil {
		panic(err)
	}
	userId, _ := strconv.Atoi(c.Query("userId"))
	identity := c.Query("token")
	user := models.UserBasic{}
	user.ID = uint(userId)
	user.Identity = identity
	ind.Execute(c.Writer, user)
}
