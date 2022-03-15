package ziface
/*
定义路由的抽象接口
通过该接口，可以让到来的请求正确处理
*/
type IRouter interface{
	//在正式处理之前预处理hook
	PreHandle(request IRequest)
	//正式处理请求
	Handle(request IRequest)
	//处理请求之后进行收尾工作
	PostHandle(request IRequest)
}