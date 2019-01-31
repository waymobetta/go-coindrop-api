package etherscan

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

const (
	// rinkebyBaseAPI is the base URL for the Etherscan API (Rinkeby)
	rinkebyBaseAPI = "https://api-rinkeby.etherscan.io/api"
	// baseAPI is the base URL for the Etherscan API (Mainnet)
	baseAPI = "https://api.etherscan.io/api"
)

// TODO:
// update structs

// ContractUser ...
type ContractUser struct {
	UserID  string `json:"user_id"`
	Address string `json:"address"`
}

// TxResponse ...
type TxResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Result  []struct {
		BlockNumber       string `json:"blockNumber"`
		TimeStamp         string `json:"timeStamp"`
		Hash              string `json:"hash"`
		Nonce             string `json:"nonce"`
		BlockHash         string `json:"blockHash"`
		TransactionIndex  string `json:"transactionIndex"`
		From              string `json:"from"`
		To                string `json:"to"`
		Value             string `json:"value"`
		Gas               string `json:"gas"`
		GasPrice          string `json:"gasPrice"`
		IsError           string `json:"isError"`
		TxreceiptStatus   string `json:"txreceipt_status"`
		Input             string `json:"input"`
		ContractAddress   string `json:"contractAddress"`
		CumulativeGasUsed string `json:"cumulativeGasUsed"`
		GasUsed           string `json:"gasUsed"`
		Confirmations     string `json:"confirmations"`
	} `json:"result"`
}

// DidInteractWithContract method proves whether an address has recently interacted with a contract
func (u *User) DidInteractWithContract(contract string) (bool, error) {
	txRes := new(TxResponse)

	etherscanAPIKey := os.Getenv("ETHERSCAN_API_KEY")

	baseURL := `https://api-rinkeby.etherscan.io/api`

	userEtherscanTxURL := fmt.Sprintf("%s?module=account&action=txlist&address=%s&startblock=0&endblock=99999999&sort=asc&apikey=%s", baseURL, u.Address, etherscanAPIKey)

	client := &http.Client{
		Timeout: time.Duration(time.Second * 10),
	}

	req, err := http.NewRequest("GET", userEtherscanTxURL, nil)
	if err != nil {
		log.Errorf("[etherscan] Error preparing GET request; %v\n", err)
		return false, err
	}

	req.Header.Set("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		log.Errorf("[etherscan] Error fetching etherscan transactions; %v\n", err)
		return false, err
	}
	defer res.Body.Close()

	byteArr, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Errorf("[etherscan] Error reading response body; %v\n", err)
		return false, err
	}

	if err := json.Unmarshal(byteArr, &txRes); err != nil {
		log.Errorf("[etherscan] Error unmarshalling JSON; %v\n", err)
		return false, err
	}

	for index := range txRes.Result {
		if strings.ToUpper(txRes.Result[index].To) == strings.ToUpper(contract) && txRes.Result[index].TxreceiptStatus == "1" {
			log.Println("[eterhscan] tx found!")
			return true, nil
		}
	}
	return false, nil
}
