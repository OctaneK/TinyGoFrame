package main

import (
	"fmt"
	"net"
	"time"
)


func main(){
	fmt.Print("connecting...")
	conn,err := net.Dial("tcp","127.0.0.1:8888")
	if err!=nil{
		fmt.Print("connected")
		return
	}
	for{
		_,e := conn.Write([]byte("hello server\n"))
		if e!=nil{
			fmt.Print("failed write")
		}
		time.Sleep(time.Second*3)
		buf := make([]byte,512)
		cnt ,er := conn.Read(buf)
		if er!=nil{
			fmt.Print("failed read")
			return
		}
		fmt.Printf("server callback write:%s%d",buf,cnt)
		time.Sleep(time.Second*3)


	}
}