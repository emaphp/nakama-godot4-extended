package match

import (
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
)

const (
	MATCH_MODULE string = "world_control"
)

type MatchRegistrar func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule) (m runtime.Match, err error)

func newMatch(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule) (m runtime.Match, err error) {
	return &Match{}, nil
}

func RegisterMatch() (string, MatchRegistrar) {
	return MATCH_MODULE, newMatch
}

func GetDefaultParams() map[string]any {
	params := map[string]any{}
	return params
}
