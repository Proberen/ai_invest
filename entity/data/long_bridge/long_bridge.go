package long_bridge

import (
	"ai_invest/conf"
	"ai_invest/container"
	"github.com/longportapp/openapi-go/config"
	"log"
)

type ILongBridgeEntity interface {
}

func initLongBridgeToken() *config.Config {
	token := conf.GetConfigToken().LongBridge

	conf, err := config.New()
	if err != nil {
		log.Fatal(err)
		return nil
	}
	conf.AppKey = token.AppKey
	conf.AppSecret = token.AppSecret
	conf.AccessToken = token.AccessToken
	conf.Language = "zh-CN"
	conf.Region = config.Region(token.Region)

	return conf
}

func NewLongBridgeEntity(c *container.Container) ILongBridgeEntity {
	return &LongBridgeEntity{
		c: c,
	}
}
