package command

import (
	"encoding/json"

	"nakama-server/pkg/match/messages"
	"nakama-server/pkg/match/model"

	"github.com/heroiclabs/nakama-common/runtime"
)

func updateJump(data runtime.MatchData, state *model.MatchState) error {
	var msg messages.UpdateJump
	content := data.GetData()
	if err := json.Unmarshal(content, &msg); err != nil {
		return err
	}

	if in, found := state.Inputs[msg.Id]; found {
		in.Jump = true
	}

	return nil
}
