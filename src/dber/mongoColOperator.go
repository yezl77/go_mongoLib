package dber

import (
	"gopkg.in/mgo.v2"
	"sync"
	"time"
)

type MongoColOperator struct {
	sync.RWMutex //加锁
	AbstractMongoBase
	consistance bool
	collection  *mgo.Collection
	mongoCom    ICommand
	mongoQue    IQuery
	wg          sync.WaitGroup
	mongoFac    ICommandFactory
	count       int
	retryNumber int //重试次数            默认尝试重试 5 次
	retryTime   int //尝试重试间隔时间 （单位:秒）  默认每隔1秒重连一次
}

func (mco *MongoColOperator) CommandFactory() ICommandFactory {
	mco.Lock()
	defer mco.Unlock()

	if mco.mongoFac == nil {
		return &MongoCommandFactory{
			AbstractMongoBase:mco.AbstractMongoBase,
			mc: mco.collection,
		}
	} else {
		return mco.mongoFac
	}
}

func (mco *MongoColOperator) doCommand(cmd ICommand) error {
	mco.wg.Add(1)
	defer mco.wg.Done()
	if mco.isAvailable {
		retryCount := 0
		err := cmd.Do()
		for err != nil {
			if mco.onError(err) {
				break
			}
			if retryCount >= mco.retryNumber {
				break
			}
			err = cmd.Do()

			if err == nil {
				break
			}
			retryCount++
			time.Sleep(time.Duration(mco.retryTime) * time.Second)       //休眠指定的时间后再重试
		}
		return FormatError(err)
	} else {
		return ColClosedError
	}
}

func (mco *MongoColOperator) Insert(data ... interface{}) error {
	mco.mongoCom = mco.CommandFactory().CreateInsertCommand(data...)
	return mco.doCommand(mco.mongoCom)
}

func (mco *MongoColOperator) InsertTest(command ICommand) error {
	return mco.doCommand(command)
}

func (mco *MongoColOperator) Remove(selector interface{}) error {
	mco.mongoCom = mco.CommandFactory().CreateRemoveCommand(false, selector)
	return mco.doCommand(mco.mongoCom)

}
func (mco *MongoColOperator) RemoveId(id interface{}) error {
	mco.mongoCom = mco.CommandFactory().CreateRemoveCommand(true, id)
	return mco.doCommand(mco.mongoCom)

}
func (mco *MongoColOperator) Update(selector interface{}, data interface{}) error {
	mco.mongoCom = mco.CommandFactory().CreateUpdateCommand(false, false, selector, data)
	return mco.doCommand(mco.mongoCom)
}
func (mco *MongoColOperator) UpdateId(id interface{}, data interface{}, upsert bool) error {
	mco.mongoCom = mco.CommandFactory().CreateUpdateCommand(true, upsert, id, data)
	return mco.doCommand(mco.mongoCom)
}

func (mco *MongoColOperator) Count(selector interface{}) (int, error ) {
	mco.mongoCom = mco.CommandFactory().CreateCountCommand(&mco.count, selector)
	return mco.count, mco.doCommand(mco.mongoCom)
}

func (mco *MongoColOperator) FindAll(selector interface{}, data *[]interface{}) error {
	mco.mongoCom = mco.CommandFactory().NewQuery(selector, false).All(data)
	return mco.doCommand(mco.mongoCom)
}
func (mco *MongoColOperator) FindOne(selector interface{}, data *interface{}) error {
	mco.mongoCom = mco.CommandFactory().NewQuery(selector, false).One(data)
	return mco.doCommand(mco.mongoCom)
}
func (mco *MongoColOperator) FindId(id interface{}, data *interface{}) error {
	mco.mongoCom = mco.CommandFactory().NewQuery(id, true).One(data)
	return mco.doCommand(mco.mongoCom)
}
func (mco *MongoColOperator) FindAndModify(selector interface{}, change interface{}, data *[]interface{}) error {
	mco.mongoCom = mco.CommandFactory().NewQuery(selector, false).Apply(change, data, true)
	return mco.doCommand(mco.mongoCom)
}

func (mco *MongoColOperator) MakeIndexes(indexes [][]string) error {
	mco.mongoCom = mco.CommandFactory().CreateMongoMakeIndexCommand(indexes)
	return mco.doCommand(mco.mongoCom)
}

func (mco *MongoColOperator) DropCol() error {
	return mco.collection.DropCollection()
}

func (mco *MongoColOperator) Close() {
	mco.wg.Wait()
	mco.close()
	mco.parent.removeSub(mco.name)
}
