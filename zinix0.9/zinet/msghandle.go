package zinet

import (
	"fmt"
	"ziMod/utils"
	"ziMod/ziface"

	
)
type MsgHandle struct{
	//为每个id的消息指向相应的方法
	APIS map[uint32] ziface.IRouter
	//业务工作池的数量
	WorkerPoolSize uint32
	//每个go程的工作队列
	TaskQueue []chan ziface.IRequest

}
func NewMsgHandle()*MsgHandle{
	return &MsgHandle{

		APIS : make(map[uint32]ziface.IRouter),
		//工作go程的数量
		WorkerPoolSize: utils.GolobalObject.WorkerPoolSize,
		//设置消息队列长度，也应该从全局变量中获取,消息队列长度应该与go程数量相同
		TaskQueue: make([]chan ziface.IRequest, utils.GolobalObject.WorkerPoolSize),

	}
}
//启动工作池
func (mh *MsgHandle)StartWokerPool(){
	for i:=0;i<int(utils.GolobalObject.WorkerPoolSize);i++{
		//当工作池被启动，则应该启动相应的worker并为管道设置最大接收请求数量
		mh.TaskQueue[i] =make(chan ziface.IRequest,utils.GolobalObject.MaxTaskLen)
		go mh.worker(uint32(i),mh.TaskQueue[i])
	}
}
//单个线程启动，阻塞等待请求
func (mh *MsgHandle)worker(workerId uint32,taskqueue chan ziface.IRequest){
	fmt.Print("Worker Id =",workerId," is starting...\n")
	//应该一直阻塞以执行相应的任务
	for{
	

		req := <-taskqueue
		mh.DoMsgHandler(req)
		
	}
}
func (mh *MsgHandle)SendMessageToQueue(msg ziface.IRequest){
	//将消息负载均衡
	workerid := msg.GetConnection().GetLinkedId()%utils.GolobalObject.WorkerPoolSize
	mh.TaskQueue[workerid]<-msg
	fmt.Print("add msg id to queue:",workerid,"\n")
}
//根据每个请求具体执行相应的方法
func (mh *MsgHandle)DoMsgHandler(zi ziface.IRequest){
	method,exist :=mh.APIS[zi.GetId()]
	if !exist{
		fmt.Print("不存在该类方法\n")
	}
	method.PreHandle(zi)
	method.Handle(zi)
	method.PostHandle(zi)

}
//为每个具体的类型id添加相应的方法router,
func (mh *MsgHandle)AddMsgHandler(id uint32,zi ziface.IRouter){

	_,exist := mh.APIS[id]
	if exist{
		panic("repeated add!")
	}
	mh.APIS[id]=zi
	fmt.Print("加入新路由id:",id,"成功\n")

}