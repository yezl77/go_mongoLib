package transaction

import (
	"github.com/pkg/errors"
)

type TransctionOps struct {
	isCommit bool
	data     interface{}
	callback ITransactionCallback
}


// 添加任务
func (this *TransctionOps) Add(transaction ITransaction) error {
	return errors.New("Not support")
}

// 每种更新操作都需要重写
// 提交任务
func (this *TransctionOps) Commit() error {
	this.isCommit = true
	return nil
}


// 检查是否已提交
func (this *TransctionOps) IsCommit() bool {
	return this.isCommit
}

// 获取数据
func (this *TransctionOps) GetData() interface{} {
	return this.data
}

// 设置数据
func (this *TransctionOps) SetData(data interface{}) error {
	if !this.isCommit {
		this.data = data
		return nil
	} else {
		return errors.New("Transaction has been committed ")
	}
}

// 获取回调
func (this *TransctionOps) GetCallback() ITransactionCallback {
	return this.callback
}

// 设置回调
func (this *TransctionOps) SetCallback(callback ITransactionCallback) error {
	if !this.isCommit {
		this.callback = callback
		return nil
	} else {
		return errors.New("Transaction has been committed ")
	}
}
