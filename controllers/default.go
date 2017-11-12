package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

//set a new controller
type FileController struct {
	beego.Controller
}

type hello struct {
	id string
}

func (this *FileController) Get() {
	//get the value testuser
	a := this.Ctx.Input.Param(":id")
	//set the json data
	this.Data["json"] = map[string]interface{}{"Test": "Hello " + a}
	this.ServeJSON()
}

func (c *MainController) Get() {
	a := c.Ctx.Input.Param(":id")
	c.Ctx.WriteString("hello " + a)
}
