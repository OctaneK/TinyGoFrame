package zinet

import (
	"fmt"
	"net"

	"ziMod/ziface"
)
type Connection struct{
	//链接绑定的socket
	Conn *net.TCPConn
	//链接的pid
	ConnectionID uint32
	//当前链接的状态
	isClose bool
	//链接绑定的方法
	HandleAPI ziface.HandleFunc
	//告知是否关闭的channel
	ExitChan chan bool
}
func NewConnection(con *net.TCPConn,id uint32,callback ziface.HandleFunc)*Connection{
	conn := &Connection{//与上面的con重名，直接编译器报错
		Conn : con,
		ConnectionID: id,
		isClose: false,
		HandleAPI: callback,
		ExitChan: make(chan bool,1),
	}
	return conn;
}
func (c *Connection)ReadMessage(){
	fmt.Print("start reading。。。",c.ConnectionID)
	defer c.Stop()
	defer fmt.Print("server is closeing...")

	for{
		buf := make([]byte,512)
		cnt,err :=c.Conn.Read(buf)
		if err!=nil{
			fmt.Print("read error")
			continue
		}
		er :=c.HandleAPI(c.Conn,buf,cnt)
		if er !=nil{
			fmt.Print(c.ConnectionID,"handleFunc error")
			break;
		}
	}
}
//启动链接,让当前链接开始工作
func (c *Connection)Start(){
	fmt.Print("Conn start....ConnID",c.ConnectionID)
	//启动从当前链接读数据的业务
	go c.ReadMessage()
	//启动从当前链接写的业务

}
//关闭链接
func (c *Connection)Stop(){
	fmt.Print("connection closed ID :",c.GetLinkedId())
	if c.isClose==true{
		fmt.Print("has closed")
	}
	c.isClose=true
	c.Conn.Close()//关闭连接的资源
	close(c.ExitChan)
}
//获取套接字conn
func (c *Connection)GetTcpConn()*net.TCPConn{
	return c.Conn
}
//获取套接字ip：port
func (c *Connection)GetAddr()net.Addr{//注意这里返回的是地址，而不是指针
	 
	 return c.Conn.RemoteAddr()
}
//获取链接id
func (c *Connection)GetLinkedId()uint32{
	return c.ConnectionID

}
//发送数据
func (c *Connection)SendMessage([]byte)error{
	return nil

}