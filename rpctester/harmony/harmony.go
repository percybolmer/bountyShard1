package harmony

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"net"
	"net/http"
	"os"
	"percybolmer/rpc-shard-testing/rpctester/contracts/devtoken"
	"percybolmer/rpc-shard-testing/rpctester/crypto"
	"percybolmer/rpc-shard-testing/rpctester/methods"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/harmony-one/harmony/core/types"

	staking "github.com/harmony-one/harmony/staking/types"
	"github.com/joho/godotenv"
)

var (
	httpClient    *http.Client
	testMetrics   []TestMetric
	benchMetrics  map[string][]TestMetric
	ethClient     *ethclient.Client
	auth          *bind.TransactOpts
	deployedToken *devtoken.Devtoken

	address                     string
	URL                         string
	smartContractDeploymentHash string
	smartContractAddress        common.Address
	// Bigint representation of 1
	ONE *big.Int
)

func init() {
	// Dont like global http Client, but in this case it might make somewhat sense since we only want to test
	httpClient = &http.Client{Timeout: time.Duration(5) * time.Second, Transport: &http.Transport{
		Dial: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 10 * time.Second,
	}}
	testMetrics = []TestMetric{}
	benchMetrics = map[string][]TestMetric{}

	ONE = big.NewInt(1000000000000000000)

	// Load .dotenv file
	environment := os.Getenv("RPCTESTER_ENVIRONMENT")

	if environment == "production" {
		if err := godotenv.Load(".env-prod"); err != nil {
			log.Fatal("Error loading .env-prod file")
		}
	} else if environment == "dev" {
		if err := godotenv.Load(".env-dev"); err != nil {
			log.Fatal("Error loading .env-dev file")
		}
	} else if environment == "test" {
		if err := godotenv.Load(".env-test"); err != nil {
			log.Fatal("Error loading .env-test file")
		}
	} else {
		if err := godotenv.Load(".env"); err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	address = os.Getenv("ADDRESS")
	URL = os.Getenv("NET_URL")
	smartContractAddr := os.Getenv("SMART_CONTRACT_ADDRESS")
	smartContractDeploymentHash = os.Getenv("SMART_CONTRACT_DEPLOY_HASH")
	// Create eth client
	ethClient, auth = crypto.NewClient()
	// Load the Smart Contract
	smartContractAddress = common.HexToAddress(smartContractAddr)
	instance, err := devtoken.NewDevtoken(smartContractAddress, ethClient)
	if err != nil {
		log.Fatal(err)
	}
	deployedToken = instance

}

// TypeResults is the results of all the tests
type TestResults struct {
	AddressUsed string       `json:"addressUsed"`
	Network     string       `json:"network"`
	Metrics     []TestMetric `json:"metrics"`
}

type TestMetric struct {
	Method   string        `json:"method"`
	Test     string        `json:"test"`
	Pass     bool          `json:"pass"`
	Duration string        `json:"duration"`
	Error    string        `json:"error,omitempty"`
	Params   []interface{} `json:"params,omitempty"`
}

// BaseRequest is the base structure of requests
type BaseRequest struct {
	ID      string `json:"id"`
	JsonRPC string `json:"jsonrpc"`
	Method  string `json:"method"`
	// Params holds the arguments
	Params []interface{} `json:"params"`
}

// BaseResponse is the base RPC response, but added extra metrics
type BaseResponse struct {
	ID      string          `json:"id"`
	JsonRPC string          `json:"jsonrpc"`
	Result  json.RawMessage `json:"result"`
	// Error is only present when errors occurs
	Error *RPCError `json:"error,omitempty"`
	// Custom data fields not part of rpc response
	Method string `json:"method"`
	// Not part of the default message
	Duration string `json:"duration"`
}

type AddressResponse struct {
	ID           string        `json:"id"`
	Balance      big.Int       `json:"balance"`
	Transactions []Transaction `json:"txs" `
	// Custom data fields not part of rpc response
	Method string `json:"method"`
	// Not part of the default message
	Duration string `json:"duration"`
}

type RPCError struct {
	Code    int64
	Message string
}

func GenerateReport() {
	results := TestResults{
		AddressUsed: address,
		Network:     URL,
		Metrics:     testMetrics,
	}
	// After all tests, Generate report
	data, err := json.Marshal(results)
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile("results.json", data, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
}

// Call will trigger a request with a payload to the RPC method given and marshal response into interface
func Call(payload []byte, method string) (*BaseResponse, error) {
	// Store request time in Response

	start := time.Now()
	resp, err := httpClient.Post(fmt.Sprintf("%s", URL), "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}
	duration := time.Since(start)

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}
	// Read data
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var br BaseResponse
	err = json.Unmarshal(body, &br)
	if err != nil {
		return nil, err
	}
	// Add Extra metrics
	br.Duration = duration.String()
	br.Method = method

	return &br, nil
}

// Addres fetches address, this does not work as the other rpc calls
func Address(id string, offset, page int, tx_view string) (*AddressResponse, error) {
	client := &http.Client{}

	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/address", URL), nil)
	req.Header.Add("Content-Type", "application/json")

	q := req.URL.Query()
	q.Add("id", id)
	q.Add("offset", strconv.Itoa(offset))
	q.Add("page", strconv.Itoa(page))
	q.Add("tx_view", tx_view)
	req.URL.RawQuery = q.Encode()
	start := time.Now()
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	duration := time.Since(start)
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}
	// Read data
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response AddressResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}
	response.Method = methods.METHOD_address
	response.Duration = duration.String()

	return &response, nil
}

// CreateRLPString is a wrapper to help with generating RLP
func CreateRLPString(to common.Address, from common.Address, amount big.Int, data []byte) (string, error) {
	nonce, err := ethClient.PendingNonceAt(context.Background(), from)
	if err != nil {
		return "", err
	}

	shardID, err := crypto.GetShardID()
	if err != nil {
		return "", err
	}
	gasPrice, err := ethClient.SuggestGasPrice(context.Background())
	if err != nil {
		return "", err
	}
	var selectedGasLimit uint64
	if data != nil {
		gasLimit, err := ethClient.EstimateGas(context.Background(), ethereum.CallMsg{
			To:   &to,
			Data: data,
		})
		if err != nil {
			return "", err
		}
		selectedGasLimit = gasLimit
	}
	if selectedGasLimit == 0 {
		selectedGasLimit = auth.GasLimit
	}
	// Create TX with Harmony Flavor
	hmy_tx := types.NewTransaction(nonce, to, uint32(shardID), &amount, selectedGasLimit, gasPrice, data)

	chainID, err := ethClient.ChainID(context.Background())
	if err != nil {
		return "", err
	}

	signedTx, err := types.SignTx(hmy_tx, types.NewEIP155Signer(chainID), crypto.GetPrivateKey())
	if err != nil {
		return "", err
	}
	ts := types.Transactions{signedTx}

	rawTxHex := hexutil.Encode(ts.GetRlp(0))
	return rawTxHex, nil
}

// CreateStakingRLPString is a wrapper to help with generating RLP for staking requests
// delgator and validator is sent as ONE accounts format ie bech32
// addr
func CreateStakingRLPString(delegator string, validator string, amount *big.Int, data []byte) (string, error) {
	nonce, err := ethClient.PendingNonceAt(context.Background(), common.HexToAddress(address))
	if err != nil {
		return "", err
	}

	gasPrice, err := ethClient.SuggestGasPrice(context.Background())
	if err != nil {
		return "", err
	}
	selectedGasLimit := auth.GasLimit

	// Create sTAKING TX with Harmony Flavor
	delegateStakePayloadMaker := func() (staking.Directive, interface{}) {
		return staking.DirectiveDelegate, staking.Delegate{
			Parse(delegator),
			Parse(validator),
			amount,
		}

	}
	_, payload := delegateStakePayloadMaker()

	data, err = rlp.EncodeToBytes(payload)
	if err != nil {
		return "", err
	}

	stakingTx, err := staking.NewStakingTransaction(nonce, selectedGasLimit, gasPrice, delegateStakePayloadMaker)

	// Currently Bugged on Testnet & Devnet REturns ChainiD 16777....?
	// Trnasaciton needs 4 or 2
	// chainID, err := ethClient.ChainID(context.Background())
	// if err != nil {
	// 	return "", err
	// }
	// HardCode to 2 for now
	signedTx, err := staking.Sign(stakingTx, staking.NewEIP155Signer(big.NewInt(2)), crypto.GetPrivateKey())
	if err != nil {
		return "", err
	}

	enc, err := rlp.EncodeToBytes(signedTx)
	if err != nil {
		return "", err
	}
	hexSignature := hexutil.Encode(enc)

	return hexSignature, nil
}
