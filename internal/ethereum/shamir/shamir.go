package shamir

import (
	"crypto/ecdsa"
	"encoding/base64"
	"log"
	"strings"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/hashicorp/vault/shamir"
)

// ReconstructPrivateKey reconstruct private key
// returns private key and private key hex encoded (using hexutil with 0x stripped)
func ReconstructPrivateKey(sss1, shareKey string) (privateKey *ecdsa.PrivateKey, privateKeyHexEncoded string, err error) {
	byteShare1, err := base64.StdEncoding.DecodeString(sss1)
	if err != nil {
		return nil, "", err
	}

	byteShareKey, err := base64.StdEncoding.DecodeString(shareKey)
	if err != nil {
		return nil, "", err
	}

	shares := [][]byte{
		byteShare1,
		byteShareKey,
	}

	// reconstruct private key
	recoveredPrivateKeyBytes, err := shamir.Combine(shares)
	if err != nil {
		return nil, "", err
	}

	// convert to private key
	privateKey, err = crypto.HexToECDSA(string(recoveredPrivateKeyBytes))
	if err != nil {
		log.Println(err)
		return nil, "", err
	}

	// convert to encoded string
	privateKeyBytes := crypto.FromECDSA(privateKey)
	privateKeyHexEncoded = strings.TrimPrefix(hexutil.Encode(privateKeyBytes), "0x")

	return privateKey, privateKeyHexEncoded, nil
}
