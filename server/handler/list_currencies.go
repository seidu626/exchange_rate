package handler

import (
	"context"
	exchange_rate "github.com/seidu626/exchange_rate/proto"
	"github.com/seidu626/exchange_rate/server/services"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func (_this *ExchangeRateServerHandler) ListCurrencies(ctx context.Context, request *exchange_rate.ListCurrencyRequest) (*exchange_rate.ListCurrencyResponse, error) {
	result, err := _this.currencyService.ListCurrencies(&services.ListCurrencyRequest{Base: request.Base, PageIndex: request.PageIndex.GetValue(), PageSize: request.PageSize.GetValue()})
	if err != nil {
		_this.logger.Error(err.Error(), zap.Error(err))
	}
	var data []*exchange_rate.Currency
	for _, s := range *result.Data {
		data = append(data, &exchange_rate.Currency{Name: s.Name, Description: s.Description})
	}

	return &exchange_rate.ListCurrencyResponse{Data: data, PageIndex: wrapperspb.Int32(result.PageIndex), PageSize: wrapperspb.Int32(result.PageSize), Total: wrapperspb.Int32(result.Total)}, nil
}
