package zinet

import (
	"fmt"
	"ziMod/ziface"
)
type MsgHandle struct{
	APIS map[uint32] ziface.IRouter

}
func NewMsgHandle()*MsgHandle{
	return &MsgHandle{APIS :make(map[uint32]ziface.IRouter)}
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