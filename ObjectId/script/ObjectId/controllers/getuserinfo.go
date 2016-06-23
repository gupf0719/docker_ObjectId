package controllers

import (
	"ObjectId/models"
	"encoding/json"

	"github.com/astaxie/beego"
)

type GetUserInfoController struct {
	beego.Controller
}

func (c *GetUserInfoController) Post() {
	params := make(map[string]interface{})        //接收数据
	callback_data := make(map[string]interface{}) //返回数据

	json.Unmarshal(c.Ctx.Input.RequestBody, &params)

	object_id := params["object_id"].(string)

	userinfo, err := models.GetinfoById(object_id)

	if err != nil {
		callback_data["state"] = "fail"
		callback_data["object_id"] = object_id
	} else {
		callback_data["state"] = "fail"
		callback_data["object_id"] = object_id
		callback_data["object_ids"] = userinfo
	}

	rdata, err := json.Marshal(callback_data)
	json.Unmarshal(rdata, callback_data)
	if err != nil {
		panic(err)
	}
	c.Data["json"] = callback_data
	c.ServeJSON()
}
