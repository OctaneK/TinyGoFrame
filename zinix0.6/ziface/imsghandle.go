package ziface

type IMsgHandle interface{
	//根据每个请求具体执行相应的方法
	DoMsgHandler(IRequest)
	//为每个具体的类型id添加相应的方法router,
	AddMsgHandler(uint32,IRouter)
}