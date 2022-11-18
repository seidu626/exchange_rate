package main

import (
	"context"
	"github.com/micro/micro/v3/service/config/env"
	exchange_rate "github.com/seidu626/exchange_rate/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	bindAddress, _ := env.NewConfig()
	err := bindAddress.Set("BIND_ADDRESS", "0.0.0.0:8080")
	if err != nil {
		return
	}

	logger, _ := zap.NewDevelopment()
	serverAddr := "0.0.0.0:8080"

	// exchangeHandler := &handler.ExchangeRateClientHandler{}
	//clientHandler := exchangeHandler.NewExchangeRatesClientHandler(logger)

	//var opts []grpc.DialOption
	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure())
	if err != nil {
		logger.Fatal("An error occurred", zap.Error(err))
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			logger.Fatal("An error occurred", zap.Error(err))
		}
	}(conn)

	logger.Info("Server running on",
		zap.String("Address: ", serverAddr))

	client := exchange_rate.NewExchangeRatesClient(conn)

	request := &exchange_rate.RateRequest{Base: exchange_rate.Currencies_USD, Destination: exchange_rate.Currencies_GBP}
	rate, err := client.GetRate(context.Background(), request)
	if err != nil {
		logger.Fatal("An error occurred", zap.Error(err))
	}
	logger.Info("Rate Response", zap.String("Rate: ", rate.String()))

}
