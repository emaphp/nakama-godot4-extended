package messages

import (
	"encoding/json"
	"nakama-server/pkg/match/model"
)

type InitialState struct {
	Positions map[string]*model.Position `json:"pos"`
	Inputs    map[string]*model.Input    `json:"inp"`
	Colors    map[string]*model.Color    `json:"col"`
	Names     map[string]string          `json:"nms"`
}

func NewInitialStateFromModel(state *model.MatchState) *InitialState {
	initialState := &InitialState{
		Positions: state.Positions,
		Inputs:    state.Inputs,
		Colors:    state.Colors,
		Names:     state.Names,
	}

	return initialState
}

func MakeInitialStatePayload(state *model.MatchState) ([]byte, error) {
	initialState := &InitialState{
		Positions: state.Positions,
		Inputs:    state.Inputs,
		Colors:    state.Colors,
		Names:     state.Names,
	}

	payload, err := json.Marshal(initialState)
	if err != nil {
		return nil, err
	}

	return payload, nil
}
