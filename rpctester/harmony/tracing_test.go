package harmony

import (
	"fmt"
	"percybolmer/rpc-shard-testing/rpctester/methods"
	"testing"
)

func (ts *testSuite) test_traceBlock(t *testing.T) {
	type testcase struct {
		name              string
		br                BaseRequest
		expectedErrorCode int64
	}

	testCases := []testcase{
		{
			name: fmt.Sprintf("%s_traceBlock", t.Name()),
			br: BaseRequest{
				ID:      "1",
				JsonRPC: "2.0",
				Method:  methods.METHOD_trace_block,
				Params: []interface{}{
					ts.LastTransactionBlockNumber,
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var result []TraceBlock

			resp, err := callAndValidateDataType(t, tc.name, tc.expectedErrorCode, tc.br, &result)
			if err != nil {
				t.Error(err)
				return
			}

			testMetrics = append(testMetrics, TestMetric{
				Method:   tc.br.Method,
				Test:     tc.name,
				Pass:     true,
				Duration: resp.Duration,
				Params:   tc.br.Params,
			})

		})
	}
}

func (ts *testSuite) test_traceTransaction(t *testing.T) {
	type testcase struct {
		name              string
		br                BaseRequest
		expectedErrorCode int64
	}

	testCases := []testcase{
		{
			name: fmt.Sprintf("%s_trace_existing_transaction", t.Name()),
			br: BaseRequest{
				ID:      "1",
				JsonRPC: "2.0",
				Method:  methods.METHOD_trace_transaction,
				Params: []interface{}{
					ts.LastTransactionHash,
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var result []TraceBlock

			resp, err := callAndValidateDataType(t, tc.name, tc.expectedErrorCode, tc.br, &result)
			if err != nil {
				t.Error(err)
				return
			}

			testMetrics = append(testMetrics, TestMetric{
				Method:   tc.br.Method,
				Test:     tc.name,
				Pass:     true,
				Duration: resp.Duration,
				Params:   tc.br.Params,
			})

		})
	}
}
