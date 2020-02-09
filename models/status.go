package models

type StatusEffect struct {
	EffectID uint16
	Param    uint16
	Duration time.Duration // float32
	ActorID  uint32        // source actor ID
}

type StatusEffectEntry struct {
	Index    uint8
	unk1     uint8
	EffectID uint16
	Param    uint16
	unk2     uint16
	Duration time.Duration // float32
	ActorID  uint32        // source actor ID
}

type StatusEffectList struct {
	ClassJobID       uint8
	LevelUnk         uint8
	Level            uint16
	CurrentHP        uint32
	MaxHP            uint32
	CurrentMP        uint16
	MaxMP            uint16
	CurrentTP        uint16
	ShieldPercentage uint8
	unk1             uint8
	Effects          [30]StatusEffect
	padding          uint32
}

type StatusEffectResult struct {
	GlobalSequence   uint32
	ActorID          uint32
	CurrentHP        uint32
	MaxHP            uint32
	CurrentMP        uint16
	CurrentTP        uint16
	MaxMP            uint16 // why are these switched Square? What are you doing?
	unk1             uint8
	ShieldPercentage uint8
	Count            uint8
	unk2             uint16
	Effects          [4]StatusEffectEntry
	unk3             uint32
}
