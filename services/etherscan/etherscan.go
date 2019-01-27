package etherscan

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

const (
	// rinkebyBaseAPI is the base URL for the Etherscan API (Rinkeby)
	rinkebyBaseAPI = "https://api-rinkeby.etherscan.io/api"
	// baseAPI is the base URL for the Etherscan API (Mainnet)
	baseAPI = "https://api.etherscan.io/api"
)

// add Etherscan struct specific for task

// DidInteractWithContract method proves whether an address has recently interacted with a contract
func DidInteractWithContract() (string, error) {
	// Etherscan API key
	etherscanAPIKey := os.Getenv("ETHERSCAN_API_KEY")

	// user's wallet address
	// current: test address
	userAddress := "0x561A370e07ba44E61D9e478aF618D7e839E674C0"

	userEtherscanTxURL := fmt.Sprintf("%s?module=account&action=txlist&address=%s&startblock=0&endblock=99999999&sort=asc&apikey=%s", rinkebyBaseAPI, userAddress, etherscanAPIKey)

	client := &http.Client{
		Timeout: time.Duration(time.Second * 10),
	}

	req, err := http.NewRequest("GET", userEtherscanTxURL, nil)
	if err != nil {
		err = fmt.Errorf("[!] Error preparing GET request\n%v", err)
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		err = fmt.Errorf("[!] Error fetching etherscan transactions\n%v", err)
		return "", err
	}
	defer res.Body.Close()

	byteArr, err := ioutil.ReadAll(res.Body)
	if err != nil {
		err = fmt.Errorf("[!] Error reading response body\n%v", err)
		return "", err
	}

	return string(byteArr), nil

	// TODO:
	// unmarshal into struct
	// prevent address from being swapped with another address that has already submitted transaction
	// proposed solution: discourage use of another user's wallet address as that is where the tokens will be sent to, and down the line potentially the ERC721 badge
}
