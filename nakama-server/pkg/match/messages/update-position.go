package messages

import "nakama-server/pkg/match/model"

type UpdatePosition struct {
	Id       string         `json:"id"`
	Position model.Position `json:"pos"`
}
