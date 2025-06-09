package service

import (
	"context"
	"crypto/ecdsa"
	"encoding/base64"
	"errors"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/hashicorp/vault/shamir"
	"github.com/segmentio/ksuid"

	apiError "celeste/internal/errors"
	"celeste/internal/ethereum/eip191"
	"celeste/internal/ethereum/eip712"
	shamirInternal "celeste/internal/ethereum/shamir"
	"celeste/internal/password"
	"celeste/module/user/domain/entity"
	"celeste/module/user/domain/repository"
	repositoryTypes "celeste/module/user/infrastructure/repository/types"
	"celeste/module/user/infrastructure/service/types"
)

// UserCommandService handles the user command service logic
type UserCommandService struct {
	repository.UserCommandRepositoryInterface
	repository.UserQueryRepositoryInterface
}

// CreateUser create a user
func (service *UserCommandService) CreateUser(ctx context.Context, data types.CreateUser) (types.CreateUserResult, error) {
	// generate wallet
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Println(err)
		return types.CreateUserResult{}, err
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)
	privateKeyEncoded := hexutil.Encode(privateKeyBytes)[2:] // strip 0x

	publicKey := privateKey.Public()
	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)
	publicAddress := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

	// apply Shamir Secret Sharing (SSS)
	bytesShares, err := shamir.Split([]byte(privateKeyEncoded), 3, 2) // 2 of 3
	if err != nil {
		log.Println(err)
		return types.CreateUserResult{}, err
	}

	var sss []string
	for _, byteShare := range bytesShares {
		sss = append(sss, base64.StdEncoding.EncodeToString(byteShare))
	}

	fmt.Println(privateKeyEncoded)
	fmt.Println("#1: ", sss[0])
	fmt.Println("#2: ", sss[1])
	fmt.Println("#3: ", sss[2])

	// TODO: remove test combine
	byteShare1, err := base64.StdEncoding.DecodeString(sss[0])
	if err != nil {
		return types.CreateUserResult{}, err
	}
	//byteShare2, err := base64.StdEncoding.DecodeString(sss[1])
	// if err != nil {
	// 	return "", err
	// }
	byteShare3, err := base64.StdEncoding.DecodeString(sss[2])
	if err != nil {
		return types.CreateUserResult{}, err
	}

	shares := [][]byte{
		byteShare1,
		byteShare3,
	}

	recovered, err := shamir.Combine(shares)
	if err != nil {
		return types.CreateUserResult{}, err
	}

	fmt.Println(string(recovered), string(recovered) == privateKeyEncoded)
	// TODO: remove above test code

	sss1 := sss[0] // for user database record
	sss2 := sss[1] // for device
	sss3 := sss[2] // for backup

	// hash password
	hashedPassword, err := password.HashPassword(data.Password)
	if err != nil {
		return types.CreateUserResult{}, err
	}

	err = service.UserCommandRepositoryInterface.InsertUser(repositoryTypes.CreateUser{
		WalletAddress: publicAddress,
		Email:         data.Email,
		Password:      hashedPassword,
		SSS1:          sss1,
		Name:          data.Name,
	})
	if err != nil {
		return types.CreateUserResult{}, err
	}

	return types.CreateUserResult{
		WalletAddress: publicAddress,
		SSS2:          sss2,
		SSS3:          sss3,
	}, nil
}

// DeactivateUser deactivates user
func (service *UserCommandService) DeactivateUser(ctx context.Context, walletAddress string) error {
	err := service.UserCommandRepositoryInterface.DeactivateUser(repositoryTypes.DeactivateUser{
		WalletAddress: walletAddress,
		Email:         fmt.Sprintf("%s@deactivated.user", walletAddress),
		Password:      "",
		SSS1:          "",
		Name:          "Deactivated User",
	})
	if err != nil {
		return err
	}

	return nil
}

// ReconstructPrivateKey reconstructs the private key
func (service *UserCommandService) ReconstructPrivateKey(ctx context.Context, data types.ReconstructPrivateKey) (string, error) {
	// reconstruct sss private key
	recoveredKey, _, err := service.recoverKey(data.WalletAddress, data.ShareKey)
	if err != nil {
		return "", err
	}

	return crypto.PubkeyToAddress(*recoveredKey.PublicKey).Hex(), nil
}

// SignEIP191 signs a message using EIP-191
func (service *UserCommandService) SignEIP191(ctx context.Context, data types.SignEIP191) (string, error) {
	// reconstruct sss private key
	recoveredKey, user, err := service.recoverKey(data.WalletAddress, data.ShareKey)
	if err != nil {
		return "", err
	}

	if crypto.PubkeyToAddress(*recoveredKey.PublicKey).Hex() != user.WalletAddress {
		return "", errors.New(apiError.UnauthorizedAccess)
	}

	// sign message
	signature, err := eip191.SignPersonalMessage(recoveredKey.PrivateKey, []byte(data.Message))
	if err != nil {
		return "", err
	}

	return signature, nil
}

// SignEIP712 signs a message using EIP-712
func (service *UserCommandService) SignEIP712(ctx context.Context, data types.SignEIP712) (string, error) {
	// reconstruct signer private key
	signerKey, user, err := service.recoverKey(data.WalletAddress, data.ShareKey)
	if err != nil {
		return "", errors.New(apiError.EthInvalidUserPrivateKey)
	}

	if crypto.PubkeyToAddress(*signerKey.PublicKey).Hex() != user.WalletAddress {
		return "", errors.New(apiError.UnauthorizedAccess)
	}

	// get signature
	signature, err := eip712.SignTypedData(data.SignerData, signerKey.PrivateKey)
	if err != nil {
		log.Println(err)
		return "", errors.New(apiError.EthInvalidTypedDataSignature)
	}

	return signature, nil
}

// UpdateUser update user by address
func (service *UserCommandService) UpdateUser(ctx context.Context, data types.UpdateUser) error {
	err := service.UserCommandRepositoryInterface.UpdateUser(repositoryTypes.UpdateUser{
		WalletAddress: data.WalletAddress,
		Name:          data.Name,
	})
	if err != nil {
		return err
	}

	return nil
}

// UpdateUserEmailVerifiedAt update user email verified at by address
func (service *UserCommandService) UpdateUserEmailVerifiedAt(ctx context.Context, email string) error {
	err := service.UserCommandRepositoryInterface.UpdateUserEmailVerifiedAt(email)
	if err != nil {
		return err
	}

	return nil
}

// UpdateUserPassword update user password by address
func (service *UserCommandService) UpdateUserPassword(ctx context.Context, data types.UpdateUserPassword) error {
	hashedPassword, err := password.HashPassword(data.Password)
	if err != nil {
		return err
	}

	err = service.UserCommandRepositoryInterface.UpdateUserPassword(repositoryTypes.UpdateUserPassword{
		WalletAddress: data.WalletAddress,
		Password:      hashedPassword,
	})
	if err != nil {
		return err
	}

	return nil
}

// generateID generates unique id
func generateID() string {
	return ksuid.New().String()
}

func (service *UserCommandService) recoverKey(walletAddress, shareKey string) (types.RecoveredKey, entity.User, error) {
	// get user by wallet address
	user, err := service.SelectUserByWalletAddress(walletAddress)
	if err != nil {
		return types.RecoveredKey{}, entity.User{}, err
	}

	// reconstruct signer private key
	recoveredPrivateKey, _, err := shamirInternal.ReconstructPrivateKey(user.SSS1, shareKey)
	if err != nil {
		return types.RecoveredKey{}, entity.User{}, errors.New(apiError.EthInvalidUserPrivateKey)
	}

	recoveredPublicKey := recoveredPrivateKey.Public()
	recoveredPublicKeyECDSA, ok := recoveredPublicKey.(*ecdsa.PublicKey)
	if !ok {
		return types.RecoveredKey{}, entity.User{}, errors.New(apiError.EthInvalidUserPublicKey)
	}

	return types.RecoveredKey{
		PrivateKey: recoveredPrivateKey,
		PublicKey:  recoveredPublicKeyECDSA,
	}, user, nil
}
