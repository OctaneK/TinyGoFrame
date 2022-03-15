package zinet

import "ziMod/ziface"
type BaseRouter struct{

}
func (br *BaseRouter)PreHandle(request ziface.IRequest){

}
//正式处理请求
func (br *BaseRouter)Handle(request ziface.IRequest){

}
//处理请求之后进行收尾工作
func (br *BaseRouter)PostHandle(request ziface.IRequest){
	
}