package service

import (
	"chatgpt/models"
	"chatgpt/utils"
	"fmt"
	"math/rand"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

// GetUserList
// @Summary 所有用户
// @Tags 用户模块
// @Success 200 {string} json{code,"message"}
// @Router /user/getUserList [get]
func GetUserList(c *gin.Context) {
	data := make([]*models.UserBasic, 10)
	data = models.GetUserList()
	c.JSON(200, gin.H{
		"message": data,
	})
}

// CreateUser
// @Summary 添加用户
// @Tags 用户模块
// @param name query string false "用户名"
// @param password query string false "密码"
// @param rePassword query string false "确认密码"
// @param phone query string false "电话号码"
// @param email query string false "邮箱"
// @Success 200 {string} json{code,"message"}
// @Router /user/createUser [get]
func CreateUser(c *gin.Context) {
	user := models.UserBasic{}
	user.Name = c.Query("name")
	user.Password = c.Query("password")
	rePassword := c.Query("rePassword")
	user.Phone = c.Query("phone")
	user.Email = c.Query("email")
	isNameExist := models.FindUserByName(user.Name)
	isPhoneExist := models.FindUserByPhone(user.Phone)
	isEmailExist := models.FindUserByEmail(user.Email)
	if isNameExist {
		c.JSON(-1, gin.H{
			"message": "用户名已存在",
		})
	} else if isPhoneExist {
		c.JSON(-2, gin.H{
			"message": "电话号码已注册",
		})
	} else if isEmailExist {
		c.JSON(-3, gin.H{
			"message": "邮箱已注册",
		})
	} else if user.Password != rePassword {
		c.JSON(-4, gin.H{
			"message": "两次密码不一致",
		})
	} else {
		user.Salt = fmt.Sprintf("%06d", rand.Int31())
		user.Password = utils.MakePassword(user.Password, user.Salt)
		models.CreateUser(user)
		c.JSON(200, gin.H{
			"message": "添加用户成功",
		})
	}

}

// DeleteUser
// @Summary 删除用户
// @Tags 用户模块
// @param id query string false "用户id"
// @Success 200 {string} json{code,"message"}
// @Router /user/deleteUser [get]
func DeleteUser(c *gin.Context) {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(c.Query("id"))
	user.ID = uint(id)

	models.DeleteUser(user)
	c.JSON(200, gin.H{
		"message": "删除用户成功",
	})
}

// UpdateUser
// @Summary 修改用户
// @Tags 用户模块
// @param id formData string false "用户id"
// @param name formData string false "用户名"
// @param password formData string false "密码"
// @param phone formData string false "电话号码"
// @param email formData string false "邮箱"
// @Success 200 {string} json{code,"message"}
// @Router /user/updateUser [post]
func UpdateUser(c *gin.Context) {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(c.PostForm("id"))
	user.ID = uint(id)
	user.Name = c.PostForm("name")
	user.Password = c.PostForm("password")
	user.Phone = c.PostForm("phone")
	user.Email = c.PostForm("email")
	fmt.Println("update: ", user)

	_, err := govalidator.ValidateStruct(user)
	if err != nil {
		fmt.Println(err)
		c.JSON(200, gin.H{
			"message": "修改参数不匹配",
		})
	} else {
		models.UpdateUser(user)
		c.JSON(200, gin.H{
			"message": "修改用户成功",
		})
	}

}
