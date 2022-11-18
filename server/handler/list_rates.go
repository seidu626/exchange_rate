package handler

import (
	"context"
	exchange_rate "github.com/seidu626/exchange_rate/proto"
	"go.uber.org/zap"
	"time"
)

func (_this *ExchangeRateServerHandler) ListRates(ctx context.Context, request *exchange_rate.ListRatesRequest) (*exchange_rate.ListRatesResponse, error) {
	_this.logger.Info("failed to fetch URL", // Structured context as strongly typed Field values.
		zap.String("base", request.Base), zap.Int("attempt", 3), zap.Duration("backoff", time.Second))

	result, err := _this.currencyService.GetCurrencyRates(request.Base)
	if err != nil {
		return nil, err
	}

	var data []*exchange_rate.CurrencyRate
	for k, v := range *result {
		data = append(data, &exchange_rate.CurrencyRate{Destination: k, Rate: float32(v), Date: _this.currencyService.GetDate().String()})
	}

	return &exchange_rate.ListRatesResponse{Base: request.Base, Date: _this.currencyService.GetDate().String(), Data: data}, nil
}
