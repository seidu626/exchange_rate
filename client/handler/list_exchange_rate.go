package handler

import (
	"context"
	exchange_rate "github.com/seidu626/exchange_rate/proto"
	"google.golang.org/grpc"
)

func (c *ExchangeRateClientHandler) ListRates(ctx context.Context, in *exchange_rate.ListRatesRequest, opts ...grpc.CallOption) (*exchange_rate.ListRatesResponse, error) {
	out := new(exchange_rate.ListRatesResponse)
	//err := c.cc.Invoke(ctx, "/exchange_rate.ExchangeRates/ListRates", in, out, opts...)
	//if err != nil {
	//	return nil, err
	//}
	return out, nil
}
