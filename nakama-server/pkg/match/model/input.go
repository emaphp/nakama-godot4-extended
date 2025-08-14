package model

type Input struct {
	Direction float64 `json:"d"`
	Jump      bool    `json:"j"`
}

func NewInput(direction float64, jump bool) *Input {
	return &Input{
		Direction: direction,
		Jump:      jump,
	}
}
