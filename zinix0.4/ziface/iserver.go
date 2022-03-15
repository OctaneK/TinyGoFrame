package ziface

type Iserver interface {//真有意思，抽象类调用抽象
	//服务器启动
	Start()
	//服务器停止
	Stop()
	//服务器进行服务
	Serve()

	AddRouter(router IRouter)//导入自己包的抽象方法不能加包名

}
