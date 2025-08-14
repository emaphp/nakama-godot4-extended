package model

import (
	"github.com/heroiclabs/nakama-common/runtime"
)

const (
	MATCH_STATE_TICK_RATE = 10
	MATCH_STATE_LABEL     = "nakama-godot4"
)

type MatchState struct {
	Presences map[string]runtime.Presence
	Positions map[string]*Position
	Inputs    map[string]*Input
	Colors    map[string]*Color
	Names     map[string]string
}

func MewMatchState() *MatchState {
	return &MatchState{
		Presences: map[string]runtime.Presence{},
		Inputs:    map[string]*Input{},
		Positions: map[string]*Position{},
		Colors:    map[string]*Color{},
		Names:     map[string]string{},
	}
}

func (state *MatchState) Delete(userId string) {
	delete(state.Presences, userId)
	delete(state.Positions, userId)
	delete(state.Inputs, userId)
	delete(state.Colors, userId)
	delete(state.Names, userId)
}
