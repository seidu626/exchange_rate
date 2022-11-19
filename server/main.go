package main

import (
	"flag"
	"fmt"
	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
	"github.com/google/gnostic/cmd/protoc-gen-openapi/generator"
	exchange_rate "github.com/seidu626/exchange_rate/proto"
	"github.com/seidu626/exchange_rate/server/handler"
	"github.com/seidu626/exchange_rate/server/services"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"
	"net"
	"time"
)

var flags flag.FlagSet

func main() {
	port := 8080
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()

	ring := redis.NewRing(&redis.RingOptions{
		Addrs: map[string]string{
			"127.0.0.1": ":6379",
		},
	})

	redisCache := cache.New(&cache.Options{
		Redis:      ring,
		LocalCache: cache.NewTinyLFU(1000, time.Minute),
	})

	currencyService := services.NewCurrencyService(logger, "ghs")
	rateServerHandler := handler.NewExchangeRateServerHandler(logger, redisCache, currencyService, 5*time.Second)

	go func() {
		OpenApiGen()
	}()

	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		logger.Fatal("failed to listen: %v", zap.Error(err))
	}
	logger.Info("Server running on", zap.Int("Port: ", port))
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	exchange_rate.RegisterExchangeRatesServer(grpcServer, rateServerHandler)
	reflection.Register(grpcServer)

	err = grpcServer.Serve(lis)
	if err != nil {
		logger.Fatal("failed to listen: %v", zap.Error(err))
	}

}

func OpenApiGen() {
	var conf = generator.Configuration{
		Version:         flags.String("version", "0.0.1", "version number text, e.g. 1.2.3"),
		Title:           flags.String("title", "", "name of the API"),
		Description:     flags.String("description", "", "description of the API"),
		Naming:          flags.String("naming", "json", `naming convention. Use "proto" for passing names directly from the proto files`),
		FQSchemaNaming:  flags.Bool("fq_schema_naming", false, `schema naming convention. If "true", generates fully-qualified schema names by prefixing them with the proto message package name`),
		EnumType:        flags.String("enum_type", "integer", `type for enum serialization. Use "string" for string-based serialization`),
		CircularDepth:   flags.Int("depth", 2, "depth of recursion for circular messages"),
		DefaultResponse: flags.Bool("default_response", true, `add default response. If "true", automatically adds a default response to operations which use the google.rpc.Status message. Useful if you use envoy or grpc-gateway to transcode as they use this type for their default error responses.`),
	}
	opts := protogen.Options{
		ParamFunc: flags.Set,
	}

	opts.Run(func(plugin *protogen.Plugin) error {
		// Enable "optional" keyword in front of type (e.g. optional string label = 1;)
		plugin.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)

		return generator.NewOpenAPIv3Generator(plugin, conf).Run()
	})
}
