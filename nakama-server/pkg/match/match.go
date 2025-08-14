package match

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"nakama-server/pkg/match/command"
	msgs "nakama-server/pkg/match/messages"
	"nakama-server/pkg/match/model"
	"nakama-server/pkg/storage"

	"github.com/heroiclabs/nakama-common/runtime"
)

const (
	MATCH_MODULE string = "world_control"

	SPAWN_POSITION_X = 1800.0
	SPAWN_POSITION_Y = 1280.0
)

type MatchRegistrar func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule) (m runtime.Match, err error)

type Match struct{}

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

// When the match is initialized. Creates empty tables in the game state that will be populated by clients.
func (m *Match) MatchInit(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, params map[string]interface{}) (interface{}, int, string) {
	matchstate := model.MewMatchState()
	tickRate := model.MATCH_STATE_TICK_RATE
	label := model.MATCH_STATE_LABEL

	return matchstate, tickRate, label
}

// When someone tries to join the match. Checks if someone is already logged in and blocks them from doing so if so.
func (m *Match) MatchJoinAttempt(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, presence runtime.Presence, metadata map[string]string) (interface{}, bool, string) {
	matchstate := state.(*model.MatchState)

	userId := presence.GetUserId()
	_, found := matchstate.Presences[userId]
	if found {
		return matchstate, false, "user already logged in"
	}

	return matchstate, true, ""
}

// When someone does join the match. Initializes their entries in the game state tables with dummy values until they spawn in.
func (m *Match) MatchJoin(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, presences []runtime.Presence) interface{} {
	matchstate := state.(*model.MatchState)

	for i, p := range presences {
		userId := p.GetUserId()

		matchstate.Presences[userId] = p
		matchstate.Positions[userId] = model.NewPostion(0, 0)
		matchstate.Inputs[userId] = model.NewInput(0, false)
		matchstate.Colors[userId] = model.NewEmptyColor()
		matchstate.Names[userId] = fmt.Sprintf("user_%d", i+1)
	}

	return matchstate
}

// When someone leaves the match. Clears their entries in the game state tables, but saves their position to storage for next time.
func (m *Match) MatchLeave(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, presences []runtime.Presence) interface{} {
	matchstate := state.(*model.MatchState)

	data := []*runtime.StorageWrite{}

	for _, p := range presences {
		playerData, err := storage.MakePlayerData(matchstate, p)
		if err != nil {
			logger.Warn("error when serializing player data: %s", err.Error())
			continue
		}
		data = append(data, playerData)

		matchstate.Delete(p.GetUserId())
	}

	if _, err := nk.StorageWrite(ctx, data); err != nil {
		return runtime.NewError("error saving data", 13)
	}

	return matchstate
}

// Called `tickrate` times per second. Handles client messages and sends game state updates. Uses
// boiler plate commands from the command pattern except when specialization is required.
func (m *Match) MatchLoop(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, messages []runtime.MatchData) interface{} {
	matchstate := state.(*model.MatchState)
	commandHandler := command.GetCommandHandler()

	for _, m := range messages {
		opcode := m.GetOpCode()
		sender := matchstate.Presences[m.GetUserId()]

		// Run boiler plate commands (state updates.)
		var msgerr error
		reducer, found := commandHandler[opcode]
		if found {
			if msgerr = reducer(m, matchstate); msgerr != nil {
				logger.Warn("error when handling OPCODE %v: %s", opcode, msgerr.Error())
			}
		} else {
			logger.Warn("unknwon OPCODE : %v", opcode)
		}

		// A client has selected a character and is spawning. Get or generate position data,
		// send them initial state, and broadcast their spawning to existing clients.
		if opcode == command.OP_DO_SPAWN {
			if msgerr != nil {
				continue
			}
			var msg msgs.DoSpawn
			_ = json.Unmarshal(m.GetData(), &msg)

			// Read stored player state
			playerData, found, err := storage.ReadhPlayerData(ctx, nk, matchstate, sender)
			if err != nil {
				logger.Warn("error retrieving player data: %s", err.Error())
			}

			// Set initial player position
			if !found {
				*matchstate.Positions[sender.GetUserId()] = *model.NewPostion(SPAWN_POSITION_X, SPAWN_POSITION_Y)
			} else {
				*matchstate.Positions[sender.GetUserId()] = *playerData
			}

			// TODO
			matchstate.Names[sender.GetUserId()] = msg.Id

			// Setup initial state
			initialState, err := msgs.MakeInitialStatePayload(matchstate)
			if err != nil {
				logger.Warn("could not encode match state: %s", err.Error())
			} else {
				if err := dispatcher.BroadcastMessage(command.OP_INITIAL_STATE, initialState, []runtime.Presence{sender}, nil, true); err != nil {
					logger.Warn("error broadcasting message: %s", err.Error())
				}
			}

			if err := dispatcher.BroadcastMessage(command.OP_DO_SPAWN, m.GetData(), nil, nil, true); err != nil {
				logger.Warn("error broadcasting message: %s", err.Error())
			}
		} else if opcode == command.OP_UPDATE_COLOR {
			if err := dispatcher.BroadcastMessage(opcode, m.GetData(), nil, nil, true); err != nil {
				logger.Warn("error broadcasting message: %s", err.Error())
			}
		}
	}

	data, err := msgs.MakeUpdateStatePayload(matchstate)
	if err != nil {
		logger.Warn("could not encode match state: %s", err.Error())
	} else {
		if err := dispatcher.BroadcastMessage(command.OP_UPDATE_STATE, data, nil, nil, true); err != nil {
			logger.Warn("error broadcasting messaage: %s", err.Error())
		}
	}

	for _, i := range matchstate.Inputs {
		i.Jump = false
	}

	return matchstate
}

// Server is shutting down. Save positions of all existing characters to storage.
func (m *Match) MatchTerminate(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, graceSeconds int) interface{} {
	matchstate := state.(*model.MatchState)

	// Save positions of all existing characters to storage
	data := []*runtime.StorageWrite{}
	for _, presence := range matchstate.Presences {
		playerData, err := storage.MakePlayerData(matchstate, presence)
		if err != nil {
			logger.Warn("error when serializing player data: %s", err.Error())
			continue
		}

		data = append(data, playerData)
	}

	if _, err := nk.StorageWrite(ctx, data); err != nil {
		return runtime.NewError("error saving data", 13)
	}

	return matchstate
}

// Called when the match handler receives a runtime signal
func (m *Match) MatchSignal(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, data string) (interface{}, string) {
	return state, data
}
