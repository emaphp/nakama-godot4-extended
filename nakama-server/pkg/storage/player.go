package storage

import (
	"context"
	"encoding/json"
	"fmt"
	"nakama-server/pkg/match/model"

	"github.com/heroiclabs/nakama-common/runtime"
)

const (
	STORAGE_COLLECTION_PLAYER   string = "player_data"
	STORAGE_KEY_POSITION_PREFIX string = "position"
)

func makePlayerPositionKey(name string) string {
	return fmt.Sprintf("%s_%s", STORAGE_KEY_POSITION_PREFIX, name)
}

// Returns a StorageWrite instance wrapping the state of a player
func MakePlayerData(state *model.MatchState, presence runtime.Presence) (*runtime.StorageWrite, error) {
	userId := presence.GetUserId()
	userName := state.Names[userId]
	position := state.Positions[userId]

	positionData, err := json.Marshal(position)
	if err != nil {
		return nil, err
	}

	return &runtime.StorageWrite{
		Collection: STORAGE_COLLECTION_PLAYER,
		Key:        makePlayerPositionKey(userName),
		UserID:     userId,
		Value:      string(positionData),
	}, nil
}

func ReadhPlayerData(ctx context.Context, nk runtime.NakamaModule, state *model.MatchState, presence runtime.Presence) (*model.Position, bool, error) {
	userId := presence.GetUserId()
	userName := state.Names[userId]

	read := &runtime.StorageRead{
		Collection: STORAGE_COLLECTION_PLAYER,
		Key:        makePlayerPositionKey(userName),
		UserID:     userId,
	}

	records, err := nk.StorageRead(ctx, []*runtime.StorageRead{read})
	if err != nil {
		return nil, false, err
	}

	if len(records) == 0 {
		return nil, false, nil
	}

	positionData := new(model.Position)
	if err := json.Unmarshal([]byte(records[0].Value), positionData); err != nil {
		return nil, false, err
	}

	return positionData, true, err
}
