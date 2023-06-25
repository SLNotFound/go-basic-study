package main

import (
	"fmt"
	"helloworld/rpc/demo2/api"
	"log"
)

type HelloService struct {
}

func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "hello:" + request
	return nil
}

func main() {
	client, err := api.DialHelloService("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing: ", err)
	}

	var reply string
	err = client.Hello("hello", &reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply)
}
