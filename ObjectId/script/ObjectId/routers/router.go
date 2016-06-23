package routers

import (
	"ObjectId/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/getobjectid", &controllers.MainController{})

	beego.Router("/adduserinfo", &controllers.AddUserInfoController{})
	beego.Router("/getuserinfo", &controllers.GetUserInfoController{})
}
