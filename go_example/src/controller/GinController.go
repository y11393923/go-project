package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HelloGin() {
	Engine.GET("/gin", func(c *gin.Context) {
		//返回一个JSON格式的字符串
		c.JSON(http.StatusOK, gin.H{
			"Blog":   "www.flysnow.org",
			"wechat": "flysnow_org",
		})
	})
}
