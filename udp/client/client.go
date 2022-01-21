package main

import (
	"fmt"
	"net"
)

func main() {
	dial, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 12000,
	})
	if err != nil {
		fmt.Println("dial udp err: ", err)
		return
	}

	defer dial.Close()
	sendData := []byte("hello world")
	_, err = dial.Write(sendData)
	if err != nil {
		fmt.Println("write data failed, err: ", err)
		return
	}

	data := make([]byte, 4086)
	n, addr, err := dial.ReadFromUDP(data)
	if err != nil {
		fmt.Println("read data from udp failed, err: ", err)
		return
	}
	fmt.Printf("recv:%v, addr:%v, count:%v\n", string(data[:n]), addr, n)
}
