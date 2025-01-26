package main

import (
	"ai_invest/container"
	rpc "ai_invest/kitex_gen/rpc"
	"context"
)

type AiInvestRPCServiceImpl struct {
	container *container.Container
}

func NewServiceHandler(container *container.Container) *AiInvestRPCServiceImpl {
	return &AiInvestRPCServiceImpl{container: container}
}

// Test implements the AiInvestRPCServiceImpl interface.
func (s *AiInvestRPCServiceImpl) Test(ctx context.Context, req *rpc.TestRequest) (resp *rpc.TestResponse, err error) {
	// TODO: Your code here...
	return
}
