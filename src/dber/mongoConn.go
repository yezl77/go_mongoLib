package dber

import (
	"gopkg.in/mgo.v2"
	"time"
)

type MongoConn struct {
	AbstractMongoBase
	url                string
	consistance        bool
	session            *mgo.Session
	isUnLimitReconnect bool //出错是否无限重连直到连接成功    默认为true ，false则按照设定重连的次数和时间
	reconnectCount     int  //重连次数            默认尝试重连 10 次
	reconnectTime      int  //尝试重连间隔时间 （单位:秒）  默认每隔1秒重连一次
}

func NewMongoConn(url string, consistance bool) (conn IConn, err error) {
	conn = &MongoConn{
		url:                url,
		consistance:        consistance,
		isUnLimitReconnect: true,
		reconnectCount:     10,
		reconnectTime:      1,
		AbstractMongoBase: AbstractMongoBase{
			isAvailable: true,
			name:        url,
			subObserver: SubObserver{
				m: make(map[string]IMongoBase)},
		},
	}
	err = conn.connect(conn)
	return
}

func (mc *MongoConn) DB(db string) (idb IDatabase, err error) {
	if mc.isAvailable {
		mdb := &MongoDatabase{
			db:       db,
			database: mc.session.DB(db),
			AbstractMongoBase: AbstractMongoBase{
				conn:        mc,
				parent:      mc,
				name:        db,
				isAvailable: true,
				status:      mc.status,
				subObserver: SubObserver{
					m: make(map[string]IMongoBase)},
			},
		}
		mc.subObserver.m[db] = mdb
		idb = mdb
	} else {
		err = ConnectClosedError
	}
	return
}

func (mc *MongoConn) connect(conn IConn) (err error) {
	mc.conn = conn
	mc.session, err = mgo.Dial(mc.url)
	if err == nil {
		mc.setConnected()
	} else {
		mc.setInvalid()
		mc.onError(err)
	}
	return FormatError(err)
}

func (mc *MongoConn) reconnect() {
	if mc.isConnected() {
		reconnectCount := 0
		if mc.isUnLimitReconnect {
			mc.setConnecting()
			for mc.session.Ping() != nil {
				mc.session.Refresh()
				reconnectCount++
			}
			mc.setConnected()
		} else {
			mc.setConnecting()
			for mc.session.Ping() != nil {
				mc.session.Refresh()
				if reconnectCount >= mc.reconnectCount {
					break
				}
				time.Sleep(time.Duration(mc.reconnectTime) * time.Second)
			}
			if mc.session.Ping() != nil {
				mc.setCanNotConnect()
			} else {
				mc.setConnected()
			}
		}
	}
}

func (mc *MongoConn) GetUrl() string {
	return mc.url
}

func (mc *MongoConn) GetSession() *mgo.Session {
	return mc.session
}

func (mc *MongoConn) GetStatus() State {
	return mc.status
}

func (mc *MongoConn) SetIsUnLimitReconnect(isUnLimitReconnect bool) {
	mc.isUnLimitReconnect = isUnLimitReconnect
}

func (mc *MongoConn) SetReconnectCount(reconnectCount int) {
	mc.reconnectCount = reconnectCount
}

func (mc *MongoConn) SetReconnectTime(reconnectTime int) {
	mc.reconnectTime = reconnectTime
}

func (mc *MongoConn) Close() {
	mc.close()
	mc.closeChild()
	defer mc.session.Close()
}
