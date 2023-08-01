package merkle_proof_verifier

import (
	"encoding/hex"
	"fmt"
	"github.com/cometbft/cometbft/libs/log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethdb/memorydb"
	"github.com/ethereum/go-ethereum/trie"
)

// VerifyProof verifies a merkle proof for Ethereum storage
func VerifyProof(logger log.Logger, proofStrs []string, rootStr string, keyStr string, valueStr string) (bool, error) {
	// Convert hex strings to common.Hash
	// that's StorageHash
	root, err := toHash(rootStr)
	if err != nil {
		return false, fmt.Errorf("failed to convert root to hash: %w", err)
	}

	// the key
	if keyStr == "0x0" {
		keyStr = "0x0000000000000000000000000000000000000000000000000000000000000000"
	}
	key, err := hex.DecodeString(keyStr[2:])
	path := crypto.Keccak256(key)

	if err != nil {
		logger.Error("failed to convert key to hash: %w", err)
		return false, err
	}
	if len(valueStr[2:])%2 != 0 {
		valueStr = "0x0" + valueStr[2:]
	}
	value, err := hex.DecodeString(valueStr[2:])
	if err != nil {
		logger.Error("failed to convert value %s to hash: %w", valueStr, err)
		return false, err
	}

	// Create a new in-memory database
	db := memorydb.New()

	// Decode each proofStr and insert it into the db
	for _, proofStr := range proofStrs {
		node, err := hex.DecodeString(proofStr[2:]) // Remove "0x" prefix and convert to bytes
		if err != nil {
			logger.Error("failed to decode proof node: %w", err)
			return false, err
		}
		key := node
		if len(node) >= 32 { // small MPT nodes are not hashed
			key = crypto.Keccak256(node)
		}
		err = db.Put(key, node)
		if err != nil {
			logger.Error("failed to insert proof node into db: %w", err)
			return false, err
		}
	}

	// Verify the proof
	val, err := trie.VerifyProof(root, path, db)
	logger.Info("VerifyProof", "val", val, "err", err)
	if err != nil {
		return false, fmt.Errorf("failed to verify proof: %w", err)
	}
	//value ==
	logger.Info("VerifyProof", "value", common.BytesToHash(value), "val", common.BytesToHash(val))
	return common.BytesToHash(val) == common.BytesToHash(val), nil
}

// toHash converts a hex string to a common.Hash
func toHash(hexStr string) (common.Hash, error) {
	bytes, err := hex.DecodeString(hexStr[2:]) // Remove "0x" prefix
	if err != nil {
		return common.Hash{}, err
	}
	return common.BytesToHash(bytes), nil
}
