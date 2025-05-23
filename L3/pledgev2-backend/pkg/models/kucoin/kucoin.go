package kucoin

import (
	"context"
	"github.com/Kucoin/kucoin-go-sdk"
	"pledgev2-backend/log"
	"pledgev2-backend/pkg/db"
)

const ApiKeyVersionV2 = "2"

var PlgrPrice = "0.0027"
var PlgrPriceChan = make(chan string, 2)

func GetExchangePrice() {
	log.Logger.Sugar().Info("GetExchangePrice")
	price, err := db.RedisGetString("plgr_price")
	if err != nil {
		log.Logger.Sugar().Error(err)
	} else {
		PlgrPrice = price
	}

	// 初始化APiService
	s := kucoin.NewApiService(
		kucoin.ApiKeyOption("key"),
		kucoin.ApiSecretOption("secret"),
		kucoin.ApiPassPhraseOption("passphrase"),
		kucoin.ApiKeyVersionOption(ApiKeyVersionV2),
	)

	// 获取WS 公共token
	rsp, err := s.WebSocketPublicToken(context.Background())
	if err != nil {
		log.Logger.Error(err.Error())
		return
	}

	// 结果反序列化,存到tk
	tk := &kucoin.WebSocketTokenModel{}
	if err := rsp.ReadData(tk); err != nil {
		log.Logger.Error(err.Error())
		return
	}

	// 使用token创建WSC
	wsClient := s.NewWebSocketClient(tk)
	// mc 接收成功的业务消息
	// ec 接收连接/订阅错误
	// err 错误
	mc, ec, err := wsClient.Connect()
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}

	// 构造订阅和退订消息
	ch := kucoin.NewSubscribeMessage("/market/ticker:PLGR-USDT", false)
	uch := kucoin.NewUnsubscribeMessage("/market/ticker:PLGR-USDT", false)

	// 发送订阅消息到KuCoin
	if err := wsClient.Subscribe(ch); err != nil {
		log.Logger.Error(err.Error())
		return
	}

	for {
		select {
		case err := <-ec:
			// 收到错误,终止WS
			wsClient.Stop()
			log.Logger.Sugar().Error(err)
			_ = wsClient.Unsubscribe(uch)
			return
		case msg := <-mc:
			// 收到消息
			// 反序列化消息,保存到t
			t := &kucoin.TickerLevel1Model{}
			if err := msg.ReadData(t); err != nil {
				log.Logger.Sugar().Errorf("msg.ReadData: %s", err.Error())
				return
			}
			// 向PlgrPriceChan推送新价格
			PlgrPriceChan <- t.Price
			// 更新价格
			PlgrPrice = t.Price
			// redis缓存
			_ = db.RedisSetString("plgr_price", PlgrPrice, 0)
		}
	}
}
