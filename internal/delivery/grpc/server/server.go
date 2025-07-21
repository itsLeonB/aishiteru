package server

import (
	"net"
	"os"
	"os/signal"
	"syscall"

	pb "github.com/itsLeonB/aishiteru/gen/go/llm"
	"github.com/itsLeonB/aishiteru/internal/config"
	"github.com/itsLeonB/aishiteru/internal/logging"
	"github.com/itsLeonB/aishiteru/internal/service/llm"
	"github.com/itsLeonB/ezutil"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedLLMServiceServer
	llmService llm.LLMService
	configs    *ezutil.Config
}

func Setup(configs *ezutil.Config) *Server {
	llmConfig, ok := configs.Generic.(*config.LLM)
	if !ok {
		logging.Logger.Fatal("generic config is not LLM config")
	}

	llmService := llm.NewLLMService(llmConfig)

	return &Server{
		llmService: llmService,
		configs:    configs,
	}
}

func (s *Server) Run() {
	port := s.configs.App.Port
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		logging.Logger.Fatalf("error listening to port %s: %v", port, err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterLLMServiceServer(grpcServer, s)

	go func() {
		logging.Logger.Infof("server started at port: %s", port)
		if err := grpcServer.Serve(listener); err != nil {
			logging.Logger.Fatalf("failed to serve: %v", err)
		}
	}()

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM)
	<-exit
	logging.Logger.Info("shutting down server...")
	grpcServer.GracefulStop()
	logging.Logger.Info("server successfully shutdown")
}
