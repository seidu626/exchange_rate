package main

import (
	"flag"
	"fmt"
	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
	"github.com/seidu626/exchange_rate/config"
	exchange_rate "github.com/seidu626/exchange_rate/proto"
	"github.com/seidu626/exchange_rate/server/handler"
	"github.com/seidu626/exchange_rate/server/services"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"time"
)

type Environment string

const (
	DEVELOPMENT Environment = "DEVELOPMENT"
	PRODUCTION              = "PRODUCTION"
)

func main() {
	conf := viper.GetViper()
	logger, _ := zap.NewDevelopment()
	config.InitConfig(conf, *logger)

	if conf.GetString("ENVIRONMENT") == PRODUCTION {
		logger, _ = zap.NewProduction()
	}
	defer logger.Sync()

	// Config file found and successfully parsed

	// Determine port for HTTP service.
	port := conf.GetString("PORT")

	ring := redis.NewRing(&redis.RingOptions{
		Addrs: map[string]string{
			conf.GetString("REDIS_HOST"): ":" + conf.GetString("REDIS_PORT"),
		},
		Password: conf.GetString("REDIS_PASSWORD"),
	})
	logger.Info("Redis Details", zap.String(conf.GetString("REDIS_HOST"), conf.GetString("REDIS_PORT")))

	redisCache := cache.New(&cache.Options{
		Redis:      ring,
		LocalCache: cache.NewTinyLFU(1000, time.Minute),
	})

	currencyService := services.NewCurrencyService(logger, "ghs")
	rateServerHandler := handler.NewExchangeRateServerHandler(logger, redisCache, currencyService, conf.GetDuration("POLLING_RATES_INTERVAL"))

	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", port))
	if err != nil {
		logger.Fatal("failed to listen: %v", zap.Error(err))
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	exchange_rate.RegisterExchangeRatesServer(grpcServer, rateServerHandler)
	reflection.Register(grpcServer)

	logger.Info("Server started.", zap.String("Port", port))
	err = grpcServer.Serve(lis)
	if err != nil {
		logger.Fatal("failed to listen: %v", zap.Error(err))
	}

}
