package long_bridge

import (
	"ai_invest/container"
)

type ILongBridgeEntity interface {
}

func NewLongBridgeEntity(c *container.Container) ILongBridgeEntity {
	return &LongBridgeEntity{
		c: c,
	}
}
