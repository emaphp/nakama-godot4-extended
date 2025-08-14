package main

import (
	"context"
	"database/sql"
	"nakama-server/pkg/match"
	"nakama-server/pkg/storage"
	"slices"

	"github.com/heroiclabs/nakama-common/runtime"
)

func InitModule(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, initializer runtime.Initializer) error {
	if err := initializer.RegisterRpc("get_world_id", getWorldId); err != nil {
		logger.Error("Unable to register: %v", err)
		return err
	}
	if err := initializer.RegisterRpc("register_character_name", registerCharacterName); err != nil {
		logger.Error("Unable to register: %v", err)
		return err
	}
	if err := initializer.RegisterRpc("remove_character_name", removeCharacterName); err != nil {
		logger.Error("Unable to register: %v", err)
		return err
	}

	return nil
}

func getWorldId(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, payload string) (string, error) {
	matches, err := nk.MatchList(ctx, 1, false, "", nil, nil, "")
	if err != nil {
		logger.Warn("get_world_id -> MatchList: %s", err.Error())
		return "", err
	}

	if len(matches) == 0 {
		m, err := nk.MatchCreate(ctx, match.MATCH_MODULE, match.GetDefaultParams())
		if err != nil {
			logger.Warn("get_world_id -> MatchCreate: %s", err.Error())
			return "", err
		}
		logger.Warn("get_world_id -> CREATE")
		return m, nil
	}

	logger.Warn("get_world_id -> FOUND")
	return matches[0].MatchId, nil
}

func registerCharacterName(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, payload string) (string, error) {
	names, err := storage.ReadNames(ctx, nk)
	if err != nil {
		return "0", err
	}
	if len(names) == 0 {
		if err := storage.WriteNames(ctx, nk, []string{payload}); err != nil {
			return "0", err
		}
		return "1", nil
	}

	if slices.Contains(names, payload) {
		return "0", nil
	}

	if err := storage.WriteNames(ctx, nk, []string{payload}); err != nil {
		return "0", err
	}

	return "1", nil
}

func removeCharacterName(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, payload string) (string, error) {
	names, err := storage.ReadNames(ctx, nk)
	if err != nil {
		return "0", err
	}
	if len(names) == 0 {
		return "1", nil
	}

	names = slices.DeleteFunc(names, func(s string) bool {
		return s == payload
	})
	if err := storage.WriteNames(ctx, nk, names); err != nil {
		return "0", err
	}

	return "1", nil
}
