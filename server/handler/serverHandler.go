package handler

import (
	"github.com/go-redis/cache/v8"
	exchange_rate "github.com/seidu626/exchange_rate/proto"
	"github.com/seidu626/exchange_rate/server/services"
	"go.uber.org/zap"
	"time"
)

type PollRates struct {
	ticker *time.Ticker
	add    chan *exchange_rate.ListRatesResponse // new rates channel
}

type ServerHandler struct {
	logger          *zap.Logger
	cache           *cache.Cache
	currencyService *services.CurrencyService
	subscriptions   *map[exchange_rate.ExchangeRates_SubscriptionServer]*[]exchange_rate.ListRatesRequest
	exchange_rate.UnimplementedExchangeRatesServer
	*PollRates
}

func (_this *ServerHandler) mustEmbedUnimplementedExchangeRatesServer() {
	//TODO implement me
	panic("implement me")
}

func NewExchangeRateServerHandler(logger *zap.Logger, cache *cache.Cache, service *services.CurrencyService, pollingInterval time.Duration) *ServerHandler {
	server := &ServerHandler{logger: logger, cache: cache, currencyService: service,
		subscriptions: &map[exchange_rate.ExchangeRates_SubscriptionServer]*[]exchange_rate.ListRatesRequest{},
		PollRates: &PollRates{
			ticker: time.NewTicker(pollingInterval),
			add:    make(chan *exchange_rate.ListRatesResponse),
		}}
	go server.StartPolling()
	return server
}
