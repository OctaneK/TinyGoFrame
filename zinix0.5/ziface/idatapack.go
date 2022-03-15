package ziface
//解决TCP粘包问题
//为message提供TLV封包拆包的服务
//拆包分为两阶段，首先进行固定长度的长度和类读取，然后根据读取到的长度再进行一次读，将实际的数据读取出来
type IDataPack interface{
	//获取包头长度
	GetHeadLen()uint32
	//提供一个封包方法,将上层的消息封装成包
	Pack(IMessage)([]byte,error)
	//提供一个拆包方法，将传来的字节流拆出具体消息
	UnPack([]byte)(IMessage,error)
}