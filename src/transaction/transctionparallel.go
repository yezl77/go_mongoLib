package transaction

import (
	"sync"
	"github.com/pkg/errors"
	"wpsgit.kingsoft.net/golibs/utils"
	"fmt"
	"runtime"
)

type TransactionParallel struct {
	sync.Mutex
	transactionArray

	over bool
}



// 添加任务
func (this *TransactionParallel) Add(transaction ITransaction) error {
	this.Lock()
	defer this.Unlock()

	return this.transactionArray.Add(transaction)
}



// 提交任务
func (this *TransactionParallel) Commit() (err error) {
	this.Lock()
	defer this.Unlock()

	if this.IsCommit() {
		return errors.New("Transaction has been commited")
	}

	transactionArray := this.getTransactionArray()

	defer func() {
		for _, ops := range transactionArray {
			this.callOnError(ops, err)
		}
		this.callOnError(this, err)
	}()

	// 生成应答队列
	var count int
	if transactionArray != nil {
		count = len(transactionArray)
	}

	retChan := make(chan error, 8*(count+1))

	// 投递任务
	if err = this.commit(retChan); err != nil {
		return
	}

	// 等待任务结束
	if err = this.waitUntilOver(count, retChan); err != nil {
		return
	}

	// 成功回调
	for _, ops := range transactionArray {
		this.callOnSuccess(ops)
	}
	this.callOnSuccess(this)

	return nil
}

// 内部提交任务
func (this *TransactionParallel) commit(ret chan error) (err error) {
	// 全局提交前回调
	if err = this.callOnCommitBegin(this); err != nil {
		return
	}

	// 并行提交
	transactionArray := this.getTransactionArray()
	if transactionArray != nil {
		// 提交任务
		for _, ops := range transactionArray {
			// 启动协程执行
			go this.runProc(ops, ret)

			// 让出执行权
			runtime.Gosched()
		}
	}

	// 回调基类    ?
	this.TransctionOps.Commit()

	// 全局提交后回调
	if err = this.callOnCommitOver(this); err != nil {
		return
	}

	return nil
}

// 并行执行线程函数
func (this *TransactionParallel) runProc(ops ITransaction, ret chan error) {
	//over 在执行完毕后则执行
	defer libutils.CatchErrorFunc(func(){
		this.over = true
		ret <- errors.New(fmt.Sprintf("Run exception: %v", ops))
	})

	err := this.runOne(ops)
	if err != nil {
		this.over = false
	}

	ret <- err
}

// 执行事务
func (this *TransactionParallel) runOne(ops ITransaction) (err error) {
	// 提交前回调
	if !this.over {
		if err = this.callOnCommitBegin(ops); err != nil {
			return
		}
	}

	// 提交
	if !this.over {
		if err = ops.Commit(); err != nil {
			return
		}
	}

	// 提交后回调
	if !this.over {
		if err = this.callOnCommitOver(ops); err != nil {
			return
		}
	}
	return nil
}

// 等待结果
func (this *TransactionParallel) waitUntilOver(count int, ret chan error) error {
	for i := 0; i < count; i++ {
		err := <- ret
		if err != nil {
			return err
		}
	}

	return nil
}
