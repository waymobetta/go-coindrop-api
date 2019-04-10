package ethereum

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"log"
	"math/big"
	"os"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/crypto/sha3"
)

var (
	// Private key of funding wallet
	PRIVATE_KEY = os.Getenv("RINKEBY_PRIVATE_KEY")
	// Rinkeby adToken address
	TOKEN_CONTRACT_ADDRESS = "0x2f9F1Bdc0EDa69853A91277D272FeaE608F3c1FB"
)

// SendEther ...
func SendEther(recipientAddress string, ethAmountInWei int64) (string, error) {
	client, err := ethclient.Dial("https://rinkeby.infura.io")
	if err != nil {
		return "", err
	}

	privateKey, err := crypto.HexToECDSA(PRIVATE_KEY)
	if err != nil {
		return "", err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return "", err
	}

	value := big.NewInt(ethAmountInWei) // in wei
	gasLimit := uint64(21000)           // in units
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return "", err
	}

	toAddress := common.HexToAddress(recipientAddress)
	var data []byte
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return "", err
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		return "", err
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return "", err
	}

	transaction := signedTx.Hash().Hex()

	return transaction, nil
}

// SendToken ...
func SendToken(tokenAmount, recipientAddress string) (string, error) {
	client, err := ethclient.Dial("https://rinkeby.infura.io")
	if err != nil {
		return "", err
	}

	privateKey, err := crypto.HexToECDSA(PRIVATE_KEY)
	if err != nil {
		return "", err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return "", err
	}

	value := big.NewInt(0) // in wei (0 eth)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return "", err
	}

	toAddress := common.HexToAddress(recipientAddress)
	tokenAddress := common.HexToAddress(TOKEN_CONTRACT_ADDRESS)

	transferFnSignature := []byte("transfer(address,uint256)")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]

	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)

	amount := new(big.Int)
	amount.SetString(tokenAmount, 10) // sets the value to 1000 tokens, in the token denomination

	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)

	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		To:   &tokenAddress,
		Data: data,
	})
	if err != nil {
		return "", err
	}

	// autoset gas for testing
	gasLimit = 200000
	gasPrice = big.NewInt(41000000000)

	tx := types.NewTransaction(nonce, tokenAddress, value, gasLimit, gasPrice, data)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return "", err
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		return "", err
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return "", err
	}

	transaction := signedTx.Hash().Hex()

	return transaction, nil
}

// VerifyAccount ...
func VerifyAccount(from, sigHex string, msg []byte) (bool, error) {
	fromAddr := common.HexToAddress(from)

	sig := hexutil.MustDecode(sigHex)
	if sig[64] != 27 && sig[64] != 28 {
		err := errors.New("invalid signature")
		return false, err
	}
	sig[64] -= 27

	pubKey, err := crypto.SigToPub(signHash(msg), sig)
	if err != nil {
		return false, err
	}

	recoveredAddr := crypto.PubkeyToAddress(*pubKey)

	return fromAddr == recoveredAddr, nil
}

func signHash(data []byte) []byte {
	msg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(data), data)
	return crypto.Keccak256([]byte(msg))
}
