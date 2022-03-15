package main

import (
	"fmt"
	"io"
	"net"
	"time"

	//"ziMod/ziface"
	"ziMod/zinet"
)


func main(){
	fmt.Print("connecting...")
	conn,err := net.Dial("tcp","127.0.0.1:8888")
	if err!=nil{
		fmt.Print("connected")
		return
	}
	for{
		/*
		_,e := conn.Write([]byte("hello server\n"))
		if e!=nil{
			fmt.Print("failed write")
		}
		//time.Sleep(time.Second*3)
		buf := make([]byte,512)
		cnt ,er := conn.Read(buf)
		if er!=nil{
			fmt.Print("failed read")
			return
		}
		fmt.Printf("server callback write:%s%d",buf,cnt)
		
		*/
		dp := zinet.NewDataPack()
		binmsg,err := dp.Pack(zinet.NewMessage(1,[]byte("hello")))
		if err!=nil{
			fmt.Print("pack error")
			return 
		}
		conn.Write(binmsg)
		
		fmt.Print("已发送信息hello\n")
		//读出包头
		headdata :=make([]byte,dp.GetHeadLen())
		_,e :=io.ReadFull(conn,headdata)
		if e!=nil{
			fmt.Print("readfull head error")
			return
		}

		msg,err :=dp.UnPack(headdata)
		if err!=nil{
			fmt.Print("unpack error")
			return
		}
		message := msg.(*zinet.Message)//将基类指针转为派生类指针，这真是太有意思了...
		if err!=nil{
			fmt.Print("dp.unPack failed")
		}
		if msg.GetMsgLen()>0{
			//同样过程，创建缓冲区并将剩余消息封装到message中
			message.Data =make([]byte, message.DataLen)
			_,er:=io.ReadFull(conn,message.Data)
			if er !=nil{
				fmt.Print("readfull message.data error")
			}
		}
		fmt.Print("recv information ID:",msg.GetMsgId(),"\ninformation:",string(msg.GetData()),"\n")
		

		time.Sleep(time.Second*3)

	}
}