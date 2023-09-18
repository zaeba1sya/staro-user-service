package main

import (
	"flag"
	"math/rand"
	"os"
	"os/signal"
	"time"
	"user/internal/config"
	"user/internal/logging"
	"user/internal/server"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func init() {
	rand.NewSource(time.Now().UnixNano())
}

func main() {
	flag.Parse()

	cfg := config.InitConfiguration()
	logger := logging.NewLogger("user.log")

	var opts []grpc.ServerOption

	grpcServer := server.NewGrpcServer(logger, &cfg.Server, opts)
	go grpcServer.Run()

	logger.Info("gRPC server started successfuly")

	signalChan := make(chan os.Signal)
	signal.Notify(signalChan, os.Interrupt, os.Kill)
	takeSignal := <-signalChan
	logger.Info("Shutting down gracefully", zap.String("signal", takeSignal.String()))
}