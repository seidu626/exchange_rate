package handler

import (
	"context"
	exchange_rate "github.com/seidu626/exchange_rate/proto"
)

func (c *ExchangeRateClientHandler) GetRate(ctx context.Context, in *exchange_rate.RateRequest) (*exchange_rate.RateResponse, error) {
	out := new(exchange_rate.RateResponse)
	//err := c.cc.Invoke(ctx, "/exchange_rate.ExchangeRates/ListRates", in, out, opts...)
	//if err != nil {
	//	return nil, err
	//}
	return out, nil
}
