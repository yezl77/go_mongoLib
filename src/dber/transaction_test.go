package dber

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/mgo.v2/bson"
	. "transaction"
)

func TestInsertTransactionParallel_Commit(t *testing.T) {
	Convey("The NewMongoConn  with a url", t, func() {
		ico, _ := NewMongoConn("mongodb://root:yzl17220@120.78.158.127:27017", false)
		db, _ := ico.DB("s_blog")
		c, _ := db.Col("mgotest5")
		//insertData := bson.M{"_id":7,"account":"admin","balance":"加ID更新前" }

		transaction := &TransactionParallel{}
		for i := 0; i < 10; i++ {
			insertData := bson.M{"_id": i, "account": "admin", "balance": "加ID更新前"}
			insertCommand := c.CommandFactory().CreateInsertCommand(insertData)
			callback := &MongoTransactionCallback{
				command: insertCommand,
			}
			object := &MongoTransaction{
				command: insertCommand,
			}
			object.SetCallback(callback)

			if err := transaction.Add(object); err != nil {
				panic(err)
			}
		}

		Convey("TestInsertTransactionParallel_Commit：事务插入，必须没有数据 应该为没有错误", func() {
			err := transaction.Commit()
			So(err, ShouldEqual, nil)
		})
		ico.Close()
	})
}

func TestInsertTransactionSerial_Commit(t *testing.T) {
	Convey("The NewMongoConn  with a url", t, func() {
		ico, _ := NewMongoConn("mongodb://root:yzl17220@120.78.158.127:27017", false)
		db, _ := ico.DB("s_blog")
		c, _ := db.Col("mgotest5")
		//insertData := bson.M{"_id":7,"account":"admin","balance":"加ID更新前" }

		transaction := &TransactionSerial{}
		for i := 546; i < 547; i++ {

			insertData := bson.M{"_id": i, "account": "admin", "balance": "加ID更新前"}
			insertCommand := c.CommandFactory().CreateInsertCommand(insertData)
			callback := &MongoTransactionCallback{
				command:insertCommand,
			}
			object := &MongoTransaction{
				command: insertCommand,
			}
			object.SetCallback(callback)


			if err := transaction.Add(object); err != nil {
				panic(err)
			}
		}
		Convey("TestInsertTransactionParallel_Commit：并行事务插入，必须没有数据 应该为没有错误", func() {
			err := transaction.Commit()
			So(err, ShouldEqual, nil)
		})
		ico.Close()
	})
}

func TestInsertTransactionMutilSerial_Commit(t *testing.T) {
	Convey("The NewMongoConn  with a url", t, func() {
		ico, _ := NewMongoConn("mongodb://root:yzl17220@120.78.158.127:27017", false)
		db, _ := ico.DB("s_blog")
		c, _ := db.Col("mgotest5")
		//insertData := bson.M{"_id":7,"account":"admin","balance":"加ID更新前" }
		id := 0
		transaction2 := &TransactionSerial{}
		for i := 0; i < 10; i++ {
			transaction1 := &TransactionSerial{}
			for i := 0; i < 10; i++ {
				transaction := &TransactionSerial{}
				for i := 0; i < 10; i++ {
					insertData := bson.M{"_id": id, "account": "admin", "balance": "加ID更新前"}
					id++
					insertCommand := c.CommandFactory().CreateInsertCommand(insertData)

					callback := &MongoTransactionCallback{
						command:insertCommand,
					}
					object := &MongoTransaction{
						command: insertCommand,
					}
					object.SetCallback(callback)


					if err := transaction.Add(object); err != nil {
						panic(err)
					}
				}
				transaction.SetCallback(&TransactionCallbackBase{})

				transaction1.Add(transaction)
			}
			transaction1.SetCallback(&TransactionCallbackBase{})

			transaction2.Add(transaction1)
		}
		transaction2.SetCallback(&TransactionCallbackBase{})

		Convey("TestInsertTransactionMutilSerial_Commit：串行事务插入,测试出错回滚 应该有错误", func() {
			err := transaction2.Commit()
			So(err, ShouldNotEqual, nil)
		})
		ico.Close()
	})
}


func TestInsertTransactionMutilParallel_Commit(t *testing.T) {
	Convey("The NewMongoConn  with a url", t, func() {
		ico, _ := NewMongoConn("mongodb://root:yzl17220@120.78.158.127:27017", false)
		db, _ := ico.DB("s_blog")
		c, _ := db.Col("mgotest5")
		//insertData := bson.M{"_id":7,"account":"admin","balance":"加ID更新前" }
		id := 0
		transaction2 := &TransactionParallel{}
		for i := 0; i < 10; i++ {
			transaction1 := &TransactionParallel{}
			for i := 0; i < 10; i++ {
				transaction := &TransactionParallel{}
				for i := 0; i < 10; i++ {
					insertData := bson.M{"_id": id, "account": "admin", "balance": "加ID更新前"}
					id++
					insertCommand := c.CommandFactory().CreateInsertCommand(insertData)

					callback := &MongoTransactionCallback{
						command:insertCommand,
					}
					object := &MongoTransaction{
						command: insertCommand,
					}
					object.SetCallback(callback)


					if err := transaction.Add(object); err != nil {
						panic(err)
					}
				}
				transaction.SetCallback(&TransactionCallbackBase{})

				transaction1.Add(transaction)
			}
			transaction1.SetCallback(&TransactionCallbackBase{})

			transaction2.Add(transaction1)
		}
		transaction2.SetCallback(&TransactionCallbackBase{})

		Convey("TestInsertTransactionMutilParallel_Commit：并行事务插入,测试出错回滚 应该有错误", func() {
			err := transaction2.Commit()
			So(err, ShouldNotEqual, nil)
		})
		ico.Close()
	})
}

func TestInsertTransactionMutilParallel_Serial_Commit(t *testing.T) {
	Convey("The NewMongoConn  with a url", t, func() {
		ico, _ := NewMongoConn("mongodb://root:yzl17220@120.78.158.127:27017", false)
		db, _ := ico.DB("s_blog")
		c, _ := db.Col("mgotest5")
		//insertData := bson.M{"_id":7,"account":"admin","balance":"加ID更新前" }
		id := 0
		transaction2 := &TransactionParallel{}
		for i := 0; i < 10; i++ {
			transaction1 := &TransactionParallel{}
			for i := 0; i < 10; i++ {
				transaction := &TransactionSerial{}
				for i := 0; i < 10; i++ {
					insertData := bson.M{"_id": id, "account": "admin", "balance": "加ID更新前"}
					id++
					insertCommand := c.CommandFactory().CreateInsertCommand(insertData)

					callback := &MongoTransactionCallback{
						command:insertCommand,
					}
					object := &MongoTransaction{
						command: insertCommand,
					}
					object.SetCallback(callback)


					if err := transaction.Add(object); err != nil {
						panic(err)
					}
				}
				transaction.SetCallback(&TransactionCallbackBase{})

				transaction1.Add(transaction)
			}
			transaction1.SetCallback(&TransactionCallbackBase{})

			transaction2.Add(transaction1)
		}
		transaction2.SetCallback(&TransactionCallbackBase{})

		Convey("TestInsertTransactionMutilParallel_Serial_Commit：并行串行事务插入,测试出错回滚 应该有错误", func() {
			err := transaction2.Commit()
			So(err, ShouldNotEqual, nil)
		})
		ico.Close()
	})
}





