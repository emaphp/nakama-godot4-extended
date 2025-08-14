package model

type Position struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

func NewPostion(x float64, y float64) *Position {
	return &Position{
		X: x,
		Y: y,
	}
}
