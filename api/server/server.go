package server

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/madmuzz05/go-final-project/internal/config"
	"github.com/madmuzz05/go-final-project/routes"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func InitServer() error {
	err := config.LoadConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("cannot load config")
	}
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	signal.Notify(ch, syscall.SIGINT)
	signal.Notify(ch, syscall.SIGTERM)
	go func() {
		oscall := <-ch
		log.Warn().Msgf("system call:%+v", oscall)
		cancel()
	}()

	err = routes.Run(ctx)

	if err != nil {
		log.Panic().Err(err)
	}
	return nil
}
