package merkle_proof_verifier

import (
	"encoding/hex"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type ProofItem struct {
	Left  bool
	Value [32]byte
}

func VerifyProof(rootHashStr string, proofStrs []string, keyStr string, valueStr string) (bool, error) {
	rootHash, err := hex.DecodeString(strings.TrimPrefix(rootHashStr, "0x"))
	if err != nil {
		return false, err
	}
	key, err := hex.DecodeString(strings.TrimPrefix(keyStr, "0x"))
	if err != nil {
		return false, err
	}
	value, err := hex.DecodeString(strings.TrimPrefix(valueStr, "0x"))
	if err != nil {
		return false, err
	}
	var proofItems []ProofItem
	for _, proofStr := range proofStrs {
		proof, err := hex.DecodeString(strings.TrimPrefix(proofStr, "0x"))
		if err != nil {
			return false, err
		}
		proofItems = append(proofItems, ProofItem{Left: true, Value: common.BytesToHash(proof)})
	}
	return verifyProof(common.BytesToHash(rootHash), proofItems, common.BytesToHash(key), value), nil
}

func verifyProof(rootHash [32]byte, proof []ProofItem, keyHash [32]byte, value []byte) bool {
	hash := crypto.Keccak256Hash(keyHash[:], value)
	for _, proofItem := range proof {
		if proofItem.Left {
			hash = crypto.Keccak256Hash(proofItem.Value[:], hash[:])
		} else {
			hash = crypto.Keccak256Hash(hash[:], proofItem.Value[:])
		}
	}
	return hash == rootHash
}
