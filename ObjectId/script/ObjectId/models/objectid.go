package models

import (
	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id_       bson.ObjectId `bson:"_id"`
	User_name string
}

//*******插入User*******
func InstertUname(user_name string) (string, error) {
	session, err := mgo.Dial(beego.AppConfig.String("url")) //连接数据库
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	db := session.DB("user")   //数据库名称
	collection := db.C("user") //如果该集合已经存在的话，则直接返回

	objectid := bson.NewObjectId()
	user := &User{
		Id_:       objectid,
		User_name: user_name,
	}
	err = collection.Insert(user)
	if err != nil {
		return "", err
	}

	return objectid.Hex(), nil
}
