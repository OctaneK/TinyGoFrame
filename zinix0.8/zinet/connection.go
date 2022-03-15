package zinet

import (
	"errors"
	"fmt"
	"io"
	"net"

	"ziMod/ziface"
)

type Connection struct {
	//链接绑定的socket
	Conn *net.TCPConn
	//链接的pid
	ConnectionID uint32
	//当前链接的状态
	isClose bool
	//链接绑定的方法
	//HandleAPI ziface.HandleFunc
	//告知是否关闭的channel
	ExitChan chan bool
	Msgchan  chan []byte

	//用于导向的路由模块集合

	MsgHandle ziface.IMsgHandle
}

func NewConnection(con *net.TCPConn, id uint32, msgHandle ziface.IMsgHandle) *Connection {
	conn := &Connection{ //与上面的con重名，直接编译器报错
		Conn:         con,
		ConnectionID: id,
		isClose:      false,
		//HandleAPI: callback,
		MsgHandle: msgHandle,

		ExitChan: make(chan bool, 1),
		Msgchan:  make(chan []byte),
	}
	return conn
}

//暂时读写合在一起，读业务启动也会让写业务同样启动
func (c *Connection) ReadMessage() {
	fmt.Print("start reading for connectionID:", c.ConnectionID, "\n")
	defer c.Stop()
	defer fmt.Print("server is closeing...\n")

	for {
		/*
			buf := make([]byte,512)
			_,err :=c.Conn.Read(buf)
			if err!=nil{
				fmt.Print("read error")
				break
			}*/
		//获取一个用于拆包封包的对象
		dp := NewDataPack()
		//首先读取头部，然后再根据长度读取数据部分
		headata := make([]byte, dp.GetHeadLen())
		_, err := io.ReadFull(c.GetTcpConn(), headata)
		if err != nil {
			fmt.Print("first readfull error")
			break
		}
		//拆包，但只将头部读取，存入id与len字段，数据部分等到后面第二次读出

		msg, er := dp.UnPack(headata)

		if er != nil {
			fmt.Print("unpack error")

		}
		message := msg.(*Message)
		//第二次读取，这一次是真正将数据读入msg的数据部分
		message.Data = make([]byte, msg.GetMsgLen())
		if msg.GetMsgLen() > 0 {
			_, e := io.ReadFull(c.GetTcpConn(), message.Data)
			if e != nil {
				fmt.Print("second readfull error")
				break
			}

		}

		fmt.Print("已读取到消息，将执行业务\n")
		/*
			er :=c.HandleAPI(c.Conn,buf,cnt)
			if er !=nil{
				fmt.Print(c.ConnectionID,"handleFunc error")
				break;
			}*/
		req := Request{
			conn: c, //有意思的地方来了，request类第一个成员是一个抽象，但是这里c是该抽象的派生实体
			msg:  msg,
		}
		//从每一个任务开启一个go程转变为将每一个任务移交给消息队列
		c.MsgHandle.SendMessageToQueue(&req)
		//go c.MsgHandle.DoMsgHandler(&req)

	}
}

//分离的写线程，链接模块启动时将挂载
func (c *Connection) WriteMessage() {
	fmt.Print("写模块已启动\n")
	defer fmt.Print(c.Conn.RemoteAddr().String(), "写模块已退出")
	//阻塞，不断向客户端写入数据
	for {
		select {
		case data := <-c.Msgchan:
			_, err := c.Conn.Write(data) //这个发送的消息是已经序列化好的消息
			if err != nil {
				fmt.Print("c.Conn.Write(data) error")
				return
			}
		case <-c.ExitChan:
			fmt.Print("读模块已经退出，写模块即将退出\n")
			return

		}

	}

}

//将要发送的数据TLV序列化后发送
func (c *Connection) SendMessage(id uint32, by []byte) error {
	if c.isClose {

		return errors.New("connection closed can not send message")
	}
	//获取打包实例
	dp := NewDataPack()
	//实际上通过message方法已经将数据打包成一个message，pack方法是为了该message转换成二进制流
	binbuf, err := dp.Pack(NewMessage(id, by))
	if err != nil {
		fmt.Print("Pack error")
	}
	/*
		_, e :=c.Conn.Write(binbuf)
		if e !=nil{
			return errors.New("write error")
		}
		fmt.Print("已成功发送数据\n")
	*/
	//从发送给客户端到发送给管道
	c.Msgchan <- binbuf

	return nil

}

//启动链接,让当前链接开始工作
func (c *Connection) Start() {
	fmt.Print("Conn start....ConnID", c.ConnectionID)
	//启动从当前链接读数据的业务
	go c.ReadMessage()
	//启动从当前链接写的业务
	go c.WriteMessage()

}



//关闭链接
func (c *Connection) Stop() {
	fmt.Print("connection closed ID :", c.GetLinkedId())
	if c.isClose {
		fmt.Print("has closed")
	}
	//关闭写模块
	c.ExitChan <- true

	c.isClose = true
	c.Conn.Close() //关闭连接的资源

	close(c.ExitChan)
	close(c.Msgchan)
}

//获取套接字conn
func (c *Connection) GetTcpConn() *net.TCPConn {
	return c.Conn
}

//获取套接字ip：port
func (c *Connection) GetAddr() net.Addr { //注意这里返回的是地址，而不是指针

	return c.Conn.RemoteAddr()
}

//获取链接id
func (c *Connection) GetLinkedId() uint32 {
	return c.ConnectionID

}
