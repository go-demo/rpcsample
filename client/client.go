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
	err = cli.Call("Foo.Hello", "Lyric", &reply)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Call result:", reply)
}
