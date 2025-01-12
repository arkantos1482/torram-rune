package types

const (
	// ModuleName defines the module name
	ModuleName = "rune"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_rune"

	// StakedRunePrefix is the prefix for staked rune storage
	StakedRunePrefix = "staked_rune:"

	// Rune Status Constants
	RuneStatusPending  = "pending"
	RuneStatusStaked   = "staked"
	RuneStatusUnstaked = "unstaked"
	RuneStatusFailed   = "failed"
)

var (
	ParamsKey = []byte("p_rune")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

// StakedRuneKey returns the store key to retrieve a StakedRune from the index fields
func StakedRuneKey(runeId string) []byte {
	return []byte(StakedRunePrefix + runeId)
}
