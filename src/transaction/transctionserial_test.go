package transaction

import (
	"testing"
	"fmt"
	"math/rand"
	"time"
	"github.com/pkg/errors"
)

// 事务任务
type testTransactionSerial struct {
	TransctionOps
	id int
	success bool
}

func (this *testTransactionSerial) Commit() error {
	if this.success {
		fmt.Printf("Success:%v \n", this.id)
		return nil
	} else {
		return errors.New(fmt.Sprintf("fail %v", this.id))
	}
}

// 类型1回调
type testTransactionSerialCallback1 struct {
	TransactionCallbackBase

	object interface{}
}

func (this *testTransactionSerialCallback1) OnSuccess(transaction ITransaction) {
	if this.object != transaction.GetData() {
		panic(fmt.Sprintf("Success check want:%v now:%v\n", this.object, transaction.GetData()))
	}

	//fmt.Printf("OnSuccess %v\n", this.idSuccess)
}

func (this *testTransactionSerialCallback1) OnError(transaction ITransaction, err error) {
	if this.object != transaction.GetData() {
		panic(fmt.Sprintf("Success check want:%v now:%v\n", this.object, transaction.GetData()))
	}
	//fmt.Printf("OnError %v\n", this.idError)
}

// 类型2回调
type testTransactionSerialCallback2 struct {
	TransactionCallbackBase
	errCount int
}

func (this *testTransactionSerialCallback2) OnError(transaction ITransaction, err error) {
	this.errCount++
}

// Case 1: 批量成功
func Test_TransactionSerial_1(t *testing.T) {
	fmt.Printf("Test_TransactionSerial_1 begin\n")

	// 随机生成次数
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	count := r.Intn(512) + 16
	fmt.Printf("Count:%v\n", count)

	// 插入事务
	transaction := &TransactionSerial{}
	for i := 0; i < count; i++ {
		object := &testTransactionSerial{
			id: i,
			success: true,
		}

		callback := &testTransactionSerialCallback1{
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
	fmt.Printf("Test_TransactionSerial_1 ok\n")
}

// Case 2: 偶发失败
func Test_TransactionSerial_2(t *testing.T) {
	fmt.Printf("Test_TransactionSerial_2 begin\n")

	// 随机生成次数
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	totalCount := r.Intn(512) + 16
	successCount := r.Intn(totalCount) + 1
	fmt.Printf("total:%v success:%v\n", totalCount, successCount)

	// 插入事务
	transaction := &TransactionSerial{}

	callback := &testTransactionSerialCallback2{
	}
	for i := 0; i < totalCount; i++ {
		object := &testTransactionSerial{
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
	if callback.errCount != successCount+1 {
		panic(fmt.Sprintf("Error check want:%v now:%v\n", successCount+1, callback.errCount))
	}

	fmt.Printf("Test_TransactionSerial_2 ok\n")
}

// Case 3: 串行+并行
func Test_TransactionSerial_3(t *testing.T) {
	fmt.Printf("Test_TransactionSerial_3 begin\n")

	// 随机生成次数
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	count := r.Intn(512) + 16
	fmt.Printf("Count:%v\n", count)

	// 插入事务
	transaction := &TransactionSerial{}
	for i := 0; i < count; i++ {
		if  i == 3 {
			subTransaction := &TransactionParallel{}
			for j := 1; j < count+1; j++ {
				object := &testTransactionSerial{
					id:      -1 * j,
					success: true,
				}

				callback := &testTransactionSerialCallback1{
					object: object,
				}

				object.SetData(object)
				object.SetCallback(callback)

				if err := subTransaction.Add(object); err != nil {
					panic(err)
				}
			}

			if err := transaction.Add(subTransaction); err != nil {
				panic(err)
			}
		} else {
			object := &testTransactionSerial{
				id:      i,
				success: true,
			}

			callback := &testTransactionSerialCallback1{
				object: object,
			}

			object.SetData(object)
			object.SetCallback(callback)

			if err := transaction.Add(object); err != nil {
				panic(err)
			}
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

	fmt.Printf("Test_TransactionSerial_3 ok\n")
}

// Case 4: 串行+并行(随机失败)
func Test_TransactionSerial_4(t *testing.T) {
	fmt.Printf("Test_TransactionSerial_4 begin\n")

	// 随机生成次数
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	totalCount := r.Intn(512) + 16
	successCount := r.Intn(totalCount) + 1
	fmt.Printf("total:%v success:%v\n", totalCount, successCount)

	// 插入事务
	transaction := &TransactionSerial{}

	callback := &testTransactionSerialCallback2{
	}

	count := successCount
	for i := 0; i < totalCount; i++ {
		if  i == 0 {
			subTransaction := &TransactionParallel{}
			subTransaction.SetCallback(callback)

			for j := 0; j < totalCount; j++ {
				object := &testTransactionSerial{
					id:      -1 * (j+1),
					success: true,
				}

				if count > 0 {
					object.success = true
					count--
				} else {
					object.success = false
				}

				object.SetData(object)
				object.SetCallback(callback)

				if err := subTransaction.Add(object); err != nil {
					panic(err)
				}
			}

			if err := transaction.Add(subTransaction); err != nil {
				panic(err)
			}
		} else {
			object := &testTransactionSerial{
				id: i,
			}

			if count > 0 {
				object.success = true
				count--
			} else {
				object.success = false
			}

			object.SetData(object)
			object.SetCallback(callback)

			if err := transaction.Add(object); err != nil {
				panic(err)
			}
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
	if callback.errCount != totalCount+1 {
		panic(fmt.Sprintf("Error check want:%v now:%v\n", totalCount+1, callback.errCount))
	}

	fmt.Printf("Test_TransactionSerial_4 ok\n")
}
