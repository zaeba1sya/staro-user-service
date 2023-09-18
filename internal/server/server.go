package server

import (
	"fmt"
	"net"
	"user/internal/config"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type GRPCServer struct {
	log *zap.Logger
	cfg *config.ServerConfig
	server *grpc.Server
}

func NewGrpcServer(logger *zap.Logger, cfg *config.ServerConfig, opts []grpc.ServerOption) *GRPCServer {
	return &GRPCServer{
		log: logger,
		cfg: cfg,
		server: grpc.NewServer(opts...),
	}
}

func (g *GRPCServer) Run() {
	g.log.Info("Starting gRPC server")
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", g.cfg.Port))
	if err != nil {
		panic(err)
	}
	g.server.Serve(lis)
}