package main

import (
	
	"fmt"
	"ziMod/ziface"
	"ziMod/zinet"
)

type PingRouter struct{
	BaseRouter zinet.BaseRouter
}
func (pr *PingRouter)PreHandle(request ziface.IRequest){

}
func (pr *PingRouter)PostHandle(request ziface.IRequest){

}
//正式处理请求,这里的请求已经将链接与消息封装成请求了
func (pr *PingRouter)Handle(request ziface.IRequest){
	fmt.Print("call router handler\n")
	fmt.Print("recv information msg id:",request.GetId(),"\n")
	fmt.Print("recv infotmation msg data:",string(request.GetData()),"\n")
	err :=request.GetConnection().SendMessage(1,[]byte("ping........"))
	if err !=nil{
		fmt.Print("request.GetConnection().SendMessage error")
	}
}
type HelloRouter struct{
	BaseRouter zinet.BaseRouter
}
func (pr *HelloRouter)PreHandle(request ziface.IRequest){

}
func (pr *HelloRouter)PostHandle(request ziface.IRequest){

}
//正式处理请求,这里的请求已经将链接与消息封装成请求了
func (pr *HelloRouter)Handle(request ziface.IRequest){
	fmt.Print("call router handler\n")
	fmt.Print("recv information msg id:",request.GetId(),"\n")
	fmt.Print("recv infotmation msg data:",string(request.GetData()),"\n")
	err :=request.GetConnection().SendMessage(1,[]byte("hello...."))
	if err !=nil{
		fmt.Print("request.GetConnection().SendMessage error")
	}
}

func main(){
	server := zinet.NewServer("SERVER")
	router := &PingRouter{}
	server.AddRouter(0,router)
	hello := &HelloRouter{}
	server.AddRouter(1,hello)
	server.Serve()
}