package models

type EffectEntry struct {
	Entries [8]ActionEffect
}

type AoEAction8 struct {
	ActionHeader

	Effects        [8]EffectEntry
	unk1           uint32
	unk2           uint16
	EffectTargetID [8]uint64
	Position       PackedPosition
	unk3           uint16
	unk4           uint32
}

type AoEAction16 struct {
	ActionHeader

	Effects        [8]EffectEntry
	unk1           uint32
	unk2           uint16
	EffectTargetID [8]uint64
	Position       PackedPosition
	unk3           uint16
	unk4           uint32
}

type AoEAction24 struct {
	ActionHeader

	Effects        [8]EffectEntry
	unk1           uint32
	unk2           uint16
	EffectTargetID [8]uint64
	Position       PackedPosition
	unk3           uint16
	unk4           uint32
}

type AoEAction32 struct {
	ActionHeader

	Effects        [8]EffectEntry
	unk1           uint32
	unk2           uint16
	EffectTargetID [8]uint64
	Position       PackedPosition
	unk3           uint16
	unk4           uint32
}
