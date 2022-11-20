package handler

import (
	"context"
	"fmt"
	"github.com/go-redis/cache/v8"
	exchange_rate "github.com/seidu626/exchange_rate/proto"
	"go.uber.org/zap"
	"time"
)

func (_this *ServerHandler) GetRate(ctx context.Context, request *exchange_rate.CurrencyRateRequest) (*exchange_rate.CurrencyRateResponse, error) {
	_this.logger.Info("GetRate", zap.String("base", request.Base), zap.String("base", request.Destination))

	cacheKey := fmt.Sprintf("%s_%s", request.Base, request.Destination)
	rate := &exchange_rate.CurrencyRateResponse{}
	err := _this.cache.Get(ctx, cacheKey, rate)
	if err == nil {
		return rate, err
	}
	_this.logger.Info("Cache Error", zap.Error(err))

	result, err := _this.currencyService.GetRate(request.Base, request.Destination)
	if err != nil {
		_this.logger.Error(err.Error(), zap.Error(err))
		return nil, err
	}

	rate = &exchange_rate.CurrencyRateResponse{Rate: &exchange_rate.CurrencyRate{Destination: request.Destination, Rate: float32(result.Rate), Inverse: float32(result.Inverse), Date: result.Date.String()}}
	if err := _this.cache.Set(&cache.Item{
		Ctx:   ctx,
		Key:   cacheKey,
		Value: rate,
		TTL:   5 * time.Minute,
	}); err != nil {
		_this.logger.Error("Cache Error", zap.Error(err))
	}

	return rate, nil
}
