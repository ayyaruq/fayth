package models

type HateListEntry struct {
	EnemyID    uint32 // The Enemy that has hate against the actor
	Percentage uint8  // how close to aggro, 100 means you have hate
	unk1       uint8  // leftovers from SE copying the HateRanking struct
	unk2       uint16
}

// HateList defines the list of targets an actor has hate on, ie coloured dots
type HateList struct {
	Count   uint8
	unk1    uint8
	unk2    uint16
	Entries [32]HateListEntry
	padding uint32
}

type HateRankingEntry struct {
	ActorID uint32 // the Actor that holds the hate
	Hate    uint32 // the total amount of hate earned by the Actor
}

// HateRanking defines the hate list for the current target, ie white bars
type HateRanking struct {
	Count   uint8
	unk1    uint8
	unk2    uint16
	Entries [32]HateRankingEntry
	padding uint32
}
