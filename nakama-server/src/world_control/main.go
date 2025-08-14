package main

import (
	"context"
	"database/sql"
	"github.com/heroiclabs/nakama-common/runtime"
	"nakama-server/pkg/match"
)

func InitModule(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, initializer runtime.Initializer) error {
	logger.Info("Hello Multiplayer!")

	moduleName, newMatch := match.RegisterMatch()

	if err := initializer.RegisterMatch(moduleName, newMatch); err != nil {
		logger.Error("[RegisterMatch] error: ", err.Error())
		return err
	}

	return nil
}
