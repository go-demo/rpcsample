package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	server := rpc.NewServer()
	err := server.Register(new(Foo))
	if err != nil {
		log.Fatal(err)
	}
	l, err := net.Listen("tcp", ":9234")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("RPC Server is running at :9234.")
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print("rpc.Serve: accept:", err.Error())
			return
		}
		go server.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}

type Foo struct {
}

func (f *Foo) Hello(name string, reply *string) error {
	*reply = fmt.Sprintf("Hello,%s!", name)
	return nil
}
