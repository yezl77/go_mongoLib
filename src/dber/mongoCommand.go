package dber

import (
	"gopkg.in/mgo.v2"
	"errors"
	"gopkg.in/mgo.v2/bson"
)

type MongoInsertCommand struct {
	consistance bool
	data        []interface{}
	id          []interface{}
	c           *mgo.Collection
}

func (mic *MongoInsertCommand) Do() (err error) {
	defer func() {
		if r := recover(); r != nil {
			switch x := r.(type) {
			case string:
				err = errors.New(x)
			case error:
				err = x
			default:
				err = UnknownPanic
			}
		}
	}()
	err = mic.c.Insert(mic.data...)
	return err
}

func (mic *MongoInsertCommand) undo() (err error) {
	for _, insertId := range mic.id {
		err = mic.c.RemoveId(insertId)
		if err != nil {
			break
		}
	}
	return
}

type MongoRemoveCommand struct {
	consistance  bool
	selectorOrId interface{}
	c            *mgo.Collection
	isId         bool
	oldData      []interface{}
}

func (mrc *MongoRemoveCommand) Do() error {
	if mrc.isId {
		return mrc.c.RemoveId(mrc.selectorOrId)
	}
	_, err := mrc.c.RemoveAll(mrc.selectorOrId)
	return err
}

func (mrc *MongoRemoveCommand) undo() (err error) {
	err = mrc.c.Insert(mrc.oldData...)
	return
}

type MongoUpdateCommand struct {
	consistance     bool
	selectorOrId    interface{}
	data            interface{}
	c               *mgo.Collection
	isId            bool
	upsert          bool
	oldData         []interface{}
	isInsertNewData bool
	insertId        interface{}
	hasNewId        bool
}

func (muc *MongoUpdateCommand) Do() error {
	if muc.isId {
		if muc.upsert {
			_, err := muc.c.UpsertId(muc.selectorOrId, muc.data)
			return err
		}
		return muc.c.UpdateId(muc.selectorOrId, muc.data)
	}
	_, err := muc.c.UpdateAll(muc.selectorOrId, muc.data)
	return err
}

func (muc *MongoUpdateCommand) undo() (err error) {

	if muc.isInsertNewData {
		err = muc.c.RemoveId(muc.insertId)
	} else {
		if muc.hasNewId{
			muc.c.RemoveId(muc.insertId)
		}else{
			for _, d := range muc.oldData {
				err=muc.c.UpdateId(d.(bson.M)["_id"], d)
				if err!=nil{
					break
				}
			}
		}
	}

	return
}

type MongoCountCommand struct {
	consistance bool
	selector    interface{}
	c           *mgo.Collection
	count       *int //count
}

func (mcc *MongoCountCommand) Do() error {
	var err error
	*mcc.count, err = mcc.c.Find(mcc.selector).Count()
	return err
}

func (mcc *MongoCountCommand) undo() (err error) {

	return
}

type MongoQueryOneCommand struct {
	consistance bool
	q           *mgo.Query
	selector    interface{}  //选择
	data        *interface{} //地址
}

func (mqoc *MongoQueryOneCommand) Do() error {
	return mqoc.q.One(mqoc.data)
}

func (mqoc *MongoQueryOneCommand) undo() (err error) {

	return
}

type MongoQueryAllCommand struct {
	consistance bool
	q           *mgo.Query
	data        *[]interface{} //地址
}

func (mqac *MongoQueryAllCommand) Do() error {
	err := mqac.q.All(mqac.data)
	return err
}

func (mqac *MongoQueryAllCommand) undo() error {

	return nil
}

type MongoQueryApplyCommand struct {
	consistance bool
	q           *mgo.Query
	data        *[]interface{} //地址
	change      mgo.Change
}

func (mqayc *MongoQueryApplyCommand) Do() error {
	// _,err:=mqayc.c.Find(mqayc.selector).Apply(mqayc.change,mqayc.data)
	_, err := mqayc.q.Apply(mqayc.change, mqayc.data)
	return err
}

func (mqayc *MongoQueryApplyCommand) undo() error {

	return nil
}

type MongoBulkRunCommand struct {
	consistance bool
	AbstractMongoBase
	conn        IConn
	b           *mgo.Bulk
	err         error
	c           *mgo.Collection
	rollback    *mgo.Bulk


}

func (mb *MongoBulkRunCommand) Do() error {
	_, err := mb.b.Run()

	if err != nil {
		for _, e := range mb.getErrors(err) {
			mb.onError(e)
		}
		_, err = mb.b.Run()
	}
	mb.err = err
	return err
}

func (mb *MongoBulkRunCommand) undo() error {
	mb.rollback.Run()
	return nil
}

func (mb *MongoBulkRunCommand) getErrors(err error) []error {
	if err != nil {
		bulkErr, ok := err.(*mgo.BulkError)
		if !ok {
			// not a bulk error
		} else {
			errs := make([]error, 0, len(bulkErr.Cases()))
			seen := make(map[error]bool)
			for _, ecase := range bulkErr.Cases() {
				msg := ecase.Err
				if !seen[msg] {
					seen[msg] = true
					errs = append(errs, msg)
				}
			}
			return errs
		}
	}
	return nil
}

type ErrorCase struct {
	Index int
	err   error
}

func (mb *MongoBulkRunCommand) GetBulkErrorCases() []ErrorCase {
	if mb.err != nil {
		bulkErr, ok := mb.err.(*mgo.BulkError)
		if !ok {
			// not a bulk error
		} else {
			var errArray []ErrorCase
			for _, errCase := range bulkErr.Cases() {
				//fmt.Println(errCase.Index,errCase.Err)
				if errCase.Err != nil {
					errArray = append(errArray, ErrorCase{errCase.Index, FormatError(errCase.Err)})
				}
			}
			return errArray
		}
	}
	return nil
}

func (mb *MongoBulkRunCommand) Unordered() {
	mb.b.Unordered()
}

func (mb *MongoBulkRunCommand) Insert(docs ...interface{}) {
	//回滚操作
	var insertIds []interface{}
	var insertDatas []interface{}
	for _, insertData := range docs {
		var id interface{}
		if insertData.(bson.M)["_id"] == nil {
			id := bson.NewObjectId()
			insertData.(bson.M)["_id"] =id
		} else {
			id= insertData.(bson.M)["_id"]
		}
		insertIds = append(insertIds, bson.M{"_id":id,})
		insertDatas= append(insertDatas,insertData)
	}
	mb.b.Remove(insertIds...)
	//结束

	mb.b.Insert(insertDatas...)
}



func (mb *MongoBulkRunCommand) Remove(selectors ...interface{}) {
	var oldDatas []interface{}
	for _,selector :=range selectors{
		var oldData []interface{}
		mb.c.Find(selector).One(&oldData)
		oldDatas= append(oldDatas,oldData)
	}
	oldDatas=RemoveRepByMap(oldDatas)
	mb.rollback.Insert(oldDatas)

	mb.b.Remove(selectors...)
}

func RemoveRepByMap(slc []interface{}) []interface{} {
	var result []interface{}
	tempMap := map[interface{}]byte{}  // 存放不重复主键
	for _, e := range slc{
		l := len(tempMap)
		tempMap[e] = 0
		if len(tempMap) != l{  // 加入map后，map长度变化，则元素不重复
			result = append(result, e)
		}
	}
	return result
}

func (mb *MongoBulkRunCommand) RemoveAll(selectors ...interface{}) {
	var oldDatas []interface{}
	for _,selector :=range selectors{
		var oldData []interface{}
		mb.c.Find(selector).All(&oldData)
		oldDatas= append(oldDatas,oldData)
	}
	oldDatas=RemoveRepByMap(oldDatas)
	mb.rollback.Insert(oldDatas)

	mb.b.RemoveAll(selectors...)
}

func (mb *MongoBulkRunCommand) Update(pairs ...interface{}) {




	mb.b.Update(pairs...)
}

func (mb *MongoBulkRunCommand) UpdateAll(pairs ...interface{}) {
	mb.b.UpdateAll(pairs...)
}

func (mb *MongoBulkRunCommand) Upsert(pairs ...interface{}) {
	mb.b.Upsert(pairs...)
}


//
type MongoMakeIndexCommand struct {
	consistance bool
	c           *mgo.Collection
	indexes     [][]string //地址
}

func (mmic *MongoMakeIndexCommand) Do() error {
	for _, index := range mmic.indexes {
		err := mmic.c.EnsureIndexKey(index...)
		if err != nil {
			return err
		}
	}
	return nil
}

func (mmic *MongoMakeIndexCommand) undo() error {

	return nil
}

