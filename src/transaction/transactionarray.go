package transaction

import (
	"github.com/pkg/errors"
)

type transactionArray struct {
	TransctionOps
}


// 添加任务
func (this *transactionArray) Add(transaction ITransaction) error {
	// 插入操作对象
	transactionArray := this.getTransactionArray()
	if transactionArray == nil {
		return errors.New("Create serial array error")
	}
	transactionArray = append(transactionArray, transaction)
	if err := this.setTransactionArray(transactionArray); err != nil {
		return err
	}
	return nil
}

// 获取队列
func (this *transactionArray) getTransactionArray() []ITransaction {
	object := this.GetData()
	if object == nil {
		array := make([]ITransaction, 0, 8)
		if err := this.setTransactionArray(array); err != nil {
			return nil
		}
		return array
	} else {
		return object.([]ITransaction)
	}
}

// 设置队列
func (this *transactionArray) setTransactionArray(array []ITransaction) error {
	return this.SetData(array)
}

// 提交前回调
func (this *transactionArray) callOnCommitBegin(transaction ITransaction) error {
	if callback := transaction.GetCallback(); callback != nil {
		return callback.OnCommitBegin(transaction)
	} else {
		return nil
	}
}

// 提交后回调
func (this *transactionArray) callOnCommitOver(transaction ITransaction) error {
	if callback := transaction.GetCallback(); callback != nil {
		return callback.OnCommitOver(transaction)
	} else {
		return nil
	}
}

// 提交成功回调
func (this *transactionArray) callOnSuccess(transaction ITransaction) {
	if callback := transaction.GetCallback(); callback != nil {
		callback.OnSuccess(transaction)
	}
}

// 提交失败回调
func (this *transactionArray) callOnError(transaction ITransaction, err error) {

	if callback := transaction.GetCallback(); callback != nil {
		callback.OnError(transaction, err)
	}
}
