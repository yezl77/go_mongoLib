package dber

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MongoCommandFactory struct {
	AbstractMongoBase
	mc *mgo.Collection
}

func (mcf *MongoCommandFactory) CreateInsertCommand(data ...interface{}) ICommand {
	//这里需要先得到插入的所有记录的id  开始
	var insertIds []interface{}
	var insertDatas []interface{}
	for _, insertData := range data {
		if insertData.(bson.M)["_id"] == nil {
			id := bson.NewObjectId()
			insertIds = append(insertIds, id)
			insertData.(bson.M)["_id"] =id
		} else {
			insertIds = append(insertIds, insertData.(bson.M)["_id"])
		}
		insertDatas= append(insertDatas,insertData)
	}
	//结束

	return &MongoInsertCommand{
		data: insertDatas,
		c:    mcf.mc,
		id:   insertIds,
	}
}

func (mcf *MongoCommandFactory) CreateRemoveCommand(isId bool, selectorOrId interface{}) ICommand {
	var oldData []interface{}
	if isId {
		mcf.mc.FindId(selectorOrId).All(&oldData)
	} else {
		mcf.mc.Find(selectorOrId).All(&oldData)
	}

	return &MongoRemoveCommand{
		selectorOrId: selectorOrId,
		c:            mcf.mc,
		isId:         isId,
		oldData:      oldData,
	}
}

func (mcf *MongoCommandFactory) CreateUpdateCommand(isId bool, upsert bool, selectorOrId interface{}, data interface{}) ICommand {
	var oldData []interface{}
	var insertId interface{}
	isInsertNewData := false
	hasNewId := false
	if isId {
		mcf.mc.FindId(selectorOrId).All(&oldData)
	} else {
		mcf.mc.Find(selectorOrId).All(&oldData)
	}

	if data.(bson.M)["_id"] != nil {
		hasNewId= true
		insertId = data.(bson.M)["_id"]
	}

	if upsert && len(oldData) == 0 {
		//为空将插入
		isInsertNewData = true
		if data.(bson.M)["_id"] == nil {
			insertId = bson.NewObjectId()
		} else {
			hasNewId= true
			insertId = data.(bson.M)["_id"]
		}
	}

	return &MongoUpdateCommand{
		selectorOrId:    selectorOrId,
		data:            data,
		c:               mcf.mc,
		isId:            isId,
		upsert:          upsert,
		oldData:         oldData,
		isInsertNewData: isInsertNewData,
		insertId:        insertId,
		hasNewId :hasNewId,
	}
}

func (mcf *MongoCommandFactory) CreateCountCommand(count *int, selector interface{}) ICommand {
	return &MongoCountCommand{
		selector: selector,
		c:        mcf.mc,
		count:    count,
	}
}

func (mcf *MongoCommandFactory) NewQuery(selector interface{}, isId bool) IQuery {
	return &MongoQuery{
		c:        mcf.mc,
		selector: selector,
		isId:     isId,
	}
}

func (mcf *MongoCommandFactory) NewBulk() IBulkCommand {
	return &MongoBulkRunCommand{
		AbstractMongoBase: mcf.AbstractMongoBase,
		b:                 mcf.mc.Bulk(),
		rollback: mcf.mc.Bulk(),
		c:mcf.mc,
	}
}

func (mcf *MongoCommandFactory) CreateMongoMakeIndexCommand(indexes [][]string) ICommand {
	return &MongoMakeIndexCommand{
		c:       mcf.mc,
		indexes: indexes,
	}
}

