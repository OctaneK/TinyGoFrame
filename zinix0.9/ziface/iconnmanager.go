package ziface

type IConnmanager interface{
	//返回当前链接个数
	Len()int
	//增加一个链接
	AddConn(IConnection)
	//删除一个链接
	DelCon(IConnection)
	//清除所有链接
	CleanCon()
	//通过链接id返回相应的链接
	GetConn(uint32)(IConnection,error)
}