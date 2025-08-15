package model

type Input struct {
	Direction float64 `json:"dir"`
	Jump      bool    `json:"jmp"`
}

func NewInput(direction float64, jump bool) *Input {
	return &Input{
		Direction: direction,
		Jump:      jump,
	}
}
