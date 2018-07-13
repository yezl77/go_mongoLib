package transaction

type ITransaction interface {
	Add(transaction ITransaction) error

	Commit() error
	IsCommit() bool

	GetData() interface{}
	SetData(data interface{}) error

	GetCallback() ITransactionCallback
	SetCallback(ITransactionCallback) error
}

type ITransactionCallback interface {
	OnCommitBegin(transaction ITransaction) error
	OnCommitOver(transaction ITransaction) error

	OnSuccess(transaction ITransaction)
	OnError(transaction ITransaction, err error)
}
