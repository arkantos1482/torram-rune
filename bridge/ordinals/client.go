package ordinals

import (
	"context"
	"encoding/json"
	"fmt"

	"torram/bridge/types"

	"github.com/btcsuite/btcd/btcec/v2/ecdsa"
	"github.com/btcsuite/btcd/btcjson"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/rpcclient"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
)

const (
	// Maximum size for Runestone data pushes
	maxRunestoneSize = 520 // Maximum push size in Bitcoin Script
	// Dust limit for Bitcoin outputs
	dustLimit = 546
	// Fee rate in satoshis per byte
	feeRate = 10
	// OP_13 opcode for Runestone
	op13 = 0x5d
	// Satoshis per Bitcoin
	satoshiPerBitcoin = 100000000
)

// Client handles communication with the Bitcoin node
type Client struct {
	rpcClient *rpcclient.Client
	network   *chaincfg.Params
}

// NewClient creates a new Bitcoin Ordinals client
func NewClient(endpoint, user, pass string) (*Client, error) {
	connCfg := &rpcclient.ConnConfig{
		Host:         endpoint,
		User:         user,
		Pass:         pass,
		HTTPPostMode: true,
		DisableTLS:   true,
	}

	client, err := rpcclient.New(connCfg, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create bitcoin rpc client: %w", err)
	}

	// Use mainnet by default
	network := &chaincfg.MainNetParams

	return &Client{
		rpcClient: client,
		network:   network,
	}, nil
}

// UpdateRuneMetadata updates the metadata of a Rune in the Bitcoin blockchain
func (c *Client) UpdateRuneMetadata(ctx context.Context, metadata types.RuneMetadata) (string, error) {
	// Convert metadata to JSON
	metadataJSON, err := json.Marshal(metadata)
	if err != nil {
		return "", fmt.Errorf("failed to marshal metadata: %w", err)
	}

	// Create the runestone data
	runestoneData := struct {
		Type     string          `json:"type"`
		Action   string          `json:"action"`
		RuneID   string          `json:"rune_id"`
		Metadata json.RawMessage `json:"metadata"`
	}{
		Type:     "rune",
		Action:   "update",
		RuneID:   metadata.RuneID,
		Metadata: metadataJSON,
	}

	runestoneJSON, err := json.Marshal(runestoneData)
	if err != nil {
		return "", fmt.Errorf("failed to marshal runestone data: %w", err)
	}

	// Create new transaction
	msgTx := wire.NewMsgTx(wire.TxVersion)

	// Get unspent outputs for the runestone
	unspent, err := c.getUnspentOutput()
	if err != nil {
		return "", fmt.Errorf("failed to get unspent output: %w", err)
	}

	// Add input
	txHash, err := chainhash.NewHashFromStr(unspent.TxID)
	if err != nil {
		return "", fmt.Errorf("failed to parse txid: %w", err)
	}
	outpoint := wire.NewOutPoint(txHash, unspent.Vout)
	txIn := wire.NewTxIn(outpoint, nil, nil)
	msgTx.AddTxIn(txIn)

	// Create Runestone script: OP_RETURN OP_13 <data>
	builder := txscript.NewScriptBuilder()
	builder.AddOp(txscript.OP_RETURN)
	builder.AddOp(op13)
	builder.AddData(runestoneJSON)

	runestoneScript, err := builder.Script()
	if err != nil {
		return "", fmt.Errorf("failed to create runestone script: %w", err)
	}
	msgTx.AddTxOut(wire.NewTxOut(0, runestoneScript))

	// Add change output
	changeScript, err := c.getChangeScript()
	if err != nil {
		return "", fmt.Errorf("failed to get change script: %w", err)
	}

	// Convert amount from BTC to satoshis and calculate change
	satoshis := int64(unspent.Amount * satoshiPerBitcoin)
	changeAmount := satoshis - c.estimateFee(msgTx)

	if changeAmount > dustLimit {
		msgTx.AddTxOut(wire.NewTxOut(changeAmount, changeScript))
	}

	// Sign the transaction
	if err := c.signTransaction(msgTx, unspent); err != nil {
		return "", fmt.Errorf("failed to sign transaction: %w", err)
	}

	// Broadcast the transaction
	txHash, err = c.rpcClient.SendRawTransaction(msgTx, true)
	if err != nil {
		return "", fmt.Errorf("failed to broadcast transaction: %w", err)
	}

	return txHash.String(), nil
}

// getUnspentOutput returns a suitable unspent output for the inscription
func (c *Client) getUnspentOutput() (*btcjson.ListUnspentResult, error) {
	// Get unspent outputs with minimum confirmations
	unspent, err := c.rpcClient.ListUnspent()
	if err != nil {
		return nil, fmt.Errorf("failed to get unspent outputs: %w", err)
	}

	// Find a suitable output
	for _, utxo := range unspent {
		if utxo.Amount > float64(dustLimit+maxRunestoneSize)/float64(btcutil.SatoshiPerBitcoin) {
			return &utxo, nil
		}
	}

	return nil, fmt.Errorf("no suitable unspent outputs found")
}

// getChangeScript returns a P2PKH script for change output
func (c *Client) getChangeScript() ([]byte, error) {
	// Get a new address for change
	addr, err := c.rpcClient.GetNewAddress("")
	if err != nil {
		return nil, fmt.Errorf("failed to get change address: %w", err)
	}

	// Create P2PKH script
	addrPubKeyHash, err := btcutil.DecodeAddress(addr.String(), c.network)
	if err != nil {
		return nil, fmt.Errorf("failed to decode address: %w", err)
	}

	return txscript.PayToAddrScript(addrPubKeyHash)
}

// signTransaction signs all inputs of the transaction
func (c *Client) signTransaction(tx *wire.MsgTx, unspent *btcjson.ListUnspentResult) error {
	// Get private key
	addr, err := btcutil.DecodeAddress(unspent.Address, c.network)
	if err != nil {
		return fmt.Errorf("failed to decode address: %w", err)
	}

	privKey, err := c.rpcClient.DumpPrivKey(addr)
	if err != nil {
		return fmt.Errorf("failed to get private key: %w", err)
	}

	// Create signing script
	script, err := txscript.PayToAddrScript(addr)
	if err != nil {
		return fmt.Errorf("failed to create script: %w", err)
	}

	// Sign each input
	for i := range tx.TxIn {
		// Calculate signature hash
		sigHash, err := txscript.CalcSignatureHash(script, txscript.SigHashAll, tx, i)
		if err != nil {
			return fmt.Errorf("failed to calculate signature hash: %w", err)
		}

		// Sign the hash
		signature := ecdsa.Sign(privKey.PrivKey, sigHash)

		// Create signature script
		sigScript, err := txscript.NewScriptBuilder().
			AddData(append(signature.Serialize(), byte(txscript.SigHashAll))).
			AddData(privKey.PrivKey.PubKey().SerializeCompressed()).
			Script()
		if err != nil {
			return fmt.Errorf("failed to create signature script: %w", err)
		}

		tx.TxIn[i].SignatureScript = sigScript
	}

	return nil
}

// estimateFee estimates the fee for the transaction
func (c *Client) estimateFee(tx *wire.MsgTx) int64 {
	// Estimate transaction size
	txSize := tx.SerializeSize()
	return int64(txSize * feeRate)
}

// Close closes the Bitcoin RPC client
func (c *Client) Close() {
	c.rpcClient.Shutdown()
}
