package service

import (
	"chatgpt/models"

	"github.com/gin-gonic/gin"
)

// GetUserList
// @Tags userservice
// @Success 200 {string} userList
// @Router /user/getUserList [get]
func GetUserList(c *gin.Context) {
	data := make([]*models.UserBasic, 10)
	data = models.GetUserList()
	c.JSON(200, gin.H{
		"message": data,
	})
}
