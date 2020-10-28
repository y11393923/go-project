package controller

import (
	"github.com/gin-gonic/gin"
	"log"
)

//实例化一个默认的gin示例
var Engine = gin.Default()

func InitWeb() {
	HelloGin()
	InitGin()
	err := Engine.Run(":8080")
	if err != nil {
		log.Fatal("init web failed ", err)
	}

	/*Router()
	beego.Run()*/
}
