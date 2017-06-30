package main

import (
"flag"
"fmt"
"git.apache.org/thrift.git/lib/go/thrift"
"os"
	"service/dan"
	"crypto/tls"
)

func Usage() {
	fmt.Fprint(os.Stderr, "Usage of ", os.Args[0], ":\n")
	flag.PrintDefaults()
	fmt.Fprint(os.Stderr, "\n")
}

func main() {
	flag.Usage = Usage
	server := flag.Bool("client", false, "Run server")
	protocol := flag.String("P", "binary", "Specify the protocol (binary, compact, json, simplejson)")
	framed := flag.Bool("framed", false, "Use framed transport")
	buffered := flag.Bool("buffered", false, "Use buffered transport")
	addr := flag.String("addr", "localhost:9090", "Address to listen to")
	secure := flag.Bool("secure", false, "Use tls secure transport")
	fmt.Println(server)
	flag.Parse()

	var protocolFactory thrift.TProtocolFactory
	switch *protocol {
	case "compact":
		protocolFactory = thrift.NewTCompactProtocolFactory()
	case "simplejson":
		protocolFactory = thrift.NewTSimpleJSONProtocolFactory()
	case "json":
		protocolFactory = thrift.NewTJSONProtocolFactory()
	case "binary", "":
		protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()
	default:
		fmt.Fprint(os.Stderr, "Invalid protocol specified", protocol, "\n")
		Usage()
		os.Exit(1)
	}

	var transportFactory thrift.TTransportFactory
	if *buffered {
		transportFactory = thrift.NewTBufferedTransportFactory(8192)
	} else {
		transportFactory = thrift.NewTTransportFactory()
	}

	if *framed {
		transportFactory = thrift.NewTFramedTransportFactory(transportFactory)
	}
	//fmt.Printf(*addr)
	//fmt.Println(*secure)
	//fmt.Println(protocolFactory)

	runClient(transportFactory, protocolFactory, *addr, *secure)
	/*
	if *server {
		if err := runServer(transportFactory, protocolFactory, *addr, *secure); err != nil {
			fmt.Println("error running server:", err)
		}
	} else {
		if err := runClient(transportFactory, protocolFactory, *addr, *secure); err != nil {
			fmt.Println("error running client:", err)
		}
	}*/

}

func runClient(transportFactory thrift.TTransportFactory, protocolFactory thrift.TProtocolFactory, addr string, secure bool) error {
	var transport thrift.TTransport
	var err error
	if secure {
		cfg := new(tls.Config)
		cfg.InsecureSkipVerify = true
		transport, err = thrift.NewTSSLSocket(addr, cfg)
	} else {
		transport, err = thrift.NewTSocket(addr)
	}
	if err != nil {
		fmt.Println("Error opening socket:", err)
		return err
	}
	transport, err = transportFactory.GetTransport(transport)
	if err != nil {
		return err
	}
	defer transport.Close()
	if err := transport.Open(); err != nil {
		return err
	}
	client:=dan.NewHelloClientFactory(transport, protocolFactory);
	client.HelloVoid()
	return err
}
