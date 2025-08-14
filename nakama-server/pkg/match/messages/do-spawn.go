package messages

import "nakama-server/pkg/match/model"

type DoSpawn struct {
	Id    string      `json:"id"`
	Color model.Color `json:"col"`
	Name  string      `json:"nm"`
}
