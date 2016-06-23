package controllers

import (
	"ObjectId/models"
	"encoding/json"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Post() {
	params := make(map[string]interface{})        //接收数据
	callback_data := make(map[string]interface{}) //返回数据

	json.Unmarshal(c.Ctx.Input.RequestBody, &params)

	user_name := params["username"].(string)

	//插入数据
	object_id, err := models.InstertUname(user_name)

	//返回数据
	if err == nil {
		callback_data["object_id"] = object_id

		rdata, err := json.Marshal(callback_data)
		json.Unmarshal(rdata, callback_data)
		if err != nil {
			panic(err)
		}

		c.Data["json"] = callback_data
		c.ServeJSON()
	}
}
