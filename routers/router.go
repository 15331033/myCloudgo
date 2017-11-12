package routers

import (
	"github.com/astaxie/beego"
	"hello/controllers"
)

func init() {
	//set the route path and choose the handler
	beego.Router("/hello/:id", &controllers.FileController{})
}
