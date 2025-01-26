package container

import (
	"ai_invest/repository/websocket"
	"github.com/google/wire"
)

var wsProviderSet = wire.NewSet(
	websocket.NewLongBridge,
)
