package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"ziMod/ziface"
)

//用于配置zinix框架的全局信息，一些信息可以通过zinix.json文件进行修改
type GolobalObj struct{
	TcpServer ziface.Iserver

	Host string
	TcpPort int
	Name string

	Version string
	MaxConn int
	MaxPackageSize uint32
	WorkerPoolSize uint32//工作池的数量
	MaxTaskLen uint32//最大工作队列数量

}
func (g *GolobalObj)Reload(){
	data,err := ioutil.ReadFile("conf/zinix.json")
	if err!=nil{
		fmt.Print("ReadFile error")
		panic(err)
	}
	json.Unmarshal(data,&GolobalObject)
}
//从json文件配置
var GolobalObject *GolobalObj
func init(){//GO语言特性，在导包的过程中就会调用init
	GolobalObject = &GolobalObj{
		Host: "127.0.0.1",
		TcpPort: 8888,
		Name: "zinixDemoServer",
		Version: "tcp4",
		MaxConn: 1000,
		MaxPackageSize: 512,
		WorkerPoolSize: 8,
		MaxTaskLen: 1024,
	}
	GolobalObject.Reload()
	fmt.Print("setting finished\n")
}