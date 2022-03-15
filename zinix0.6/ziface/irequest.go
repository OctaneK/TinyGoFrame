package ziface

//request模块将请求数据和请求方法绑定在一起
type IRequest interface{
	//得到当前的链接模块的索引
	GetConnection() IConnection//非常有意思，这个返回的是一个抽象基类方法
	//得到当前的数据，该数据是一个被封装好的消息
	GetData()[]byte
	//获取该请求的类型，使用一个数字表示
	GetId()uint32
}
