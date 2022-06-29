package harmony

import (
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"percybolmer/rpc-shard-testing/rpctester/methods"
	"testing"
)

func (ts *testSuite) test_V1_getStakingTransactionByBlockHashAndIndex(t *testing.T) {
	// https://github.com/harmony-one/bounties/issues/117#issuecomment-1170274370
	// Skip unless devnet
	if URL != "https://api.s0.b.hmny.io/" && URL != "https://api.s0.pops.one/" {
		t.Skip("only run staking on testnet")
	}
	type testcase struct {
		name              string
		br                BaseRequest
		expectedErrorCode int64
	}

	testCases := []testcase{
		{
			name: fmt.Sprintf("%s_stake_transaction_exists", t.Name()),
			br: BaseRequest{
				ID:      "1",
				JsonRPC: "2.0",
				Method:  methods.METHOD_transaction_V1_getStakingTransactionByBlockHashAndIndex,
				Params: []interface{}{
					ts.NetworkHeader.BlockHash,
					"0x0",
				},
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

func (ts *testSuite) test_V2_getStakingTransactionByBlockHashAndIndex(t *testing.T) {
	// https://github.com/harmony-one/bounties/issues/117#issuecomment-1170274370
	// Skip unless devnet
	if URL != "https://api.s0.b.hmny.io/" && URL != "https://api.s0.pops.one/" {
		t.Skip("only run staking on testnet")
	}
	type testcase struct {
		name              string
		br                BaseRequest
		expectedErrorCode int64
	}

	testCases := []testcase{
		{
			name: fmt.Sprintf("%s_stake_transaction_exists", t.Name()),
			br: BaseRequest{
				ID:      "1",
				JsonRPC: "2.0",
				Method:  methods.METHOD_transaction_V2_getStakingTransactionByBlockHashAndIndex,
				Params: []interface{}{
					ts.NetworkHeader.BlockHash,
					0,
				},
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

func (ts *testSuite) test_V1_getStakingTransactionByBlockNumberAndIndex(t *testing.T) {
	// https://github.com/harmony-one/bounties/issues/117#issuecomment-1170274370
	// Skip unless devnet
	if URL != "https://api.s0.b.hmny.io/" && URL != "https://api.s0.pops.one/" {
		t.Skip("only run staking on testnet")
	}
	type testcase struct {
		name              string
		br                BaseRequest
		expectedErrorCode int64
	}

	testCases := []testcase{
		{
			name: fmt.Sprintf("%s_staking_by_Block_and_Index", t.Name()),
			br: BaseRequest{
				ID:      "1",
				JsonRPC: "2.0",
				Method:  methods.METHOD_transaction_V1_getStakingTransactionByBlockNumberAndIndex,
				Params: []interface{}{
					ts.NetworkHeader.BlockNumber,
					"0x0",
				},
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

func (ts *testSuite) test_V2_getStakingTransactionByBlockNumberAndIndex(t *testing.T) {
	// https://github.com/harmony-one/bounties/issues/117#issuecomment-1170274370
	// Skip unless devnet
	if URL != "https://api.s0.b.hmny.io/" && URL != "https://api.s0.pops.one/" {
		t.Skip("only run staking on testnet")
	}
	type testcase struct {
		name              string
		br                BaseRequest
		expectedErrorCode int64
	}

	testCases := []testcase{
		{
			name: fmt.Sprintf("%s_staking_by_Block_and_Index", t.Name()),
			br: BaseRequest{
				ID:      "1",
				JsonRPC: "2.0",
				Method:  methods.METHOD_transaction_V2_getStakingTransactionByBlockNumberAndIndex,
				Params: []interface{}{
					ts.NetworkHeader.BlockNumber,
					1,
				},
			},
		},
		{
			name:              fmt.Sprintf("%s_transaction index 0", t.Name()),
			expectedErrorCode: -32000,
			br: BaseRequest{
				ID:      "1",
				JsonRPC: "2.0",
				Method:  methods.METHOD_transaction_V2_getStakingTransactionByBlockNumberAndIndex,
				Params: []interface{}{
					ts.NetworkHeader.BlockNumber,
					0,
				},
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

func (ts *testSuite) test_V2_getStakingTransactionByHash(t *testing.T) {
	// https://github.com/harmony-one/bounties/issues/117#issuecomment-1170274370
	// Skip unless devnet
	if URL != "https://api.s0.b.hmny.io/" && URL != "https://api.s0.pops.one/" {
		t.Skip("only run staking on testnet")
	}
	type testcase struct {
		name              string
		br                BaseRequest
		expectedErrorCode int64
	}

	testCases := []testcase{
		{
			name: fmt.Sprintf("%s_staking_by_hash", t.Name()),
			br: BaseRequest{
				ID:      "1",
				JsonRPC: "2.0",
				Method:  methods.METHOD_transaction_V2_getStakingTransactionByHash,
				Params: []interface{}{
					ts.LastStakingTransactionHash,
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			data, err := json.Marshal(tc.br)
			if err != nil {
				testMetrics = append(testMetrics, TestMetric{
					Method: tc.br.Method,
					Test:   tc.name,
					Pass:   false,
					Error:  err.Error(),
					Params: tc.br.Params,
				})
				t.Error(err)
				return
			}
			// Perform the RPC Call
			resp, err := Call(data, tc.br.Method)
			if err != nil {
				testMetrics = append(testMetrics, TestMetric{
					Method:   tc.br.Method,
					Test:     tc.name,
					Pass:     false,
					Duration: resp.Duration,
					Error:    err.Error(),
					Params:   tc.br.Params,
				})
				t.Error(err)
				return
			}
			if resp.Error != nil {
				if resp.Error.Code != tc.expectedErrorCode {
					testMetrics = append(testMetrics, TestMetric{
						Method:   tc.br.Method,
						Test:     tc.name,
						Pass:     false,
						Duration: resp.Duration,
						Error:    resp.Error.Message,
						Params:   tc.br.Params,
					})
					t.Error(resp.Error.Message)
					return
				}
			}
			// This step validates that the returned response is the correct data type
			if resp.Result != nil {
				log.Println(string(resp.Result))
				var s string // Wrong Data type, but func does not work
				err = json.Unmarshal(resp.Result, &s)
				if err != nil {
					testMetrics = append(testMetrics, TestMetric{
						Method:   tc.br.Method,
						Test:     tc.name,
						Pass:     false,
						Duration: resp.Duration,
						Error:    err.Error(),
						Params:   tc.br.Params,
					})
					t.Error(err)
					return
				}
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

func (ts *testSuite) test_V1_getStakingTransactionByHash(t *testing.T) {
	// https://github.com/harmony-one/bounties/issues/117#issuecomment-1170274370
	// Skip unless devnet
	if URL != "https://api.s0.b.hmny.io/" && URL != "https://api.s0.pops.one/" {
		t.Skip("only run staking on testnet")
	}
	type testcase struct {
		name              string
		br                BaseRequest
		expectedErrorCode int64
	}

	testCases := []testcase{
		{
			name: fmt.Sprintf("%s_staking_by_hash", t.Name()),
			br: BaseRequest{
				ID:      "1",
				JsonRPC: "2.0",
				Method:  methods.METHOD_transaction_V1_getStakingTransactionByHash,
				Params: []interface{}{
					ts.LastStakingTransactionHash,
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			data, err := json.Marshal(tc.br)
			if err != nil {
				testMetrics = append(testMetrics, TestMetric{
					Method: tc.br.Method,
					Test:   tc.name,
					Pass:   false,
					Error:  err.Error(),
					Params: tc.br.Params,
				})
				t.Error(err)
				return
			}
			// Perform the RPC Call
			resp, err := Call(data, tc.br.Method)
			if err != nil {
				testMetrics = append(testMetrics, TestMetric{
					Method:   tc.br.Method,
					Test:     tc.name,
					Pass:     false,
					Duration: resp.Duration,
					Error:    err.Error(),
					Params:   tc.br.Params,
				})
				t.Error(err)
				return
			}
			if resp.Error != nil {
				if resp.Error.Code != tc.expectedErrorCode {
					testMetrics = append(testMetrics, TestMetric{
						Method:   tc.br.Method,
						Test:     tc.name,
						Pass:     false,
						Duration: resp.Duration,
						Error:    resp.Error.Message,
						Params:   tc.br.Params,
					})
					t.Error(resp.Error.Message)
					return
				}
			}
			// This step validates that the returned response is the correct data type
			if resp.Result != nil {
				var s []string // Wrong Data type, but func does not work
				err = json.Unmarshal(resp.Result, &s)
				if err != nil {
					testMetrics = append(testMetrics, TestMetric{
						Method:   tc.br.Method,
						Test:     tc.name,
						Pass:     false,
						Duration: resp.Duration,
						Error:    err.Error(),
						Params:   tc.br.Params,
					})
					t.Error(err)
					return
				}
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

func test_V1_getCurrentTransactionErrorSink(t *testing.T) {
	type testcase struct {
		name              string
		br                BaseRequest
		expectedErrorCode int64
	}

	testCases := []testcase{
		{
			name: fmt.Sprintf("%s_staking_error_sink", t.Name()),
			br: BaseRequest{
				ID:      "1",
				JsonRPC: "2.0",
				Method:  methods.METHOD_transaction_V1_getCurrentTransactionErrorSink,
				Params:  []interface{}{},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			data, err := json.Marshal(tc.br)
			if err != nil {
				testMetrics = append(testMetrics, TestMetric{
					Method: tc.br.Method,
					Test:   tc.name,
					Pass:   false,
					Error:  err.Error(),
					Params: tc.br.Params,
				})
				t.Error(err)
				return
			}
			// Perform the RPC Call
			resp, err := Call(data, tc.br.Method)
			if err != nil {
				testMetrics = append(testMetrics, TestMetric{
					Method:   tc.br.Method,
					Test:     tc.name,
					Pass:     false,
					Duration: resp.Duration,
					Error:    err.Error(),
					Params:   tc.br.Params,
				})
				t.Error(err)
				return
			}
			if resp.Error != nil {
				if resp.Error.Code != tc.expectedErrorCode {
					testMetrics = append(testMetrics, TestMetric{
						Method:   tc.br.Method,
						Test:     tc.name,
						Pass:     false,
						Duration: resp.Duration,
						Error:    resp.Error.Message,
						Params:   tc.br.Params,
					})
					t.Error(resp.Error.Message)
					return
				}
			}
			// This step validates that the returned response is the correct data type
			if resp.Result != nil {
				var s []ErrorSinkLog // Wrong Data type, but func does not work
				err = json.Unmarshal(resp.Result, &s)
				if err != nil {
					testMetrics = append(testMetrics, TestMetric{
						Method:   tc.br.Method,
						Test:     tc.name,
						Pass:     false,
						Duration: resp.Duration,
						Error:    err.Error(),
						Params:   tc.br.Params,
					})
					t.Error(err)
					return
				}
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

func test_V2_getCurrentTransactionErrorSink(t *testing.T) {
	type testcase struct {
		name              string
		br                BaseRequest
		expectedErrorCode int64
	}

	testCases := []testcase{
		{
			name: fmt.Sprintf("%s_staking_error_sink", t.Name()),
			br: BaseRequest{
				ID:      "1",
				JsonRPC: "2.0",
				Method:  methods.METHOD_transaction_V2_getCurrentTransactionErrorSink,
				Params:  []interface{}{},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			data, err := json.Marshal(tc.br)
			if err != nil {
				testMetrics = append(testMetrics, TestMetric{
					Method: tc.br.Method,
					Test:   tc.name,
					Pass:   false,
					Error:  err.Error(),
					Params: tc.br.Params,
				})
				t.Error(err)
				return
			}
			// Perform the RPC Call
			resp, err := Call(data, tc.br.Method)
			if err != nil {
				testMetrics = append(testMetrics, TestMetric{
					Method:   tc.br.Method,
					Test:     tc.name,
					Pass:     false,
					Duration: resp.Duration,
					Error:    err.Error(),
					Params:   tc.br.Params,
				})
				t.Error(err)
				return
			}
			if resp.Error != nil {
				if resp.Error.Code != tc.expectedErrorCode {
					testMetrics = append(testMetrics, TestMetric{
						Method:   tc.br.Method,
						Test:     tc.name,
						Pass:     false,
						Duration: resp.Duration,
						Error:    resp.Error.Message,
						Params:   tc.br.Params,
					})
					t.Error(resp.Error.Message)
					return
				}
			}
			// This step validates that the returned response is the correct data type
			if resp.Result != nil {
				var s []ErrorSinkLog
				err = json.Unmarshal(resp.Result, &s)
				if err != nil {
					testMetrics = append(testMetrics, TestMetric{
						Method:   tc.br.Method,
						Test:     tc.name,
						Pass:     false,
						Duration: resp.Duration,
						Error:    err.Error(),
						Params:   tc.br.Params,
					})
					t.Error(err)
					return
				}
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

func test_V1_getPendingCrossLinks(t *testing.T) {
	type testcase struct {
		name              string
		br                BaseRequest
		expectedErrorCode int64
	}

	testCases := []testcase{
		{
			name: fmt.Sprintf("%s_pending_cross_links", t.Name()),
			br: BaseRequest{
				ID:      "1",
				JsonRPC: "2.0",
				Method:  methods.METHOD_transaction_V1_getPendingCrossLinks,
				Params:  []interface{}{},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			data, err := json.Marshal(tc.br)
			if err != nil {
				testMetrics = append(testMetrics, TestMetric{
					Method: tc.br.Method,
					Test:   tc.name,
					Pass:   false,
					Error:  err.Error(),
					Params: tc.br.Params,
				})
				t.Error(err)
				return
			}
			// Perform the RPC Call
			resp, err := Call(data, tc.br.Method)
			if err != nil {
				testMetrics = append(testMetrics, TestMetric{
					Method:   tc.br.Method,
					Test:     tc.name,
					Pass:     false,
					Duration: resp.Duration,
					Error:    err.Error(),
					Params:   tc.br.Params,
				})
				t.Error(err)
				return
			}
			if resp.Error != nil {
				if resp.Error.Code != tc.expectedErrorCode {
					testMetrics = append(testMetrics, TestMetric{
						Method:   tc.br.Method,
						Test:     tc.name,
						Pass:     false,
						Duration: resp.Duration,
						Error:    resp.Error.Message,
						Params:   tc.br.Params,
					})
					t.Error(resp.Error.Message)
					return
				}
			}
			// This step validates that the returned response is the correct data type
			if resp.Result != nil {
				var s LatestHeader // Wrong Data type, but func does not work
				err = json.Unmarshal(resp.Result, &s)
				if err != nil {
					testMetrics = append(testMetrics, TestMetric{
						Method:   tc.br.Method,
						Test:     tc.name,
						Pass:     false,
						Duration: resp.Duration,
						Error:    err.Error(),
						Params:   tc.br.Params,
					})
					t.Error(err)
					return
				}
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

func test_V2_getPendingCrossLinks(t *testing.T) {
	type testcase struct {
		name              string
		br                BaseRequest
		expectedErrorCode int64
	}

	testCases := []testcase{
		{
			name: fmt.Sprintf("%s_pending_cross_links", t.Name()),
			br: BaseRequest{
				ID:      "1",
				JsonRPC: "2.0",
				Method:  methods.METHOD_transaction_V2_getPendingCrossLinks,
				Params:  []interface{}{},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			data, err := json.Marshal(tc.br)
			if err != nil {
				testMetrics = append(testMetrics, TestMetric{
					Method: tc.br.Method,
					Test:   tc.name,
					Pass:   false,
					Error:  err.Error(),
					Params: tc.br.Params,
				})
				t.Error(err)
				return
			}
			// Perform the RPC Call
			resp, err := Call(data, tc.br.Method)
			if err != nil {
				testMetrics = append(testMetrics, TestMetric{
					Method:   tc.br.Method,
					Test:     tc.name,
					Pass:     false,
					Duration: resp.Duration,
					Error:    err.Error(),
					Params:   tc.br.Params,
				})
				t.Error(err)
				return
			}
			if resp.Error != nil {
				if resp.Error.Code != tc.expectedErrorCode {
					testMetrics = append(testMetrics, TestMetric{
						Method:   tc.br.Method,
						Test:     tc.name,
						Pass:     false,
						Duration: resp.Duration,
						Error:    resp.Error.Message,
						Params:   tc.br.Params,
					})
					t.Error(resp.Error.Message)
					return
				}
			}
			// This step validates that the returned response is the correct data type
			if resp.Result != nil {
				var s LatestHeader // Wrong Data type, but func does not work
				err = json.Unmarshal(resp.Result, &s)
				if err != nil {
					testMetrics = append(testMetrics, TestMetric{
						Method:   tc.br.Method,
						Test:     tc.name,
						Pass:     false,
						Duration: resp.Duration,
						Error:    err.Error(),
						Params:   tc.br.Params,
					})
					t.Error(err)
					return
				}
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

func test_V1_getPendingCXReceipts(t *testing.T) {
	type testcase struct {
		name              string
		br                BaseRequest
		expectedErrorCode int64
	}

	testCases := []testcase{
		{
			name: fmt.Sprintf("%s_pendingCX_receipts", t.Name()),
			br: BaseRequest{
				ID:      "1",
				JsonRPC: "2.0",
				Method:  methods.METHOD_transaction_V1_getPendingCXReceipts,
				Params:  []interface{}{},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			data, err := json.Marshal(tc.br)
			if err != nil {
				testMetrics = append(testMetrics, TestMetric{
					Method: tc.br.Method,
					Test:   tc.name,
					Pass:   false,
					Error:  err.Error(),
					Params: tc.br.Params,
				})
				t.Error(err)
				return
			}
			// Perform the RPC Call
			resp, err := Call(data, tc.br.Method)
			if err != nil {
				testMetrics = append(testMetrics, TestMetric{
					Method:   tc.br.Method,
					Test:     tc.name,
					Pass:     false,
					Duration: resp.Duration,
					Error:    err.Error(),
					Params:   tc.br.Params,
				})
				t.Error(err)
				return
			}
			if resp.Error != nil {
				if resp.Error.Code != tc.expectedErrorCode {
					testMetrics = append(testMetrics, TestMetric{
						Method:   tc.br.Method,
						Test:     tc.name,
						Pass:     false,
						Duration: resp.Duration,
						Error:    resp.Error.Message,
						Params:   tc.br.Params,
					})
					t.Error(resp.Error.Message)
					return
				}
			}
			// This step validates that the returned response is the correct data type
			if resp.Result != nil {
				var s []Reciept // Wrong Data type, but func does not work
				err = json.Unmarshal(resp.Result, &s)
				if err != nil {
					testMetrics = append(testMetrics, TestMetric{
						Method:   tc.br.Method,
						Test:     tc.name,
						Pass:     false,
						Duration: resp.Duration,
						Error:    err.Error(),
						Params:   tc.br.Params,
					})
					t.Error(err)
					return
				}
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
func test_V2_getPendingCXReceipts(t *testing.T) {
	type testcase struct {
		name              string
		br                BaseRequest
		expectedErrorCode int64
	}

	testCases := []testcase{
		{
			name: fmt.Sprintf("%s_pendingCX_receipts", t.Name()),
			br: BaseRequest{
				ID:      "1",
				JsonRPC: "2.0",
				Method:  methods.METHOD_transaction_V2_getPendingCXReceipts,
				Params:  []interface{}{},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			data, err := json.Marshal(tc.br)
			if err != nil {
				testMetrics = append(testMetrics, TestMetric{
					Method: tc.br.Method,
					Test:   tc.name,
					Pass:   false,
					Error:  err.Error(),
					Params: tc.br.Params,
				})
				t.Error(err)
				return
			}
			// Perform the RPC Call
			resp, err := Call(data, tc.br.Method)
			if err != nil {
				testMetrics = append(testMetrics, TestMetric{
					Method:   tc.br.Method,
					Test:     tc.name,
					Pass:     false,
					Duration: resp.Duration,
					Error:    err.Error(),
					Params:   tc.br.Params,
				})
				t.Error(err)
				return
			}
			if resp.Error != nil {
				if resp.Error.Code != tc.expectedErrorCode {
					testMetrics = append(testMetrics, TestMetric{
						Method:   tc.br.Method,
						Test:     tc.name,
						Pass:     false,
						Duration: resp.Duration,
						Error:    resp.Error.Message,
						Params:   tc.br.Params,
					})
					t.Error(resp.Error.Message)
					return
				}
			}
			// This step validates that the returned response is the correct data type
			if resp.Result != nil {
				var s []Reciept // Wrong Data type, but func does not work
				err = json.Unmarshal(resp.Result, &s)
				if err != nil {
					testMetrics = append(testMetrics, TestMetric{
						Method:   tc.br.Method,
						Test:     tc.name,
						Pass:     false,
						Duration: resp.Duration,
						Error:    err.Error(),
						Params:   tc.br.Params,
					})
					t.Error(err)
					return
				}
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

func test_V1_getCXReceiptByHash(t *testing.T) {
	type testcase struct {
		name              string
		br                BaseRequest
		expectedErrorCode int64
	}

	testCases := []testcase{
		{
			name: fmt.Sprintf("%s_CX_receipts_Hash", t.Name()),
			br: BaseRequest{
				ID:      "1",
				JsonRPC: "2.0",
				Method:  methods.METHOD_transaction_V1_getCXReceiptByHash,
				Params: []interface{}{
					"0x6b106dc5619c86b6c0cb64b17e5c464e8008e08cf0f1bb0e3fa2657fb42daade",
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			data, err := json.Marshal(tc.br)
			if err != nil {
				testMetrics = append(testMetrics, TestMetric{
					Method: tc.br.Method,
					Test:   tc.name,
					Pass:   false,
					Error:  err.Error(),
					Params: tc.br.Params,
				})
				t.Error(err)
				return
			}
			// Perform the RPC Call
			resp, err := Call(data, tc.br.Method)
			if err != nil {
				testMetrics = append(testMetrics, TestMetric{
					Method:   tc.br.Method,
					Test:     tc.name,
					Pass:     false,
					Duration: resp.Duration,
					Error:    err.Error(),
					Params:   tc.br.Params,
				})
				t.Error(err)
				return
			}
			if resp.Error != nil {
				if resp.Error.Code != tc.expectedErrorCode {
					testMetrics = append(testMetrics, TestMetric{
						Method:   tc.br.Method,
						Test:     tc.name,
						Pass:     false,
						Duration: resp.Duration,
						Error:    resp.Error.Message,
						Params:   tc.br.Params,
					})
					t.Error(resp.Error.Message)
					return
				}
			}
			// This step validates that the returned response is the correct data type
			if resp.Result != nil {
				var s []Reciept
				err = json.Unmarshal(resp.Result, &s)
				if err != nil {
					testMetrics = append(testMetrics, TestMetric{
						Method:   tc.br.Method,
						Test:     tc.name,
						Pass:     false,
						Duration: resp.Duration,
						Error:    err.Error(),
						Params:   tc.br.Params,
					})
					t.Error(err)
					return
				}
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
func test_V2_getCXReceiptByHash(t *testing.T) {
	type testcase struct {
		name              string
		br                BaseRequest
		expectedErrorCode int64
	}

	testCases := []testcase{
		{
			name: fmt.Sprintf("%s_CX_receipts_by_Hash", t.Name()),
			br: BaseRequest{
				ID:      "1",
				JsonRPC: "2.0",
				Method:  methods.METHOD_transaction_V2_getCXReceiptByHash,
				Params: []interface{}{
					"0x6b106dc5619c86b6c0cb64b17e5c464e8008e08cf0f1bb0e3fa2657fb42daade",
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			data, err := json.Marshal(tc.br)
			if err != nil {
				testMetrics = append(testMetrics, TestMetric{
					Method: tc.br.Method,
					Test:   tc.name,
					Pass:   false,
					Error:  err.Error(),
					Params: tc.br.Params,
				})
				t.Error(err)
				return
			}
			// Perform the RPC Call
			resp, err := Call(data, tc.br.Method)
			if err != nil {
				testMetrics = append(testMetrics, TestMetric{
					Method:   tc.br.Method,
					Test:     tc.name,
					Pass:     false,
					Duration: resp.Duration,
					Error:    err.Error(),
					Params:   tc.br.Params,
				})
				t.Error(err)
				return
			}
			if resp.Error != nil {
				if resp.Error.Code != tc.expectedErrorCode {
					testMetrics = append(testMetrics, TestMetric{
						Method:   tc.br.Method,
						Test:     tc.name,
						Pass:     false,
						Duration: resp.Duration,
						Error:    resp.Error.Message,
						Params:   tc.br.Params,
					})
					t.Error(resp.Error.Message)
					return
				}
			}
			// This step validates that the returned response is the correct data type
			if resp.Result != nil {
				var s []Reciept // Wrong Data type, but func does not work
				err = json.Unmarshal(resp.Result, &s)
				if err != nil {
					testMetrics = append(testMetrics, TestMetric{
						Method:   tc.br.Method,
						Test:     tc.name,
						Pass:     false,
						Duration: resp.Duration,
						Error:    err.Error(),
						Params:   tc.br.Params,
					})
					t.Error(err)
					return
				}
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

func test_V1_pendingTransactions(t *testing.T) {
	type testcase struct {
		name              string
		br                BaseRequest
		expectedErrorCode int64
	}

	testCases := []testcase{
		{
			name: fmt.Sprintf("%s_pending_Transactions", t.Name()),
			br: BaseRequest{
				ID:      "1",
				JsonRPC: "2.0",
				Method:  methods.METHOD_transaction_V1_pendingTransactions,
				Params:  []interface{}{},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			data, err := json.Marshal(tc.br)
			if err != nil {
				testMetrics = append(testMetrics, TestMetric{
					Method: tc.br.Method,
					Test:   tc.name,
					Pass:   false,
					Error:  err.Error(),
					Params: tc.br.Params,
				})
				t.Error(err)
				return
			}
			// Perform the RPC Call
			resp, err := Call(data, tc.br.Method)
			if err != nil {
				testMetrics = append(testMetrics, TestMetric{
					Method:   tc.br.Method,
					Test:     tc.name,
					Pass:     false,
					Duration: resp.Duration,
					Error:    err.Error(),
					Params:   tc.br.Params,
				})
				t.Error(err)
				return
			}
			if resp.Error != nil {
				if resp.Error.Code != tc.expectedErrorCode {
					testMetrics = append(testMetrics, TestMetric{
						Method:   tc.br.Method,
						Test:     tc.name,
						Pass:     false,
						Duration: resp.Duration,
						Error:    resp.Error.Message,
						Params:   tc.br.Params,
					})
					t.Error(resp.Error.Message)
					return
				}
			}
			// This step validates that the returned response is the correct data type
			if resp.Result != nil {
				var s []TransactionByHashV1
				err = json.Unmarshal(resp.Result, &s)
				if err != nil {
					testMetrics = append(testMetrics, TestMetric{
						Method:   tc.br.Method,
						Test:     tc.name,
						Pass:     false,
						Duration: resp.Duration,
						Error:    err.Error(),
						Params:   tc.br.Params,
					})
					t.Error(err)
					return
				}
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
func test_V2_pendingTransactions(t *testing.T) {
	type testcase struct {
		name              string
		br                BaseRequest
		expectedErrorCode int64
	}

	testCases := []testcase{
		{
			name: fmt.Sprintf("%s_pending_Transactions", t.Name()),
			br: BaseRequest{
				ID:      "1",
				JsonRPC: "2.0",
				Method:  methods.METHOD_transaction_V2_pendingTransactions,
				Params:  []interface{}{},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			data, err := json.Marshal(tc.br)
			if err != nil {
				testMetrics = append(testMetrics, TestMetric{
					Method: tc.br.Method,
					Test:   tc.name,
					Pass:   false,
					Error:  err.Error(),
					Params: tc.br.Params,
				})
				t.Error(err)
				return
			}
			// Perform the RPC Call
			resp, err := Call(data, tc.br.Method)
			if err != nil {
				testMetrics = append(testMetrics, TestMetric{
					Method:   tc.br.Method,
					Test:     tc.name,
					Pass:     false,
					Duration: resp.Duration,
					Error:    err.Error(),
					Params:   tc.br.Params,
				})
				t.Error(err)
				return
			}
			if resp.Error != nil {
				if resp.Error.Code != tc.expectedErrorCode {
					testMetrics = append(testMetrics, TestMetric{
						Method:   tc.br.Method,
						Test:     tc.name,
						Pass:     false,
						Duration: resp.Duration,
						Error:    resp.Error.Message,
						Params:   tc.br.Params,
					})
					t.Error(resp.Error.Message)
					return
				}
			}
			// This step validates that the returned response is the correct data type
			if resp.Result != nil {
				var s []TransactionByHashV2
				err = json.Unmarshal(resp.Result, &s)
				if err != nil {
					testMetrics = append(testMetrics, TestMetric{
						Method:   tc.br.Method,
						Test:     tc.name,
						Pass:     false,
						Duration: resp.Duration,
						Error:    err.Error(),
						Params:   tc.br.Params,
					})
					t.Error(err)
					return
				}
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

func (ts *testSuite) test_sendRawStakingTransaction(t *testing.T) {
	// https://github.com/harmony-one/bounties/issues/117#issuecomment-1170274370
	// Skip unless devnet
	if URL != "https://api.s0.b.hmny.io/" && URL != "https://api.s0.pops.one/" {
		t.Skip("only run staking on testnet")
	}
	if len(ts.ElectedValidators) == 0 {
		t.Skip("cannot stake without validators")
	}
	type testcase struct {
		name              string
		br                BaseRequest
		expectedErrorCode int64
	}

	testCases := []testcase{
		{
			name: fmt.Sprintf("%s_staking_transaction", t.Name()),
			br: BaseRequest{
				ID:      "1",
				JsonRPC: "2.0",
				Method:  methods.METHOD_transaction_sendRawStakingTransaction,
				Params:  []interface{}{},
			},
		},
	}

	if len(ts.ValidatorsV2.Validators) == 0 {
		t.Skip("Cant perfrom without validators")
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Build a Staking Transaction

			payload, err := CreateStakingRLPString(address, ts.ElectedValidators[0], big.NewInt(0).Mul(ONE, big.NewInt(101)), nil)
			if err != nil {
				testMetrics = append(testMetrics, TestMetric{
					Method: tc.br.Method,
					Test:   tc.name,
					Pass:   false,
					Error:  err.Error(),
					Params: tc.br.Params,
				})
				t.Error(err)
				return
			}
			tc.br.Params = append(tc.br.Params, payload)

			var result string

			resp, err := callAndValidateDataType(t, tc.name, tc.expectedErrorCode, tc.br, &result)
			if err != nil {
				testMetrics = append(testMetrics, TestMetric{
					Method: tc.br.Method,
					Test:   tc.name,
					Pass:   false,
					Error:  err.Error(),
					Params: tc.br.Params,
				})
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

func (ts *testSuite) test_V1_getTransactionHistory(t *testing.T) {
	type testcase struct {
		name              string
		br                BaseRequest
		expectedErrorCode int64
	}

	testCases := []testcase{
		{
			name: fmt.Sprintf("%s_transaction_history", t.Name()),
			br: BaseRequest{
				ID:      "1",
				JsonRPC: "2.0",
				Method:  methods.METHOD_transaction_V1_getTransactionHistory,
				Params: []interface{}{
					TransactionArguments{
						Address:   address,
						TxType:    "ALL",
						FullTx:    true,
						PageSize:  100,
						PageIndex: 0,
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			data, err := json.Marshal(tc.br)
			if err != nil {
				testMetrics = append(testMetrics, TestMetric{
					Method: tc.br.Method,
					Test:   tc.name,
					Pass:   false,
					Error:  err.Error(),
					Params: tc.br.Params,
				})
				t.Error(err)
				return
			}
			// Perform the RPC Call
			resp, err := Call(data, tc.br.Method)
			if err != nil {
				testMetrics = append(testMetrics, TestMetric{
					Method:   tc.br.Method,
					Test:     tc.name,
					Pass:     false,
					Duration: resp.Duration,
					Error:    err.Error(),
					Params:   tc.br.Params,
				})
				t.Error(err)
				return
			}
			if resp.Error != nil {
				if resp.Error.Code != tc.expectedErrorCode {
					testMetrics = append(testMetrics, TestMetric{
						Method:   tc.br.Method,
						Test:     tc.name,
						Pass:     false,
						Duration: resp.Duration,
						Error:    resp.Error.Message,
						Params:   tc.br.Params,
					})
					t.Error(resp.Error.Message)
					return
				}
			}
			// This step validates that the returned response is the correct data type
			if resp.Result != nil {
				var s TransactionHistoryV1
				err = json.Unmarshal(resp.Result, &s)
				if err != nil {
					testMetrics = append(testMetrics, TestMetric{
						Method:   tc.br.Method,
						Test:     tc.name,
						Pass:     false,
						Duration: resp.Duration,
						Error:    err.Error(),
						Params:   tc.br.Params,
					})
					t.Error(err)
					return
				}
				// Verify that this Transaction is present in the TS
				var found bool
				for _, txHist := range s.Transactions {
					if txHist.Hash == ts.LastTransactionHash {
						found = true
					}
				}
				if !found {
					testMetrics = append(testMetrics, TestMetric{
						Method:   tc.br.Method,
						Test:     tc.name,
						Pass:     false,
						Duration: resp.Duration,
						Error:    "test transaction was not found in history",
						Params:   tc.br.Params,
					})
					t.Error("test transaction was not found in history")
					return
				}
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
func (ts *testSuite) test_V2_getTransactionHistory(t *testing.T) {
	type testcase struct {
		name              string
		br                BaseRequest
		expectedErrorCode int64
	}

	testCases := []testcase{
		{
			name: fmt.Sprintf("%s_transaction_history", t.Name()),
			br: BaseRequest{
				ID:      "1",
				JsonRPC: "2.0",
				Method:  methods.METHOD_transaction_V2_getTransactionHistory,
				Params: []interface{}{
					TransactionArguments{
						Address:   address,
						TxType:    "ALL",
						FullTx:    true,
						PageSize:  100,
						PageIndex: 0,
					},
				},
			},
		},
		{
			name:              fmt.Sprintf("%s_bad_txArgs", t.Name()),
			expectedErrorCode: -32602,
			br: BaseRequest{
				ID:      "1",
				JsonRPC: "2.0",
				Method:  methods.METHOD_transaction_V2_getTransactionHistory,
				Params: []interface{}{
					address,
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			data, err := json.Marshal(tc.br)
			if err != nil {
				testMetrics = append(testMetrics, TestMetric{
					Method: tc.br.Method,
					Test:   tc.name,
					Pass:   false,
					Error:  err.Error(),
					Params: tc.br.Params,
				})
				t.Error(err)
				return
			}
			// Perform the RPC Call
			resp, err := Call(data, tc.br.Method)
			if err != nil {
				testMetrics = append(testMetrics, TestMetric{
					Method:   tc.br.Method,
					Test:     tc.name,
					Pass:     false,
					Duration: resp.Duration,
					Error:    err.Error(),
					Params:   tc.br.Params,
				})
				t.Error(err)
				return
			}
			if resp.Error != nil {
				if resp.Error.Code != tc.expectedErrorCode {
					testMetrics = append(testMetrics, TestMetric{
						Method:   tc.br.Method,
						Test:     tc.name,
						Pass:     false,
						Duration: resp.Duration,
						Error:    resp.Error.Message,
						Params:   tc.br.Params,
					})
					t.Error(resp.Error.Message)
					return
				}
			}
			// This step validates that the returned response is the correct data type
			if resp.Result != nil {
				var s TransactionHistoryV2
				err = json.Unmarshal(resp.Result, &s)
				if err != nil {
					testMetrics = append(testMetrics, TestMetric{
						Method:   tc.br.Method,
						Test:     tc.name,
						Pass:     false,
						Duration: resp.Duration,
						Error:    err.Error(),
						Params:   tc.br.Params,
					})
					t.Error(err)
					return
				}
				// Verify that this Transaction is present in the TS
				var found bool
				for _, txHist := range s.Transactions {
					if txHist.Hash == ts.LastTransactionHash {
						found = true
					}
				}
				if !found {
					testMetrics = append(testMetrics, TestMetric{
						Method:   tc.br.Method,
						Test:     tc.name,
						Pass:     false,
						Duration: resp.Duration,
						Error:    "test transaction was not found in history",
						Params:   tc.br.Params,
					})
					t.Error("test transaction was not found in history")
					return
				}
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

func (ts *testSuite) test_V1_getTransactionReceipt(t *testing.T) {
	type testcase struct {
		name              string
		br                BaseRequest
		expectedErrorCode int64
	}

	testCases := []testcase{
		{
			name: fmt.Sprintf("%s_transaction_receipt", t.Name()),
			br: BaseRequest{
				ID:      "1",
				JsonRPC: "2.0",
				Method:  methods.METHOD_transaction_V1_getTransactionReceipt,
				Params: []interface{}{
					ts.LastTransactionHash,
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			data, err := json.Marshal(tc.br)
			if err != nil {
				testMetrics = append(testMetrics, TestMetric{
					Method: tc.br.Method,
					Test:   tc.name,
					Pass:   false,
					Error:  err.Error(),
					Params: tc.br.Params,
				})
				t.Error(err)
				return
			}
			// Perform the RPC Call
			resp, err := Call(data, tc.br.Method)
			if err != nil {
				testMetrics = append(testMetrics, TestMetric{
					Method:   tc.br.Method,
					Test:     tc.name,
					Pass:     false,
					Duration: resp.Duration,
					Error:    err.Error(),
					Params:   tc.br.Params,
				})
				t.Error(err)
				return
			}
			if resp.Error != nil {
				if resp.Error.Code != tc.expectedErrorCode {
					testMetrics = append(testMetrics, TestMetric{
						Method:   tc.br.Method,
						Test:     tc.name,
						Pass:     false,
						Duration: resp.Duration,
						Error:    resp.Error.Message,
						Params:   tc.br.Params,
					})
					t.Error(resp.Error.Message)
					return
				}
			}
			// This step validates that the returned response is the correct data type
			if resp.Result != nil {
				var s TransactionReceipt_V1

				err = json.Unmarshal(resp.Result, &s)
				if err != nil {
					testMetrics = append(testMetrics, TestMetric{
						Method:   tc.br.Method,
						Test:     tc.name,
						Pass:     false,
						Duration: resp.Duration,
						Error:    err.Error(),
						Params:   tc.br.Params,
					})
					t.Error(err)
					return
				}
				ts.LastTransactionReceiptV1 = s
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

func (ts *testSuite) test_V2_getTransactionReceipt(t *testing.T) {
	type testcase struct {
		name              string
		br                BaseRequest
		expectedErrorCode int64
	}

	testCases := []testcase{
		{
			name: fmt.Sprintf("%s_transaction_receipt", t.Name()),
			br: BaseRequest{
				ID:      "1",
				JsonRPC: "2.0",
				Method:  methods.METHOD_transaction_V2_getTransactionReceipt,
				Params: []interface{}{
					ts.LastTransactionHash,
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			data, err := json.Marshal(tc.br)
			if err != nil {
				testMetrics = append(testMetrics, TestMetric{
					Method: tc.br.Method,
					Test:   tc.name,
					Pass:   false,
					Error:  err.Error(),
					Params: tc.br.Params,
				})
				t.Error(err)
				return
			}
			// Perform the RPC Call
			resp, err := Call(data, tc.br.Method)
			if err != nil {
				testMetrics = append(testMetrics, TestMetric{
					Method:   tc.br.Method,
					Test:     tc.name,
					Pass:     false,
					Duration: resp.Duration,
					Error:    err.Error(),
					Params:   tc.br.Params,
				})
				t.Error(err)
				return
			}
			if resp.Error != nil {
				if resp.Error.Code != tc.expectedErrorCode {
					testMetrics = append(testMetrics, TestMetric{
						Method:   tc.br.Method,
						Test:     tc.name,
						Pass:     false,
						Duration: resp.Duration,
						Error:    resp.Error.Message,
						Params:   tc.br.Params,
					})
					t.Error(resp.Error.Message)
					return
				}
			}
			// This step validates that the returned response is the correct data type
			if resp.Result != nil {
				var s TransactionReceipt_V2
				err = json.Unmarshal(resp.Result, &s)
				if err != nil {
					testMetrics = append(testMetrics, TestMetric{
						Method:   tc.br.Method,
						Test:     tc.name,
						Pass:     false,
						Duration: resp.Duration,
						Error:    err.Error(),
						Params:   tc.br.Params,
					})
					t.Error(err)
					return
				}
				ts.LastTransactionBlockHash = s.BlockHash
				ts.LastTransactionBlockNumber = s.BlockNumber
				ts.LastTransactionReceiptV2 = s

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

func (ts *testSuite) test_V1_getBlockTransactionCountByHash(t *testing.T) {
	type testcase struct {
		name              string
		br                BaseRequest
		expectedErrorCode int64
	}

	testCases := []testcase{
		{
			name: fmt.Sprintf("%s_transaction_hash_count", t.Name()),
			br: BaseRequest{
				ID:      "1",
				JsonRPC: "2.0",
				Method:  methods.METHOD_transaction_V1_getBlockTransactionCountByHash,
				Params: []interface{}{
					ts.LastTransactionHash,
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			data, err := json.Marshal(tc.br)
			if err != nil {
				testMetrics = append(testMetrics, TestMetric{
					Method: tc.br.Method,
					Test:   tc.name,
					Pass:   false,
					Error:  err.Error(),
					Params: tc.br.Params,
				})
				t.Error(err)
				return
			}
			// Perform the RPC Call
			resp, err := Call(data, tc.br.Method)
			if err != nil {
				testMetrics = append(testMetrics, TestMetric{
					Method:   tc.br.Method,
					Test:     tc.name,
					Pass:     false,
					Duration: resp.Duration,
					Error:    err.Error(),
					Params:   tc.br.Params,
				})
				t.Error(err)
				return
			}
			if resp.Error != nil {
				if resp.Error.Code != tc.expectedErrorCode {
					testMetrics = append(testMetrics, TestMetric{
						Method:   tc.br.Method,
						Test:     tc.name,
						Pass:     false,
						Duration: resp.Duration,
						Error:    resp.Error.Message,
						Params:   tc.br.Params,
					})
					t.Error(resp.Error.Message)
					return
				}
			}
			// This step validates that the returned response is the correct data type
			if resp.Result != nil {
				var s string
				err = json.Unmarshal(resp.Result, &s)
				if err != nil {
					testMetrics = append(testMetrics, TestMetric{
						Method:   tc.br.Method,
						Test:     tc.name,
						Pass:     false,
						Duration: resp.Duration,
						Error:    err.Error(),
						Params:   tc.br.Params,
					})
					t.Error(err)
					return
				}
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

func (ts *testSuite) test_V2_getBlockTransactionCountByHash(t *testing.T) {
	type testcase struct {
		name              string
		br                BaseRequest
		expectedErrorCode int64
	}

	testCases := []testcase{
		{
			name: fmt.Sprintf("%s_transaction_hash_count", t.Name()),
			br: BaseRequest{
				ID:      "1",
				JsonRPC: "2.0",
				Method:  methods.METHOD_transaction_V2_getBlockTransactionCountByHash,
				Params: []interface{}{
					ts.LastTransactionHash,
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			data, err := json.Marshal(tc.br)
			if err != nil {
				testMetrics = append(testMetrics, TestMetric{
					Method: tc.br.Method,
					Test:   tc.name,
					Pass:   false,
					Error:  err.Error(),
					Params: tc.br.Params,
				})
				t.Error(err)
				return
			}
			// Perform the RPC Call
			resp, err := Call(data, tc.br.Method)
			if err != nil {
				testMetrics = append(testMetrics, TestMetric{
					Method:   tc.br.Method,
					Test:     tc.name,
					Pass:     false,
					Duration: resp.Duration,
					Error:    err.Error(),
					Params:   tc.br.Params,
				})
				t.Error(err)
				return
			}
			if resp.Error != nil {
				if resp.Error.Code != tc.expectedErrorCode {
					testMetrics = append(testMetrics, TestMetric{
						Method:   tc.br.Method,
						Test:     tc.name,
						Pass:     false,
						Duration: resp.Duration,
						Error:    resp.Error.Message,
						Params:   tc.br.Params,
					})
					t.Error(resp.Error.Message)
					return
				}
			}
			// This step validates that the returned response is the correct data type
			if resp.Result != nil {
				var s int64
				err = json.Unmarshal(resp.Result, &s)
				if err != nil {
					testMetrics = append(testMetrics, TestMetric{
						Method:   tc.br.Method,
						Test:     tc.name,
						Pass:     false,
						Duration: resp.Duration,
						Error:    err.Error(),
						Params:   tc.br.Params,
					})
					t.Error(err)
					return
				}
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

func (ts *testSuite) test_V1_getBlockTransactionCountByNumber(t *testing.T) {
	type testcase struct {
		name              string
		br                BaseRequest
		expectedErrorCode int64
	}

	testCases := []testcase{
		{
			name: fmt.Sprintf("%s_transaction_hash_count", t.Name()),
			br: BaseRequest{
				ID:      "1",
				JsonRPC: "2.0",
				Method:  methods.METHOD_transaction_V1_getBlockTransactionCountByNumber,
				Params: []interface{}{
					ts.LastTransactionBlockNumber,
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			data, err := json.Marshal(tc.br)
			if err != nil {
				testMetrics = append(testMetrics, TestMetric{
					Method: tc.br.Method,
					Test:   tc.name,
					Pass:   false,
					Error:  err.Error(),
					Params: tc.br.Params,
				})
				t.Error(err)
				return
			}
			// Perform the RPC Call
			resp, err := Call(data, tc.br.Method)
			if err != nil {
				testMetrics = append(testMetrics, TestMetric{
					Method:   tc.br.Method,
					Test:     tc.name,
					Pass:     false,
					Duration: resp.Duration,
					Error:    err.Error(),
					Params:   tc.br.Params,
				})
				t.Error(err)
				return
			}
			if resp.Error != nil {
				if resp.Error.Code != tc.expectedErrorCode {
					testMetrics = append(testMetrics, TestMetric{
						Method:   tc.br.Method,
						Test:     tc.name,
						Pass:     false,
						Duration: resp.Duration,
						Error:    resp.Error.Message,
						Params:   tc.br.Params,
					})
					t.Error(resp.Error.Message)
					return
				}
			}
			// This step validates that the returned response is the correct data type
			if resp.Result != nil {
				var s string
				err = json.Unmarshal(resp.Result, &s)
				if err != nil {
					testMetrics = append(testMetrics, TestMetric{
						Method:   tc.br.Method,
						Test:     tc.name,
						Pass:     false,
						Duration: resp.Duration,
						Error:    err.Error(),
						Params:   tc.br.Params,
					})
					t.Error(err)
					return
				}
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

func (ts *testSuite) test_V2_getBlockTransactionCountByNumber(t *testing.T) {
	type testcase struct {
		name              string
		br                BaseRequest
		expectedErrorCode int64
	}

	testCases := []testcase{
		{
			name: fmt.Sprintf("%s_transaction_hash_count", t.Name()),
			br: BaseRequest{
				ID:      "1",
				JsonRPC: "2.0",
				Method:  methods.METHOD_transaction_V2_getBlockTransactionCountByNumber,
				Params: []interface{}{
					ts.LastTransactionBlockNumber,
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			data, err := json.Marshal(tc.br)
			if err != nil {
				testMetrics = append(testMetrics, TestMetric{
					Method: tc.br.Method,
					Test:   tc.name,
					Pass:   false,
					Error:  err.Error(),
					Params: tc.br.Params,
				})
				t.Error(err)
				return
			}
			// Perform the RPC Call
			resp, err := Call(data, tc.br.Method)
			if err != nil {
				testMetrics = append(testMetrics, TestMetric{
					Method:   tc.br.Method,
					Test:     tc.name,
					Pass:     false,
					Duration: resp.Duration,
					Error:    err.Error(),
					Params:   tc.br.Params,
				})
				t.Error(err)
				return
			}
			if resp.Error != nil {
				if resp.Error.Code != tc.expectedErrorCode {
					testMetrics = append(testMetrics, TestMetric{
						Method:   tc.br.Method,
						Test:     tc.name,
						Pass:     false,
						Duration: resp.Duration,
						Error:    resp.Error.Message,
						Params:   tc.br.Params,
					})
					t.Error(resp.Error.Message)
					return
				}
			}
			// This step validates that the returned response is the correct data type
			if resp.Result != nil {
				var s int64
				err = json.Unmarshal(resp.Result, &s)
				if err != nil {
					testMetrics = append(testMetrics, TestMetric{
						Method:   tc.br.Method,
						Test:     tc.name,
						Pass:     false,
						Duration: resp.Duration,
						Error:    err.Error(),
						Params:   tc.br.Params,
					})
					t.Error(err)
					return
				}
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

func (ts *testSuite) test_V1_getTransactionByHash(t *testing.T) {
	type testcase struct {
		name              string
		br                BaseRequest
		expectedErrorCode int64
	}

	testCases := []testcase{
		{
			name: fmt.Sprintf("%s_transaction_should_exist", t.Name()),
			br: BaseRequest{
				ID:      "1",
				JsonRPC: "2.0",
				Method:  methods.METHOD_transaction_V1_getTransactionByHash,
				Params: []interface{}{
					ts.LastTransactionHash,
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			data, err := json.Marshal(tc.br)
			if err != nil {
				testMetrics = append(testMetrics, TestMetric{
					Method: tc.br.Method,
					Test:   tc.name,
					Pass:   false,
					Error:  err.Error(),
					Params: tc.br.Params,
				})
				t.Error(err)
				return
			}
			// Perform the RPC Call
			resp, err := Call(data, tc.br.Method)
			if err != nil {
				testMetrics = append(testMetrics, TestMetric{
					Method:   tc.br.Method,
					Test:     tc.name,
					Pass:     false,
					Duration: resp.Duration,
					Error:    err.Error(),
					Params:   tc.br.Params,
				})
				t.Error(err)
				return
			}
			if resp.Error != nil {
				if resp.Error.Code != tc.expectedErrorCode {
					testMetrics = append(testMetrics, TestMetric{
						Method:   tc.br.Method,
						Test:     tc.name,
						Pass:     false,
						Duration: resp.Duration,
						Error:    resp.Error.Message,
						Params:   tc.br.Params,
					})
					t.Error(resp.Error.Message)
					return
				}
			}
			// This step validates that the returned response is the correct data type
			if resp.Result != nil {
				var s TransactionByHashV1
				err = json.Unmarshal(resp.Result, &s)
				if err != nil {
					testMetrics = append(testMetrics, TestMetric{
						Method:   tc.br.Method,
						Test:     tc.name,
						Pass:     false,
						Duration: resp.Duration,
						Error:    err.Error(),
						Params:   tc.br.Params,
					})
					t.Error(err)
					return
				}
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

func (ts *testSuite) test_V2_getTransactionByHash(t *testing.T) {
	type testcase struct {
		name              string
		br                BaseRequest
		expectedErrorCode int64
	}

	testCases := []testcase{
		{
			name: fmt.Sprintf("%s_transaction_should_exist", t.Name()),
			br: BaseRequest{
				ID:      "1",
				JsonRPC: "2.0",
				Method:  methods.METHOD_transaction_V2_getTransactionByHash,
				Params: []interface{}{
					ts.LastTransactionHash,
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			data, err := json.Marshal(tc.br)
			if err != nil {
				testMetrics = append(testMetrics, TestMetric{
					Method: tc.br.Method,
					Test:   tc.name,
					Pass:   false,
					Error:  err.Error(),
					Params: tc.br.Params,
				})
				t.Error(err)
				return
			}
			// Perform the RPC Call
			resp, err := Call(data, tc.br.Method)
			if err != nil {
				testMetrics = append(testMetrics, TestMetric{
					Method:   tc.br.Method,
					Test:     tc.name,
					Pass:     false,
					Duration: resp.Duration,
					Error:    err.Error(),
					Params:   tc.br.Params,
				})
				t.Error(err)
				return
			}
			if resp.Error != nil {
				if resp.Error.Code != tc.expectedErrorCode {
					testMetrics = append(testMetrics, TestMetric{
						Method:   tc.br.Method,
						Test:     tc.name,
						Pass:     false,
						Duration: resp.Duration,
						Error:    resp.Error.Message,
						Params:   tc.br.Params,
					})
					t.Error(resp.Error.Message)
					return
				}
			}
			// This step validates that the returned response is the correct data type
			if resp.Result != nil {
				var s TransactionByHashV2
				err = json.Unmarshal(resp.Result, &s)
				if err != nil {
					testMetrics = append(testMetrics, TestMetric{
						Method:   tc.br.Method,
						Test:     tc.name,
						Pass:     false,
						Duration: resp.Duration,
						Error:    err.Error(),
						Params:   tc.br.Params,
					})
					t.Error(err)
					return
				}
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

func (ts *testSuite) test_V1_getTransactionByBlockNumberAndIndex(t *testing.T) {
	type testcase struct {
		name              string
		br                BaseRequest
		expectedErrorCode int64
	}

	testCases := []testcase{
		{
			name: fmt.Sprintf("%s_transaction_should_exist", t.Name()),
			br: BaseRequest{
				ID:      "1",
				JsonRPC: "2.0",
				Method:  methods.METHOD_transaction_V1_getTransactionByBlockNumberAndIndex,
				Params: []interface{}{
					ts.LastTransactionReceiptV1.BlockNumber,
					ts.LastTransactionReceiptV1.TransactionIndex,
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			data, err := json.Marshal(tc.br)
			if err != nil {
				testMetrics = append(testMetrics, TestMetric{
					Method: tc.br.Method,
					Test:   tc.name,
					Pass:   false,
					Error:  err.Error(),
					Params: tc.br.Params,
				})
				t.Error(err)
				return
			}
			// Perform the RPC Call
			resp, err := Call(data, tc.br.Method)
			if err != nil {
				testMetrics = append(testMetrics, TestMetric{
					Method:   tc.br.Method,
					Test:     tc.name,
					Pass:     false,
					Duration: resp.Duration,
					Error:    err.Error(),
					Params:   tc.br.Params,
				})
				t.Error(err)
				return
			}
			if resp.Error != nil {
				if resp.Error.Code != tc.expectedErrorCode {
					testMetrics = append(testMetrics, TestMetric{
						Method:   tc.br.Method,
						Test:     tc.name,
						Pass:     false,
						Duration: resp.Duration,
						Error:    resp.Error.Message,
						Params:   tc.br.Params,
					})
					t.Error(resp.Error.Message)
					return
				}
			}
			// This step validates that the returned response is the correct data type
			if resp.Result != nil {
				var s TransactionByHashV1
				err = json.Unmarshal(resp.Result, &s)
				if err != nil {
					testMetrics = append(testMetrics, TestMetric{
						Method:   tc.br.Method,
						Test:     tc.name,
						Pass:     false,
						Duration: resp.Duration,
						Error:    err.Error(),
						Params:   tc.br.Params,
					})
					t.Error(err)
					return
				}
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
func (ts *testSuite) test_V2_getTransactionByBlockNumberAndIndex(t *testing.T) {
	type testcase struct {
		name              string
		br                BaseRequest
		expectedErrorCode int64
	}

	testCases := []testcase{
		{
			name: fmt.Sprintf("%s_transaction_should_exist", t.Name()),
			br: BaseRequest{
				ID:      "1",
				JsonRPC: "2.0",
				Method:  methods.METHOD_transaction_V2_getTransactionByBlockNumberAndIndex,
				Params: []interface{}{
					ts.LastTransactionReceiptV2.BlockNumber,
					ts.LastTransactionReceiptV2.TransactionIndex,
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			data, err := json.Marshal(tc.br)
			if err != nil {
				testMetrics = append(testMetrics, TestMetric{
					Method: tc.br.Method,
					Test:   tc.name,
					Pass:   false,
					Error:  err.Error(),
					Params: tc.br.Params,
				})
				t.Error(err)
				return
			}
			// Perform the RPC Call
			resp, err := Call(data, tc.br.Method)
			if err != nil {
				testMetrics = append(testMetrics, TestMetric{
					Method:   tc.br.Method,
					Test:     tc.name,
					Pass:     false,
					Duration: resp.Duration,
					Error:    err.Error(),
					Params:   tc.br.Params,
				})
				t.Error(err)
				return
			}
			if resp.Error != nil {
				if resp.Error.Code != tc.expectedErrorCode {
					testMetrics = append(testMetrics, TestMetric{
						Method:   tc.br.Method,
						Test:     tc.name,
						Pass:     false,
						Duration: resp.Duration,
						Error:    resp.Error.Message,
						Params:   tc.br.Params,
					})
					t.Error(resp.Error.Message)
					return
				}
			}
			// This step validates that the returned response is the correct data type
			if resp.Result != nil {
				var s TransactionByHashV2
				err = json.Unmarshal(resp.Result, &s)
				if err != nil {
					testMetrics = append(testMetrics, TestMetric{
						Method:   tc.br.Method,
						Test:     tc.name,
						Pass:     false,
						Duration: resp.Duration,
						Error:    err.Error(),
						Params:   tc.br.Params,
					})
					t.Error(err)
					return
				}
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

func (ts *testSuite) test_V1_getTransactionByBlockHashAndIndex(t *testing.T) {
	type testcase struct {
		name              string
		br                BaseRequest
		expectedErrorCode int64
	}

	testCases := []testcase{
		{
			name: fmt.Sprintf("%s_transaction_should_exist", t.Name()),
			br: BaseRequest{
				ID:      "1",
				JsonRPC: "2.0",
				Method:  methods.METHOD_transaction_V1_getTransactionByBlockHashAndIndex,
				Params: []interface{}{
					ts.LastTransactionReceiptV1.BlockHash,
					ts.LastTransactionReceiptV1.TransactionIndex,
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			data, err := json.Marshal(tc.br)
			if err != nil {
				testMetrics = append(testMetrics, TestMetric{
					Method: tc.br.Method,
					Test:   tc.name,
					Pass:   false,
					Error:  err.Error(),
					Params: tc.br.Params,
				})
				t.Error(err)
				return
			}
			// Perform the RPC Call
			resp, err := Call(data, tc.br.Method)
			if err != nil {
				testMetrics = append(testMetrics, TestMetric{
					Method:   tc.br.Method,
					Test:     tc.name,
					Pass:     false,
					Duration: resp.Duration,
					Error:    err.Error(),
					Params:   tc.br.Params,
				})
				t.Error(err)
				return
			}
			if resp.Error != nil {
				if resp.Error.Code != tc.expectedErrorCode {
					testMetrics = append(testMetrics, TestMetric{
						Method:   tc.br.Method,
						Test:     tc.name,
						Pass:     false,
						Duration: resp.Duration,
						Error:    resp.Error.Message,
						Params:   tc.br.Params,
					})
					t.Error(resp.Error.Message)
					return
				}
			}
			// This step validates that the returned response is the correct data type
			if resp.Result != nil {
				var s TransactionByHashV1
				err = json.Unmarshal(resp.Result, &s)
				if err != nil {
					testMetrics = append(testMetrics, TestMetric{
						Method:   tc.br.Method,
						Test:     tc.name,
						Pass:     false,
						Duration: resp.Duration,
						Error:    err.Error(),
						Params:   tc.br.Params,
					})
					t.Error(err)
					return
				}
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

func (ts *testSuite) test_V2_getTransactionByBlockHashAndIndex(t *testing.T) {
	type testcase struct {
		name              string
		br                BaseRequest
		expectedErrorCode int64
	}

	testCases := []testcase{
		{
			name: fmt.Sprintf("%s_transaction_should_exist", t.Name()),
			br: BaseRequest{
				ID:      "1",
				JsonRPC: "2.0",
				Method:  methods.METHOD_transaction_V2_getTransactionByBlockHashAndIndex,
				Params: []interface{}{
					ts.LastTransactionReceiptV2.BlockHash,
					ts.LastTransactionReceiptV2.TransactionIndex,
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			data, err := json.Marshal(tc.br)
			if err != nil {
				testMetrics = append(testMetrics, TestMetric{
					Method: tc.br.Method,
					Test:   tc.name,
					Pass:   false,
					Error:  err.Error(),
					Params: tc.br.Params,
				})
				t.Error(err)
				return
			}
			// Perform the RPC Call
			resp, err := Call(data, tc.br.Method)
			if err != nil {
				testMetrics = append(testMetrics, TestMetric{
					Method:   tc.br.Method,
					Test:     tc.name,
					Pass:     false,
					Duration: resp.Duration,
					Error:    err.Error(),
					Params:   tc.br.Params,
				})
				t.Error(err)
				return
			}
			if resp.Error != nil {
				if resp.Error.Code != tc.expectedErrorCode {
					testMetrics = append(testMetrics, TestMetric{
						Method:   tc.br.Method,
						Test:     tc.name,
						Pass:     false,
						Duration: resp.Duration,
						Error:    resp.Error.Message,
						Params:   tc.br.Params,
					})
					t.Error(resp.Error.Message)
					return
				}
			}
			// This step validates that the returned response is the correct data type
			if resp.Result != nil {
				var s TransactionByHashV2
				err = json.Unmarshal(resp.Result, &s)
				if err != nil {
					testMetrics = append(testMetrics, TestMetric{
						Method:   tc.br.Method,
						Test:     tc.name,
						Pass:     false,
						Duration: resp.Duration,
						Error:    err.Error(),
						Params:   tc.br.Params,
					})
					t.Error(err)
					return
				}
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

func (ts *testSuite) test_V1_getBlockByNumber(t *testing.T) {
	type testcase struct {
		name              string
		br                BaseRequest
		expectedErrorCode int64
	}

	testCases := []testcase{
		{
			name: fmt.Sprintf("%s_existing_block", t.Name()),
			br: BaseRequest{
				ID:      "1",
				JsonRPC: "2.0",
				Method:  methods.METHOD_transaction_V1_getBlockByNumber,
				Params: []interface{}{
					ts.LastTransactionReceiptV1.BlockNumber,
					true,
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			data, err := json.Marshal(tc.br)
			if err != nil {
				testMetrics = append(testMetrics, TestMetric{
					Method: tc.br.Method,
					Test:   tc.name,
					Pass:   false,
					Error:  err.Error(),
					Params: tc.br.Params,
				})
				t.Error(err)
				return
			}
			// Perform the RPC Call
			resp, err := Call(data, tc.br.Method)
			if err != nil {
				testMetrics = append(testMetrics, TestMetric{
					Method:   tc.br.Method,
					Test:     tc.name,
					Pass:     false,
					Duration: resp.Duration,
					Error:    err.Error(),
					Params:   tc.br.Params,
				})
				t.Error(err)
				return
			}
			if resp.Error != nil {
				if resp.Error.Code != tc.expectedErrorCode {
					testMetrics = append(testMetrics, TestMetric{
						Method:   tc.br.Method,
						Test:     tc.name,
						Pass:     false,
						Duration: resp.Duration,
						Error:    resp.Error.Message,
						Params:   tc.br.Params,
					})
					t.Error(resp.Error.Message)
					return
				}
			}
			// This step validates that the returned response is the correct data type
			if resp.Result != nil {
				var s BlockV1
				err = json.Unmarshal(resp.Result, &s)
				if err != nil {
					testMetrics = append(testMetrics, TestMetric{
						Method:   tc.br.Method,
						Test:     tc.name,
						Pass:     false,
						Duration: resp.Duration,
						Error:    err.Error(),
						Params:   tc.br.Params,
					})
					t.Error(err)
					return
				}
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

func (ts *testSuite) test_V2_getBlockByNumber(t *testing.T) {
	type testcase struct {
		name              string
		br                BaseRequest
		expectedErrorCode int64
	}

	type include struct {
		FullTx         bool `json:"fullTx"`
		WithSigners    bool `json:"withSigner"`
		IncludeSigners bool `json:"includeSigners"`
	}

	testCases := []testcase{
		{
			name: fmt.Sprintf("%s_existing_block", t.Name()),
			br: BaseRequest{
				ID:      "1",
				JsonRPC: "2.0",
				Method:  methods.METHOD_transaction_V2_getBlockByNumber,
				Params: []interface{}{
					ts.LastTransactionReceiptV2.BlockNumber,
					include{
						FullTx:         true,
						WithSigners:    false,
						IncludeSigners: false,
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			data, err := json.Marshal(tc.br)
			if err != nil {
				testMetrics = append(testMetrics, TestMetric{
					Method: tc.br.Method,
					Test:   tc.name,
					Pass:   false,
					Error:  err.Error(),
					Params: tc.br.Params,
				})
				t.Error(err)
				return
			}
			// Perform the RPC Call
			resp, err := Call(data, tc.br.Method)
			if err != nil {
				testMetrics = append(testMetrics, TestMetric{
					Method:   tc.br.Method,
					Test:     tc.name,
					Pass:     false,
					Duration: resp.Duration,
					Error:    err.Error(),
					Params:   tc.br.Params,
				})
				t.Error(err)
				return
			}
			if resp.Error != nil {
				if resp.Error.Code != tc.expectedErrorCode {
					testMetrics = append(testMetrics, TestMetric{
						Method:   tc.br.Method,
						Test:     tc.name,
						Pass:     false,
						Duration: resp.Duration,
						Error:    resp.Error.Message,
						Params:   tc.br.Params,
					})
					t.Error(resp.Error.Message)
					return
				}
			}
			// This step validates that the returned response is the correct data type
			if resp.Result != nil {
				var s BlockV2
				err = json.Unmarshal(resp.Result, &s)
				if err != nil {
					testMetrics = append(testMetrics, TestMetric{
						Method:   tc.br.Method,
						Test:     tc.name,
						Pass:     false,
						Duration: resp.Duration,
						Error:    err.Error(),
						Params:   tc.br.Params,
					})
					t.Error(err)
					return
				}
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

func (ts *testSuite) test_V1_getBlockByHash(t *testing.T) {
	type testcase struct {
		name              string
		br                BaseRequest
		expectedErrorCode int64
	}

	testCases := []testcase{
		{
			name: fmt.Sprintf("%s_existing_block", t.Name()),
			br: BaseRequest{
				ID:      "1",
				JsonRPC: "2.0",
				Method:  methods.METHOD_transaction_V1_getBlockByHash,
				Params: []interface{}{
					ts.LastTransactionReceiptV1.BlockHash,
					true,
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			data, err := json.Marshal(tc.br)
			if err != nil {
				testMetrics = append(testMetrics, TestMetric{
					Method: tc.br.Method,
					Test:   tc.name,
					Pass:   false,
					Error:  err.Error(),
					Params: tc.br.Params,
				})
				t.Error(err)
				return
			}
			// Perform the RPC Call
			resp, err := Call(data, tc.br.Method)
			if err != nil {
				testMetrics = append(testMetrics, TestMetric{
					Method:   tc.br.Method,
					Test:     tc.name,
					Pass:     false,
					Duration: resp.Duration,
					Error:    err.Error(),
					Params:   tc.br.Params,
				})
				t.Error(err)
				return
			}
			if resp.Error != nil {
				if resp.Error.Code != tc.expectedErrorCode {
					testMetrics = append(testMetrics, TestMetric{
						Method:   tc.br.Method,
						Test:     tc.name,
						Pass:     false,
						Duration: resp.Duration,
						Error:    resp.Error.Message,
						Params:   tc.br.Params,
					})
					t.Error(resp.Error.Message)
					return
				}
			}
			// This step validates that the returned response is the correct data type
			if resp.Result != nil {
				var s BlockV1
				err = json.Unmarshal(resp.Result, &s)
				if err != nil {
					testMetrics = append(testMetrics, TestMetric{
						Method:   tc.br.Method,
						Test:     tc.name,
						Pass:     false,
						Duration: resp.Duration,
						Error:    err.Error(),
						Params:   tc.br.Params,
					})
					t.Error(err)
					return
				}
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

func (ts *testSuite) test_V2_getBlockByHash(t *testing.T) {
	type testcase struct {
		name              string
		br                BaseRequest
		expectedErrorCode int64
	}

	type include struct {
		FullTx         bool `json:"fullTx"`
		WithSigners    bool `json:"withSigner"`
		IncludeSigners bool `json:"includeSigners"`
	}

	testCases := []testcase{
		{
			name: fmt.Sprintf("%s_existing_block", t.Name()),
			br: BaseRequest{
				ID:      "1",
				JsonRPC: "2.0",
				Method:  methods.METHOD_transaction_V2_getBlockByHash,
				Params: []interface{}{
					ts.LastTransactionReceiptV2.BlockHash,
					include{
						FullTx:         true,
						WithSigners:    false,
						IncludeSigners: false,
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			data, err := json.Marshal(tc.br)
			if err != nil {
				testMetrics = append(testMetrics, TestMetric{
					Method: tc.br.Method,
					Test:   tc.name,
					Pass:   false,
					Error:  err.Error(),
					Params: tc.br.Params,
				})
				t.Error(err)
				return
			}
			// Perform the RPC Call
			resp, err := Call(data, tc.br.Method)
			if err != nil {
				testMetrics = append(testMetrics, TestMetric{
					Method:   tc.br.Method,
					Test:     tc.name,
					Pass:     false,
					Duration: resp.Duration,
					Error:    err.Error(),
					Params:   tc.br.Params,
				})
				t.Error(err)
				return
			}
			if resp.Error != nil {
				if resp.Error.Code != tc.expectedErrorCode {
					testMetrics = append(testMetrics, TestMetric{
						Method:   tc.br.Method,
						Test:     tc.name,
						Pass:     false,
						Duration: resp.Duration,
						Error:    resp.Error.Message,
						Params:   tc.br.Params,
					})
					t.Error(resp.Error.Message)
					return
				}
			}
			// This step validates that the returned response is the correct data type
			if resp.Result != nil {
				var s BlockV2
				err = json.Unmarshal(resp.Result, &s)
				if err != nil {
					testMetrics = append(testMetrics, TestMetric{
						Method:   tc.br.Method,
						Test:     tc.name,
						Pass:     false,
						Duration: resp.Duration,
						Error:    err.Error(),
						Params:   tc.br.Params,
					})
					t.Error(err)
					return
				}
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

func (ts *testSuite) test_V1_getBlocks(t *testing.T) {
	type testcase struct {
		name              string
		br                BaseRequest
		expectedErrorCode int64
	}

	testCases := []testcase{
		{
			name: fmt.Sprintf("%s_existing_block", t.Name()),
			br: BaseRequest{
				ID:      "1",
				JsonRPC: "2.0",
				Method:  methods.METHOD_transaction_V1_getBlocks,
				Params: []interface{}{
					ts.LastTransactionReceiptV1.BlockNumber,
					ts.LastTransactionReceiptV1.BlockNumber,
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			data, err := json.Marshal(tc.br)
			if err != nil {
				testMetrics = append(testMetrics, TestMetric{
					Method: tc.br.Method,
					Test:   tc.name,
					Pass:   false,
					Error:  err.Error(),
					Params: tc.br.Params,
				})
				t.Error(err)
				return
			}
			// Perform the RPC Call
			resp, err := Call(data, tc.br.Method)
			if err != nil {
				testMetrics = append(testMetrics, TestMetric{
					Method:   tc.br.Method,
					Test:     tc.name,
					Pass:     false,
					Duration: resp.Duration,
					Error:    err.Error(),
					Params:   tc.br.Params,
				})
				t.Error(err)
				return
			}
			if resp.Error != nil {
				if resp.Error.Code != tc.expectedErrorCode {
					testMetrics = append(testMetrics, TestMetric{
						Method:   tc.br.Method,
						Test:     tc.name,
						Pass:     false,
						Duration: resp.Duration,
						Error:    resp.Error.Message,
						Params:   tc.br.Params,
					})
					t.Error(resp.Error.Message)
					return
				}
			}
			// This step validates that the returned response is the correct data type
			if resp.Result != nil {
				var s []BlockV1
				err = json.Unmarshal(resp.Result, &s)
				if err != nil {
					testMetrics = append(testMetrics, TestMetric{
						Method:   tc.br.Method,
						Test:     tc.name,
						Pass:     false,
						Duration: resp.Duration,
						Error:    err.Error(),
						Params:   tc.br.Params,
					})
					t.Error(err)
					return
				}
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

func (ts *testSuite) test_V2_getBlocks(t *testing.T) {
	type testcase struct {
		name              string
		br                BaseRequest
		expectedErrorCode int64
	}

	type include struct {
		FullTx         bool `json:"fullTx"`
		WithSigners    bool `json:"withSigner"`
		IncludeSigners bool `json:"includeSigners"`
	}

	testCases := []testcase{
		{
			name: fmt.Sprintf("%s_existing_block", t.Name()),
			br: BaseRequest{
				ID:      "1",
				JsonRPC: "2.0",
				Method:  methods.METHOD_transaction_V2_getBlocks,
				Params: []interface{}{
					ts.LastTransactionReceiptV2.BlockNumber,
					ts.LastTransactionReceiptV2.BlockNumber,
					include{
						FullTx:         true,
						WithSigners:    false,
						IncludeSigners: false,
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			data, err := json.Marshal(tc.br)
			if err != nil {
				testMetrics = append(testMetrics, TestMetric{
					Method: tc.br.Method,
					Test:   tc.name,
					Pass:   false,
					Error:  err.Error(),
					Params: tc.br.Params,
				})
				t.Error(err)
				return
			}
			// Perform the RPC Call
			resp, err := Call(data, tc.br.Method)
			if err != nil {
				testMetrics = append(testMetrics, TestMetric{
					Method:   tc.br.Method,
					Test:     tc.name,
					Pass:     false,
					Duration: resp.Duration,
					Error:    err.Error(),
					Params:   tc.br.Params,
				})
				t.Error(err)
				return
			}
			if resp.Error != nil {
				if resp.Error.Code != tc.expectedErrorCode {
					testMetrics = append(testMetrics, TestMetric{
						Method:   tc.br.Method,
						Test:     tc.name,
						Pass:     false,
						Duration: resp.Duration,
						Error:    resp.Error.Message,
						Params:   tc.br.Params,
					})
					t.Error(resp.Error.Message)
					return
				}
			}
			// This step validates that the returned response is the correct data type
			if resp.Result != nil {
				var s []BlockV2
				err = json.Unmarshal(resp.Result, &s)
				if err != nil {
					testMetrics = append(testMetrics, TestMetric{
						Method:   tc.br.Method,
						Test:     tc.name,
						Pass:     false,
						Duration: resp.Duration,
						Error:    err.Error(),
						Params:   tc.br.Params,
					})
					t.Error(err)
					return
				}
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

func (ts *testSuite) test_tx(t *testing.T) {
	type testcase struct {
		name              string
		br                BaseRequest
		expectedErrorCode int64
	}

	testCases := []testcase{
		{
			name: fmt.Sprintf("%s_existing_tx", t.Name()),
			br: BaseRequest{
				ID:      "1",
				JsonRPC: "2.0",
				Method:  methods.METHOD_transaction_tx,
				Params: []interface{}{
					ts.LastTransactionHash,
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			data, err := json.Marshal(tc.br)
			if err != nil {
				testMetrics = append(testMetrics, TestMetric{
					Method: tc.br.Method,
					Test:   tc.name,
					Pass:   false,
					Error:  err.Error(),
					Params: tc.br.Params,
				})
				t.Error(err)
				return
			}
			// Perform the RPC Call
			resp, err := Call(data, tc.br.Method)
			if err != nil {
				testMetrics = append(testMetrics, TestMetric{
					Method:   tc.br.Method,
					Test:     tc.name,
					Pass:     false,
					Duration: resp.Duration,
					Error:    err.Error(),
					Params:   tc.br.Params,
				})
				t.Error(err)
				return
			}
			if resp.Error != nil {
				if resp.Error.Code != tc.expectedErrorCode {
					testMetrics = append(testMetrics, TestMetric{
						Method:   tc.br.Method,
						Test:     tc.name,
						Pass:     false,
						Duration: resp.Duration,
						Error:    resp.Error.Message,
						Params:   tc.br.Params,
					})
					t.Error(resp.Error.Message)
					return
				}
			}
			// This step validates that the returned response is the correct data type
			if resp.Result != nil {
				var s Transaction
				err = json.Unmarshal(resp.Result, &s)
				if err != nil {
					testMetrics = append(testMetrics, TestMetric{
						Method:   tc.br.Method,
						Test:     tc.name,
						Pass:     false,
						Duration: resp.Duration,
						Error:    err.Error(),
						Params:   tc.br.Params,
					})
					t.Error(err)
					return
				}
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
