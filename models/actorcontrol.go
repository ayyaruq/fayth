package models

type ActorControl struct {
	Category uint16
	padding1 uint16
	Params   [4]uint32
	padding2 uint32
}

type ActorControlSelf struct {
	Category uint16
	padding1 uint16
	Params   [6]uint32
	padding2 uint32
}

type ActorControlTarget struct {
	Category uint16
	padding1 uint16
	Params   [4]uint32
	padding2 uint32
	TargetID uint64
}
