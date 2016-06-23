package controllers

import (
	"ObjectId/models"
	"encoding/json"

	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2/bson"
)

type AddUserInfoController struct {
	beego.Controller
}

func (c *AddUserInfoController) Post() {
	params := make(map[string]interface{})        //接收数据
	callback_data := make(map[string]interface{}) //返回数据

	json.Unmarshal(c.Ctx.Input.RequestBody, &params)

	object_id := params["object_id"].(string)
	o_id := bson.ObjectIdHex(object_id)

	object_ids := params["object_ids"].(map[string]interface{})
	object_id1 := object_ids["object_id1"].(string)
	object_id2 := object_ids["object_id2"].(string)
	object_id3 := object_ids["object_id3"].(string)
	object_id4 := object_ids["object_id4"].(string)
	object_id5 := object_ids["object_id5"].(string)

	userinfo := models.Userinfo{Id_: o_id, Object_id1: object_id1, Object_id2: object_id2, Object_id3: object_id3,
		Object_id4: object_id4, Object_id5: object_id5}

	err := models.Adduserinfo(&userinfo)

	//返回数据
	if err != nil {
		callback_data["state"] = "fail"
	} else {
		callback_data["state"] = "success"
	}

	rdata, err := json.Marshal(callback_data)
	json.Unmarshal(rdata, callback_data)
	if err != nil {
		panic(err)
	}
	c.Data["json"] = callback_data
	c.ServeJSON()
}
