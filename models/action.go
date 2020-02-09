package models

type ActionHeader struct {
	AnimationTargetID uint32 `json:"targetID"` // who the animation targets
	unk1              uint32
	ActionID          uint32  `json:"actionID"` // what action was cast, as shown in battle log
	GlobalSequence    uint32  `json:"globalSequence"`
	AnimationLockTime float32 `json:"animationLockTime"` // doesn't seem to do anything
	someTargetID      uint32  // always 0x0E000000, the internal def for INVALID TARGET ID
	Hidden            uint16  `json:"hiddenAnimation"` // show animation if 0, counts up for each animation not shown
	Direction         uint16  `json:"direction"`       // 0x0000 ~ 0xFFFF, NWSE <=> 0,0x4000,0x8000,0xC000
	AnimationID       uint16  `json:"animationID"`     // animation played on casting character
	Animationvariant  uint8   `json:"animationVariant"`
	EffectDisplayType uint8   `json:"effectDisplayType"`
	unk2              uint8
	Count             uint8 `json:"count"` // ignores effects if 0, otherwise renders all
	unk3              uint16
	unk4              uint32
	unk5              uint16
}

type ActionEffect struct {
	EffectType  uint8
	HitSeverity uint8
	padding     uint8
	Percentage  uint8
	Multiplier  uint8 // Total Damage = 65535 * Multiplier * (Flags & 0x40) + Damage
	Flags       uint8 // (Flags & 0xA0) means attacker receives damage instead
	Damage      uint16
}

type Action struct {
	ActionHeader

	Effects        [8]ActionEffect
	unk1           uint32
	unk2           uint16
	EffectTargetID uint32 // who the effect targets
	EffectFlags    uint32 // do nothing if >0, just shows animation with no log or UI text
	unk3           uint32
}
