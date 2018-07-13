package dber



//
//func TestServerNotReachError(t *testing.T) {
//	Convey("The dber.NewMongoConn  with a url", t, func() {
//
//		ico, _ := NewMongoConn("mongodb://root:yzl17220@120.78.158.127:27017", false)
//		db1, _ := ico.DB("s_blog")
//		c1, _ := db1.Col("mgotest2")
//
//		ctrl := gomock.NewController(t)
//		defer ctrl.Finish()
//		com := test.NewMockICommand(ctrl)
//
//		netOpErr := errors.New("no reachable servers")
//		com.EXPECT().Do(false).Return(netOpErr)
//		Convey("mock 测试网络错误 返回no reachable servers 错误 应该错误 ！=nil,", func() {
//			err := c1.InsertTest(com)
//			t.Log(err)
//			So(err, ShouldNotEqual, nil)
//		})
//
//	})
//}

//测试结果：正确
//func TestServerNotReachError(t *testing.T) {
//	Convey("The dber.NewMongoConn  with a url", t, func() {
//		ico, _ := NewMongoConn("mongodb://root:yzl17220@120.78.158.127:27017", false)
//		db1, _ := ico.DB("s_blog")
//		c1, _ := db1.Col("mgotest2")
//		ctrl := gomock.NewController(t)
//		defer ctrl.Finish()
//		com := NewMockICommand(ctrl)
//		netOpErr := errors.New("no reachable servers")
//		com.EXPECT().Do(false).Return(netOpErr)
//		Convey("mock 测试网络错误 返回no reachable servers 错误 应该错误 ！=nil,", func() {
//			err := c1.doCommand(com,false)
//			t.Log(err)
//			So(err, ShouldNotEqual, nil)
//		})
//
//	})
//}
//
////测试结果：暂时存在注入错误报错的情况
//func TestNetOpError(t *testing.T) {
//	Convey("The dber.NewMongoConn  with a url", t, func() {
//		ico, _ := NewMongoConn("mongodb://root:yzl17220@120.78.158.127:27017", false)
//		db1, _ := ico.DB("s_blog")
//		c1, _ := db1.Col("mgotest2")
//
//		ctrl := gomock.NewController(t)
//		defer ctrl.Finish()
//		com := NewMockICommand(ctrl)
//		Convey(" 测试网络错误*net.OpError应该错误！=nil,", func() {
//			var opErr error
//			opErr = &net.OpError{Err: errors.New("net.OpError")}
//			com.EXPECT().Do(false).Return(opErr)
//
//			Convey("关闭数据集后尝试插入数据 应该错误不为nil,", func() {
//				err := c1.doCommand(com,false)
//				t.Log(err)
//				So(err, ShouldNotEqual, nil)
//			})
//		})
//	})
//}
