package handler

import (
	exchange_rate "github.com/seidu626/exchange_rate/proto"
	"go.uber.org/zap"
)

type ExchangeRateClientHandler struct {
	logger *zap.Logger
	exchange_rate.ExchangeRatesClient
}

func (_this *ExchangeRateClientHandler) NewExchangeRatesClientHandler(logger *zap.Logger) *ExchangeRateClientHandler {
	return &ExchangeRateClientHandler{logger: logger}
}
