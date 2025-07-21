package main

import (
	"github.com/itsLeonB/aishiteru/internal/config"
	"github.com/itsLeonB/aishiteru/internal/delivery/grpc/server"
	"github.com/itsLeonB/aishiteru/internal/logging"
	"github.com/itsLeonB/ezutil"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	configs := ezutil.LoadConfig(config.Defaults())
	logging.InitLogger(configs.App)
	s := server.Setup(configs)
	s.Run()
}
