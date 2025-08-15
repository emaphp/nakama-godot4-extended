package main

import (
	"context"
	"database/sql"

	"nakama-server/pkg/match"

	"github.com/heroiclabs/nakama-common/runtime"
)

func InitModule(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, initializer runtime.Initializer) error {
	moduleName, newMatch := match.RegisterMatch()

	if err := initializer.RegisterMatch(moduleName, newMatch); err != nil {
		logger.Error("[RegisterMatch] error: ", err.Error())
		return err
	}

	return nil
}
