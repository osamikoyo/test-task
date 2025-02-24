package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/osamikoyo/test-task/internal/config"
	"github.com/osamikoyo/test-task/internal/data"
	"github.com/osamikoyo/test-task/internal/handler"
	"github.com/osamikoyo/test-task/internal/service"
	"github.com/osamikoyo/test-task/pkg/loger"
)

func main() {
	logger := loger.New()
	cfg := config.Load()

	logger.Info().Msg("starting initialization")

	repo, err := data.New(cfg)
	if err != nil{
		logger.Error().Err(err)
	}
	service := service.NewSongService(repo, logger)
	handler := handler.NewSongHandler(service, logger)
	
	r := gin.Default()

	handler.RegisterRoutes(r)

	if err = r.Run(fmt.Sprintf("%s:%s", cfg.SERVER_HOST, cfg.SERVER_PORT)); err != nil{
		logger.Error().Err(err)
	}

	logger.Info().Msg("server started succesfully!")
}