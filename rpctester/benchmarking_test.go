package main

import (
	"context"
	"testing"
)

func BenchmarkProtocolisLastBlock100(b *testing.B) {
	// run the isLastBlock function b.N times
	lastBlock, err := ethClient.BlockNumber(context.Background())
	if err != nil {
		testMetrics = append(testMetrics, TestMetric{
			Method: "isLastBlock",
			Test:   "isLastBlock",
			Pass:   false,
			Error:  err.Error(),
		})
		b.Error(err)
		return
	}
	for n := 0; n < b.N; n++ {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				func() {
					var result bool
					resp, err := benchmarkCall("Benchmarking_isLastBlock", BaseRequest{
						ID:      "1",
						JsonRPC: "2.0",
						Method:  METHOD_protocol_isLastBlock,
						Params: []interface{}{
							lastBlock,
						},
					}, &result)
					if err != nil {
						b.Error(err)
						return
					}

					if resp.Error != nil {
						b.Error(resp.Error.Message)
						return
					}
				}()
			}
		})

	}

	b.Log(benchMetrics)
}
