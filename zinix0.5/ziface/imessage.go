package ziface
//将请求的消息封装到一个message中
type IMessage interface{
	GetMsgId()	uint32
	GetMsgLen() uint32
	GetData() []byte//获取消息内容，有意思的是返回的是切片..切片是指针吗？

	SetMsgId(uint32)
	SetData([]byte)//设置消息内容....同样传递切片
	SetDMsgLen(uint32)

	
}