package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"torram/bridge"
)

func main() {
	// Parse command line flags
	cosmosGRPC := flag.String("cosmos-grpc", "localhost:9090", "Cosmos gRPC endpoint")
	btcRPC := flag.String("btc-rpc", "localhost:8332", "Bitcoin RPC endpoint")
	btcUser := flag.String("btc-user", "", "Bitcoin RPC username")
	btcPass := flag.String("btc-pass", "", "Bitcoin RPC password")
	flag.Parse()

	// Create bridge service
	bridgeService, err := bridge.NewBridgeService(bridge.Config{
		CosmosGRPCEndpoint: *cosmosGRPC,
		BitcoinRPCEndpoint: *btcRPC,
		BitcoinRPCUser:     *btcUser,
		BitcoinRPCPass:     *btcPass,
	})
	if err != nil {
		log.Fatalf("Failed to create bridge service: %v", err)
	}

	// Setup context with cancellation
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Handle shutdown gracefully
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-sigCh
		log.Printf("Received signal %v, shutting down...", sig)
		cancel()
	}()

	// Start the bridge service
	if err := bridgeService.Start(ctx); err != nil {
		log.Fatalf("Bridge service failed: %v", err)
	}
}
