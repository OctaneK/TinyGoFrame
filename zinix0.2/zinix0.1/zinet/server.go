package zinet

import (
	"fmt"
	"net"

	"ziMod/ziface"
)

type Server struct {
	Name      string
	IPversion string
	IP        string
	Port      int
}

//只有完全继承了三个方法才能在后面new方法返回抽象层方法
func (server *Server) Start() {
	go func() {
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
			conn, e := listener.Accept()
			if e != nil {
				fmt.Print("listen error")
				continue
			}
			fmt.Print("new  connection established\n")
			go func() {
				for {
					buf := make([]byte, 512)
					cnt, err := conn.Read(buf)
					if err != nil {
						fmt.Print("read error")
						break
					}
					
					_, er := conn.Write(buf[0:cnt])
					if er != nil {
						fmt.Print("write error")
						continue
					}

				}

			}()

		}
	}()

}
func (server *Server) Stop() {
	

}
func (server *Server) Serve() {
	//异步的启动服务器
	server.Start()
	//做一些启动服务器之外的业务
	select{

	}
}

//非常有意思的写法，返回的是抽象层的接口方法而不是实际的server指针
func NewServer(name string) ziface.Iserver {
	s := &Server{
		Name:      name,
		IPversion: "tcp",
		IP:        "127.0.0.1",
		Port:      8888,
	}
	return s
}
