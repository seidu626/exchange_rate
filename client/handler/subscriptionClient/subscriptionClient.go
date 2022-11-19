package subscriptionClient

import (
	exchange_rate "github.com/seidu626/exchange_rate/proto"
	"go.uber.org/zap"
)

type Handler struct {
	logger *zap.Logger
	exchange_rate.ExchangeRates_SubscriptionClient
}
