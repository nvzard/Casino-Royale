package main

import (
	"github.com/nvzard/casino-royale/server"
	"github.com/nvzard/casino-royale/utils"
	"go.uber.org/zap"
)

var logger *zap.SugaredLogger

func init() {
	logger = utils.GetLogger()
}

func main() {
	logger.Info("Casino Royale Service started")

	r := server.SetupApiServer()
	err := r.Run()
	if err != nil {
		logger.Fatalw("Failed to start server", "error", err.Error())
	}
}
