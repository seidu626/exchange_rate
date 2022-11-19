package ratesClient

import (
	exchange_rate "github.com/seidu626/exchange_rate/proto"
	"go.uber.org/zap"
)

type ClientHandler struct {
	logger *zap.Logger
	exchange_rate.ExchangeRatesClient
}

func (_this *ClientHandler) NewExchangeRatesClientHandler(logger *zap.Logger) *ClientHandler {
	return &ClientHandler{logger: logger}
}
