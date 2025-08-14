package command

import (
	"encoding/json"

	"nakama-server/pkg/match/messages"
	"nakama-server/pkg/match/model"

	"github.com/heroiclabs/nakama-common/runtime"
)

func updatePosition(data runtime.MatchData, state *model.MatchState) error {
	var msg messages.UpdatePosition
	content := data.GetData()
	if err := json.Unmarshal(content, &msg); err != nil {
		return err
	}

	if pos, found := state.Positions[msg.Id]; found {
		*pos = msg.Position
	}

	return nil
}
