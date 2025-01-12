package bridge

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"torram/bridge/cosmos"
	"torram/bridge/ordinals"
	"torram/bridge/types"
)

// Config holds the configuration for the bridge service
type Config struct {
	CosmosGRPCEndpoint string
	BitcoinRPCEndpoint string
	BitcoinRPCUser     string
	BitcoinRPCPass     string
}

// BridgeService handles the communication between Cosmos and Bitcoin networks
type BridgeService struct {
	config        Config
	cosmosClient  *cosmos.Client
	ordinalClient *ordinals.Client
}

// NewBridgeService creates a new instance of BridgeService
func NewBridgeService(config Config) (*BridgeService, error) {
	cosmosClient, err := cosmos.NewClient(config.CosmosGRPCEndpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to create cosmos client: %w", err)
	}

	ordinalClient, err := ordinals.NewClient(config.BitcoinRPCEndpoint, config.BitcoinRPCUser, config.BitcoinRPCPass)
	if err != nil {
		return nil, fmt.Errorf("failed to create ordinal client: %w", err)
	}

	return &BridgeService{
		config:        config,
		cosmosClient:  cosmosClient,
		ordinalClient: ordinalClient,
	}, nil
}

// Start begins listening for events and processing them
func (s *BridgeService) Start(ctx context.Context) error {
	// Create channels for events
	stakeCh := make(chan types.StakeEvent)
	unstakeCh := make(chan types.UnstakeEvent)
	errCh := make(chan error)

	// Start listening for events
	go s.cosmosClient.SubscribeStakeEvents(ctx, stakeCh, errCh)
	go s.cosmosClient.SubscribeUnstakeEvents(ctx, unstakeCh, errCh)

	// Process events
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()

		case err := <-errCh:
			log.Printf("Error in event subscription: %v", err)
			continue

		case event := <-stakeCh:
			if err := s.handleStakeEvent(ctx, event); err != nil {
				log.Printf("Failed to handle stake event: %v", err)
			}

		case event := <-unstakeCh:
			if err := s.handleUnstakeEvent(ctx, event); err != nil {
				log.Printf("Failed to handle unstake event: %v", err)
			}
		}
	}
}

// handleStakeEvent processes a stake event
func (s *BridgeService) handleStakeEvent(ctx context.Context, event types.StakeEvent) error {
	// Parse metadata from event
	var metadata types.RuneMetadata
	if err := json.Unmarshal([]byte(event.Metadata), &metadata); err != nil {
		return fmt.Errorf("failed to parse metadata: %w", err)
	}

	// Update Ordinal metadata on Bitcoin
	txHash, err := s.ordinalClient.UpdateRuneMetadata(ctx, metadata)
	if err != nil {
		// Update status to failed in Cosmos
		if updateErr := s.cosmosClient.UpdateRuneStatus(ctx, event.RuneID, "failed"); updateErr != nil {
			log.Printf("Failed to update rune status to failed: %v", updateErr)
		}
		return fmt.Errorf("failed to update ordinal metadata: %w", err)
	}

	// Update status to staked in Cosmos
	if err := s.cosmosClient.UpdateRuneStatus(ctx, event.RuneID, "staked"); err != nil {
		return fmt.Errorf("failed to update rune status: %w", err)
	}

	log.Printf("Successfully processed stake event for rune %s, Bitcoin tx: %s", event.RuneID, txHash)
	return nil
}

// handleUnstakeEvent processes an unstake event
func (s *BridgeService) handleUnstakeEvent(ctx context.Context, event types.UnstakeEvent) error {
	// Parse metadata from event
	var metadata types.RuneMetadata
	if err := json.Unmarshal([]byte(event.Metadata), &metadata); err != nil {
		return fmt.Errorf("failed to parse metadata: %w", err)
	}

	// Update Ordinal metadata on Bitcoin
	txHash, err := s.ordinalClient.UpdateRuneMetadata(ctx, metadata)
	if err != nil {
		// Update status to failed in Cosmos
		if updateErr := s.cosmosClient.UpdateRuneStatus(ctx, event.RuneID, "failed"); updateErr != nil {
			log.Printf("Failed to update rune status to failed: %v", updateErr)
		}
		return fmt.Errorf("failed to update ordinal metadata: %w", err)
	}

	// Update status to unstaked in Cosmos
	if err := s.cosmosClient.UpdateRuneStatus(ctx, event.RuneID, "unstaked"); err != nil {
		return fmt.Errorf("failed to update rune status: %w", err)
	}

	log.Printf("Successfully processed unstake event for rune %s, Bitcoin tx: %s", event.RuneID, txHash)
	return nil
}
