package main

import (
	"src/service/dan"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
)

const (
	NetworkAddr = "127.0.0.1:7791"
)

type HelloImpl struct {
}
func NewHello() *HelloImpl {
	return &HelloImpl{}
}
func (this *HelloImpl) HelloVoid() (err error){
	fmt.Println("hello")
	return
}

func (this *HelloImpl) HelloBoolean(para bool) (p bool,err error){
	return false,err
}

func (this *HelloImpl) HelloString(para string) (p string ,err error){
	fmt.Println("hello")
	return
}

func (this *HelloImpl) HelloInt(para int32) (p int32 ,err error){
	fmt.Println("hello123")
	return para,err
}

func (this *HelloImpl) HelloNull() (p string,err error){
	fmt.Println("hello")
	return
}

func main() {//TProtocolFactory
	var addr string;
	addr = "127.0.0.1:7911";
	/*
	var protocolFactory thrift.TProtocolFactory
	protocolFactory = thrift.NewTCompactProtocolFactory()
	var transportFactory thrift.TTransportFactory
	transportFactory = thrift.NewTTransportFactory()
	*/
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	if err := runServer(transportFactory, protocolFactory, addr); err != nil {
		fmt.Println("error running server:", err)
	}
}

func runServer(transportFactory thrift.TTransportFactory, protocolFactory thrift.TProtocolFactory, addr string) error {
	var transport thrift.TServerTransport
	var err error
	transport, err = thrift.NewTServerSocket(addr)

	if err != nil {
		return err
	}
	fmt.Printf("%T\n", transport)
	handler := NewHello()
	//handler.HelloVoid()
	processor := dan.NewHelloProcessor(handler)
	server := thrift.NewTSimpleServer4(processor, transport, transportFactory, protocolFactory)

	fmt.Println("Starting the simple server... on ", addr)
	return server.Serve()
}