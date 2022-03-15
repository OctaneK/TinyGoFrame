package zinet

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"ziMod/utils"
	"ziMod/ziface"
)

type DataPack struct{}
//该函数只是用于创建一个实例，借助这个实例调用拆包封包方法
func NewDataPack()*DataPack{
	return &DataPack{}
}

func (dp *DataPack)GetHeadLen()uint32{
	return 8//四字节长度，四字节类型
}
//提供一个封包方法,将上层的消息封装成包
func (dp *DataPack)Pack(msg ziface.IMessage)([]byte,error){
	//创建一个新缓冲
	dataBuff := bytes.NewBuffer([]byte{})
	//依次将消息长度，消息类型，消息内容封装到一起
	//向数组中以小端方式写入数据
	err := binary.Write(dataBuff,binary.LittleEndian,msg.GetMsgLen())
	if err !=nil{
		return nil,err
	}
	er := binary.Write(dataBuff,binary.LittleEndian,msg.GetMsgId())
	if er !=nil{
		return nil,er
	}
	e := binary.Write(dataBuff,binary.LittleEndian,msg.GetData())
	if e !=nil{
		return nil,e
	}
	return dataBuff.Bytes(),nil


}
//提供一个拆包方法，将传来的字节流拆出具体消息
func (dp *DataPack)UnPack(by []byte)(ziface.IMessage,error){
	//依然是创建一个缓冲，不过这次创建的是一个读缓冲
	databuff := bytes.NewReader(by)
	//创建存放信息的message类
	msg := &Message{}
	//按照封包的过程依次将信息读到相应的message位置
	err := binary.Read(databuff,binary.LittleEndian,&msg.DataLen)//注意这里要将data的地址传入
	if err !=nil{
		fmt.Print("binary.Read(databuff,binary.LittleEndian,msg.DataLen) error\n")
		return nil,err
	}
	er := binary.Read(databuff,binary.LittleEndian,&msg.Id)
	if er !=nil{
		fmt.Print("er := binary.Read(databuff,binary.LittleEndian,msg.Id)\n")
		return nil,er
	}
	//如果规定了包的大小，并且读出来的包长度大于给定的值，则抛出异常
	if utils.GolobalObject.MaxPackageSize>0 && utils.GolobalObject.MaxPackageSize<msg.DataLen{
		fmt.Print("utils.GolobalObject.MaxPackageSize>0 && utils.GolobalObject.MaxPackageSize<msg.DataLen\n")
		return nil,errors.New("package is too large")
	}
	return msg,nil//返回消息，后续只需将数据读出到msg之中即可

}