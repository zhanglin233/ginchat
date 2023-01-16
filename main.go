package main

import (
	"chatgpt/models"
	"chatgpt/utils"
)

func main() {
	utils.InitConfig()
	utils.InitMysql()
	models.GetUserList()
	// r := router.Router()

	// r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
