package main

import (
	"fmt"
	"google.golang.org/grpc"
	"grpc-go/mygrpc"
	"net"
)

func main() {

	listen, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("listen failed...")
	}

	myService := mygrpc.MyService{}

	server := grpc.NewServer()

	mygrpc.RegisterAddServiceServer(server, &myService)

	_ = server.Serve(listen)

}
