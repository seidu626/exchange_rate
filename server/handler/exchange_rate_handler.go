package handler

import (
	exchange_rate "github.com/seidu626/exchange_rate/proto"
	"github.com/seidu626/exchange_rate/server/services"
	"go.uber.org/zap"
)

type ExchangeRateServerHandler struct {
	logger          *zap.Logger
	currencyService *services.CurrencyService
	exchange_rate.UnimplementedExchangeRatesServer
}

func (_this *ExchangeRateServerHandler) mustEmbedUnimplementedExchangeRatesServer() {
	//TODO implement me
	panic("implement me")
}

func NewExchangeRateServerHandler(logger *zap.Logger, service *services.CurrencyService) *ExchangeRateServerHandler {
	return &ExchangeRateServerHandler{logger: logger, currencyService: service}
}
