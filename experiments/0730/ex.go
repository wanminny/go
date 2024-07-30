package main

import (
	"log"
	"net"
)

func main() {
	udpA, err := net.ResolveUDPAddr("udp", "8.8.8.8:6677")
	if err != nil {
		log.Println(err.Error())
	}
	log.Println(udpA)
	tcpA, err := net.ResolveTCPAddr("tcp", "192.168.0.1:5432")
	if err != nil {
		log.Println(err.Error())
	}
	log.Println(tcpA)
	log.Println(tcpA.String())

}
