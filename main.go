package main

import (
	"ToDo/dao"
	"ToDo/models"
	"ToDo/router"
	"fmt"
)

func main() {
	//创建数据库
	//连接数据库
	err := dao.InitMysql()
	if err != nil {
		fmt.Println("gorm.open failed:", err.Error())
	}
	//defer DB.Close()
	//绑定模型
	dao.DB.AutoMigrate(&models.Todo{})
	r := router.SetupRouter()
	r.Run(":9999")
}
