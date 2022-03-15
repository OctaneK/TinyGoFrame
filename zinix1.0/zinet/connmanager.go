package zinet

import (
	"errors"
	"fmt"
	"sync"
	"ziMod/ziface"
)

type ConnManager struct{
	//通过id寻找相应的链接
	Connections map[uint32]ziface.IConnection
	//为操作connections提供锁
	ConnLock sync.RWMutex
}
func NewConnManager()*ConnManager{
	return &ConnManager{
		Connections: make(map[uint32]ziface.IConnection),
	}
}
func (c *ConnManager)AddConn(con ziface.IConnection){
	c.ConnLock.Lock()
	defer c.ConnLock.Unlock()
	c.Connections[con.GetLinkedId()]=con
	
	fmt.Print("成功加入链接 ID：",con.GetLinkedId()," 当前链接总数：",c.Len(),"\n")
}
func (c *ConnManager)DelCon(con ziface.IConnection){
	c.ConnLock.Lock()
	defer c.ConnLock.Unlock()
	delete(c.Connections,con.GetLinkedId())
	fmt.Print("成功删除链接 ID：",con.GetLinkedId()," 当前链接总数：",c.Len(),"\n")
}
func (c *ConnManager)GetConn(id uint32)(ziface.IConnection,error){
	c.ConnLock.RLock()
	defer c.ConnLock.RUnlock()
	
	conn ,ok:= c.Connections[id]
	if !ok {
		return nil,errors.New("该链接不存在")
	}else{
		return conn,nil
	}

}
func (c *ConnManager)CleanCon(){
	for id,con := range c.Connections{
		//关闭各个链接的资源，并将其从链接模块中删除
		con.Stop()
		delete(c.Connections,id)
	}
	fmt.Print("已删除所有连接\n")
}
func (c *ConnManager)Len()int{
	return len(c.Connections)
}