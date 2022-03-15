package main
import "ziMod/zinet"


func main(){
	server := zinet.NewServer("SERVER")
	server.Serve()
}