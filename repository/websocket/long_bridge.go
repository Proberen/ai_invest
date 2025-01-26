package websocket

import (
	"ai_invest/conf/logs"
	"context"
	"github.com/longportapp/openapi-go/config"
	"github.com/longportapp/openapi-go/quote"
	"time"
)

type LongBridgeWS interface {
	//GetHistoryKLine 获取历史k线
	GetHistoryKLine(ctx context.Context,
		symbol string,
		period quote.Period,
		adjustType quote.AdjustType,
		startTime time.Time,
		endTime time.Time,
	) ([]quote.Candlestick, error)
}

type longBridgeWS struct{}

func NewLongBridge() LongBridgeWS {
	return &longBridgeWS{}
}

func (l *longBridgeWS) GetHistoryKLine(ctx context.Context,
	symbol string,
	period quote.Period,
	adjustType quote.AdjustType,
	startTime time.Time,
	endTime time.Time) ([]quote.Candlestick, error) {
	allCandlesticks := make([]quote.Candlestick, 0)

	conf, _ := config.New()
	quoteContext, err := quote.NewFromCfg(conf)
	if err != nil {
		logs.Error("[GetHistoryKLine] NewFromCfg err:%+v", err)
		return allCandlesticks, nil
	}
	defer quoteContext.Close()

	s, err := quoteContext.HistoryCandlesticksByDate(ctx, symbol, period, adjustType, &startTime, &endTime)
	if err != nil {
		logs.Error("[GetHistoryKLine] HistoryCandlesticksByDate err:%+v", err)
		return allCandlesticks, err
	}
	for _, v := range s {
		logs.Error("[GetHistoryKLine] res:%+v", *v)
		allCandlesticks = append(allCandlesticks, *v)
	}

	return allCandlesticks, nil
}
