package ziface

import "net"
type IConnection interface{
	//启动链接
	Start()
	//关闭链接
	Stop()
	//获取套接字conn
	GetTcpConn()*net.TCPConn
	//获取套接字ip：port
	GetAddr()net.Addr
	//获取链接id
	GetLinkedId()uint32
	//发送数据
	SendMessage(uint32,[]byte)error
}
//定义一个处理链接的业务的方法，看起来很像static方法，游离于类之外
type HandleFunc func(*net.TCPConn,[]byte,int)error