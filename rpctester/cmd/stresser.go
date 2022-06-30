package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"percybolmer/rpc-shard-testing/rpctester/benchmarker"
	"percybolmer/rpc-shard-testing/rpctester/harmony"
	"percybolmer/rpc-shard-testing/rpctester/methods"
	"runtime"
	"time"

	"github.com/spf13/cobra"
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

func stressStakingMethods() {
	log.Println("Benchmarking Staking methods")

}

func stressProtocolMethods() {
	log.Println("Benchmarking Protocol methods")
	benchmarkMethod(methods.METHOD_protocol_isLastBlock, BuildRequestGenerator(methods.METHOD_protocol_isLastBlock, []interface{}{"0x1"}))
	benchmarkMethod(methods.METHOD_protocol_epochLastBlock, BuildRequestGenerator(methods.METHOD_protocol_epochLastBlock, []interface{}{"0x1"}))
	benchmarkMethod(methods.METHOD_protocol_lastestHeader, BuildRequestGenerator(methods.METHOD_protocol_lastestHeader, []interface{}{}))
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
	responesRecieved, totalDuration, failures := bencher.Consumer(respChan)

	average := totalDuration / responesRecieved
	averageString := time.Duration(average)
	result.Methods[method] = MethodResult{
		Average:   averageString.String(),
		Failures:  failures,
		Responses: responesRecieved,
	}
}
