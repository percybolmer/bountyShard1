package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"os"
	"percybolmer/rpc-shard-testing/rpctester/benchmarker"
	"percybolmer/rpc-shard-testing/rpctester/harmony"
	"percybolmer/rpc-shard-testing/rpctester/methods"
	"runtime"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/sha3"
)

func init() {
	rootCmd.AddCommand(stressCMD)
}

// BaseRequest is the base structure of requests
type BaseRequest struct {
	ID      string `json:"id"`
	JsonRPC string `json:"jsonrpc"`
	Method  string `json:"method"`
	// Params holds the arguments
	Params []interface{} `json:"params"`
}

var stressCMD = &cobra.Command{
	Use:   "stress",
	Short: "Stresser will stress test each endpoint available",
	Long:  ``,
	Run:   stressTest,
}

var (
	requestsToSend int
	concurrent     int
	result         Result
	/**
	Global data shared acrross the stressers
	*/
	networkHeader            harmony.NetworkHeader
	SelectedValidatorAddress string
	TransactionHash          string
	StakingTransactionHash   string
	TransactionReceipt       harmony.TransactionReceipt_V2
	TransactionReceiptV1     harmony.TransactionReceipt_V1
)

type Result struct {
	AddressUsed string `json:"addressUsed"`
	Network     string `json:"network"`
	Methods     map[string]MethodResult
}

// Each method contains their result
type MethodResult struct {
	Average   string // in nano
	Responses int64
	Failures  int64
	Data      []byte `json:"-"`
}

// init makes sure that we add the needed flags for this particular cobra command
func init() {
	rootCmd.AddCommand(stressCMD)

	stressCMD.Flags().IntVarP(&requestsToSend, "requests", "r", 1000, "The amount of requests to send to every endpoint")
	stressCMD.Flags().IntVarP(&concurrent, "concurrent", "c", 100, "The concurrent amount of requests to send")

	stressCMD.MarkFlagRequired("requests")
	stressCMD.MarkFlagRequired("concurrent")

	result = Result{
		Methods: make(map[string]MethodResult),
	}
}

func stressTest(cmd *cobra.Command, args []string) {
	// Max out CPU
	runtime.GOMAXPROCS(runtime.NumCPU())

	// Apply result data
	result.AddressUsed = harmony.TestAddress
	result.Network = harmony.URL

	stressProtocolMethods()
	stressStakingMethods()
	stressContractMethods()
	stressAccountMethods()
	stressFilterMethods()
	stressTransactionMethods()
	stressTraceMethods()
	// After all tests, Generate report
	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile("stress-result.json", data, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
}

func stressTraceMethods() {
	log.Println("Benchmarking Trance methods")
	benchmarkMethod(methods.METHOD_trace_block, BuildRequestGenerator(methods.METHOD_trace_block, []interface{}{networkHeader.BlockNumber}))
	benchmarkMethod(methods.METHOD_trace_transaction, BuildRequestGenerator(methods.METHOD_trace_transaction, []interface{}{TransactionHash}))

}

func stressTransactionMethods() {
	log.Println("Benchmarking Transaction Methods")
	benchmarkMethod(methods.METHOD_staking_V2_getElectedValidatorAddresses, BuildRequestGenerator(methods.METHOD_staking_V2_getElectedValidatorAddresses, []interface{}{}))
	// Transfer 0.001 ONE
	toAddr := common.HexToAddress(harmony.TestAddress)
	fromAddr := common.HexToAddress(harmony.TestAddress)

	rlp, err := harmony.CreateRLPString(toAddr, fromAddr, *big.NewInt(0).Div(harmony.ONE, big.NewInt(1000)), nil)
	if err != nil {
		log.Fatal("Failedd to create RLP string for transaction benching")
	}
	benchmarkMethod(methods.METHOD_transaction_sendRawTransaction, BuildRequestGenerator(methods.METHOD_transaction_sendRawTransaction, []interface{}{rlp}))
	if err := GetMethodResponse(methods.METHOD_transaction_sendRawTransaction, &TransactionHash); err != nil {
		log.Println("Failed transaction test", err.Error())
	}
	payload, err := harmony.CreateStakingRLPString(harmony.TestAddress, SelectedValidatorAddress, big.NewInt(0).Mul(harmony.ONE, big.NewInt(101)), nil)
	if err != nil {
		log.Println("Failed to create Staking transaction")
	} else {
		// Only delegate on Devnet
		if harmony.URL == "https://api.s0.ps.hmny.io/" || harmony.URL == "https://api.s0.pops.one/" {
			benchmarkMethod(methods.METHOD_transaction_sendRawStakingTransaction, BuildRequestGenerator(methods.METHOD_transaction_sendRawStakingTransaction, []interface{}{payload}))
			if err := GetMethodResponse(methods.METHOD_transaction_sendRawStakingTransaction, &StakingTransactionHash); err != nil {
				log.Println("Failed transaction test", err.Error())
			}
		}
	}
	benchmarkMethod(methods.METHOD_transaction_V1_getCurrentTransactionErrorSink, BuildRequestGenerator(methods.METHOD_transaction_V1_getCurrentTransactionErrorSink, []interface{}{}))
	benchmarkMethod(methods.METHOD_transaction_V2_getCurrentTransactionErrorSink, BuildRequestGenerator(methods.METHOD_transaction_V2_getCurrentTransactionErrorSink, []interface{}{}))
	benchmarkMethod(methods.METHOD_transaction_V1_getPendingCrossLinks, BuildRequestGenerator(methods.METHOD_transaction_V1_getPendingCrossLinks, []interface{}{}))
	benchmarkMethod(methods.METHOD_transaction_V2_getPendingCrossLinks, BuildRequestGenerator(methods.METHOD_transaction_V2_getPendingCrossLinks, []interface{}{}))
	benchmarkMethod(methods.METHOD_transaction_V1_getPendingCXReceipts, BuildRequestGenerator(methods.METHOD_transaction_V1_getPendingCXReceipts, []interface{}{}))
	benchmarkMethod(methods.METHOD_transaction_V2_getPendingCXReceipts, BuildRequestGenerator(methods.METHOD_transaction_V2_getPendingCXReceipts, []interface{}{}))
	benchmarkMethod(methods.METHOD_transaction_V1_getCXReceiptByHash, BuildRequestGenerator(methods.METHOD_transaction_V1_getCXReceiptByHash, []interface{}{"0x6b106dc5619c86b6c0cb64b17e5c464e8008e08cf0f1bb0e3fa2657fb42daade"}))
	benchmarkMethod(methods.METHOD_transaction_V2_getCXReceiptByHash, BuildRequestGenerator(methods.METHOD_transaction_V2_getCXReceiptByHash, []interface{}{"0x6b106dc5619c86b6c0cb64b17e5c464e8008e08cf0f1bb0e3fa2657fb42daade"}))
	benchmarkMethod(methods.METHOD_transaction_V1_pendingTransactions, BuildRequestGenerator(methods.METHOD_transaction_V1_pendingTransactions, []interface{}{}))
	benchmarkMethod(methods.METHOD_transaction_V2_pendingTransactions, BuildRequestGenerator(methods.METHOD_transaction_V2_pendingTransactions, []interface{}{}))
	benchmarkMethod(methods.METHOD_transaction_V1_getStakingTransactionByBlockHashAndIndex, BuildRequestGenerator(methods.METHOD_transaction_V1_getStakingTransactionByBlockHashAndIndex, []interface{}{networkHeader.BlockHash, "0x0"}))
	benchmarkMethod(methods.METHOD_transaction_V2_getStakingTransactionByBlockHashAndIndex, BuildRequestGenerator(methods.METHOD_transaction_V2_getStakingTransactionByBlockHashAndIndex, []interface{}{networkHeader.BlockHash, 0}))
	benchmarkMethod(methods.METHOD_transaction_V1_getStakingTransactionByBlockNumberAndIndex, BuildRequestGenerator(methods.METHOD_transaction_V1_getStakingTransactionByBlockNumberAndIndex, []interface{}{networkHeader.BlockNumber, "0x0"}))
	benchmarkMethod(methods.METHOD_transaction_V2_getStakingTransactionByBlockNumberAndIndex, BuildRequestGenerator(methods.METHOD_transaction_V2_getStakingTransactionByBlockNumberAndIndex, []interface{}{networkHeader.BlockNumber, 1}))
	benchmarkMethod(methods.METHOD_transaction_V1_getStakingTransactionByHash, BuildRequestGenerator(methods.METHOD_transaction_V1_getStakingTransactionByHash, []interface{}{StakingTransactionHash}))
	benchmarkMethod(methods.METHOD_transaction_V2_getStakingTransactionByHash, BuildRequestGenerator(methods.METHOD_transaction_V2_getStakingTransactionByHash, []interface{}{StakingTransactionHash}))
	benchmarkMethod(methods.METHOD_transaction_V1_getTransactionHistory, BuildRequestGenerator(methods.METHOD_transaction_V1_getTransactionHistory, []interface{}{harmony.TransactionArguments{
		Address:   harmony.TestAddress,
		TxType:    "ALL",
		FullTx:    true,
		PageSize:  100,
		PageIndex: 0,
	}}))
	benchmarkMethod(methods.METHOD_transaction_V2_getTransactionHistory, BuildRequestGenerator(methods.METHOD_transaction_V2_getTransactionHistory, []interface{}{harmony.TransactionArguments{
		Address:   harmony.TestAddress,
		TxType:    "ALL",
		FullTx:    true,
		PageSize:  100,
		PageIndex: 0,
	}}))
	benchmarkMethod(methods.METHOD_transaction_V1_getTransactionReceipt, BuildRequestGenerator(methods.METHOD_transaction_V1_getTransactionReceipt, []interface{}{TransactionHash}))
	if err := GetMethodResponse(methods.METHOD_transaction_V1_getTransactionReceipt, &TransactionReceiptV1); err != nil {
		log.Println("Failed transaction Receipt", err.Error())
	}
	benchmarkMethod(methods.METHOD_transaction_V2_getTransactionReceipt, BuildRequestGenerator(methods.METHOD_transaction_V2_getTransactionReceipt, []interface{}{TransactionHash}))
	if err := GetMethodResponse(methods.METHOD_transaction_V2_getTransactionReceipt, &TransactionReceipt); err != nil {
		log.Println("Failed transaction Receipt", err.Error())
	}
	type include struct {
		FullTx         bool `json:"fullTx"`
		WithSigners    bool `json:"withSigner"`
		IncludeSigners bool `json:"includeSigners"`
	}

	benchmarkMethod(methods.METHOD_transaction_V1_getBlockTransactionCountByHash, BuildRequestGenerator(methods.METHOD_transaction_V1_getBlockTransactionCountByHash, []interface{}{TransactionHash}))
	benchmarkMethod(methods.METHOD_transaction_V2_getBlockTransactionCountByHash, BuildRequestGenerator(methods.METHOD_transaction_V2_getBlockTransactionCountByHash, []interface{}{TransactionHash}))
	benchmarkMethod(methods.METHOD_transaction_V1_getBlockTransactionCountByNumber, BuildRequestGenerator(methods.METHOD_transaction_V1_getBlockTransactionCountByNumber, []interface{}{TransactionReceipt.BlockNumber}))
	benchmarkMethod(methods.METHOD_transaction_V2_getBlockTransactionCountByNumber, BuildRequestGenerator(methods.METHOD_transaction_V2_getBlockTransactionCountByNumber, []interface{}{TransactionReceipt.BlockNumber}))
	benchmarkMethod(methods.METHOD_transaction_V1_getTransactionByHash, BuildRequestGenerator(methods.METHOD_transaction_V1_getTransactionByHash, []interface{}{TransactionHash}))
	benchmarkMethod(methods.METHOD_transaction_V2_getTransactionByHash, BuildRequestGenerator(methods.METHOD_transaction_V2_getTransactionByHash, []interface{}{TransactionHash}))
	benchmarkMethod(methods.METHOD_transaction_V1_getTransactionByBlockNumberAndIndex, BuildRequestGenerator(methods.METHOD_transaction_V1_getTransactionByBlockNumberAndIndex, []interface{}{TransactionReceiptV1.BlockNumber, TransactionReceiptV1.TransactionIndex}))
	benchmarkMethod(methods.METHOD_transaction_V2_getTransactionByBlockNumberAndIndex, BuildRequestGenerator(methods.METHOD_transaction_V2_getTransactionByBlockNumberAndIndex, []interface{}{TransactionReceipt.BlockNumber, TransactionReceipt.TransactionIndex}))
	benchmarkMethod(methods.METHOD_transaction_V1_getTransactionByBlockHashAndIndex, BuildRequestGenerator(methods.METHOD_transaction_V1_getTransactionByBlockHashAndIndex, []interface{}{TransactionReceiptV1.BlockHash, TransactionReceiptV1.TransactionIndex}))
	benchmarkMethod(methods.METHOD_transaction_V2_getTransactionByBlockHashAndIndex, BuildRequestGenerator(methods.METHOD_transaction_V2_getTransactionByBlockHashAndIndex, []interface{}{TransactionReceipt.BlockHash, TransactionReceipt.TransactionIndex}))
	benchmarkMethod(methods.METHOD_transaction_V1_getBlockByNumber, BuildRequestGenerator(methods.METHOD_transaction_V1_getBlockByNumber, []interface{}{TransactionReceiptV1.BlockNumber, true}))
	benchmarkMethod(methods.METHOD_transaction_V2_getBlockByNumber, BuildRequestGenerator(methods.METHOD_transaction_V2_getBlockByNumber, []interface{}{TransactionReceipt.BlockNumber, include{
		FullTx:         true,
		WithSigners:    false,
		IncludeSigners: false,
	}}))
	benchmarkMethod(methods.METHOD_transaction_V1_getBlockByHash, BuildRequestGenerator(methods.METHOD_transaction_V1_getBlockByHash, []interface{}{TransactionReceiptV1.BlockHash, true}))
	benchmarkMethod(methods.METHOD_transaction_V2_getBlockByHash, BuildRequestGenerator(methods.METHOD_transaction_V2_getBlockByHash, []interface{}{TransactionReceipt.BlockHash, include{FullTx: true, WithSigners: false, IncludeSigners: false}}))
	benchmarkMethod(methods.METHOD_transaction_V1_getBlocks, BuildRequestGenerator(methods.METHOD_transaction_V1_getBlocks, []interface{}{TransactionReceiptV1.BlockNumber, TransactionReceiptV1.BlockNumber}))
	benchmarkMethod(methods.METHOD_transaction_V2_getBlocks, BuildRequestGenerator(methods.METHOD_transaction_V2_getBlocks, []interface{}{TransactionReceipt.BlockNumber, TransactionReceipt.BlockNumber, include{FullTx: true, WithSigners: false, IncludeSigners: false}}))
	benchmarkMethod(methods.METHOD_transaction_tx, BuildRequestGenerator(methods.METHOD_transaction_tx, []interface{}{TransactionHash}))

}

func stressFilterMethods() {
	log.Println("Benchmarking Filter methods")
	benchmarkMethod(methods.METHOD_filter_newFilter, BuildRequestGenerator(methods.METHOD_filter_newFilter, []interface{}{harmony.Filter{FromBlock: "0x1", ToBlock: "0x2", Address: harmony.TestAddress, Topics: []string{"0x000000000000000000000000a94f5374fce5edbc8e2a8697c15331677e6ebf0b"}}}))
	var filterid string
	if err := GetMethodResponse(methods.METHOD_filter_newFilter, &filterid); err != nil {
		log.Println("Missing filter id ", err.Error())
	}

	benchmarkMethod(methods.METHOD_filter_getFilterLogs, BuildRequestGenerator(methods.METHOD_filter_getFilterLogs, []interface{}{filterid}))
	benchmarkMethod(methods.METHOD_filter_newPendingtransactionFilter, BuildRequestGenerator(methods.METHOD_filter_newPendingtransactionFilter, []interface{}{}))
	benchmarkMethod(methods.METHOD_filter_newBlockFilter, BuildRequestGenerator(methods.METHOD_filter_newBlockFilter, []interface{}{}))
	benchmarkMethod(methods.METHOD_filter_getFilterChanges, BuildRequestGenerator(methods.METHOD_filter_getFilterChanges, []interface{}{filterid}))
	benchmarkMethod(methods.METHOD_filter_getLogs, BuildRequestGenerator(methods.METHOD_filter_getLogs, []interface{}{harmony.Filter{Address: harmony.TestAddress}}))

}

func stressAccountMethods() {
	log.Println("Benchmarking Address Methods")
	benchmarkMethod(methods.METHOD_V2_getBalanceByBlockNumber, BuildRequestGenerator(methods.METHOD_V2_getBalanceByBlockNumber, []interface{}{harmony.TestAddress, "1"}))
	benchmarkMethod(methods.METHOD_V1_getBalanceByBlockNumber, BuildRequestGenerator(methods.METHOD_V1_getBalanceByBlockNumber, []interface{}{harmony.TestAddress, "0x01"}))
	benchmarkMethod(methods.METHOD_V1_getTransactionCount, BuildRequestGenerator(methods.METHOD_V1_getTransactionCount, []interface{}{harmony.TestAddress, "latest"}))
	benchmarkMethod(methods.METHOD_V2_getTransactionCount, BuildRequestGenerator(methods.METHOD_V2_getTransactionCount, []interface{}{harmony.TestAddress, 1}))
	benchmarkMethod(methods.METHOD_V1_getBalance, BuildRequestGenerator(methods.METHOD_V1_getBalance, []interface{}{harmony.TestAddress, "latest"}))
	benchmarkMethod(methods.METHOD_V2_getBalance, BuildRequestGenerator(methods.METHOD_V2_getBalance, []interface{}{harmony.TestAddress}))

}

func stressContractMethods() {
	log.Println("Benchmarking Contract Methods")
	benchmarkMethod(methods.METHOD_contract_getStorageAt, BuildRequestGenerator(methods.METHOD_contract_getStorageAt, []interface{}{harmony.SmartContractAddress, "0x0", "latest"}))
	benchmarkMethod(methods.METHOD_contract_getCode, BuildRequestGenerator(methods.METHOD_contract_getCode, []interface{}{harmony.SmartContractAddress, "latest"}))
	/**
	Create Contract call
	*/
	transferFnSignature := []byte("owner()")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]
	data := hexutil.Bytes{}
	data = append(data, methodID...)

	rpcCall := harmony.RpcCallArgs{
		From: common.HexToAddress(harmony.TestAddress).String(),
		To:   harmony.SmartContractAddress.String(),
		Data: data.String(),
	}

	benchmarkMethod(methods.METHOD_contract_call, BuildRequestGenerator(methods.METHOD_contract_call, []interface{}{rpcCall, "latest"}))
	var owner string
	if err := GetMethodResponse(methods.METHOD_contract_call, &owner); err != nil {
		log.Println("Call to smart contract failed", err.Error())
	}

	benchmarkMethod(methods.METHOD_contract_estimateGas, BuildRequestGenerator(methods.METHOD_contract_estimateGas, []interface{}{rpcCall}))

}

func stressStakingMethods() {
	log.Println("Benchmarking Staking methods")
	benchmarkMethod(methods.METHOD_staking_V1_getValidators, BuildRequestGenerator(methods.METHOD_staking_V1_getValidators, []interface{}{networkHeader.Epoch}))
	benchmarkMethod(methods.METHOD_staking_V2_getValidators, BuildRequestGenerator(methods.METHOD_staking_V2_getValidators, []interface{}{networkHeader.Epoch}))
	// Fetch Validators
	var validators harmony.GetValidatorsV2
	if err := GetMethodResponse(methods.METHOD_staking_V2_getValidators, &validators); err != nil {
		log.Fatal("Cant proceed without valid network header: ", err.Error())
	}

	if len(validators.Validators) == 0 {
		log.Fatal("Cant proceed staking methods without validator")
	}
	SelectedValidatorAddress = validators.Validators[0].Address

	benchmarkMethod(methods.METHOD_staking_getCirculatingSupply, BuildRequestGenerator(methods.METHOD_staking_getCirculatingSupply, []interface{}{}))
	benchmarkMethod(methods.METHOD_staking_getTotalSupply, BuildRequestGenerator(methods.METHOD_staking_getTotalSupply, []interface{}{}))
	benchmarkMethod(methods.METHOD_staking_getStakingNetworkInfo, BuildRequestGenerator(methods.METHOD_staking_getStakingNetworkInfo, []interface{}{}))
	benchmarkMethod(methods.METHOD_staking_getAllValidatorInformation, BuildRequestGenerator(methods.METHOD_staking_getAllValidatorInformation, []interface{}{}))
	benchmarkMethod(methods.METHOD_staking_getAllValidatorInformationByBlockNumber, BuildRequestGenerator(methods.METHOD_staking_getAllValidatorInformationByBlockNumber, []interface{}{0, networkHeader.BlockNumber}))
	benchmarkMethod(methods.METHOD_staking_getCurrentUtilityMetrics, BuildRequestGenerator(methods.METHOD_staking_getCurrentUtilityMetrics, []interface{}{}))
	benchmarkMethod(methods.METHOD_staking_getDelegationsByValidator, BuildRequestGenerator(methods.METHOD_staking_getDelegationsByValidator, []interface{}{SelectedValidatorAddress}))
	benchmarkMethod(methods.METHOD_staking_getDelegationsByDelegatorAndValidator, BuildRequestGenerator(methods.METHOD_staking_getDelegationsByDelegatorAndValidator, []interface{}{harmony.TestAddress, SelectedValidatorAddress}))
	benchmarkMethod(methods.METHOD_staking_getDelegationsByDelegator, BuildRequestGenerator(methods.METHOD_staking_getDelegationsByDelegator, []interface{}{harmony.TestAddress}))
	benchmarkMethod(methods.METHOD_staking_getValidatorMetrics, BuildRequestGenerator(methods.METHOD_staking_getValidatorMetrics, []interface{}{SelectedValidatorAddress}))
	benchmarkMethod(methods.METHOD_staking_getMedianRawStakeSnapshot, BuildRequestGenerator(methods.METHOD_staking_getMedianRawStakeSnapshot, []interface{}{}))
	benchmarkMethod(methods.METHOD_staking_getActiveValidatorAddresses, BuildRequestGenerator(methods.METHOD_staking_getActiveValidatorAddresses, []interface{}{}))
	benchmarkMethod(methods.METHOD_staking_V1_getAllValidatorAddresses, BuildRequestGenerator(methods.METHOD_staking_V1_getAllValidatorAddresses, []interface{}{}))
	benchmarkMethod(methods.METHOD_staking_V2_getAllValidatorAddresses, BuildRequestGenerator(methods.METHOD_staking_V2_getAllValidatorAddresses, []interface{}{}))
	benchmarkMethod(methods.METHOD_staking_V1_getCurrentStakingErrorSink, BuildRequestGenerator(methods.METHOD_staking_V1_getCurrentStakingErrorSink, []interface{}{}))
	benchmarkMethod(methods.METHOD_staking_V2_getCurrentStakingErrorSink, BuildRequestGenerator(methods.METHOD_staking_V2_getCurrentStakingErrorSink, []interface{}{}))
	benchmarkMethod(methods.METHOD_staking_getValidatorInformation, BuildRequestGenerator(methods.METHOD_staking_getValidatorInformation, []interface{}{SelectedValidatorAddress}))
	benchmarkMethod(methods.METHOD_staking_getSignedBlocks, BuildRequestGenerator(methods.METHOD_staking_getSignedBlocks, []interface{}{SelectedValidatorAddress}))
	benchmarkMethod(methods.METHOD_staking_V1_isBlockSigner, BuildRequestGenerator(methods.METHOD_staking_V1_isBlockSigner, []interface{}{"0x0", SelectedValidatorAddress}))
	benchmarkMethod(methods.METHOD_staking_V2_isBlockSigner, BuildRequestGenerator(methods.METHOD_staking_V2_isBlockSigner, []interface{}{networkHeader.BlockNumber, SelectedValidatorAddress}))
	benchmarkMethod(methods.METHOD_staking_V1_getBlockSigners, BuildRequestGenerator(methods.METHOD_staking_V1_getBlockSigners, []interface{}{"0x1"}))
	benchmarkMethod(methods.METHOD_staking_V2_getBlockSigners, BuildRequestGenerator(methods.METHOD_staking_V2_getBlockSigners, []interface{}{networkHeader.BlockNumber}))

}

func stressProtocolMethods() {
	log.Println("Benchmarking Protocol methods")
	benchmarkMethod(methods.METHOD_protocol_lastestHeader, BuildRequestGenerator(methods.METHOD_protocol_lastestHeader, []interface{}{}))
	// Method should contain network data, fetch it

	if err := GetMethodResponse(methods.METHOD_protocol_lastestHeader, &networkHeader); err != nil {
		log.Fatal("Cant proceed without valid network header: ", err.Error())
	}
	benchmarkMethod(methods.METHOD_protocol_isLastBlock, BuildRequestGenerator(methods.METHOD_protocol_isLastBlock, []interface{}{"0x1"}))
	benchmarkMethod(methods.METHOD_protocol_epochLastBlock, BuildRequestGenerator(methods.METHOD_protocol_epochLastBlock, []interface{}{"0x1"}))
	benchmarkMethod(methods.METHOD_protocol_getShardingStructure, BuildRequestGenerator(methods.METHOD_protocol_getShardingStructure, []interface{}{}))
	benchmarkMethod(methods.METHOD_protocol_V1_blockNumber, BuildRequestGenerator(methods.METHOD_protocol_V1_blockNumber, []interface{}{}))
	benchmarkMethod(methods.METHOD_protocol_V2_blockNumber, BuildRequestGenerator(methods.METHOD_protocol_V2_blockNumber, []interface{}{}))
	benchmarkMethod(methods.METHOD_protocol_syncing, BuildRequestGenerator(methods.METHOD_protocol_syncing, []interface{}{}))
	benchmarkMethod(methods.METHOD_protocol_V1_gasPrice, BuildRequestGenerator(methods.METHOD_protocol_V1_gasPrice, []interface{}{}))
	benchmarkMethod(methods.METHOD_protocol_V2_gasPrice, BuildRequestGenerator(methods.METHOD_protocol_V2_gasPrice, []interface{}{}))
	benchmarkMethod(methods.METHOD_protocol_peerCount, BuildRequestGenerator(methods.METHOD_protocol_peerCount, []interface{}{}))
	benchmarkMethod(methods.METHOD_protocol_V1_getEpoch, BuildRequestGenerator(methods.METHOD_protocol_V1_getEpoch, []interface{}{}))
	benchmarkMethod(methods.METHOD_protocol_V2_getEpoch, BuildRequestGenerator(methods.METHOD_protocol_V2_getEpoch, []interface{}{}))
	benchmarkMethod(methods.METHOD_protocol_getLeader, BuildRequestGenerator(methods.METHOD_protocol_getLeader, []interface{}{}))

}

func GetMethodResponse(method string, into interface{}) error {
	response := result.Methods[method].Data
	var baseresp harmony.BaseResponse
	err := json.Unmarshal(response, &baseresp)
	if err != nil {
		log.Fatal("Cannot continue stress tester without network header data from LatestHeader")
	}

	return json.Unmarshal(baseresp.Result, into)
}

// BuildRequestGenerator is a wrapper util to generate requests
func BuildRequestGenerator(method string, params []interface{}) benchmarker.GenerateRequestFunc {
	return func() *http.Request {
		br := BaseRequest{
			ID:      "1",
			JsonRPC: "2.0",
			Method:  method,
			Params:  params,
		}
		payload, err := json.Marshal(br)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(string(payload))
		req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s", harmony.URL), bytes.NewBuffer(payload))
		if err != nil {
			log.Fatal("Your request constructer is broken")
		}
		req.Header.Add("content-type", "application/json")
		return req
	}
}

// benchmarkMethod is general wrapper around calling an benchmark
func benchmarkMethod(method string, requestGen benchmarker.GenerateRequestFunc) {
	bencher := benchmarker.NewBenchmarker(requestsToSend, concurrent)
	runtime.GOMAXPROCS(runtime.NumCPU())
	reqChan := make(chan *http.Request)
	respChan := make(chan benchmarker.Response)

	go bencher.Dispatcher(reqChan, requestGen)
	go bencher.WorkerPool(reqChan, respChan)
	responesRecieved, totalDuration, failures, data := bencher.Consumer(respChan)

	average := totalDuration / responesRecieved
	averageString := time.Duration(average)
	result.Methods[method] = MethodResult{
		Average:   averageString.String(),
		Failures:  failures,
		Responses: responesRecieved,
		Data:      data,
	}
}
