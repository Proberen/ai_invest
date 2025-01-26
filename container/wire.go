//go:build wireinject
// +build wireinject

package container

import (
	"ai_invest/repository/websocket"
	"github.com/google/wire"
)

type Container struct {
	LongBridgeWSClient websocket.LongBridgeWS
}

func (c *Container) GetLongBridgeWSClient() websocket.LongBridgeWS {
	return c.LongBridgeWSClient
}

func Init() (*Container, error) {
	wire.Build(
		wire.Struct(new(Container), "*"),
		wsProviderSet,
	)
	return &Container{}, nil
}
