package main

import (
	"ai_invest/container"
	"context"
)

type AgentImpl struct {
	container container.Container
}

func (*AgentImpl) NewAgent(ctx context.Context) {
	return
}
