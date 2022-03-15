package ziface

//request模块将请求数据和请求方法绑定在一起
type IRequest interface{
	//得到当前的链接
	GetConnection() IConnection//非常有意思，这个返回的是一个抽象基类方法
	//得到当前的数据
	GetData()[]byte

	GetId()uint32
}
