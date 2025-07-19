package main

import (
	"context"
	"database/sql"
	"github.com/heroiclabs/nakama-common/runtime"
	"nakama-server/pkg/server"
)

func InitModule(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, initializer runtime.Initializer) error {
	greeter := server.NewHelloWorldGreeter()
	logger.Info(greeter.Greet())
	return nil
}
