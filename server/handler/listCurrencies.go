package handler

import (
	"context"
	"fmt"
	"github.com/go-redis/cache/v8"
	exchange_rate "github.com/seidu626/exchange_rate/proto"
	"github.com/seidu626/exchange_rate/server/services"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"time"
)

func (_this *ServerHandler) ListCurrencies(ctx context.Context, request *exchange_rate.ListCurrencyRequest) (*exchange_rate.ListCurrencyResponse, error) {
	_this.logger.Info("ListCurrencies")
	cacheKey := fmt.Sprintf("ListCurrencies_PageIndex:%v_PageSize:%v", request.PageIndex, request.PageSize)
	currencies := &exchange_rate.ListCurrencyResponse{}
	err := _this.cache.Get(ctx, cacheKey, currencies)
	if err == nil {
		return currencies, err
	}
	_this.logger.Info("Cache Error", zap.Error(err))

	result, err := _this.currencyService.ListCurrencies(&services.ListCurrencyRequest{PageIndex: request.PageIndex.GetValue(), PageSize: request.PageSize.GetValue()})
	if err != nil {
		_this.logger.Error(err.Error(), zap.Error(err))
	}
	var data []*exchange_rate.Currency
	for _, s := range *result.Data {
		data = append(data, &exchange_rate.Currency{Name: s.Name, Description: s.Description})
	}

	currencies = &exchange_rate.ListCurrencyResponse{Data: data, PageIndex: wrapperspb.Int32(result.PageIndex), PageSize: wrapperspb.Int32(result.PageSize), Total: wrapperspb.Int32(result.Total)}

	if err := _this.cache.Set(&cache.Item{
		Ctx:   ctx,
		Key:   cacheKey,
		Value: currencies,
		TTL:   1 * time.Hour,
	}); err != nil {
		_this.logger.Error("Cache Error", zap.Error(err))
	}

	return currencies, nil
}
