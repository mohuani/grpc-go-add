package mygrpc

import (
	"context"
	"fmt"
)

type MyService struct {
}

func (s *MyService) Add(ctx context.Context, request *AddRequest) (*AddReply, error) {
	fmt.Println(request.String())
	return &AddReply{
		Res: myAdd(request.A, request.B),
	}, nil
}

func myAdd(a int32, b int32) int32 {
	return a + b
}
