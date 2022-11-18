package main

import (
	"flag"
	"fmt"
	exchange_rate "github.com/seidu626/exchange_rate/proto"
	"github.com/seidu626/exchange_rate/server/handler"
	"github.com/seidu626/exchange_rate/server/services"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

func main() {
	port := 8080
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()
	currencyService := services.NewCurrencyService(logger, "ghs")
	rateHandler := handler.NewExchangeRateServerHandler(logger, currencyService)

	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		logger.Fatal("failed to listen: %v",
			zap.Error(err))
	}
	logger.Info("Server running on",
		zap.Int("Port: ", port))

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	exchange_rate.RegisterExchangeRatesServer(grpcServer, rateHandler)
	reflection.Register(grpcServer)

	err = grpcServer.Serve(lis)
	if err != nil {
		logger.Fatal("failed to listen: %v",
			zap.Error(err))
	}

}
