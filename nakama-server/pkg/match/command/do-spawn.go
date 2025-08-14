package command

import (
	"encoding/json"

	"nakama-server/pkg/match/messages"
	"nakama-server/pkg/match/model"

	"github.com/heroiclabs/nakama-common/runtime"
)

func doSpawn(data runtime.MatchData, state *model.MatchState) error {
	var msg messages.DoSpawn
	content := data.GetData()
	if err := json.Unmarshal(content, &msg); err != nil {
		return err
	}

	if color, found := state.Colors[msg.Id]; found {
		*color = msg.Color
	}

	return nil
}
