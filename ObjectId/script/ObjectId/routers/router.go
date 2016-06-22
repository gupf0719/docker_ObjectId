package routers

import (
	"ObjectId/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/getobjectid", &controllers.MainController{})
}
