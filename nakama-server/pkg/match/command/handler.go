package command

import (
	"nakama-server/pkg/match/model"

	"github.com/heroiclabs/nakama-common/runtime"
)

const (
	OP_UPDATE_POSITION int64 = iota + 1
	OP_UPDATE_INPUT
	OP_UPDATE_STATE
	OP_UPDATE_JUMP
	OP_DO_SPAWN
	OP_UPDATE_COLOR
	OP_INITIAL_STATE
)

type MatchStateReducer func(data runtime.MatchData, state *model.MatchState) error

type CommandHandler map[int64]MatchStateReducer

var (
	commandHandler CommandHandler
)

func init() {
	commandHandler = make(CommandHandler)
	commandHandler[OP_UPDATE_POSITION] = updatePosition
	commandHandler[OP_UPDATE_INPUT] = updateInput
	commandHandler[OP_UPDATE_JUMP] = updateJump
	commandHandler[OP_DO_SPAWN] = doSpawn
	commandHandler[OP_UPDATE_COLOR] = updateColor
}

func GetCommandHandler() CommandHandler {
	return commandHandler
}
