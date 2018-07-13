package dber

import "gopkg.in/mgo.v2"



type MongoQuery struct {
	selector interface{}
	q        *mgo.Query
	c        *mgo.Collection
	isId     bool
}

func (mq *MongoQuery) Sort(key ...string) IQuery {
	mq.q = mq.c.Find(mq.selector).Sort(key...)
	return mq
}

func (mq *MongoQuery) Limit(limit int) IQuery {
	mq.q = mq.c.Find(mq.selector).Limit(limit)
	return mq
}

func (mq *MongoQuery) One(data *interface{}) ICommand {
	if mq.isId {
		mq.q = mq.c.FindId(mq.selector)
	} else {
		mq.q = mq.c.Find(mq.selector)
	}
	return &MongoQueryOneCommand{
		data: data,
		q:    mq.q,
	}
}


func (mq *MongoQuery) All(data *[]interface{}) ICommand {
	mq.q = mq.c.Find(mq.selector)
	return &MongoQueryAllCommand{
		data: data,
		q:    mq.q,
	}
}


func (mq *MongoQuery) Apply(change interface{}, data *[]interface{}, returnNew bool) ICommand {
	mq.q = mq.c.Find(mq.selector)
	return &MongoQueryApplyCommand{
		data: data,
		q:    mq.q,
		change: mgo.Change{
			Update:    change,
			Upsert:    true,
			Remove:    false,
			ReturnNew: returnNew,
		},
	}
}




