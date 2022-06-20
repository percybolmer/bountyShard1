package rpctester

import (
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// address is used in tests to find related addr
var address = flag.String("address", "0xA5241513DA9F4463F1d4874b548dFBAC29D91f34", "The Address to use")
var url = flag.String("url", "http://localhost:9500", "the network rpc URL")

const (
	METHOD_V1_getBalanceByBlockNumber = "hmy_getBalanceByBlockNumber"
	METHOD_V2_getBalanceByBlockNumber = "hmyv2_getBalanceByBlockNumber"
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

type TestMetric struct {
	Method   string `json:"method"`
	Test     string `json:"test"`
	Pass     bool   `json:"pass"`
	Duration string `json:"duration"`
	Error    string `json:"error,omitempty"`
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

type RPCError struct {
	Code    int64
	Message string
}

// Call will trigger a request with a payload to the RPC method given and marshal response into interface
func Call(payload []byte, method string) (*BaseResponse, error) {
	// Store request time in Response

	start := time.Now()
	resp, err := httpClient.Post(fmt.Sprintf("%s/%s", *url, method), "application/json", bytes.NewBuffer(payload))
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
