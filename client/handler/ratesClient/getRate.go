package ratesClient

import (
	"context"
	exchange_rate "github.com/seidu626/exchange_rate/proto"
)

func (c *ClientHandler) GetRate(ctx context.Context, in *exchange_rate.CurrencyRateRequest) (*exchange_rate.CurrencyRateResponse, error) {
	out := new(exchange_rate.CurrencyRateResponse)
	//err := c.cc.Invoke(ctx, "/exchange_rate.ExchangeRates/ListRates", in, out, opts...)
	//if err != nil {
	//	return nil, err
	//}
	return out, nil
}
