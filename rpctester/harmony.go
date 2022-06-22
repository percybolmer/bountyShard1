package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"os"
	"strconv"
	"time"
)

const (
	METHOD_V1_getBalanceByBlockNumber         = "hmy_getBalanceByBlockNumber"
	METHOD_V2_getBalanceByBlockNumber         = "hmyv2_getBalanceByBlockNumber"
	METHOD_V1_getTransactionCount             = "hmy_getTransactionCount"
	METHOD_V2_getTransactionCount             = "hmyv2_getTransactionCount"
	METHOD_V1_getBalance                      = "hmy_getBalance"
	METHOD_V2_getBalance                      = "hmyv2_getBalance"
	METHOD_address                            = "address"
	METHOD_filter_getFilterLogs               = "hmy_getFilterLogs"
	METHOD_filter_newFilter                   = "hmy_newFilter"
	METHOD_filter_newPendingTranscationFilter = "hmy_newPendingTransactionFilter"
	METHOD_filter_newBlockFilter              = "hmy_newBlockFilter"
	METHOD_filter_getFilterChanges            = "hmy_getFilterChanges"
	METHOD_filter_getLogs                     = "hmy_getLogs"
)

var (
	httpClient  *http.Client
	testMetrics []TestMetric
)

func init() {
	// Dont like global http Client, but in this case it might make somewhat sense since we only want to test
	httpClient = &http.Client{Timeout: time.Duration(5) * time.Second}
	testMetrics = []TestMetric{}

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
