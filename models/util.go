package models

type Coordinate uint16

type Position struct {
	X, Y, Z float32
}

type PackedPosition struct {
	X, Y, Z Coordinate
}

func (c Coordinate) Float() float32 {
	return (float32(c) * 2000 / 65535) - 1000
}

func (p PackedPosition) Unpack() Position {
	return Position{p.X.Float(), p.Y.Float(), p.Z.Float()}
}
