package dber

import (
	"errors"
	"net"
	"gopkg.in/mgo.v2"
	"fmt"
)

//网络错误
type DBError struct {
	Code    int
	Message string
}

var (
	SessionClosedError = errors.New("Session already closed, cannot do anything ")
	ConnectClosedError = errors.New("Connect already closed, cannot do anything ")
	DBClosedError      = errors.New("DataBase already closed, cannot do anything ")
	ColClosedError     = errors.New("Collection already closed, cannot do anything ")
	NetError           = errors.New("Net temporary error，start reconnect... ")
	NotReachError      = errors.New("No reachable servers，start reconnect... ")
	UnknownPanic       = errors.New("Error:Unknown panic ")
	NotFoundError      = errors.New("Not Found ") //找不到文档
	LastError          = errors.New("Last Error ")
)

func FormatError(err error) error {
	if err != nil {
		_, ok := err.(*net.OpError)
		v, ok2 := err.(*mgo.BulkError)
		_, ok3 := err.(*mgo.LastError)
		v4, ok4 := err.(*mgo.QueryError)
		if ok {
			fmt.Println("net temporary error，start reconnect...")
			return NetError
		} else if ok2 {
			return errors.New(v.Error())
		} else if err.Error() == "no reachable servers" {
			return NotReachError
		} else if err.Error() == "Session already closed" {
			return SessionClosedError
		} else if err == mgo.ErrNotFound { //not found 不需要重试
			return NotFoundError
		} else if ok3 {
			return LastError
		} else if ok4 {
			return errors.New("Query Error:" + v4.Error())
		} else {
			//fmt.Println(reflect.TypeOf(err), "other error, no solve")
			return errors.New(err.Error())
		}

	}
	return nil
}

//type NetError struct {
//	errorInfo string
// }
//
//func (ne *NetError)GetError() string{
//	return ne.errorInfo
//}
//
//type QueryError struct {
//	errorInfo string
//}
//func (ne *QueryError)GetError() string{
//	return ne.errorInfo
//}
//
//type OtherError struct {
//	errorInfo string
//}
//func (ne *OtherError)GetError() string{
//	return ne.errorInfo
//}
//
//type ServerError struct{
//	errorInfo string
//}
//func (ne *ServerError)GetError() string{
//	return ne.errorInfo
//}
//type ClientError struct {
//	errorInfo string
//}
//func (ne *ClientError)GetError() string{
//	return ne.errorInfo
//}
//type OperatorInvalidError struct {
//	errorInfo string
//}
//
//func (ne *OperatorInvalidError)GetError() string{
//	return ne.errorInfo
//}
//
//type DataInvalidError struct {
//	errorInfo string
//}
//func (ne *DataInvalidError)GetError() string{
//	return ne.errorInfo
//}
//
//func FormatError(err error)IDBError{
//	if err != nil {
//		_, ok := err.(*net.OpError)
//		_, ok2 := err.(*mgo.BulkError)
//		if ok || ok2 {
//			fmt.Println()
//			return &NetError{errorInfo:"net temporary error"}
//		} else if err.Error() == "no reachable servers" {
//			return &NetError{errorInfo:"no reachable servers"}
//		} else if err.Error() == "Session already closed" {
//			return &ClientError{errorInfo:"Session already closed" }
//		} else if err == mgo.ErrNotFound { //not found 不需要重试
//			return &QueryError{errorInfo:"not found" }
//		} else {
//			_, ok := err.(*mgo.LastError)
//			if ok {
//				return &OperatorInvalidError{"Update or Upsert LastError"}
//			} else {
//				return &OtherError{"otherError!!"}
//			}
//		}
//	}
//	return nil
//
//}
