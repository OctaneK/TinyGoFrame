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
	MsgHandle ziface.IMsgHandle
}
//非常有意思的写法，返回的是抽象层的接口方法而不是实际的server指针
func NewServer(name string) *Server {
	s := &Server{
		Name:      utils.GolobalObject.Name,
		IPversion: "tcp4",
		IP:        utils.GolobalObject.Host,
		Port:      utils.GolobalObject.TcpPort,
		MsgHandle: NewMsgHandle(),
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
		//阻塞的等待客户端链接，处理客户端链接业务
		for {
			//有客户链接请求，则返回
			conn, e := listener.AcceptTCP()
			if e != nil {
				fmt.Print("listen error")
				continue
			}
			var id uint32 = 0
			fmt.Print("new  connection established\n")
			//将该链接与业务方法绑定，并启动链接模块为其服务
			//将服务器的router复制给该链接
			delCon := NewConnection(conn, id, server.MsgHandle)
			id++
			go delCon.Start()//启动连接模块

		}
	}()

}
func (server *Server) Stop() {

}
func (server *Server) Serve() {
	//异步的启动服务器
	server.Start()
	fmt.Print("server name:",server.Name,"\n")
	//做一些启动服务器之外的业务
	select {}
}


