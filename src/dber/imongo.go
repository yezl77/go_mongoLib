package dber

type IConn interface {
	connect(conn IConn) (err error)
	reconnect()
	DB(db string) (IDatabase, error)
	Close()
}

type IDatabase interface {
	Col(col string) (IColOperator, error)
	DropDatabase() error
	Close()
}

type IColOperator interface {
	CommandFactory() ICommandFactory
	Insert(data ...interface{}) error
	Remove(selector interface{}) error
	RemoveId(id interface{}) error
	Update(selector interface{}, data interface{}) error
	UpdateId(id interface{}, data interface{}, upsert bool) error
	FindAll(selector interface{}, data *[]interface{}) error
	FindOne(selector interface{}, data *interface{}) error
	FindId(id interface{}, data *interface{}) error
	FindAndModify(selector interface{}, change interface{}, data *[]interface{}) error
	Count(selector interface{}) (int, error)
	Close()
	MakeIndexes(indexes [][]string) error
	DropCol() error
	doCommand(cmd ICommand) error
}

type IMongoBase interface {
	GetName() string
	GetStatus() State
	Close()
	close()
	onChildClose(conn IMongoBase)
	closeChild()
	removeSub(name string)
	onError(err error) bool //返回true 表示不需要进行错误已经处理，不需要做错误处理如重试Command
}

type ICommandFactory interface {
	CreateInsertCommand(data ... interface{}) ICommand
	CreateRemoveCommand(isId bool, data interface{}) ICommand
	CreateUpdateCommand(isId bool, upsert bool, selectorOrId interface{}, data interface{}) ICommand
	CreateCountCommand(count *int, selector interface{}) ICommand
	NewQuery(selector interface{}, isId bool) IQuery
	NewBulk() IBulkCommand
	CreateMongoMakeIndexCommand(indexes [][]string) ICommand
	CreateDropColCommand() ICommand
}

type IQuery interface {
	Sort(key ...string) IQuery
	Limit(limit int) IQuery
	One(data *interface{}) ICommand
	All(data *[]interface{}) ICommand
	Apply(change interface{}, data *[]interface{}, returnNew bool) ICommand
}

//批量操作
type IBulkCommand interface {
	ICommand
	Unordered()  //设置批量操作为无序
	Insert(docs ...interface{})
	Remove(selectors ...interface{})
	RemoveAll(selectors ...interface{})
	Update(pairs ...interface{})
	UpdateAll(pairs ...interface{})
	Upsert(pairs ...interface{})
	GetBulkErrorCases() []ErrorCase
}

type ICommand interface {
	Do() error
	undo() error
	//GetData()[]interface{}
	////SetData(ids []interface{})
	//GetId()[]interface{}
	//SetId(ids []interface{})
}
