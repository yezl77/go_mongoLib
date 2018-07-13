package dber

import (
	"testing"
	"gopkg.in/mgo.v2/bson"
	. "github.com/smartystreets/goconvey/convey"
	"time"
	"sync"
	"fmt"
)

type person struct {
	Username string
	Age      int
}

//TODO mock
func TestNewMongoConn(t *testing.T) {
	Convey("The NewMongoConn  with a url", t, func() {
		url := "mongodb://root:yzl17220@120.78.158.127:27017"
		iso, err := NewMongoConn(url, false)
		iso.Close()
		Convey("The NewMongoConn 连接应该为没有错误", func() {
			So(err, ShouldEqual, nil)
		})
	})
}

func TestDB(t *testing.T) {
	Convey("The NewMongoConn  with a url", t, func() {
		url := "mongodb://root:yzl17220@120.78.158.127:27017"

		ico, _ := NewMongoConn(url, false)
		Convey("TestDB：连接数据库 应该为没有错误", func() {
			_, err := ico.DB("s_blog")
			So(err, ShouldEqual, nil)
		})

	})
}

func TestCol(t *testing.T) {
	Convey("The NewMongoConn  with a url", t, func() {
		url := "mongodb://root:yzl17220@120.78.158.127:27017"
		ico, _ := NewMongoConn(url, false)
		db, _ := ico.DB("s_blog")

		Convey("TestCol：连接集合表 应该为没有错误", func() {
			_, err := db.Col("mgotest")
			So(err, ShouldEqual, nil)
		})

	})
}

func TestInsert(t *testing.T) {

	Convey("The NewMongoConn  with a url", t, func() {
		ico, _ := NewMongoConn("mongodb://root:yzl17220@120.78.158.127:27017", false)
		db, _ := ico.DB("s_blog")
		c, _ := db.Col("mgotest")
		Convey("TestInsert：插入数据 应该为没有错误", func() {
			err := c.Insert(&person{
				Username: "admin",
				Age:      10,
			})
			So(err, ShouldEqual, nil)
		})
		ico.Close()
	})
}

func TestInsertSameId(t *testing.T) {
	Convey("The NewMongoConn  with a url", t, func() {
		ico, _ := NewMongoConn("mongodb://root:yzl17220@120.78.158.127:27017", false)
		db, _ := ico.DB("s_blog")
		c, _ := db.Col("mgotest")

		var data interface{}
		c.FindOne(bson.M{"age": 99999}, &data)
		Convey("TestInsertSameId：插入相同ID数据 应该报错误", func() {
			err := c.Insert(data)
			So(err, ShouldNotEqual, nil)
		})
		ico.Close()
	})
}

func TestRemove(t *testing.T) {
	Convey("The NewMongoConn  with a url", t, func() {
		ico, _ := NewMongoConn("mongodb://root:yzl17220@120.78.158.127:27017", false)
		db, _ := ico.DB("s_blog")
		c, _ := db.Col("mgotest")
		Convey("TestRemove：删除数据 应该错误为nil", func() {
			err := c.Remove(bson.M{"username": "admin"})
			So(err, ShouldEqual, nil)
		})
		ico.Close()
	})
}

func TestRemoveId(t *testing.T) {
	Convey("The NewMongoConn  with a url", t, func() {
		ico, _ := NewMongoConn("mongodb://root:yzl17220@120.78.158.127:27017", false)
		db, _ := ico.DB("s_blog")
		c, _ := db.Col("mgotest")
		Convey("TestRemove：删除数据 应该错误为nil", func() {
			err := c.RemoveId(bson.ObjectIdHex("5b34a1379903e2b5b0903302"))
			So(err, ShouldEqual, nil)
		})
		ico.Close()
	})
}

func TestUpdate(t *testing.T) {
	Convey("The NewMongoConn  with a url", t, func() {
		ico, _ := NewMongoConn("mongodb://root:yzl17220@120.78.158.127:27017", false)
		db, _ := ico.DB("s_blog")
		c, _ := db.Col("mgotest")
		Convey("TestUpdate：更新数据 应该错误为nil", func() {
			err := c.Update(bson.M{"username": "admin"}, bson.M{"$set": bson.M{"age": 22222222222222222}})
			So(err, ShouldEqual, nil)
		})
		ico.Close()
	})
}

func TestUpdateId(t *testing.T) {
	Convey("The NewMongoConn  with a url", t, func() {
		ico, _ := NewMongoConn("mongodb://root:yzl17220@120.78.158.127:27017", false)
		db, _ := ico.DB("s_blog")
		c, _ := db.Col("mgotest")
		Convey("TestUpdateId：更新数据Id 应该错误为nil", func() {
			err := c.UpdateId(bson.ObjectIdHex("5b3335cb9903e2b5b08f8335"), bson.M{"$set": bson.M{"age": 99999}}, true)
			So(err, ShouldEqual, nil)
		})
		ico.Close()
	})
}

func TestCount(t *testing.T) {
	Convey("The NewMongoConn  with a url", t, func() {
		ico, _ := NewMongoConn("mongodb://root:yzl17220@120.78.158.127:27017", false)
		db, _ := ico.DB("s_blog")
		c, _ := db.Col("mgotest")
		c.Insert(&person{
			Username: "admin",
			Age:      10,
		})
		Convey("TestCount：查询数量 应该错误为nil,count不为0", func() {
			count, err := c.Count(bson.M{"username": "admin"})
			So(err, ShouldEqual, nil)
			So(count, ShouldNotEqual, 0)
		})

	})
}

func TestFindAll(t *testing.T) {
	Convey("The NewMongoConn  with a url", t, func() {
		ico, _ := NewMongoConn("mongodb://root:yzl17220@120.78.158.127:27017", false)
		db, _ := ico.DB("s_blog")
		c, _ := db.Col("mgotest")

		Convey("TestFindAll：查询全部 应该错误为nil", func() {
			var data []interface{}
			err := c.FindAll(bson.M{"age": 5454}, &data)
			So(err, ShouldEqual, nil)
			t.Log("TestFindAll:", data)
		})

	})
}

func TestFindOne(t *testing.T) {
	Convey("The NewMongoConn  with a url", t, func() {
		ico, _ := NewMongoConn("mongodb://root:yzl17220@120.78.158.127:27017", false)
		db, _ := ico.DB("s_blog")
		c, _ := db.Col("mgotest")
		Convey("TestFindOne：查询一个 应该错误为nil,", func() {
			var data interface{}
			err := c.FindOne(bson.M{"age": 99999}, &data)
			So(err, ShouldEqual, nil)
			t.Log("TestFindAll:", data)
		})
	})
}

func TestFindId(t *testing.T) {
	Convey("The NewMongoConn  with a url", t, func() {
		ico, _ := NewMongoConn("mongodb://root:yzl17220@120.78.158.127:27017", false)
		db, _ := ico.DB("s_blog")
		c, _ := db.Col("mgotest")

		Convey("TestFindId：查找数据 应该错误为nil,", func() {
			var dataId interface{}
			err := c.FindId(bson.ObjectIdHex("5b3335cb9903e2b5b08f8335"), &dataId)
			So(err, ShouldEqual, nil)
			t.Log("TestFindAll:", dataId)
		})
	})
}

func TestFindAndModify(t *testing.T) {
	Convey("The NewMongoConn  with a url", t, func() {
		ico, _ := NewMongoConn("mongodb://root:yzl17220@120.78.158.127:27017", false)
		db, _ := ico.DB("s_blog")
		c, _ := db.Col("mgotest")
		Convey("TestFindAndModify：查找数据 应该错误为nil,", func() {
			var dataM []interface{}
			err := c.FindAndModify(bson.M{"age": 54545}, &person{Username: "admin", Age: 4444445555,}, &dataM)
			So(err, ShouldEqual, nil)
			t.Log("TestFindAll:", dataM)
		})

	})
}

func TestFindAllSort(t *testing.T) {
	Convey("The NewMongoConn  with a url", t, func() {
		ico, _ := NewMongoConn("mongodb://root:yzl17220@120.78.158.127:27017", false)
		db, _ := ico.DB("s_blog")
		c, _ := db.Col("mgotest")
		Convey("TestFindAll：查询全部并根据key排序 应该错误为nil", func() {
			var data []interface{}
			q := c.CommandFactory().NewQuery(bson.M{"username": "fdfsffsdf"}, false).Sort("age").All(&data)
			err :=q.Do( )
			So(err, ShouldEqual, nil)
			t.Log("TestFindAllSort:", data)
		})
	})
}

func TestFindAllLimit(t *testing.T) {
	Convey("The NewMongoConn  with a url", t, func() {
		ico, _ := NewMongoConn("mongodb://root:yzl17220@120.78.158.127:27017", false)
		db, _ := ico.DB("s_blog")
		c, _ := db.Col("mgotest")
		Convey("TestFindAll：查询全部并限制个数 应该错误为nil", func() {
			var data []interface{}
			command := c.CommandFactory().NewQuery(bson.M{"username": "fdfsffsdf"}, false).Limit(2).All(&data)
			err :=command.Do()
			So(err, ShouldEqual, nil)
			t.Log("TestFindAllSort:", data)
		})
	})
}

func TestMakeIndexes(t *testing.T) {
	Convey("The NewMongoConn  with a url", t, func() {
		ico, _ := NewMongoConn("mongodb://root:yzl17220@120.78.158.127:27017", false)
		db, _ := ico.DB("s_blog")
		c, _ := db.Col("mgotest")
		indexes := [][]string{
			{"99999", "5b3335cb9903e2b5b08f8335"}, {"a", "b", "d", "e"}, {"a", "b", "d", "e"},
		}

		Convey("TestMakeIndexes：创建索引 应该错误为nil,", func() {
			err := c.MakeIndexes(indexes)
			So(err, ShouldEqual, nil)
		})
	})
}

func TestDropCol(t *testing.T) {
	Convey("The NewMongoConn  with a url", t, func() {
		ico, _ := NewMongoConn("mongodb://root:yzl17220@120.78.158.127:27017", false)
		db, _ := ico.DB("s_blog")
		c, _ := db.Col("item")
		c.Insert(&person{
			Username: "admin",
			Age:      10,
		})

		Convey("TestDropCol：删除数据集 应该错误为nil,", func() {
			err := c.DropCol()
			So(err, ShouldEqual, nil)
		})

	})
}

func TestInsertBulk(t *testing.T) {
	Convey("The NewMongoConn  with a url", t, func() {
		ico, _ := NewMongoConn("mongodb://root:yzl17220@120.78.158.127:27017", false)
		db, _ := ico.DB("s_blog")
		c, _ := db.Col("mgotest3")
		c2, _ := db.Col("mgotest21")
		var data []interface{}
		c.FindAll(bson.M{"username": "fdfsffsdf"}, &data)
		fmt.Println(data)

		Convey("TestInsertBulk：批量插入数据集 应该错误为nil,", func() {
			b := c2.CommandFactory().NewBulk()
			b.Insert(data...)
			err:=b.Do()
			So(err, ShouldEqual, nil)
		})
	})
}
func TestRemoveAllBulk(t *testing.T) {
	Convey("The NewMongoConn  with a url", t, func() {
		ico, _ := NewMongoConn("mongodb://root:yzl17220@120.78.158.127:27017", false)
		db, _ := ico.DB("s_blog")
		c2, _ := db.Col("mgotest3")
		Convey("TestRemoveAllBulk：批量删除所有数据集 应该错误为nil,", func() {
			b := c2.CommandFactory().NewBulk()

			b.RemoveAll(bson.M{"username": "admin"},bson.M{"age": 499})

			err:=b.Do()
			So(err, ShouldEqual, nil)
		})
	})
}

func TestRemoveBulk(t *testing.T) {
	Convey("The NewMongoConn  with a url", t, func() {
		ico, _ := NewMongoConn("mongodb://root:yzl17220@120.78.158.127:27017", false)
		db, _ := ico.DB("s_blog")
		c2, _ := db.Col("mgotest3")
		Convey("TestRemoveBulk：批量删除数据集 应该错误为nil,", func() {
			b := c2.CommandFactory().NewBulk()
			b.Remove(bson.M{"username": "admin"},bson.M{"age": 499})
			err:=b.Do()
			So(err, ShouldEqual, nil)
		})
	})
}

func TestUpdateAllBulk(t *testing.T) {
	Convey("The NewMongoConn  with a url", t, func() {
		ico, _ := NewMongoConn("mongodb://root:yzl17220@120.78.158.127:27017", false)
		db, _ := ico.DB("s_blog")
		c2, _ := db.Col("mgotest3")
		Convey("TestInsertBulk：批量更新所有数据集 应该错误为nil,", func() {
			b := c2.CommandFactory().NewBulk()
			b.UpdateAll(bson.M{"username": "root"}, bson.M{"$set": bson.M{"age": 22222222222222222}},bson.M{"age": 7777}, bson.M{"$set": bson.M{"username": "root","age": 777777777}},)
			err:=b.Do()
			So(err, ShouldEqual, nil)
		})
	})
}


func TestUpdateBulk(t *testing.T) {
	Convey("The NewMongoConn  with a url", t, func() {
		ico, _ := NewMongoConn("mongodb://root:yzl17220@120.78.158.127:27017", false)
		db, _ := ico.DB("s_blog")
		c2, _ := db.Col("mgotest3")
		Convey("TestUpdateBulk：批量更新数据集 应该错误为nil,", func() {
			b := c2.CommandFactory().NewBulk()
			b.Update(bson.M{"username": "root"}, bson.M{"$set": bson.M{"age": 4242454554}},bson.M{"age": 7777}, bson.M{"$set": bson.M{"username": "root","age": 777777777}},)
			err:=b.Do()
			So(err, ShouldEqual, nil)
		})
	})
}


func TestUpsertBulk(t *testing.T) {
	Convey("The NewMongoConn  with a url", t, func() {
		ico, _ := NewMongoConn("mongodb://root:yzl17220@120.78.158.127:27017", false)
		db, _ := ico.DB("s_blog")
		c2, _ := db.Col("mgotest3")
		Convey("TestUpsertBulk：批量更新或插入数据集 应该错误为nil,", func() {
			b := c2.CommandFactory().NewBulk()
			b.Upsert(bson.M{"username": "bai"}, bson.M{"$set": bson.M{"age": 4242454554}},bson.M{"age": 4414141}, bson.M{"$set": bson.M{"username": "root","age": 777777777}},)
			time.Sleep(10 * time.Second)
			err:=b.Do()
			So(err, ShouldEqual, nil)
		})
	})
}

func TestBulk(t *testing.T) {
	Convey("The NewMongoConn  with a url", t, func() {
		ico, _ := NewMongoConn("mongodb://root:yzl17220@120.78.158.127:27017", false)
		db, _ := ico.DB("s_blog")
		c2, _ := db.Col("mgotest3")
		Convey("TestInsertBulk：批量插入数据集 应该错误为nil,", func() {
			b := c2.CommandFactory().NewBulk()
			//b.Upsert("bai", bson.M{"$set": bson.M{"age": 4242454554}},bson.M{"age": 4414141}, bson.M{"$set": bson.M{"username": "root","age": 777777777}},)

			b.Update(bson.M{"username": "root"}, bson.M{"$set": bson.M{"age": 4242454554}},bson.M{"age": 7777}, bson.M{"$set": bson.M{"username": "root","age": 777777777}},)
			b.UpdateAll(bson.M{"username": "root"}, bson.M{"$set": bson.M{"age": 22222222222222222}},bson.M{"age": 7777}, bson.M{"$set": bson.M{"username": "root","age": 777777777}},)
			b.Remove(bson.M{"username": "bai"},bson.M{"age": 499})
			p1:=&person{Age:777,Username:"uuu",}
			p2:=&person{Age:777,Username:"rrr",}
			p3:=&person{Age:414,Username:"rrr",}
			p4:=&person{Age:777,Username:"rrr",}
			//ps:=new person[p1,p2,p3,p4]
			b.Upsert("bai", bson.M{"$set": bson.M{"age": 4242454554}},bson.M{"age": 4414141}, bson.M{"$set": bson.M{"username": "root","age": 777777777}},)
			b.Insert(p1,p2,p3,p4)
			b.Unordered()
			b.RemoveAll(bson.M{"username": "rrr"},bson.M{"age": 414})
			time.Sleep(10 * time.Second)
			err:=b.Do()
			So(err, ShouldEqual, nil)
		})
	})
}

func TestBulkError(t *testing.T) {
	Convey("The NewMongoConn  with a url", t, func() {
		ico, _ := NewMongoConn("mongodb://root:yzl17220@120.78.158.127:27017", false)
		db, _ := ico.DB("s_blog")
		c2, _ := db.Col("mgotest2")
		Convey("TestInsertBulk：批量插入数据集 应该错误为nil,", func() {
			b := c2.CommandFactory().NewBulk()
			//time.Sleep(10 * time.Second)
			b.Unordered()

			b.Update(bson.M{"username": "root"}, bson.M{"$set": bson.M{"age": 4242454554}})

			var data []interface{}
			c2.FindAll(bson.M{"username": "testsameid"}, &data)
			b.Insert(data...)

			b.Update(bson.M{"username": "root"}, bson.M{"$set": bson.M{"age": 4242454554}})

			b.Insert(data...)

			b.Upsert(bson.M{"username": "bai"}, bson.M{"$set": bson.M{"age": 11111111}})

			err:=b.Do()
			for _,cas := range b.GetBulkErrorCases(){
				fmt.Println(cas.Index," ",cas.err)
			}


			fmt.Println("========err======\n",err)
			So(err, ShouldEqual, nil)
		})
	})
}

func TestUnOrderBulk(t *testing.T) {
	Convey("The NewMongoConn  with a url", t, func() {
		ico, _ := NewMongoConn("mongodb://root:yzl17220@120.78.158.127:27017", false)
		db, _ := ico.DB("s_blog")
		c2, _ := db.Col("mgotest3")
		Convey("TestInsertBulk：批量插入数据集 应该错误为nil,", func() {
			b := c2.CommandFactory().NewBulk()

			b.Update(bson.M{"username": "root"}, bson.M{"$set": bson.M{"age": 4242454554}},bson.M{"age": 7777}, bson.M{"$set": bson.M{"username": "root","age": 777777777}},)
			b.UpdateAll(bson.M{"username": "root"}, bson.M{"$set": bson.M{"age": 22222222222222222}},bson.M{"age": 7777}, bson.M{"$set": bson.M{"username": "root","age": 777777777}},)
			b.Remove(bson.M{"username": "bai"},bson.M{"age": 499})
			p1:=&person{Age:777,Username:"rrr",}
			p2:=&person{Age:777,Username:"rrr",}
			p3:=&person{Age:414,Username:"rrr",}
			p4:=&person{Age:777,Username:"uuu",}
			//ps:=new person[p1,p2,p3,p4]
			time.Sleep(10 * time.Second)
			b.Upsert(bson.M{"username": "bai"}, bson.M{"$set": bson.M{"age": 4242454554}},bson.M{"age": 4414141}, bson.M{"$set": bson.M{"username": "root","age": 777777777}},)
			b.Insert(p1,"dfdf",p2,p3,p4)
			b.RemoveAll(bson.M{"username": "uuu"},bson.M{"age": 414})
			b.Unordered()

			time.Sleep(10 * time.Second)
			err:=b.Do()
			So(err, ShouldEqual, nil)
		})
	})
}

func TestIcoCloseInsert(t *testing.T) {
	Convey("The NewMongoConn  with a url", t, func() {
		ico, _ := NewMongoConn("mongodb://root:yzl17220@120.78.158.127:27017", false)
		db1, _ := ico.DB("s_blog")
		c1, _ := db1.Col("mgotest2")

		ico.Close()
		Convey("关闭连接ico 后尝试插入数据 应该错误不为nil,", func() {
			err := c1.Insert(&person{
				Username: "admin",
				Age:      10,})
			So(err, ShouldNotEqual, nil)
		})
	})
}

func TestDbCloseInsert(t *testing.T) {
	Convey("The NewMongoConn  with a url", t, func() {
		ico, _ := NewMongoConn("mongodb://root:yzl17220@120.78.158.127:27017", false)
		db1, _ := ico.DB("s_blog")
		c1, _ := db1.Col("mgotest2")
		db1.Close()
		Convey("关闭数据库db1 后尝试插入数据 应该错误不为nil,", func() {
			err := c1.Insert(&person{
				Username: "admin",
				Age:      10,})
			So(err, ShouldNotEqual, nil)
		})
	})
}

func TestColCloseInsert(t *testing.T) {
	Convey("The NewMongoConn  with a url", t, func() {
		ico, _ := NewMongoConn("mongodb://root:yzl17220@120.78.158.127:27017", false)
		db1, _ := ico.DB("s_blog")
		c1, _ := db1.Col("mgotest2")
		c1.Close()
		Convey("关闭数据集c1 后尝试插入数据 应该错误不为nil,", func() {
			err := c1.Insert(&person{
				Username: "admin",
				Age:      10,})
			So(err, ShouldNotEqual, nil)
		})
	})
}

func TestDbCloseCol(t *testing.T) {
	Convey("The NewMongoConn  with a url", t, func() {
		ico, _ := NewMongoConn("mongodb://root:yzl17220@120.78.158.127:27017", false)
		db1, _ := ico.DB("s_blog")
		db1.Close()
		Convey("关闭数据库db1  后尝试重新连接数据集s_blog 应该错误！=nil,", func() {
			_, err := db1.Col("mgotest2")
			So(err, ShouldNotEqual, nil)
		})
	})
}

func TestIcoCloseDB(t *testing.T) {
	Convey("The NewMongoConn  with a url", t, func() {
		ico, _ := NewMongoConn("mongodb://root:yzl17220@120.78.158.127:27017", false)
		ico.Close()
		Convey("关闭数据库db1  后尝试重新连接数据集s_blog 应该错误！=nil,", func() {
			_, err := ico.DB("s_blog")
			So(err, ShouldNotEqual, nil)
		})
	})
}

func TestIcoCloseCOl(t *testing.T) {
	Convey("The NewMongoConn  with a url", t, func() {
		ico, _ := NewMongoConn("mongodb://root:yzl17220@120.78.158.127:27017", false)
		db1, _ := ico.DB("s_blog")
		ico.Close()
		Convey("关闭数据库ico  后尝试重新连接数据集s_blog 应该错误！=nil,", func() {
			_, err := db1.Col("mgotest2")
			So(err, ShouldNotEqual, nil)
		})
	})
}

//测试网络错误对多操作的影响
//测试结果：
//批量插入出现其他错误，如queryError,(没有查找到对应的记录)。不会再重试
//连接可以得到重用，没有出现连接溢出。连接数保持在1-4条
func TestMutilProcessANDReconnect(t *testing.T) {
	Convey("The NewMongoConn  with a url", t, func() {
		ico, _ := NewMongoConn("mongodb://root:yzl17220@120.78.158.127:27017", false)
		db, _ := ico.DB("s_blog")
		c, _ := db.Col("mgotest3")
		var wg sync.WaitGroup
		Convey("测试中途断网 重新联网后依然完成 应该错误==nil,", func() {
			for i := 1; i <= 10000; i++ {
				fmt.Println("=====================测试第几次：", i)
				for i := 1; i <= 10; i++ {
					wg.Add(1)
					go func() {
						time.Sleep(10 * time.Second)
						err := c.Insert(&person{
							Username: "admin",
							Age:      10,
						})
						if err == nil {
							fmt.Println("测试成功")
						} else {
							fmt.Println("测试失败")
						}
						wg.Done()
						//time.Sleep(10 * time.Second)
						//
					}()
				}
				c.DropCol()
				wg.Wait()
			}
		})
	})
}


