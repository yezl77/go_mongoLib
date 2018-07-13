package dber

import (
	"sync"
	"net"
	"gopkg.in/mgo.v2"
	"fmt"
)

type SubObserver struct {
	sync.RWMutex //加锁
	m map[string]IMongoBase
}

type State int

const (
	Invalid       State = iota // 初始状态
	Connected                  // 数据库连接正常连接
	Closed                     // 数据库没有连接或者连接关闭状态
	Connecting                 // 数据库连接正在重连
	CanNotConnect              // 无法连接对应数据库
)

type AbstractMongoBase struct {
	sync.RWMutex //加锁
	conn        IConn
	parent      IMongoBase
	subObserver SubObserver
	name        string
	isAvailable bool  //是否可用
	status      State //状态机
}

func (amb *AbstractMongoBase) GetStatus() (s State) {
	return amb.status
}

func (amb *AbstractMongoBase) GetName() string {
	return amb.name
}

func (amb *AbstractMongoBase) removeSub(name string) {
	delete(amb.subObserver.m, name)
}

func (amb *AbstractMongoBase) onChildClose(conn IMongoBase) {
	amb.parent.removeSub(conn.GetName())
}

func (amb *AbstractMongoBase) closeChild() {
	amb.subObserver.Lock()
	defer amb.subObserver.Unlock()

	if len(amb.subObserver.m) != 0 {
		for k, v := range amb.subObserver.m {
			v.Close()
			amb.removeSub(k)
		}
	}
}

func (amb *AbstractMongoBase) close() {
	amb.Lock()
	defer amb.Unlock()

	if amb.isAvailable {
		amb.isAvailable = false
		if amb.status != CanNotConnect { //处于CanNotConnect状态不能转换为Closed
			amb.status = Closed
		}
	}
}

func (amb *AbstractMongoBase) isConnected() bool {
	amb.Lock()
	defer amb.Unlock()
	return amb.status == Connected
}

func (amb *AbstractMongoBase) setInvalid() {
	amb.Lock()
	defer amb.Unlock()
	amb.status = Invalid
}

func (amb *AbstractMongoBase) setConnecting() {
	amb.Lock()
	defer amb.Unlock()
	amb.status = Connecting
}
func (amb *AbstractMongoBase) setConnected() {
	amb.Lock()
	defer amb.Unlock()
	amb.status = Connected
}

func (amb *AbstractMongoBase) setCanNotConnect() {
	amb.Lock()
	defer amb.Unlock()
	amb.status = CanNotConnect
}

func (amb *AbstractMongoBase) Close() {
	amb.close()
	amb.onChildClose(amb)
}

func (amb *AbstractMongoBase) onError(err error) bool {
	//fmt.Println(reflect.TypeOf(err))
	//fmt.Println("测试",err.Error())
	if err != nil {
		_, ok := err.(*net.OpError)
		_, ok2 := err.(*mgo.BulkError)
		if ok || ok2 {
			fmt.Println("net temporary error，start reconnect...")
			amb.conn.reconnect()
			return false
		} else if err.Error() == "no reachable servers" {
			fmt.Println("no reachable servers，start reconnect...")
			amb.conn.reconnect()
			return false
		} else if err.Error() == "Session already closed" {
			fmt.Println(SessionClosedError.Error())
			return true
		} else if err == mgo.ErrNotFound { //not found 不需要重试
			return true
		} else {
			_, ok := err.(*mgo.LastError)
			if ok {
				fmt.Println("*mgo.LastError：not retry")
				return true
			} else {
				//fmt.Println(reflect.TypeOf(err),"other error, no solve")
				return true
			}
		}
	}
	return true
}

//func (amb *AbstractMongoBase) onError(err error) bool {
//	//fmt.Println(reflect.TypeOf(err))
//	//fmt.Println("测试",err.Error())
//	if err != nil {
//		if err == NetError {
//			fmt.Println("net temporary error，start reconnect...")
//			amb.conn.reconnect()
//			return false
//		} else if err == NotReachError {
//			fmt.Println("no reachable servers，start reconnect...")
//			amb.conn.reconnect()
//			return false
//		} else if err == SessionClosedError {
//			fmt.Println("Session already closed")
//			return true
//		} else if err == NotFoundError { //not found 不需要重试
//			return true
//		} else if err == LastError {
//			fmt.Println("*mgo.LastError：not retry")
//			return true
//		} else if err == OtherError {
//			return true
//		} else {
//			fmt.Println(reflect.TypeOf(err), "other error, no solve")
//			return true
//		}
//	}
//	return true
//}
