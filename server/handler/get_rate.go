package handler

import (
	"context"
	exchange_rate "github.com/seidu626/exchange_rate/proto"
	"go.uber.org/zap"
	"time"
)

func (_this *ExchangeRateServerHandler) GetRate(ctx context.Context, request *exchange_rate.CurrencyRateRequest) (*exchange_rate.CurrencyRateResponse, error) {
	_this.logger.Info("failed to fetch URL", // Structured context as strongly typed Field values.
		zap.String("base", request.Base), zap.String("base", request.Destination), zap.Int("attempt", 3), zap.Duration("backoff", time.Second))

	result, err := _this.currencyService.GetRate(request.Base, request.Destination)
	if err != nil {
		_this.logger.Error(err.Error(), zap.Error(err))
		return nil, err
	}

	rate := &exchange_rate.CurrencyRateResponse{Rate: &exchange_rate.CurrencyRate{Destination: request.Destination, Rate: float32(result.Rate), Inverse: float32(result.Inverse), Date: result.Date.String()}}

	return rate, nil
}
