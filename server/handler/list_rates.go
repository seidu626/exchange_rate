package handler

import (
	"context"
	"github.com/go-redis/cache/v8"
	exchange_rate "github.com/seidu626/exchange_rate/proto"
	"go.uber.org/zap"
	"time"
)

func (_this *ServerHandler) ListRates(ctx context.Context, request *exchange_rate.ListRatesRequest) (*exchange_rate.ListRatesResponse, error) {
	_this.logger.Info("failed to fetch URL", // Structured context as strongly typed Field values.
		zap.String("base", request.Base), zap.Int("attempt", 3), zap.Duration("backoff", time.Second))

	cacheKey := "__" + request.Base + "__"
	rates := &exchange_rate.ListRatesResponse{}
	err := _this.cache.Get(ctx, cacheKey, rates)
	if err == nil {
		return rates, err
	}
	_this.logger.Info("Cache Error", zap.Error(err))

	result, err := _this.currencyService.GetCurrencyRates(request.Base)
	if err != nil {
		return nil, err
	}

	var data []*exchange_rate.CurrencyRate
	for k, v := range *result {
		data = append(data, &exchange_rate.CurrencyRate{Destination: k, Rate: float32(v), Date: _this.currencyService.GetDate().String()})
	}

	rates = &exchange_rate.ListRatesResponse{Base: request.Base, Date: _this.currencyService.GetDate().String(), Data: data}
	if err := _this.cache.Set(&cache.Item{
		Ctx:   ctx,
		Key:   cacheKey,
		Value: rates,
		TTL:   1 * time.Hour,
	}); err != nil {
		_this.logger.Error("Cache Error", zap.Error(err))
	}
	return rates, nil
}
