package test

import (
	"ai_invest/conf/config"
	"ai_invest/conf/logs"
	"ai_invest/container"
	"context"
	"github.com/longportapp/openapi-go/quote"
	"testing"
	"time"
)

func TestName(t *testing.T) {
	config.InitConfig()
	c, _ := container.Init()
	logs.InitLogger()
	ctx := context.Context(context.Background())
	c.LongBridgeWSClient.GetHistoryKLine(ctx, "AAPL.US", quote.PeriodDay, quote.AdjustTypeForward, time.Date(2024, time.December, 1, 0, 0, 0, 0, time.UTC), time.Now())
}
