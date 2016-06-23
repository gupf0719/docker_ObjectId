package models

import (
	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Userinfo struct {
	Id_       bson.ObjectId `bson:"_id"`
	Object_id1 string
	Object_id2 string
	Object_id3 string
	Object_id4 string
	Object_id5 string
}

func Adduserinfo(m *Userinfo) error{
	session, err := mgo.Dial(beego.AppConfig.String("url")) //连接数据库
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	db := session.DB("user")     //数据库名称
	collection := db.C("userinfo") //如果该集合已经存在的话，则直接返回

	err = collection.Insert(m)

	if err != nil {
		return err
	}

	return nil
}

func GetinfoById(object_id string) (*Userinfo,error){
	session, err := mgo.Dial(beego.AppConfig.String("url")) //连接数据库
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	db := session.DB("user")     //数据库名称
	collection := db.C("userinfo") //如果该集合已经存在的话，则直接返回

	result := Userinfo{}
	err = collection.Find(bson.M{"_id": bson.ObjectIdHex(object_id)}).One(&result)

	if err !=nil{
		return &result,err
	}

	return &result,nil
}
