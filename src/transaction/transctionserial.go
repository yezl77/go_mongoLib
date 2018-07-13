package transaction

import (
	"sync"
	"github.com/pkg/errors"
)

type TransactionSerial struct {
	sync.Mutex
	transactionArray
}

// 添加任务
func (this *TransactionSerial) Add(transaction ITransaction) error {
	this.Lock()
	defer this.Unlock()
	return this.transactionArray.Add(transaction)
}

// 提交任务
func (this *TransactionSerial) Commit() (err error) {
	this.Lock()
	defer this.Unlock()

	successCount := -1
	defer func() {
		if err != nil {
			// 失败回调
			transactionArray := this.getTransactionArray()
			if successCount >= 0 && transactionArray != nil {
				var firstId int
				if successCount + 1 >= len(transactionArray) {
					firstId = len(transactionArray) - 1
				} else {
					firstId = successCount + 1
				}

				for i := firstId; i >= 0; i-- {
					ops := transactionArray[i]
					this.callOnError(ops, err)
				}
				this.callOnError(this, err)
			}
		}
	} ()

	if this.IsCommit() {
		return errors.New("Transaction has been commited")
	}

	// 全局提交前回调
	if err = this.callOnCommitBegin(this); err != nil {
		return
	}

	// 串行提交
	transactionArray := this.getTransactionArray()
	if transactionArray != nil {
		// 提交任务
		for _, ops := range transactionArray {
			// 提交前回调
			if err = this.callOnCommitBegin(ops); err != nil {
				return
			}

			// 提交
			if err = ops.Commit(); err != nil {
				return
			}

			// 提交后回调
			if err = this.callOnCommitOver(ops); err != nil {
				return
			}

			successCount++
		}

		// 成功回调
		for _, ops := range transactionArray {
			this.callOnSuccess(ops)
		}
	}

	// 回调基类
	this.TransctionOps.Commit()

	// 全局提交后回调
	if err = this.callOnCommitOver(this); err != nil {
		return
	}

	// 成功回调
	this.callOnSuccess(this)

	return err
}
