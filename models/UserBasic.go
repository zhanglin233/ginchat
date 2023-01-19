package models

import (
	"chatgpt/utils"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type UserBasic struct {
	gorm.Model
	Name          string
	Password      string
	Salt          string
	Phone         string `valid:"matches(^1[3-9]{1}\\d{9}$)"`
	Email         string `valid:"email"` //valid:后面不能有空格
	Identity      string
	ClientIp      string
	ClientPort    string
	LoginTime     time.Time
	HeartbeatTime time.Time
	LogoutTIME    time.Time
	IsLogout      bool
	DeviceInfo    string
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}

func GetUserList() []*UserBasic {
	DB := utils.DB
	data := make([]*UserBasic, 10)
	DB.Find(&data)
	for _, v := range data {
		fmt.Println(v)
	}
	return data
}

func FindUserByName(name string) UserBasic {
	user := UserBasic{}
	utils.DB.Where("name = ?", name).First(&user)
	return user
}

func FindUserByPhone(phone string) UserBasic {
	user := UserBasic{}
	utils.DB.Where("phone = ?", phone).First(&user)
	return user
}

func FindUserByEmail(email string) UserBasic {
	user := UserBasic{}
	utils.DB.Where("email = ?", email).First(&user)
	return user
}

func FindUserByNameAndPwd(name, password string) UserBasic {
	user := UserBasic{}
	utils.DB.Where("name = ? and pass_word= ?", name, password).First(&user)

	temp := utils.MD5Encode(fmt.Sprintf("%d", time.Now().Unix()))

	utils.DB.Model(&user).Where("id = ?", user.ID).Update("identity", temp)
	return user
}

func CreateUser(user UserBasic) *gorm.DB {

	return utils.DB.Create(&user)
}

func DeleteUser(user UserBasic) *gorm.DB {
	return utils.DB.Delete(&user)
}

func UpdateUser(user UserBasic) *gorm.DB {
	return utils.DB.Model(&user).Updates(UserBasic{Name: user.Name, Password: user.Password, Phone: user.Phone, Email: user.Email})
}
