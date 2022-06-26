package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"os"
	"percybolmer/rpc-shard-testing/rpctester/contracts/devtoken"
	"percybolmer/rpc-shard-testing/rpctester/crypto"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/harmony-one/harmony/core/types"
	"github.com/joho/godotenv"
)

const (
	METHOD_V1_getBalanceByBlockNumber = "hmy_getBalanceByBlockNumber"
	METHOD_V2_getBalanceByBlockNumber = "hmyv2_getBalanceByBlockNumber"
	METHOD_V1_getTransactionCount     = "hmy_getTransactionCount"
	METHOD_V2_getTransactionCount     = "hmyv2_getTransactionCount"
	METHOD_V1_getBalance              = "hmy_getBalance"
	METHOD_V2_getBalance              = "hmyv2_getBalance"
	METHOD_address                    = "address"
	/**
	Filter related methods

	*/
	METHOD_filter_getFilterLogs               = "hmy_getFilterLogs"
	METHOD_filter_newFilter                   = "hmy_newFilter"
	METHOD_filter_newPendingtransactionFilter = "hmy_newPendingTransactionFilter"
	METHOD_filter_newBlockFilter              = "hmy_newBlockFilter"
	METHOD_filter_getFilterChanges            = "hmy_getFilterChanges"
	METHOD_filter_getLogs                     = "hmy_getLogs"
	/*
		transactions related methods,
		Get help with the staking transactions
		Since they all fail,
	*/
	METHOD_transaction_V1_getStakingTransactionByBlockHashAndIndex   = "hmy_getStakingTransactionByBlockHashAndIndex"
	METHOD_transaction_V2_getStakingTransactionByBlockHashAndIndex   = "hmyv2_getStakingTransactionByBlockHashAndIndex"
	METHOD_transaction_V1_getStakingTransactionByBlockNumberAndIndex = "hmy_getStakingTransactionByBlockNumberAndIndex"
	METHOD_transaction_V2_getStakingTransactionByBlockNumberAndIndex = "hmyv2_getStakingTransactionByBlockNumberAndIndex"
	METHOD_transaction_V1_getStakingTransactionByHash                = "hmy_getStakingTransactionByHash"
	METHOD_transaction_V2_getStakingTransactionByHash                = "hmyv2_getStakingTransactionByHash"
	METHOD_transaction_V1_getCurrentTransactionErrorSink             = "hmy_getCurrentTransactionErrorSink"
	METHOD_transaction_V2_getCurrentTransactionErrorSink             = "hmyv2_getCurrentTransactionErrorSink"
	//
	METHOD_transaction_V1_getPendingCrossLinks = "hmy_getPendingCrossLinks"
	METHOD_transaction_V2_getPendingCrossLinks = "hmyv2_getPendingCrossLinks"
	METHOD_transaction_V1_getPendingCXReceipts = "hmy_getPendingCXReceipts"
	METHOD_transaction_V2_getPendingCXReceipts = "hmyv2_getPendingCXReceipts"
	METHOD_transaction_V1_getCXReceiptByHash   = "hmy_getCXReceiptByHash"
	METHOD_transaction_V2_getCXReceiptByHash   = "hmyv2_getCXReceiptByHash"
	METHOD_transaction_V1_pendingTransactions  = "hmy_pendingTransactions"
	METHOD_transaction_V2_pendingTransactions  = "hmyv2_pendingTransactions"
	// TODO How do I format the RawStaking Transaction?
	METHOD_transaction_sendRawStakingTransaction              = "hmy_sendRawStakingTransaction"
	METHOD_transaction_sendRawTransaction                     = "hmy_sendRawTransaction"
	METHOD_transaction_V1_getTransactionHistory               = "hmy_getTransactionsHistory"
	METHOD_transaction_V2_getTransactionHistory               = "hmyv2_getTransactionsHistory"
	METHOD_transaction_V1_getTransactionReceipt               = "hmy_getTransactionReceipt"
	METHOD_transaction_V2_getTransactionReceipt               = "hmyv2_getTransactionReceipt"
	METHOD_transaction_V1_getBlockTransactionCountByHash      = "hmy_getBlockTransactionCountByHash"
	METHOD_transaction_V2_getBlockTransactionCountByHash      = "hmyv2_getBlockTransactionCountByHash"
	METHOD_transaction_V1_getBlockTransactionCountByNumber    = "hmy_getBlockTransactionCountByNumber"
	METHOD_transaction_V2_getBlockTransactionCountByNumber    = "hmyv2_getBlockTransactionCountByNumber"
	METHOD_transaction_V1_getTransactionByHash                = "hmy_getTransactionByHash"
	METHOD_transaction_V2_getTransactionByHash                = "hmyv2_getTransactionByHash"
	METHOD_transaction_V1_getTransactionByBlockNumberAndIndex = "hmy_getTransactionByBlockNumberAndIndex"
	METHOD_transaction_V2_getTransactionByBlockNumberAndIndex = "hmyv2_getTransactionByBlockNumberAndIndex"
	METHOD_transaction_V1_getTransactionByBlockHashAndIndex   = "hmy_getTransactionByBlockHashAndIndex"
	METHOD_transaction_V2_getTransactionByBlockHashAndIndex   = "hmyv2_getTransactionByBlockHashAndIndex"
	METHOD_transaction_V1_getBlockByNumber                    = "hmy_getBlockByNumber"
	METHOD_transaction_V2_getBlockByNumber                    = "hmyv2_getBlockByNumber"
	METHOD_transaction_V1_getBlockByHash                      = "hmy_getBlockByHash"
	METHOD_transaction_V2_getBlockByHash                      = "hmyv2_getBlockByHash"
	METHOD_transaction_V1_getBlocks                           = "hmy_getBlocks"
	METHOD_transaction_V2_getBlocks                           = "hmyv2_getBlocks"
	METHOD_transaction_tx                                     = "tx"

	/**
	Contract related
	*/
	METHOD_contract_getStorageAt = "hmy_getStorageAt"
	METHOD_contract_getCode      = "hmy_getCode"
	METHOD_contract_call         = "hmy_call"
	METHOD_contract_estimateGas  = "hmy_estimateGas"

	/**
	Protocol Related
	*/
	METHOD_protocol_isLastBlock          = "hmy_isLastBlock"
	METHOD_protocol_epochLastBlock       = "hmy_epochLastBlock"
	METHOD_protocol_lastestHeader        = "hmy_latestHeader"
	METHOD_protocol_getShardingStructure = "hmy_getShardingStructure"
	METHOD_protocol_V1_blockNumber       = "hmy_blockNumber"
	METHOD_protocol_V2_blockNumber       = "hmyv2_blockNumber"
	METHOD_protocol_syncing              = "hmy_syncing"
	METHOD_protocol_V1_gasPrice          = "hmy_gasPrice"
	METHOD_protocol_V2_gasPrice          = "hmyv2_gasPrice"
	METHOD_protocol_peerCount            = "net_peerCount"
	METHOD_protocol_V1_getEpoch          = "hmy_getEpoch"
	METHOD_protocol_V2_getEpoch          = "hmyv2_getEpoch"
	METHOD_protocol_getLeader            = "hmy_getLeader"

	/**
	Staking Related Methods
	*/
	METHOD_staking_getCirculatingSupply                    = "hmy_getCirculatingSupply"
	METHOD_staking_getTotalSupply                          = "hmy_getTotalSupply"
	METHOD_staking_getStakingNetworkInfo                   = "hmy_getStakingNetworkInfo"
	METHOD_staking_getAllValidatorInformation              = "hmy_getAllValidatorInformation"
	METHOD_staking_getAllValidatorInformationByBlockNumber = "hmy_getAllValidatorInformationByBlockNumber"
	METHOD_staking_getCurrentUtilityMetrics                = "hmy_getCurrentUtilityMetrics"
	// getDelegationsByValidator has the wrong data type in the DOCS
	// for Undelegations, should I update or not
	METHOD_staking_getDelegationsByValidator             = "hmy_getDelegationsByValidator"
	METHOD_staking_getDelegationsByDelegatorAndValidator = "hmy_getDelegationsByDelegatorAndValidator"
	METHOD_staking_getDelegationsByDelegator             = "hmy_getDelegationsByDelegator"
	METHOD_staking_getValidatorMetrics                   = "hmy_getValidatorMetrics"
	METHOD_staking_getMedianRawStakeSnapshot             = "hmy_getMedianRawStakeSnapshot"
	METHOD_staking_getActiveValidatorAddresses           = "hmy_getActiveValidatorAddresses"
	METHOD_staking_getAllValidatorAddresses              = "hmy_getAllValidatorAddresses"
	METHOD_staking_V1_getCurrentStakingErrorSink         = "hmy_getCurrentStakingErrorSink"
	METHOD_staking_V2_getCurrentStakingErrorSink         = "hmyv2_getCurrentStakingErrorSink"
	METHOD_staking_getValidatorInformation               = "hmy_getValidatorInformation"
	METHOD_staking_V1_getValidators                      = "hmy_getValidators"
	METHOD_staking_V2_getValidators                      = "hmyv2_getValidators"
	METHOD_staking_getSignedBlocks                       = "hmy_getSignedBlocks"
	METHOD_staking_V1_isBlockSigner                      = "hmy_isBlockSigner"
	METHOD_staking_V2_isBlockSigner                      = "hmyv2_isBlockSigner"
	METHOD_staking_V1_getBlockSigners                    = "hmy_getBlockSigners"
	METHOD_staking_V2_getBlockSigners                    = "hmyv2_getBlockSigners"

	/**
	Tracing methods
	*/
	METHOD_trace_block       = "trace_block"
	METHOD_trace_transaction = "trace_transaction"
)

var (
	httpClient    *http.Client
	testMetrics   []TestMetric
	ethClient     *ethclient.Client
	auth          *bind.TransactOpts
	deployedToken *devtoken.Devtoken

	address                     string
	url                         string
	smartContractDeploymentHash string
	smartContractAddress        common.Address
	// Bigint representation of 1
	ONE *big.Int
)

func init() {
	// Dont like global http Client, but in this case it might make somewhat sense since we only want to test
	httpClient = &http.Client{Timeout: time.Duration(5) * time.Second}
	testMetrics = []TestMetric{}

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
	} else {
		if err := godotenv.Load(".env"); err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	address = os.Getenv("ADDRESS")
	url = os.Getenv("NET_URL")
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
		Network:     url,
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
	resp, err := httpClient.Post(fmt.Sprintf("%s", url), "application/json", bytes.NewBuffer(payload))
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

	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/address", url), nil)
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
	response.Method = METHOD_address
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
