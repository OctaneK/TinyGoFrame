package zinet

import "ziMod/ziface"

type Request struct{
	conn ziface.IConnection//可以理解为一根虚函数指针
	data []byte
}
//返回另一个抽象基类的虚函数指针，通过这个虚函数指针再去调用其派生类的方法
func (re *Request)GetConnection()ziface.IConnection{
	return re.conn
}
func (re *Request)GetData()[]byte{
	return re.data
}
