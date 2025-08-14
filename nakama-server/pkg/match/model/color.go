package model

type Color struct {
	Red   float64 `json:"r"`
	Green float64 `json:"g"`
	Blue  float64 `json:"b"`
	Alpha float64 `json:"a"`
}

func NewEmptyColor() *Color {
	return &Color{0, 0, 0, 0}
}
