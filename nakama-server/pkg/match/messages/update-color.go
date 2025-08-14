package messages

import "nakama-server/pkg/match/model"

type UpdateColor struct {
	Id    string      `json:"id"`
	Color model.Color `json:"color"`
}
