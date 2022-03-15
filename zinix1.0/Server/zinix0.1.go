package main

import (
	
	"fmt"
	"ziMod/ziface"
	"ziMod/zinet"
)

type PingRouter struct{
	
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


func MyHookStart(con ziface.IConnection){
	fmt.Print("this is hookstart func\n")
	con.SendMessage(0,[]byte("欢迎上线！\n"))
}
func MyHookStop(con ziface.IConnection){
	fmt.Print("this is hookStop func")
	fmt.Print(con.GetLinkedId(),"已下线\n")
}
func main(){
	server := zinet.NewServer("SERVER")
	//添加钩子函数
	server.SetHookStart(MyHookStart)
	server.SetHookStop(MyHookStop)
	//添加路由功能
	router := &PingRouter{}
	server.AddRouter(0,router)
	hello := &HelloRouter{}
	server.AddRouter(1,hello)

	//启动服务器
	server.Serve()
}