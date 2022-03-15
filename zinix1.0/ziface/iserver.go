package ziface

type Iserver interface {//真有意思，抽象类调用抽象
	//服务器启动
	Start()
	//服务器停止
	Stop()
	//服务器进行服务
	Serve()

	AddRouter(uint32,IRouter)//导入自己包的抽象方法不能加包名
	GetconMgr()IConnmanager//获取链接管理模块
	//设置钩子函数
	SetHookStart(func(IConnection))
	SetHookStop(func(IConnection))
	//调用钩子函数
	CallHookStart(IConnection)
	CallHookStop(IConnection)

}
