package controllers

import (
	"fmt"
	"encoding/json"
	"ObjectId/models"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Post() {
	params := make(map[string]interface{})        //接收数据
	callback_data := make(map[string]interface{}) //返回数据

	json.Unmarshal(c.Ctx.Input.RequestBody, &params)
	fmt.Println(c.Ctx.Input.RequestBody)
	fmt.Println(params["username"].(string))
	user_name := params["username"].(string)

	//插入数据
	objectid, err := models.InstertUname(user_name)

	//返回数据
	if err == nil {
		callback_data["objectid"] = objectid

		rdata, err := json.Marshal(callback_data)
		json.Unmarshal(rdata, callback_data)
		if err != nil {
			panic(err)
		}

		c.Data["json"] = callback_data
		c.ServeJSON()
	}
}
