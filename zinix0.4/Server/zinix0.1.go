package main

import (
	"fmt"
	"ziMod/ziface"
	"ziMod/zinet"
)

type PingRouter struct{
	BaseRouter zinet.BaseRouter
}
func (pr *PingRouter)PreHandle(request ziface.IRequest){//使用外面包的抽象需要加包名
	fmt.Print("prehandle ping......")
	_,err := request.GetConnection().GetTcpConn().Write([]byte("PRE:message from server:ping.....\n"))
	if err!=nil{
		fmt.Print("PREhandle error")
		return
	}

}
//正式处理请求
func (pr *PingRouter)Handle(request ziface.IRequest){
	fmt.Print("handle ping......")
	_,err := request.GetConnection().GetTcpConn().Write([]byte("handle:message from server:ping.....\n"))
	if err!=nil{
		fmt.Print("handle error")
		return
	}
}
//处理请求之后进行收尾工作
func (pr *PingRouter)PostHandle(request ziface.IRequest){
	fmt.Print("Posthandle ping......")
	_,err := request.GetConnection().GetTcpConn().Write([]byte("POST:message from server:ping.....\n"))
	if err!=nil{
		fmt.Print("POST error")
		return
	}
}
func main(){
	server := zinet.NewServer("SERVER")
	router := &PingRouter{}
	server.AddRouter(router)
	server.Serve()
}