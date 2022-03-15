package zinet

import (
	//"errors"
	"fmt"
	"net"

	"ziMod/utils"
	"ziMod/ziface"
)

type Server struct {
	Name      string
	IPversion string
	IP        string
	Port      int
	MsgHandle ziface.IMsgHandle//消息管理模块
	ConnManager ziface.IConnmanager//链接管理模块

	HookStart func(ziface.IConnection)//建立链接之前业务
	HookStop func(ziface.IConnection)//销毁链接之前业务
}
//非常有意思的写法，返回的是抽象层的接口方法而不是实际的server指针
func NewServer(name string) *Server {
	s := &Server{
		Name:      utils.GolobalObject.Name,
		IPversion: "tcp4",
		IP:        utils.GolobalObject.Host,
		Port:      utils.GolobalObject.TcpPort,
		MsgHandle: NewMsgHandle(),
		ConnManager: NewConnManager(),
		HookStart: nil,
		HookStop: nil,
	}
	return s
}
/*
//这个为服务器在收到数据之后做的业务，接受数据已经在此之前完成
func CallBackToClient(conn *net.TCPConn, buf []byte, cnt int) error {

	_, er := conn.Write(buf[0:cnt])
	if er != nil {
		fmt.Print("write error")
		return errors.New("CallBackToClient error")
	}
	return nil

}
*/
//添加路由功能，为每一个请求执行对应的方法
func (server *Server)AddRouter(id uint32,Router ziface.IRouter){
	server.MsgHandle.AddMsgHandler(id,Router)
}
//只有完全继承了三个方法才能在后面new方法返回抽象层方法
func (server *Server) Start() {
	go func() {
		fmt.Print("当前链接最大为： ",utils.GolobalObject.MaxConn,"\n")
		//开启工作池和消息队列
		server.MsgHandle.StartWokerPool()

		//获取一个tcp链接addr
		addr, err := net.ResolveTCPAddr(server.IPversion, fmt.Sprintf("%s:%d", server.IP, server.Port))
		if err != nil {
			fmt.Print("reslovetcp error")
			return
		}
		//监听服务器地址
		listener, er := net.ListenTCP(server.IPversion, addr)
		if er != nil {
			fmt.Print("Listen error")
			return
		}
		fmt.Print("listening....\n")
		var id uint32 = 0
		//阻塞的等待客户端链接，处理客户端链接业务
		for {
			//有客户链接请求，则返回
			conn, e := listener.AcceptTCP()
			if e != nil {
				fmt.Print("listen error")
				continue
			}
			//在链接接入之前需要判断服务器是否已经到达最大连接个数
			if server.ConnManager.Len()>=utils.GolobalObject.MaxConn{
				conn.Close()//将打开的连接关闭
				fmt.Print("已经达到最大连接个数\n")
				continue
			}
			
			//fmt.Print("new  connection established\n")
			//将该链接与业务方法绑定，并启动链接模块为其服务
			//将服务器的router复制给该链接
			delCon := NewConnection(server,conn, id, server.MsgHandle)
			id++
			go delCon.Start()//启动连接模块

		}
	}()

}
func (server *Server)GetconMgr()ziface.IConnmanager{
	return server.ConnManager
}
func (server *Server) Stop() {
	fmt.Print("服务器已关闭\n")
	//回收所有资源
	server.ConnManager.CleanCon()

}
func (server *Server) Serve() {
	//异步的启动服务器
	server.Start()
	fmt.Print("server name:",server.Name,"\n")
	//做一些启动服务器之外的业务
	select {}
}
//设置钩子函数
func (server*Server)SetHookStart(hook func(ziface.IConnection)){
	server.HookStart=hook
}
func (server*Server)SetHookStop(hook func(ziface.IConnection)){
	server.HookStop=hook
}
//调用钩子函数
func (server*Server)CallHookStart(con ziface.IConnection){
	if server.HookStart!=nil{
		fmt.Print("hookstart.........\n")
		server.HookStart(con)
	}
}
func (server*Server)CallHookStop(con ziface.IConnection){
	if server.HookStop !=nil{
		fmt.Print("hookstop...........\n")
		server.HookStop(con)
	}
}

