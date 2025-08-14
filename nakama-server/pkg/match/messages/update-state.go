package messages

import (
	"encoding/json"
	"nakama-server/pkg/match/model"
)

type UpdateState struct {
	Positions map[string]*model.Position `json:"pos"`
	Inputs    map[string]*model.Input    `json:"inp"`
}

func NewUpdateStateFromModel(state *model.MatchState) *UpdateState {
	return &UpdateState{
		Positions: state.Positions,
		Inputs:    state.Inputs,
	}
}

func MakeUpdateStatePayload(state *model.MatchState) ([]byte, error) {
	matchstate := &UpdateState{
		Positions: state.Positions,
		Inputs:    state.Inputs,
	}

	payload, err := json.Marshal(matchstate)
	if err != nil {
		return nil, err
	}

	return payload, nil
}
