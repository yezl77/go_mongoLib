// Code generated by MockGen. DO NOT EDIT.
// Source: dber\IMongo.go

// Package dber is a generated GoMock package.
package dber

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockIConn is a mock of IConn interface
type MockIConn struct {
	ctrl     *gomock.Controller
	recorder *MockIConnMockRecorder
}

// MockIConnMockRecorder is the mock recorder for MockIConn
type MockIConnMockRecorder struct {
	mock *MockIConn
}

// NewMockIConn creates a new mock instance
func NewMockIConn(ctrl *gomock.Controller) *MockIConn {
	mock := &MockIConn{ctrl: ctrl}
	mock.recorder = &MockIConnMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIConn) EXPECT() *MockIConnMockRecorder {
	return m.recorder
}

// connect mocks base method
func (m *MockIConn) connect(conn IConn) error {
	ret := m.ctrl.Call(m, "connect", conn)
	ret0, _ := ret[0].(error)
	return ret0
}

// connect indicates an expected call of connect
func (mr *MockIConnMockRecorder) connect(conn interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "connect", reflect.TypeOf((*MockIConn)(nil).connect), conn)
}

// reconnect mocks base method
func (m *MockIConn) reconnect() {
	m.ctrl.Call(m, "reconnect")
}

// reconnect indicates an expected call of reconnect
func (mr *MockIConnMockRecorder) reconnect() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "reconnect", reflect.TypeOf((*MockIConn)(nil).reconnect))
}

// DB mocks base method
func (m *MockIConn) DB(db string) (IDatabase, error) {
	ret := m.ctrl.Call(m, "DB", db)
	ret0, _ := ret[0].(IDatabase)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DB indicates an expected call of DB
func (mr *MockIConnMockRecorder) DB(db interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DB", reflect.TypeOf((*MockIConn)(nil).DB), db)
}

// Close mocks base method
func (m *MockIConn) Close() {
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close
func (mr *MockIConnMockRecorder) Close() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockIConn)(nil).Close))
}

// MockIDatabase is a mock of IDatabase interface
type MockIDatabase struct {
	ctrl     *gomock.Controller
	recorder *MockIDatabaseMockRecorder
}

// MockIDatabaseMockRecorder is the mock recorder for MockIDatabase
type MockIDatabaseMockRecorder struct {
	mock *MockIDatabase
}

// NewMockIDatabase creates a new mock instance
func NewMockIDatabase(ctrl *gomock.Controller) *MockIDatabase {
	mock := &MockIDatabase{ctrl: ctrl}
	mock.recorder = &MockIDatabaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIDatabase) EXPECT() *MockIDatabaseMockRecorder {
	return m.recorder
}

// Col mocks base method
func (m *MockIDatabase) Col(col string) (IColOperator, error) {
	ret := m.ctrl.Call(m, "Col", col)
	ret0, _ := ret[0].(IColOperator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Col indicates an expected call of Col
func (mr *MockIDatabaseMockRecorder) Col(col interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Col", reflect.TypeOf((*MockIDatabase)(nil).Col), col)
}

// DropDatabase mocks base method
func (m *MockIDatabase) DropDatabase() error {
	ret := m.ctrl.Call(m, "DropDatabase")
	ret0, _ := ret[0].(error)
	return ret0
}

// DropDatabase indicates an expected call of DropDatabase
func (mr *MockIDatabaseMockRecorder) DropDatabase() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DropDatabase", reflect.TypeOf((*MockIDatabase)(nil).DropDatabase))
}

// Close mocks base method
func (m *MockIDatabase) Close() {
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close
func (mr *MockIDatabaseMockRecorder) Close() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockIDatabase)(nil).Close))
}

// MockIColOperator is a mock of IColOperator interface
type MockIColOperator struct {
	ctrl     *gomock.Controller
	recorder *MockIColOperatorMockRecorder
}

// MockIColOperatorMockRecorder is the mock recorder for MockIColOperator
type MockIColOperatorMockRecorder struct {
	mock *MockIColOperator
}

// NewMockIColOperator creates a new mock instance
func NewMockIColOperator(ctrl *gomock.Controller) *MockIColOperator {
	mock := &MockIColOperator{ctrl: ctrl}
	mock.recorder = &MockIColOperatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIColOperator) EXPECT() *MockIColOperatorMockRecorder {
	return m.recorder
}

// commandFactory mocks base method
func (m *MockIColOperator) commandFactory() ICommandFactory {
	ret := m.ctrl.Call(m, "commandFactory")
	ret0, _ := ret[0].(ICommandFactory)
	return ret0
}

// commandFactory indicates an expected call of commandFactory
func (mr *MockIColOperatorMockRecorder) commandFactory() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "commandFactory", reflect.TypeOf((*MockIColOperator)(nil).commandFactory))
}

// Insert mocks base method
func (m *MockIColOperator) Insert(data ...interface{}) error {
	varargs := []interface{}{}
	for _, a := range data {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Insert", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert
func (mr *MockIColOperatorMockRecorder) Insert(data ...interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockIColOperator)(nil).Insert), data...)
}

// Remove mocks base method
func (m *MockIColOperator) Remove(selector interface{}) error {
	ret := m.ctrl.Call(m, "Remove", selector)
	ret0, _ := ret[0].(error)
	return ret0
}

// Remove indicates an expected call of Remove
func (mr *MockIColOperatorMockRecorder) Remove(selector interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockIColOperator)(nil).Remove), selector)
}

// RemoveId mocks base method
func (m *MockIColOperator) RemoveId(id interface{}) error {
	ret := m.ctrl.Call(m, "RemoveId", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveId indicates an expected call of RemoveId
func (mr *MockIColOperatorMockRecorder) RemoveId(id interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveId", reflect.TypeOf((*MockIColOperator)(nil).RemoveId), id)
}

// Update mocks base method
func (m *MockIColOperator) Update(selector, data interface{}) error {
	ret := m.ctrl.Call(m, "Update", selector, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockIColOperatorMockRecorder) Update(selector, data interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockIColOperator)(nil).Update), selector, data)
}

// UpdateId mocks base method
func (m *MockIColOperator) UpdateId(id, data interface{}, upsert bool) error {
	ret := m.ctrl.Call(m, "UpdateId", id, data, upsert)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateId indicates an expected call of UpdateId
func (mr *MockIColOperatorMockRecorder) UpdateId(id, data, upsert interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateId", reflect.TypeOf((*MockIColOperator)(nil).UpdateId), id, data, upsert)
}

// FindAll mocks base method
func (m *MockIColOperator) FindAll(selector interface{}, data *[]interface{}) error {
	ret := m.ctrl.Call(m, "FindAll", selector, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// FindAll indicates an expected call of FindAll
func (mr *MockIColOperatorMockRecorder) FindAll(selector, data interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockIColOperator)(nil).FindAll), selector, data)
}

// FindOne mocks base method
func (m *MockIColOperator) FindOne(selector interface{}, data *interface{}) error {
	ret := m.ctrl.Call(m, "FindOne", selector, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// FindOne indicates an expected call of FindOne
func (mr *MockIColOperatorMockRecorder) FindOne(selector, data interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOne", reflect.TypeOf((*MockIColOperator)(nil).FindOne), selector, data)
}

// FindId mocks base method
func (m *MockIColOperator) FindId(id interface{}, data *interface{}) error {
	ret := m.ctrl.Call(m, "FindId", id, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// FindId indicates an expected call of FindId
func (mr *MockIColOperatorMockRecorder) FindId(id, data interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindId", reflect.TypeOf((*MockIColOperator)(nil).FindId), id, data)
}

// FindAndModify mocks base method
func (m *MockIColOperator) FindAndModify(selector, change interface{}, data *[]interface{}) error {
	ret := m.ctrl.Call(m, "FindAndModify", selector, change, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// FindAndModify indicates an expected call of FindAndModify
func (mr *MockIColOperatorMockRecorder) FindAndModify(selector, change, data interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAndModify", reflect.TypeOf((*MockIColOperator)(nil).FindAndModify), selector, change, data)
}

// Count mocks base method
func (m *MockIColOperator) Count(selector interface{}) (int, error) {
	ret := m.ctrl.Call(m, "Count", selector)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count
func (mr *MockIColOperatorMockRecorder) Count(selector interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockIColOperator)(nil).Count), selector)
}

// Close mocks base method
func (m *MockIColOperator) Close() {
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close
func (mr *MockIColOperatorMockRecorder) Close() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockIColOperator)(nil).Close))
}

// MakeIndexes mocks base method
func (m *MockIColOperator) MakeIndexes(indexes [][]string) error {
	ret := m.ctrl.Call(m, "MakeIndexes", indexes)
	ret0, _ := ret[0].(error)
	return ret0
}

// MakeIndexes indicates an expected call of MakeIndexes
func (mr *MockIColOperatorMockRecorder) MakeIndexes(indexes interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MakeIndexes", reflect.TypeOf((*MockIColOperator)(nil).MakeIndexes), indexes)
}

// DropCol mocks base method
func (m *MockIColOperator) DropCol() error {
	ret := m.ctrl.Call(m, "DropCol")
	ret0, _ := ret[0].(error)
	return ret0
}

// DropCol indicates an expected call of DropCol
func (mr *MockIColOperatorMockRecorder) DropCol() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DropCol", reflect.TypeOf((*MockIColOperator)(nil).DropCol))
}

// doCommand mocks base method
func (m *MockIColOperator) doCommand(cmd ICommand, consistance bool) error {
	ret := m.ctrl.Call(m, "doCommand", cmd, consistance)
	ret0, _ := ret[0].(error)
	return ret0
}

// doCommand indicates an expected call of doCommand
func (mr *MockIColOperatorMockRecorder) doCommand(cmd, consistance interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "doCommand", reflect.TypeOf((*MockIColOperator)(nil).doCommand), cmd, consistance)
}

// MockIMongoBase is a mock of IMongoBase interface
type MockIMongoBase struct {
	ctrl     *gomock.Controller
	recorder *MockIMongoBaseMockRecorder
}

// MockIMongoBaseMockRecorder is the mock recorder for MockIMongoBase
type MockIMongoBaseMockRecorder struct {
	mock *MockIMongoBase
}

// NewMockIMongoBase creates a new mock instance
func NewMockIMongoBase(ctrl *gomock.Controller) *MockIMongoBase {
	mock := &MockIMongoBase{ctrl: ctrl}
	mock.recorder = &MockIMongoBaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIMongoBase) EXPECT() *MockIMongoBaseMockRecorder {
	return m.recorder
}

// GetName mocks base method
func (m *MockIMongoBase) GetName() string {
	ret := m.ctrl.Call(m, "GetName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetName indicates an expected call of GetName
func (mr *MockIMongoBaseMockRecorder) GetName() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetName", reflect.TypeOf((*MockIMongoBase)(nil).GetName))
}

// GetStatus mocks base method
func (m *MockIMongoBase) GetStatus() State {
	ret := m.ctrl.Call(m, "GetStatus")
	ret0, _ := ret[0].(State)
	return ret0
}

// GetStatus indicates an expected call of GetStatus
func (mr *MockIMongoBaseMockRecorder) GetStatus() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStatus", reflect.TypeOf((*MockIMongoBase)(nil).GetStatus))
}

// Close mocks base method
func (m *MockIMongoBase) Close() {
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close
func (mr *MockIMongoBaseMockRecorder) Close() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockIMongoBase)(nil).Close))
}

// close mocks base method
func (m *MockIMongoBase) close() {
	m.ctrl.Call(m, "close")
}

// close indicates an expected call of close
func (mr *MockIMongoBaseMockRecorder) close() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "close", reflect.TypeOf((*MockIMongoBase)(nil).close))
}

// onChildClose mocks base method
func (m *MockIMongoBase) onChildClose(conn IMongoBase) {
	m.ctrl.Call(m, "onChildClose", conn)
}

// onChildClose indicates an expected call of onChildClose
func (mr *MockIMongoBaseMockRecorder) onChildClose(conn interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "onChildClose", reflect.TypeOf((*MockIMongoBase)(nil).onChildClose), conn)
}

// closeChild mocks base method
func (m *MockIMongoBase) closeChild() {
	m.ctrl.Call(m, "closeChild")
}

// closeChild indicates an expected call of closeChild
func (mr *MockIMongoBaseMockRecorder) closeChild() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "closeChild", reflect.TypeOf((*MockIMongoBase)(nil).closeChild))
}

// removeSub mocks base method
func (m *MockIMongoBase) removeSub(name string) {
	m.ctrl.Call(m, "removeSub", name)
}

// removeSub indicates an expected call of removeSub
func (mr *MockIMongoBaseMockRecorder) removeSub(name interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "removeSub", reflect.TypeOf((*MockIMongoBase)(nil).removeSub), name)
}

// onError mocks base method
func (m *MockIMongoBase) onError(err error) bool {
	ret := m.ctrl.Call(m, "onError", err)
	ret0, _ := ret[0].(bool)
	return ret0
}

// onError indicates an expected call of onError
func (mr *MockIMongoBaseMockRecorder) onError(err interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "onError", reflect.TypeOf((*MockIMongoBase)(nil).onError), err)
}

// MockICommandFactory is a mock of ICommandFactory interface
type MockICommandFactory struct {
	ctrl     *gomock.Controller
	recorder *MockICommandFactoryMockRecorder
}

// MockICommandFactoryMockRecorder is the mock recorder for MockICommandFactory
type MockICommandFactoryMockRecorder struct {
	mock *MockICommandFactory
}

// NewMockICommandFactory creates a new mock instance
func NewMockICommandFactory(ctrl *gomock.Controller) *MockICommandFactory {
	mock := &MockICommandFactory{ctrl: ctrl}
	mock.recorder = &MockICommandFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockICommandFactory) EXPECT() *MockICommandFactoryMockRecorder {
	return m.recorder
}

// CreateInsertCommand mocks base method
func (m *MockICommandFactory) CreateInsertCommand(data ...interface{}) ICommand {
	varargs := []interface{}{}
	for _, a := range data {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateInsertCommand", varargs...)
	ret0, _ := ret[0].(ICommand)
	return ret0
}

// CreateInsertCommand indicates an expected call of CreateInsertCommand
func (mr *MockICommandFactoryMockRecorder) CreateInsertCommand(data ...interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateInsertCommand", reflect.TypeOf((*MockICommandFactory)(nil).CreateInsertCommand), data...)
}

// CreateRemoveCommand mocks base method
func (m *MockICommandFactory) CreateRemoveCommand(isId bool, data interface{}) ICommand {
	ret := m.ctrl.Call(m, "CreateRemoveCommand", isId, data)
	ret0, _ := ret[0].(ICommand)
	return ret0
}

// CreateRemoveCommand indicates an expected call of CreateRemoveCommand
func (mr *MockICommandFactoryMockRecorder) CreateRemoveCommand(isId, data interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRemoveCommand", reflect.TypeOf((*MockICommandFactory)(nil).CreateRemoveCommand), isId, data)
}

// CreateUpdateCommand mocks base method
func (m *MockICommandFactory) CreateUpdateCommand(isId, upsert bool, selectorOrId, data interface{}) ICommand {
	ret := m.ctrl.Call(m, "CreateUpdateCommand", isId, upsert, selectorOrId, data)
	ret0, _ := ret[0].(ICommand)
	return ret0
}

// CreateUpdateCommand indicates an expected call of CreateUpdateCommand
func (mr *MockICommandFactoryMockRecorder) CreateUpdateCommand(isId, upsert, selectorOrId, data interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUpdateCommand", reflect.TypeOf((*MockICommandFactory)(nil).CreateUpdateCommand), isId, upsert, selectorOrId, data)
}

// CreateCountCommand mocks base method
func (m *MockICommandFactory) CreateCountCommand(count *int, selector interface{}) ICommand {
	ret := m.ctrl.Call(m, "CreateCountCommand", count, selector)
	ret0, _ := ret[0].(ICommand)
	return ret0
}

// CreateCountCommand indicates an expected call of CreateCountCommand
func (mr *MockICommandFactoryMockRecorder) CreateCountCommand(count, selector interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCountCommand", reflect.TypeOf((*MockICommandFactory)(nil).CreateCountCommand), count, selector)
}

// NewQuery mocks base method
func (m *MockICommandFactory) NewQuery(selector interface{}, isId bool) IQuery {
	ret := m.ctrl.Call(m, "NewQuery", selector, isId)
	ret0, _ := ret[0].(IQuery)
	return ret0
}

// NewQuery indicates an expected call of NewQuery
func (mr *MockICommandFactoryMockRecorder) NewQuery(selector, isId interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewQuery", reflect.TypeOf((*MockICommandFactory)(nil).NewQuery), selector, isId)
}

// CreateMongoMakeIndexCommand mocks base method
func (m *MockICommandFactory) CreateMongoMakeIndexCommand(indexes [][]string) ICommand {
	ret := m.ctrl.Call(m, "CreateMongoMakeIndexCommand", indexes)
	ret0, _ := ret[0].(ICommand)
	return ret0
}

// CreateMongoMakeIndexCommand indicates an expected call of CreateMongoMakeIndexCommand
func (mr *MockICommandFactoryMockRecorder) CreateMongoMakeIndexCommand(indexes interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMongoMakeIndexCommand", reflect.TypeOf((*MockICommandFactory)(nil).CreateMongoMakeIndexCommand), indexes)
}

// CreateDropColCommand mocks base method
func (m *MockICommandFactory) CreateDropColCommand() ICommand {
	ret := m.ctrl.Call(m, "CreateDropColCommand")
	ret0, _ := ret[0].(ICommand)
	return ret0
}

// CreateDropColCommand indicates an expected call of CreateDropColCommand
func (mr *MockICommandFactoryMockRecorder) CreateDropColCommand() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDropColCommand", reflect.TypeOf((*MockICommandFactory)(nil).CreateDropColCommand))
}

// MockIQuery is a mock of IQuery interface
type MockIQuery struct {
	ctrl     *gomock.Controller
	recorder *MockIQueryMockRecorder
}

// MockIQueryMockRecorder is the mock recorder for MockIQuery
type MockIQueryMockRecorder struct {
	mock *MockIQuery
}

// NewMockIQuery creates a new mock instance
func NewMockIQuery(ctrl *gomock.Controller) *MockIQuery {
	mock := &MockIQuery{ctrl: ctrl}
	mock.recorder = &MockIQueryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIQuery) EXPECT() *MockIQueryMockRecorder {
	return m.recorder
}

// Sort mocks base method
func (m *MockIQuery) Sort(key ...string) IQuery {
	varargs := []interface{}{}
	for _, a := range key {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Sort", varargs...)
	ret0, _ := ret[0].(IQuery)
	return ret0
}

// Sort indicates an expected call of Sort
func (mr *MockIQueryMockRecorder) Sort(key ...interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sort", reflect.TypeOf((*MockIQuery)(nil).Sort), key...)
}

// Limit mocks base method
func (m *MockIQuery) Limit(limit int) IQuery {
	ret := m.ctrl.Call(m, "Limit", limit)
	ret0, _ := ret[0].(IQuery)
	return ret0
}

// Limit indicates an expected call of Limit
func (mr *MockIQueryMockRecorder) Limit(limit interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Limit", reflect.TypeOf((*MockIQuery)(nil).Limit), limit)
}

// One mocks base method
func (m *MockIQuery) One(data *interface{}) ICommand {
	ret := m.ctrl.Call(m, "One", data)
	ret0, _ := ret[0].(ICommand)
	return ret0
}

// One indicates an expected call of One
func (mr *MockIQueryMockRecorder) One(data interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "One", reflect.TypeOf((*MockIQuery)(nil).One), data)
}

// All mocks base method
func (m *MockIQuery) All(data *[]interface{}) ICommand {
	ret := m.ctrl.Call(m, "All", data)
	ret0, _ := ret[0].(ICommand)
	return ret0
}

// All indicates an expected call of All
func (mr *MockIQueryMockRecorder) All(data interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "All", reflect.TypeOf((*MockIQuery)(nil).All), data)
}

// Apply mocks base method
func (m *MockIQuery) Apply(change interface{}, data *[]interface{}, returnNew bool) ICommand {
	ret := m.ctrl.Call(m, "Apply", change, data, returnNew)
	ret0, _ := ret[0].(ICommand)
	return ret0
}

// Apply indicates an expected call of Apply
func (mr *MockIQueryMockRecorder) Apply(change, data, returnNew interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Apply", reflect.TypeOf((*MockIQuery)(nil).Apply), change, data, returnNew)
}

// MockICommand is a mock of ICommand interface
type MockICommand struct {
	ctrl     *gomock.Controller
	recorder *MockICommandMockRecorder
}

// MockICommandMockRecorder is the mock recorder for MockICommand
type MockICommandMockRecorder struct {
	mock *MockICommand
}

// NewMockICommand creates a new mock instance
func NewMockICommand(ctrl *gomock.Controller) *MockICommand {
	mock := &MockICommand{ctrl: ctrl}
	mock.recorder = &MockICommandMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockICommand) EXPECT() *MockICommandMockRecorder {
	return m.recorder
}

// Do mocks base method
func (m *MockICommand) Do(consistance bool) error {
	ret := m.ctrl.Call(m, "Do", consistance)
	ret0, _ := ret[0].(error)
	return ret0
}

// Do indicates an expected call of Do
func (mr *MockICommandMockRecorder) Do(consistance interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Do", reflect.TypeOf((*MockICommand)(nil).Do), consistance)
}
