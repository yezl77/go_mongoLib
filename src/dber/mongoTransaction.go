package dber

import (
	. "transaction"
)

type MongoTransaction struct {
	command ICommand
	TransctionOps
	id      int //TODO
}

// 重写提交任务
func (this *MongoTransaction) Commit() (err error) {
	err = this.command.Do()
	if err == nil {
		this.TransctionOps.Commit()
	}
	return
}


// 操作回调
type MongoTransactionCallback struct {
	TransactionCallbackBase
	command  ICommand
}



func (this *MongoTransactionCallback) OnError(transaction ITransaction, err error) {
	if transaction.IsCommit()&&err!=nil{
		this.command.undo()
	}
}




