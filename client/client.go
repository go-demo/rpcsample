package main

import (
	"log"
	"net/rpc/jsonrpc"
)

func main() {
	cli, err := jsonrpc.Dial("tcp", ":9234")
	if err != nil {
		log.Fatal(err)
	}
	var reply string
	md := map[string]string{
		"Foo": "bar",
	}
	err = cli.Call("Foo.HelloMap", md, &reply)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Call result:", reply)
}
