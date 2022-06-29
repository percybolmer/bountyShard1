package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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

func stressTest(cmd *cobra.Command, args []string) {

	bencher := benchmarker.NewBenchmarker(100000, 100)
	runtime.GOMAXPROCS(runtime.NumCPU())
	reqChan := make(chan *http.Request)
	respChan := make(chan benchmarker.Response)
	start := time.Now()
	go bencher.Dispatcher(reqChan, func() *http.Request {
		br := BaseRequest{
			ID:      "1",
			JsonRPC: "2.0",
			Method:  methods.METHOD_protocol_isLastBlock,
			Params: []interface{}{
				"0x1",
			},
		}
		payload, err := json.Marshal(br)
		if err != nil {
			log.Fatal(err)
		}
		req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s", harmony.URL), bytes.NewBuffer(payload))
		if err != nil {
			log.Fatal("Your request constructer is broken")
		}
		return req
	})
	go bencher.WorkerPool(reqChan, respChan)
	conns, size := bencher.Consumer(respChan)
	took := time.Since(start)
	ns := took.Nanoseconds()
	av := ns / conns
	average, err := time.ParseDuration(fmt.Sprintf("%d", av) + "ns")
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("Connections:\t%d\nConcurrent:\t%d\nTotal size:\t%d bytes\nTotal time:\t%s\nAverage time:\t%s\n", conns, max, size, took, average)

}
