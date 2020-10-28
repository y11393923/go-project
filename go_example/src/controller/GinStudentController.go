package controller

import (
	"config"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"model"
	"net/http"
	"strconv"
)

func InitGin() {
	student := Engine.Group("/student")
	{
		student.POST("/", insertStudent)
		student.PUT("/", updateStudent)
		student.DELETE("/:id", deleteStudentById)
		student.GET("/:id", selectById)
		student.GET("/", selectAll)
	}
}

func insertStudent(c *gin.Context) {
	fmt.Println(c.FullPath())
	var stu model.Student
	//获取json参数
	if err := c.BindJSON(&stu); err != nil {
		log.Fatal(err.Error())
		return
	}
	_, err := model.InsertStudent(config.DB, stu)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func updateStudent(c *gin.Context) {
	fmt.Println(c.FullPath())
	//获取json参数
	var stu model.Student
	if err := c.BindJSON(&stu); err != nil {
		log.Fatal(err.Error())
		return
	}
	_, err := model.UpdateStudent(config.DB, stu)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func deleteStudentById(c *gin.Context) {
	fmt.Println(c.FullPath())
	//获取url参数并转换类型
	id, _ := strconv.Atoi(c.Param("id"))
	_, err := model.DeleteStudentById(config.DB, id)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func selectById(c *gin.Context) {
	fmt.Println(c.FullPath())
	//获取url参数并转换类型
	id, _ := strconv.Atoi(c.Param("id"))
	stu, err := model.SelectById(config.DB, id)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	c.JSON(http.StatusOK, stu)
}

func selectAll(c *gin.Context) {
	fmt.Println(c.FullPath())
	stus, err := model.SelectAll(config.DB)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	c.JSON(http.StatusOK, stus)
}
