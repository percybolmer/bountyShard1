package main

import (
	"context"
	"fmt"
	"math/big"
	"testing"
)

func test_isLastBlock(t *testing.T) {
	type testcase struct {
		name              string
		br                BaseRequest
		expectedErrorCode int64
		expectedReturn    bool
	}
	lastBlock, err := ethClient.BlockNumber(context.Background())
	if err != nil {
		testMetrics = append(testMetrics, TestMetric{
			Method: "isLastBlock",
			Test:   "isLastBlock",
			Pass:   false,
			Error:  err.Error(),
		})
		t.Error(err)
		return
	}

	testCases := []testcase{
		{
			name:           fmt.Sprintf("%s_not_the_last_block", t.Name()),
			expectedReturn: false,
			br: BaseRequest{
				ID:      "1",
				JsonRPC: "2.0",
				Method:  METHOD_protocol_isLastBlock,
				Params: []interface{}{
					10,
				},
			},
		}, {
			name:           fmt.Sprintf("%s_is_last", t.Name()),
			expectedReturn: true,
			br: BaseRequest{
				ID:      "1",
				JsonRPC: "2.0",
				Method:  METHOD_protocol_isLastBlock,
				Params: []interface{}{
					lastBlock,
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var result bool

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

func test_epochLastBlock(t *testing.T) {
	type testcase struct {
		name              string
		br                BaseRequest
		expectedErrorCode int64
		expectedReturn    bool
	}

	testCases := []testcase{
		{
			name:           fmt.Sprintf("%s_epochLastBlock", t.Name()),
			expectedReturn: false,
			br: BaseRequest{
				ID:      "1",
				JsonRPC: "2.0",
				Method:  METHOD_protocol_epochLastBlock,
				Params: []interface{}{
					10,
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var result int64

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

func (ts *testSuite) test_latestHeader(t *testing.T) {
	type testcase struct {
		name              string
		br                BaseRequest
		expectedErrorCode int64
		expectedReturn    bool
	}

	testCases := []testcase{
		{
			name:           fmt.Sprintf("%s_latestHeader", t.Name()),
			expectedReturn: false,
			br: BaseRequest{
				ID:      "1",
				JsonRPC: "2.0",
				Method:  METHOD_protocol_lastestHeader,
				Params:  []interface{}{},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var result NetworkHeader

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
			ts.NetworkHeader = result

		})
	}
}

func test_getShardingStructure(t *testing.T) {
	type testcase struct {
		name              string
		br                BaseRequest
		expectedErrorCode int64
		expectedReturn    bool
	}

	testCases := []testcase{
		{
			name:           fmt.Sprintf("%s_shardingStructure", t.Name()),
			expectedReturn: false,
			br: BaseRequest{
				ID:      "1",
				JsonRPC: "2.0",
				Method:  METHOD_protocol_getShardingStructure,
				Params:  []interface{}{},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var result shardingStructure

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

func test_V1_blockNumber(t *testing.T) {
	type testcase struct {
		name              string
		br                BaseRequest
		expectedErrorCode int64
	}

	testCases := []testcase{
		{
			name: fmt.Sprintf("%s_block_number", t.Name()),
			br: BaseRequest{
				ID:      "1",
				JsonRPC: "2.0",
				Method:  METHOD_protocol_V1_blockNumber,
				Params:  []interface{}{},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var result string

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
func test_V2_blockNumber(t *testing.T) {
	type testcase struct {
		name              string
		br                BaseRequest
		expectedErrorCode int64
	}

	testCases := []testcase{
		{
			name: fmt.Sprintf("%s_V2_block_number", t.Name()),
			br: BaseRequest{
				ID:      "1",
				JsonRPC: "2.0",
				Method:  METHOD_protocol_V2_blockNumber,
				Params:  []interface{}{},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var result int64

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
func test_syncing(t *testing.T) {
	type testcase struct {
		name              string
		br                BaseRequest
		expectedErrorCode int64
	}

	testCases := []testcase{
		{
			name: fmt.Sprintf("%s_syncing", t.Name()),
			br: BaseRequest{
				ID:      "1",
				JsonRPC: "2.0",
				Method:  METHOD_protocol_syncing,
				Params:  []interface{}{},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var result bool

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

func test_V1_gasPrice(t *testing.T) {
	type testcase struct {
		name              string
		br                BaseRequest
		expectedErrorCode int64
	}

	testCases := []testcase{
		{
			name: fmt.Sprintf("%s_gasPrice_V1", t.Name()),
			br: BaseRequest{
				ID:      "1",
				JsonRPC: "2.0",
				Method:  METHOD_protocol_V1_gasPrice,
				Params:  []interface{}{},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var result string

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
func test_V2_gasPrice(t *testing.T) {
	type testcase struct {
		name              string
		br                BaseRequest
		expectedErrorCode int64
	}

	testCases := []testcase{
		{
			name: fmt.Sprintf("%s_V2_gas_Price", t.Name()),
			br: BaseRequest{
				ID:      "1",
				JsonRPC: "2.0",
				Method:  METHOD_protocol_V2_gasPrice,
				Params:  []interface{}{},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var result big.Int

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

func test_peerCount(t *testing.T) {
	type testcase struct {
		name              string
		br                BaseRequest
		expectedErrorCode int64
	}

	testCases := []testcase{
		{
			name: fmt.Sprintf("%s_peerCount", t.Name()),
			br: BaseRequest{
				ID:      "1",
				JsonRPC: "2.0",
				Method:  METHOD_protocol_peerCount,
				Params:  []interface{}{},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var result string

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

func test_V1_getEpoch(t *testing.T) {
	type testcase struct {
		name              string
		br                BaseRequest
		expectedErrorCode int64
	}

	testCases := []testcase{
		{
			name: fmt.Sprintf("%s_getEpoch_V1", t.Name()),
			br: BaseRequest{
				ID:      "1",
				JsonRPC: "2.0",
				Method:  METHOD_protocol_V1_getEpoch,
				Params:  []interface{}{},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var result string

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
func test_V2_getEpoch(t *testing.T) {
	type testcase struct {
		name              string
		br                BaseRequest
		expectedErrorCode int64
	}

	testCases := []testcase{
		{
			name: fmt.Sprintf("%s_V2_get_Epoch", t.Name()),
			br: BaseRequest{
				ID:      "1",
				JsonRPC: "2.0",
				Method:  METHOD_protocol_V2_getEpoch,
				Params:  []interface{}{},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var result big.Int

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

func test_getLeader(t *testing.T) {
	type testcase struct {
		name              string
		br                BaseRequest
		expectedErrorCode int64
	}

	testCases := []testcase{
		{
			name: fmt.Sprintf("%s_get_leader", t.Name()),
			br: BaseRequest{
				ID:      "1",
				JsonRPC: "2.0",
				Method:  METHOD_protocol_getLeader,
				Params:  []interface{}{},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var result string

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
