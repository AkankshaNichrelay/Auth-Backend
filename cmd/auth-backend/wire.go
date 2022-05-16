//go:build wireinject
// +build wireinject

package main

import (
	"context"
	"log"

	"github.com/AkankshaNichrelay/Auth-Backend/internal/auth"
	"github.com/AkankshaNichrelay/Auth-Backend/internal/config"
	"github.com/AkankshaNichrelay/Auth-Backend/internal/db"
	"github.com/AkankshaNichrelay/Auth-Backend/internal/handler"
	"github.com/google/wire"
)

func InitializeAndRun(ctx context.Context, logger *log.Logger, configFile string) (*handler.Handler, error) {
	panic(
		wire.Build(
			config.New,
			config.NewDBConfig,
			db.New,
			auth.New,
			handler.New,
		),
	)
}
