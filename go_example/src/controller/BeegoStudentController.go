package controller

import (
	"config"
	"encoding/json"
	"github.com/astaxie/beego"
	"log"
	"model"
	"strconv"
)

type StudentsController struct {
	beego.Controller
}

type ParamController struct {
	beego.Controller
}

func (this *StudentsController) Get() {
	stus, err := model.SelectAll(config.DB)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	//返回json格式
	this.Data["json"] = stus
	this.ServeJSON()
}

func (this *ParamController) Get() {
	id, _ := strconv.Atoi(this.GetString(":id"))
	stu, err := model.SelectById(config.DB, id)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	this.Data["json"] = stu
	this.ServeJSON()
}

func (this *StudentsController) Post() {
	var stu model.Student
	//获取json参数并绑定  要在app.conf中设置copyrequestbody=true
	e := json.Unmarshal(this.Ctx.Input.RequestBody, &stu)
	if e != nil {
		log.Fatal(e.Error())
		return
	}
	_, err := model.InsertStudent(config.DB, stu)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	this.Ctx.WriteString("ok")
}

func (this *StudentsController) Put() {
	var stu model.Student
	//获取json参数并绑定
	e := json.Unmarshal(this.Ctx.Input.RequestBody, &stu)
	if e != nil {
		log.Fatal(e.Error())
		return
	}
	_, err := model.UpdateStudent(config.DB, stu)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	this.Ctx.WriteString("ok")
}

func (this *StudentsController) Delete() {
	id, _ := strconv.Atoi(this.GetString(":id"))
	_, err := model.DeleteStudentById(config.DB, id)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	this.Ctx.WriteString("ok")
}

func Router() {
	beego.Router("/student", &StudentsController{})
	beego.Router("/student/:id", &ParamController{})
}
