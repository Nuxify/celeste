package types

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/signer/core/apitypes"
)

type CreateUser struct {
	Email    string
	Password string
	Name     string
}

type CreateUserResult struct {
	WalletAddress string
	SSS2          string
	SSS3          string
}

type ReconstructPrivateKey struct {
	ShareKey      string
	WalletAddress string
}

type SignEIP191 struct {
	ShareKey      string
	WalletAddress string
	Message       string
}

type SignEIP712 struct {
	ShareKey      string
	WalletAddress string
	SignerData    apitypes.TypedData
}

type UpdateUser struct {
	WalletAddress string
	Name          string
}

type UpdateUserPassword struct {
	WalletAddress string
	Password      string
}

type ReconstructPrivateKeyResult struct {
	PublicKeyToAddress string
	PrivateKey         *ecdsa.PrivateKey
	PublicKey          *ecdsa.PublicKey
}
