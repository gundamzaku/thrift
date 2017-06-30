package main
import (
	"service/dan"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"

)

func main() {
	var addr string;
	var transportFactory thrift.TTransportFactory
	addr = "localhost:9090";
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	//transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	/*
	之前问题一直调不通，原因就出在这个上面，其实用下面的方式来赋值transportFactory
	我不知道为什么用了上面的方式，结果导致请求堵塞。
	 */
	transportFactory = thrift.NewTTransportFactory()
	var err error

	if(err!=nil){
		fmt.Println("something error")
	}

	if err := runClient(transportFactory, protocolFactory, addr); err != nil {
		fmt.Println("error running client:", err)
	}

}

func runClient(transportFactory thrift.TTransportFactory, protocolFactory thrift.TProtocolFactory, addr string) error {
	var transport thrift.TTransport
	var err error

	transport, err = thrift.NewTSocket(addr)

	if err != nil {
		fmt.Println("Error opening socket:", err)
		return err
	}
	transport,err = transportFactory.GetTransport(transport)
	defer transport.Close()
	if err := transport.Open(); err != nil {
		return err
	}
	client := dan.NewHelloClientFactory(transport, protocolFactory)
	client.HelloVoid()
	rs,err := client.HelloInt(50)
	fmt.Println(rs)
	return err
}