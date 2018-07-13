package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"fmt"
)

func main(){
	testRemoveselector()
}

func testRemoveId(){
	ico, _ := mgo.Dial("mongodb://root:yzl17220@120.78.158.127:27017")
	db := ico.DB("s_blog")
	mgotest4 := db.C("mgotest4")
	id:=bson.ObjectIdHex("5b3f27fc2f8e6b1898a27c58")
	//id:=bson.M{"_id":bson.ObjectIdHex("5b3f27fc2f8e6b1898a27c58")}
	var data []interface{}
	mgotest4.FindId(id).All(&data)
	fmt.Println(data)
	mgotest4.RemoveId(id)
	mgotest4.Insert(data...)

}

func testRemoveselector(){
	ico, _ := mgo.Dial("mongodb://root:yzl17220@120.78.158.127:27017")
	db := ico.DB("s_blog")
	mgotest4 := db.C("mgotest4")
	selector:=bson.M{"account":"admin"}
	var data []interface{}
	mgotest4.Find(selector).All(&data)
	fmt.Println(data)
	mgotest4.Remove(selector)
	mgotest4.Insert(data...)
}

func testUpdateid(){

}


func insert(){
	ico, _ := mgo.Dial("mongodb://root:yzl17220@120.78.158.127:27017")
	db := ico.DB("s_blog")
	mgotest4 := db.C("mgotest4")
	var id  interface{}

	//id = bson.NewObjectId()
	//fmt.Println("bson.NewObjectId() ID:",id)
	//insertData := bson.M{"_id":id ,"account":"admin","balance":"更新前" }
	insertData := bson.M{"account":"admin","balance":"加ID更新前" }
	if insertData["_id"]==nil {
		id = bson.NewObjectId()
		insertData["_id"] =id
		fmt.Println("没有ID------------新建ID:",id)
	}else{
		fmt.Println("有ID+++++++++++")
		id = insertData["_id"]
		fmt.Println("bson.NewObjectId() ID:",id)
	}
	err:=mgotest4.Insert(insertData)
	fmt.Println(err)


	selector:=bson.M{"account":"admin"}

	var data []interface{}
	mgotest4.Find(selector).All(&data)
	fmt.Println("Remove: ",data)
	mgotest4.Remove(selector)
	mgotest4.Insert(data...)
	//mgotest4.RemoveId(id)
	updateData :=bson.M{"$set": bson.M{"_id":bson.ObjectIdHex("77777"),"account":"change","balance": "加ID更新后"}}
	mgotest4.UpdateAll(selector,updateData)
	for _,d :=range data{
		fmt.Println(d.(bson.M)["_id"])
		mgotest4.UpdateId(d.(bson.M)["_id"],d)
	}

}
