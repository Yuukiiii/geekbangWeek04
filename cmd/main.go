package main

import (
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

var logger, _ = zap.NewProduction()

func main() {
	if err := cmd.Execute(); err != nil {
		logger.Fatal("err", zap.Error(err))
	}
}
