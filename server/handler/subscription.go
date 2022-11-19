package handler

import (
	"context"
	exchange_rate "github.com/seidu626/exchange_rate/proto"
	"go.uber.org/zap"
	"io"
)

func (_this *ServerHandler) Subscription(stream exchange_rate.ExchangeRates_SubscriptionServer) error {
	// handle client messages
	for {
		rr, err := stream.Recv() // Recv is a blocking method which returns on client data
		// io.EOF signals that the client has closed the connection
		if err == io.EOF {
			_this.logger.Info("Client has closed connection")
			// if connection closed, then we remove client from subscribers
			delete(*_this.subscriptions, stream)
			break
		}

		// any other error means the transport between the server and client is unavailable
		if err != nil {
			_this.logger.Error("Unable to read from client", zap.Error(err))
			// if get any kind of error, then we remove client from subscribers
			delete(*_this.subscriptions, stream)
			return err
		}

		_this.logger.Info("Handle client request", zap.String("request_base", rr.Base))

		rrs, ok := (*_this.subscriptions)[stream]
		if !ok {
			rrs = &[]exchange_rate.ListRatesRequest{}
		}

		*rrs = append(*rrs, *rr)
		(*_this.subscriptions)[stream] = rrs
	}

	return nil
}

// GetSubscriptions returns clients subscriptions request
func (_this *ServerHandler) GetSubscriptions(ctx context.Context) (*map[exchange_rate.ExchangeRates_SubscriptionServer][]exchange_rate.ListRatesRequest, error) {
	cacheKey := "__subscriptions__"
	rates := &map[exchange_rate.ExchangeRates_SubscriptionServer][]exchange_rate.ListRatesRequest{}
	err := _this.cache.Get(ctx, cacheKey, rates)
	if err == nil {
		return rates, err
	}
	_this.logger.Info("Cache Error", zap.Error(err))
	return rates, err
}
