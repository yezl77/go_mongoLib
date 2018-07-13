package dber

import (
	"gopkg.in/mgo.v2"
)



type MongoDatabase struct {
	AbstractMongoBase
	db       string
	database *mgo.Database
}

func (mdb *MongoDatabase) Col(col string) (mco IColOperator, err error) {
	if mdb.isAvailable{
		collection := mdb.database.C(col)
		ico := &MongoColOperator{
			collection:  collection,
			consistance: false,
			retryNumber: 5,
			retryTime:   1,
			AbstractMongoBase: AbstractMongoBase{
				conn:               mdb.conn,
				isAvailable:        true,
				status:             mdb.status,
				parent:             mdb,
				name:               col,
				subObserver: SubObserver{
					m: make(map[string]IMongoBase)},
			},
		}
		mdb.subObserver.m[col] = ico
		mco = ico
	}else{
	err = DBClosedError
	}
	return
}

func (mdb *MongoDatabase) DropDatabase() error {
	err := mdb.database.DropDatabase()
	mdb.onError(err)
	return FormatError(err)
}

func (mdb *MongoDatabase) Close()  { //先关掉下面的子类，再通知上层响应自己删掉
	mdb.close()
	mdb.closeChild()
	mdb.onChildClose(mdb)
}


