package cosmos

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/cometbft/cometbft/rpc/client/http"
	cmttypes "github.com/cometbft/cometbft/types"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"torram/bridge/types"
	runetypes "torram/x/rune/types"
)

// Client handles communication with the Cosmos chain
type Client struct {
	conn             *grpc.ClientConn
	msgClient        runetypes.MsgClient
	cometWsClient    *http.HTTP
	cometBftEndpoint string
}

// NewClient creates a new Cosmos client
func NewClient(grpcEndpoint string) (*Client, error) {
	// Setup gRPC connection
	conn, err := grpc.NewClient(
		grpcEndpoint,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to grpc: %w", err)
	}

	// Extract CometBFT endpoint from gRPC endpoint
	// Assuming format: localhost:9090 -> http://localhost:26657
	parts := strings.Split(grpcEndpoint, ":")
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid grpc endpoint format")
	}
	cometBftEndpoint := fmt.Sprintf("http://%s:26657", parts[0])

	// Setup CometBFT WebSocket client
	wsClient, err := http.New(cometBftEndpoint, "/websocket")
	if err != nil {
		return nil, fmt.Errorf("failed to create websocket client: %w", err)
	}

	// Start the WebSocket client
	if err := wsClient.Start(); err != nil {
		return nil, fmt.Errorf("failed to start websocket client: %w", err)
	}
	log.Printf("WebSocket client connected to %s", cometBftEndpoint)

	msgClient := runetypes.NewMsgClient(conn)

	return &Client{
		conn:             conn,
		msgClient:        msgClient,
		cometWsClient:    wsClient,
		cometBftEndpoint: cometBftEndpoint,
	}, nil
}

// SubscribeStakeEvents subscribes to stake events from the Cosmos chain
func (c *Client) SubscribeStakeEvents(ctx context.Context, stakeCh chan<- types.StakeEvent, errCh chan<- error) {
	// Subscribe to all transactions and filter in code
	query := "tm.event='Tx'"
	log.Printf("Subscribing to all transactions with query: %s", query)

	eventCh, err := c.cometWsClient.Subscribe(ctx, "bridge-client", query)
	if err != nil {
		errCh <- fmt.Errorf("failed to subscribe to stake events: %w", err)
		return
	}
	log.Println("Successfully subscribed to transactions")

	go func() {
		defer c.cometWsClient.Unsubscribe(ctx, "bridge-client", query)

		for {
			select {
			case <-ctx.Done():
				return

			case resultEvent := <-eventCh:
				// Type assert to TxEvent
				txEvent, ok := resultEvent.Data.(cmttypes.EventDataTx)
				if !ok {
					errCh <- fmt.Errorf("received non-tx event: %T", resultEvent.Data)
					continue
				}

				// Extract events from transaction
				events := txEvent.Result.Events
				log.Printf("Received transaction with %d events", len(events))

				for _, event := range events {
					log.Printf("Event Type: %s", event.Type)
					log.Printf("Event Attributes: %+v", event.Attributes)

					// Check if this is a stake event
					if event.Type != "rune_stake_initiated" {
						continue
					}

					log.Printf("Found stake event: %+v", event)

					// Parse event attributes
					var (
						runeID   string
						owner    string
						amount   uint64
						metadata string
					)

					for _, attr := range event.Attributes {
						key := string(attr.Key)
						value := string(attr.Value)
						log.Printf("Processing attribute - Key: %s, Value: %s", key, value)

						switch key {
						case "rune_id":
							runeID = value
						case "owner":
							owner = value
						case "amount":
							parsedAmount, err := strconv.ParseUint(value, 10, 64)
							if err != nil {
								errCh <- fmt.Errorf("failed to parse amount: %w", err)
								continue
							}
							amount = parsedAmount
						case "metadata":
							metadata = value
						}
					}

					// Validate required fields
					if runeID == "" || owner == "" || metadata == "" {
						errCh <- fmt.Errorf("missing required event attributes")
						continue
					}

					log.Printf("Sending stake event - RuneID: %s, Owner: %s, Amount: %d", runeID, owner, amount)
					// Send stake event to channel
					stakeCh <- types.StakeEvent{
						RuneID:   runeID,
						Owner:    owner,
						Amount:   amount,
						Metadata: metadata,
					}
				}
			}
		}
	}()
}

// SubscribeUnstakeEvents subscribes to unstake events from the Cosmos chain
func (c *Client) SubscribeUnstakeEvents(ctx context.Context, unstakeCh chan<- types.UnstakeEvent, errCh chan<- error) {
	// Subscribe to all transactions and filter in code
	query := "tm.event='Tx'"
	log.Printf("Subscribing to all transactions with query: %s", query)

	eventCh, err := c.cometWsClient.Subscribe(ctx, "bridge-client", query)
	if err != nil {
		errCh <- fmt.Errorf("failed to subscribe to unstake events: %w", err)
		return
	}
	log.Println("Successfully subscribed to transactions")

	go func() {
		defer c.cometWsClient.Unsubscribe(ctx, "bridge-client", query)

		for {
			select {
			case <-ctx.Done():
				return

			case resultEvent := <-eventCh:
				// Type assert to TxEvent
				txEvent, ok := resultEvent.Data.(cmttypes.EventDataTx)
				if !ok {
					errCh <- fmt.Errorf("received non-tx event: %T", resultEvent.Data)
					continue
				}

				// Extract events from transaction
				events := txEvent.Result.Events
				log.Printf("Received transaction with %d events", len(events))

				for _, event := range events {
					log.Printf("Event Type: %s", event.Type)
					log.Printf("Event Attributes: %+v", event.Attributes)

					// Check if this is an unstake event
					if event.Type != "rune_unstake_initiated" {
						continue
					}

					log.Printf("Found unstake event: %+v", event)

					// Parse event attributes
					var (
						runeID   string
						owner    string
						amount   uint64
						metadata string
					)

					for _, attr := range event.Attributes {
						key := string(attr.Key)
						value := string(attr.Value)
						log.Printf("Processing attribute - Key: %s, Value: %s", key, value)

						switch key {
						case "rune_id":
							runeID = value
						case "owner":
							owner = value
						case "amount":
							parsedAmount, err := strconv.ParseUint(value, 10, 64)
							if err != nil {
								errCh <- fmt.Errorf("failed to parse amount: %w", err)
								continue
							}
							amount = parsedAmount
						case "metadata":
							metadata = value
						}
					}

					// Validate required fields
					if runeID == "" || owner == "" || metadata == "" {
						errCh <- fmt.Errorf("missing required event attributes")
						continue
					}

					log.Printf("Sending unstake event - RuneID: %s, Owner: %s, Amount: %d", runeID, owner, amount)
					// Send unstake event to channel
					unstakeCh <- types.UnstakeEvent{
						RuneID:   runeID,
						Owner:    owner,
						Amount:   amount,
						Metadata: metadata,
					}
				}
			}
		}
	}()
}

// UpdateRuneStatus updates the status of a Rune in the Cosmos chain
func (c *Client) UpdateRuneStatus(ctx context.Context, runeID string, status string) error {
	msg := &runetypes.MsgUpdateRuneStatus{
		Authority: "bridge_module", // This should be the bridge module account address
		RuneId:    runeID,
		Status:    status,
	}

	// TODO: Sign the message with the bridge module account key
	_, err := c.msgClient.UpdateRuneStatus(ctx, msg)
	if err != nil {
		return fmt.Errorf("failed to update rune status: %w", err)
	}

	return nil
}

// Close closes all connections
func (c *Client) Close() error {
	if err := c.conn.Close(); err != nil {
		return fmt.Errorf("failed to close grpc connection: %w", err)
	}
	if err := c.cometWsClient.Stop(); err != nil {
		return fmt.Errorf("failed to close websocket connection: %w", err)
	}
	return nil
}
