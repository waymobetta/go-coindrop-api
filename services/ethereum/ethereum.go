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
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/waymobetta/go-coindrop-api/services/ethereum/erc721"
	"golang.org/x/crypto/sha3"
)

var (
	// Private key of funding wallet
	PRIVATE_KEY = os.Getenv("RINKEBY_PRIVATE_KEY")
	// Private key for LOCAL_NODE
	// PRIVATE_KEY = "0f98477e4825d6b682184acefc51d4ab7ecd6b86ce29f3c984534fc1d53786e7"
	// Infura API key
	INFURA_API_KEY = os.Getenv("INFURA_API_KEY")
	// Infura Rinkeby URL
	INFURA_RINKEBY = fmt.Sprintf("https://rinkeby.infura.io/v3/%s", INFURA_API_KEY)
	// Infura Mainnet URL
	INFURA_MAINNET = fmt.Sprintf("https://mainnet.infura.io/v3/%s", INFURA_API_KEY)
	// Local Ethereum Node
	LOCAL_NODE = "http://localhost:8545"
	// Ethereum Client endpoint
	ETHEREUM_CLIENT_URL = INFURA_RINKEBY
	// Rinkeby adToken address
	TOKEN_CONTRACT_ADDRESS = "0x2f9F1Bdc0EDa69853A91277D272FeaE608F3c1FB"
	// Custom Errors
	publicKeyError = errors.New("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
)

// DeployERC721Contract is an external method that deploys a new ERC-721 token contract given the contract name and token symbol
func DeployERC721Contract(tokenName, tokenSymbol string) (common.Address, error) {
	client, err := ethclient.Dial(ETHEREUM_CLIENT_URL)
	if err != nil {
		return common.HexToAddress(""), err
	}

	privateKey, err := crypto.HexToECDSA(PRIVATE_KEY)
	if err != nil {
		return common.HexToAddress(""), err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return common.HexToAddress(""), publicKeyError
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return common.HexToAddress(""), err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return common.HexToAddress(""), err
	}

	auth := bind.NewKeyedTransactor(privateKey)

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = gasPrice

	address, tx, _, err := erc721.DeployErc721(
		auth,
		client,
		tokenName,
		tokenSymbol,
	)

	if err != nil {
		return common.HexToAddress(""), err
	}
	fmt.Printf("https://rinkeby.etherscan.io/tx/%s\n", tx.Hash().Hex())
	return address, nil
}

// MintERC721ForUser is an external method that mints a unique ERC-721 token from a given ERC-721 contract address for a given recipient address
func MintERC721Token(
	tokenId int,
	contractAddress,
	recipientAddress common.Address,
	tokenURI string,
) (string, error) {
	tId := big.NewInt(int64(tokenId))

	log.Println("[+] connecting to ethereum..")
	client, err := ethclient.Dial(ETHEREUM_CLIENT_URL)
	if err != nil {
		return "", err
	}
	log.Println("[*] connected to ethereum")

	privateKey, err := crypto.HexToECDSA(PRIVATE_KEY)
	if err != nil {
		return "", err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return "", publicKeyError
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	log.Println("[+] setting transaction details..")

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return "", err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return "", err
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(200000) // in units
	auth.Nonce = big.NewInt(int64(nonce))
	auth.GasPrice = gasPrice

	log.Println("[*] transaction details set")

	instance, err := erc721.NewErc721(contractAddress, client)
	if err != nil {
		return "", err
	}

	log.Println("[+] minting new ERC-721 token for:", recipientAddress.Hex())

	tx, err := instance.Erc721Transactor.MintWithTokenURI(
		auth,
		recipientAddress,
		tId,
		tokenURI,
	)
	if err != nil {
		return "", err
	}
	log.Printf("[TRANSACTION] https://rinkeby.etherscan.io/tx/%s\n", tx.Hash().Hex())
	log.Println("[*] new ERC-721 successfully minted for:", recipientAddress.Hex())
	return tx.Hash().Hex(), nil
}

// GetTotalSupply is an external method that retrieves the total token supply at a given ERC-721 contract address
func GetTotalSupply(contractAddress common.Address) (*big.Int, error) {
	log.Println("[+] connecting to ethereum..")
	client, err := ethclient.Dial(ETHEREUM_CLIENT_URL)
	if err != nil {
		return big.NewInt(0), err
	}
	log.Println("[*] connected to ethereum")
	instance, err := erc721.NewErc721(contractAddress, client)
	if err != nil {
		return big.NewInt(0), err
	}

	log.Println("[+] getting total supply of ERC-721 tokens..")

	totalSupply, err := instance.Erc721Caller.TotalSupply(nil)
	if err != nil {
		return big.NewInt(0), err
	}
	log.Println("[*] total supply of ERC-721 tokens:", totalSupply)
	return totalSupply, nil
}

// GetBalanceOf is an external method that retrieves the token balance at a given ERC-721 contract address by a given owner address
func GetBalanceOf(ownerAddress, contractAddress common.Address) (*big.Int, error) {
	log.Println("[+] connecting to ethereum..")
	client, err := ethclient.Dial(ETHEREUM_CLIENT_URL)
	if err != nil {
		return big.NewInt(0), err
	}
	log.Println("[*] connected to ethereum")

	instance, err := erc721.NewErc721(contractAddress, client)
	if err != nil {
		return big.NewInt(0), err
	}

	log.Println("[+] getting total token balance of wallet:", ownerAddress.Hex())

	ownerSupply, err := instance.Erc721Caller.BalanceOf(nil, ownerAddress)
	if err != nil {
		return big.NewInt(0), err
	}
	log.Printf("[*] %s has a total ERC-721 token balance of: %v", ownerAddress.Hex(), ownerSupply)
	return ownerSupply, nil
}

// GetTokenSymbol is an external method that retrieves the ERC-721 token symbol at a given ERC-721 contract address
func GetTokenSymbol(contractAddress common.Address) (string, error) {
	log.Println("[+] connecting to ethereum..")
	client, err := ethclient.Dial(ETHEREUM_CLIENT_URL)
	if err != nil {
		return "", err
	}
	log.Println("[*] connected to ethereum")

	instance, err := erc721.NewErc721(contractAddress, client)
	if err != nil {
		return "", err
	}

	log.Println("[+] getting token symbol of contract:", contractAddress.Hex())

	tokenSymbol, err := instance.Erc721Caller.Symbol(nil)
	if err != nil {
		return "", err
	}
	return tokenSymbol, nil
}

// GetTokenName is an external method that retrieves the ERC-721 token name at a given ERC-721 contract address
func GetTokenName(contractAddress common.Address) (string, error) {
	log.Println("[+] connecting to ethereum..")
	client, err := ethclient.Dial(ETHEREUM_CLIENT_URL)
	if err != nil {
		return "", err
	}
	log.Println("[*] connected to ethereum")

	instance, err := erc721.NewErc721(contractAddress, client)
	if err != nil {
		return "", err
	}

	log.Println("[+] getting name of contract:", contractAddress.Hex())

	tokenName, err := instance.Erc721Caller.Name(nil)
	if err != nil {
		return "", err
	}
	return tokenName, nil
}

// SendEther is an external method that sends a given number of ether (in Wei) to a given recipient
func SendEther(recipientAddress string, ethAmountInWei int64) (string, error) {
	client, err := ethclient.Dial(ETHEREUM_CLIENT_URL)
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
		return "", publicKeyError
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

// SendERC20Token is an external method that sends a given number of ERC-20 tokens from a pre-deployed contract to a given recipient
func SendERC20Token(tokenAmount, recipientAddress string) (string, error) {
	client, err := ethclient.Dial(ETHEREUM_CLIENT_URL)
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

// VerifyAddress verifies the ownership of a given ethereum address
func VerifyAddress(from, sigHex string, msg []byte) error {
	fromAddr := common.HexToAddress(from)

	sig := hexutil.MustDecode(sigHex)
	if sig[64] != 27 && sig[64] != 28 {
		err := errors.New("invalid signature")
		return err
	}
	sig[64] -= 27

	pubKey, err := crypto.SigToPub(signHash(msg), sig)
	if err != nil {
		return err
	}

	recoveredAddr := crypto.PubkeyToAddress(*pubKey)

	if fromAddr != recoveredAddr {
		return errors.New("from address and recovered address do not match")
	}

	return nil
}

// signHash is an internal method that re-signs the message in order to compare its validity
func signHash(data []byte) []byte {
	msg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(data), data)
	return crypto.Keccak256([]byte(msg))
}
