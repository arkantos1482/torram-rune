package types

import "time"

// RuneMetadata represents the metadata stored in Bitcoin Ordinals
type RuneMetadata struct {
	RuneID         string    `json:"rune_id"`
	State          string    `json:"state"`
	Chain          string    `json:"chain"`
	Owner          string    `json:"owner"`
	TxHash         string    `json:"tx_hash"`
	StakeTimestamp time.Time `json:"stake_timestamp"`
}

// StakeEvent represents a stake event from the Cosmos chain
type StakeEvent struct {
	RuneID   string
	Owner    string
	Amount   uint64
	Metadata string
}

// UnstakeEvent represents an unstake event from the Cosmos chain
type UnstakeEvent struct {
	RuneID   string
	Owner    string
	Amount   uint64
	Metadata string
}
