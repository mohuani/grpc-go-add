package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc-go/mygrpc"
)

func main() {

	clientConn, err := grpc.Dial(":9999", grpc.WithInsecure())
	if err != nil {
		fmt.Println("clientConn failed ...")
	}

	addServiceClient := mygrpc.NewAddServiceClient(clientConn)

	addReq := mygrpc.AddRequest{
		A: 13,
		B: 14,
	}
	addReply, err := addServiceClient.Add(context.Background(), &addReq)
	fmt.Println(addReply)

}
