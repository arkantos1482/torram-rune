# Torram: Bitcoin Ordinal RUNE Bridge & Staking Platform

A Cosmos SDK-based application that enables staking of Bitcoin Ordinal RUNEs.

## Architecture Overview

The application consists of two main components:

1. **Rune Module (`x/rune`)**: Handles RUNE staking operations and state management on the Cosmos chain
2. **Bridge Service (`bridge/`)**: Manages synchronization between Cosmos chain and Bitcoin Ordinal Runes

### Key Features

#### RUNE Staking
- State tracking with multiple status types (pending, staked, unstaked, failed)
- Event-driven architecture for stake/unstake operations

#### Bitcoin Bridge
- Unidirectional bridge between Cosmos chain and Bitcoin network
- Automatic synchronization of RUNE metadata with Bitcoin Ordinals

## Technical Implementation

### Staking Flow
1. User initiates stake transaction through Cosmos SDK messages
2. `MsgStakeRune` creates a staking record with pending status
3. Bridge service listens for stake events
4. Ordinal metadata is updated on Bitcoin network
5. Stake status is confirmed on Cosmos chain

### Bridge Service
```go
type BridgeService struct {
    config        Config
    cosmosClient  *cosmos.Client
    ordinalClient *ordinals.Client
}
```

- Maintains connections to both Cosmos and Bitcoin networks
- Processes stake/unstake events
- Updates RUNE metadata
- Handles failure scenarios with automatic status updates

### RUNE Metadata Structure
```go
type RuneMetadata struct {
    RuneID         string    
    State          string    
    Chain          string    
    Owner          string    
    TxHash         string    
    StakeTimestamp time.Time 
}
```

## Event System

The platform uses an event-driven architecture:
- `rune_stake_initiated`: Emitted when a stake is initiated
- `rune_unstake_initiated`: Emitted when an unstake is requested
- Events include metadata for cross-chain synchronization


## Dependencies

- Cosmos SDK
- Bitcoin Core RPC
- Ordinals Protocol Support
