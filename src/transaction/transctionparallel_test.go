package transaction

import (
	"testing"
	"fmt"
	"math/rand"
	"time"
	"github.com/pkg/errors"
)

// 事务任务
type testTransactionParallel struct {
	TransctionOps

	id int
	success bool
}



func (this *testTransactionParallel) Commit() error {
	//Icommand . DO
	if this.success {
		fmt.Printf("Success:%v \n", this.id)
		return nil
	} else {
		return errors.New(fmt.Sprintf("fail %v", this.id))
	}

}

// 类型1回调
type testTransactionParallelCallback1 struct {
	TransactionCallbackBase

	object interface{}
}

func (this *testTransactionParallelCallback1) OnSuccess(transaction ITransaction) {
	if this.object != transaction.GetData() {
		panic(fmt.Sprintf("Success check want:%v now:%v\n", this.object, transaction.GetData()))
	}
	// fmt.Printf("OnSuccess %v\n", this.idSuccess)
}

func (this *testTransactionParallelCallback1) OnError(transaction ITransaction, err error) {
	if this.object != transaction.GetData() {
		panic(fmt.Sprintf("Success check want:%v now:%v\n", this.object, transaction.GetData()))
	}
	// fmt.Printf("OnError %v\n", this.idError)
}

// 类型2回调
type testTransactionParallelCallback2 struct {
	TransactionCallbackBase

	errCount int
}

func (this *testTransactionParallelCallback2) OnError(transaction ITransaction, err error) {
	this.errCount++
}

// Case 1: 批量成功
func Test_TransactionParallel_1(t *testing.T) {
	fmt.Printf("Test_TransactionParallel_1 begin\n")

	// 随机生成次数
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	count := r.Intn(512) + 16
	fmt.Printf("Count:%v\n", count)

	// 插入事务
	transaction := &TransactionParallel{}
	for i := 0; i < count; i++ {
		object := &testTransactionParallel{
			id: i,
			success: true,
		}

		callback := &testTransactionParallelCallback1{
			object: object,
		}

		object.SetData(object)
		object.SetCallback(callback)

		if err := transaction.Add(object); err != nil {
			panic(err)
		}
	}

	// 提交，预期成功
	if err := transaction.Commit(); err != nil {
		panic(err)
	}

	// 提交，预期失败
	if err := transaction.Commit(); err == nil {
		panic(fmt.Sprintf("Multicommit success"))
	}
	fmt.Printf("Test_TransactionParallel_1 ok\n")
}

// Case 2: 偶发失败
func Test_TransactionParallel_2(t *testing.T) {
	fmt.Printf("Test_TransactionParallel_2 begin\n")

	// 随机生成次数
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	totalCount := r.Intn(512) + 16
	successCount := r.Intn(totalCount) + 1
	fmt.Printf("total:%v success:%v\n", totalCount, successCount)

	// 插入事务
	transaction := &TransactionParallel{}

	callback := &testTransactionParallelCallback2{
	}
	for i := 0; i < totalCount; i++ {
		object := &testTransactionParallel{
			id: i,
		}
		if i < successCount {
			object.success = true
		} else {
			object.success = false
		}

		object.SetData(object)
		object.SetCallback(callback)

		if err := transaction.Add(object); err != nil {
			panic(err)
		}
	}

	// 提交，预期成功
	err := transaction.Commit()
	if totalCount == successCount {
		if err != nil {
			panic(err)
		}
	} else {
		if err == nil {
			panic("Success all")
		}
	}
	// 检查错误次数
	if callback.errCount != totalCount {
		panic(fmt.Sprintf("Error check want:%v now:%v\n", totalCount, callback.errCount))
	}

	fmt.Printf("Test_TransactionParallel_2 ok\n")
}
