package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	localcrypto "percybolmer/rpc-shard-testing/rpctester/crypto"
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/harmony-one/harmony/core/types"
)

func test_V1_getStakingTransactionByBlockHashAndIndex(t *testing.T) {
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
				Method:  METHOD_transaction_V1_getStakingTransactionByBlockHashAndIndex,
				Params: []interface{}{
					"0x428ead93e632d5255ea3d1fb61b02ab8493cf562a398af2159c33ecd53c62c16",
					"0x0",
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

func test_V2_getStakingTransactionByBlockHashAndIndex(t *testing.T) {
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
				Method:  METHOD_transaction_V2_getStakingTransactionByBlockHashAndIndex,
				Params: []interface{}{
					"0x428ead93e632d5255ea3d1fb61b02ab8493cf562a398af2159c33ecd53c62c16",
					0,
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

func test_V1_getStakingTransactionByBlockNumberAndIndex(t *testing.T) {
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
				Method:  METHOD_transaction_V1_getStakingTransactionByBlockNumberAndIndex,
				Params: []interface{}{
					"0x4",
					"0x0",
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

func test_V2_getStakingTransactionByBlockNumberAndIndex(t *testing.T) {
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
				Method:  METHOD_transaction_V2_getStakingTransactionByBlockNumberAndIndex,
				Params: []interface{}{
					4,
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
				Method:  METHOD_transaction_V2_getStakingTransactionByBlockNumberAndIndex,
				Params: []interface{}{
					4,
					0,
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

func test_V2_getStakingTransactionByHash(t *testing.T) {
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
				Method:  METHOD_transaction_V2_getStakingTransactionByHash,
				Params: []interface{}{
					"0x1dff358dad4d0fc95b11acc9826b190d8b7971ac26b3f7ebdee83c10cafaf86f",
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

func test_V1_getStakingTransactionByHash(t *testing.T) {
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
				Method:  METHOD_transaction_V1_getStakingTransactionByHash,
				Params: []interface{}{
					"0x1dff358dad4d0fc95b11acc9826b190d8b7971ac26b3f7ebdee83c10cafaf86f",
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
				Method:  METHOD_transaction_V1_getCurrentTransactionErrorSink,
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
				Method:  METHOD_transaction_V2_getCurrentTransactionErrorSink,
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
				Method:  METHOD_transaction_V1_getPendingCrossLinks,
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
				Method:  METHOD_transaction_V2_getPendingCrossLinks,
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
				Method:  METHOD_transaction_V1_getPendingCXReceipts,
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
				Method:  METHOD_transaction_V2_getPendingCXReceipts,
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
				Method:  METHOD_transaction_V1_getCXReceiptByHash,
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
				Method:  METHOD_transaction_V2_getCXReceiptByHash,
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
				Method:  METHOD_transaction_V1_pendingTransactions,
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
				Method:  METHOD_transaction_V2_pendingTransactions,
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

func test_sendRawStakingTransaction(t *testing.T) {
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
				Method:  METHOD_transaction_sendRawStakingTransaction,
				Params: []interface{}{
					"", // TODO, How to Sign and Format this transaction?
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
func test_sendRawTransaction(t *testing.T) {
	type testcase struct {
		name              string
		br                BaseRequest
		expectedErrorCode int64
	}

	testCases := []testcase{
		{
			name: fmt.Sprintf("%s_raw_transaction", t.Name()),
			br: BaseRequest{
				ID:      "1",
				JsonRPC: "2.0",
				Method:  METHOD_transaction_sendRawTransaction,
				Params:  []interface{}{},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			// Build the Transaction usign the deploy smart contract
			if deployedToken == nil {
				t.Skip("Skipping test because no smart contract deployed")
			}
			// Create a Transaction and Sign it for a transfer
			fromAddress := localcrypto.GetAddress()

			gasPrice, err := ethClient.SuggestGasPrice(context.Background())
			if err != nil {
				log.Fatal(err)
			}

			txData := Txdata{
				AccountNonce: auth.Nonce.Uint64(),
				Price:        gasPrice,
				GasLimit:     auth.GasLimit,
				ShardID:      2,
				ToShardID:    2,
				Recipient:    &fromAddress,
				Amount:       big.NewInt(10),
			}
			bdata, err := json.Marshal(txData)
			if err != nil {
				log.Fatal(err)
			}
			tx := types.NewTransaction(auth.Nonce.Uint64(), fromAddress, 0, big.NewInt(100), auth.GasLimit, gasPrice, bdata)
			signedTx, err := types.SignTx(tx, types.NewEIP155Signer(big.NewInt(2)), localcrypto.GetPrivateKey())
			if err != nil {
				log.Fatal(err)
			}
			// Convert into a harmony Transaction since they use their own with ShardID

			ts := types.Transactions{signedTx}

			rlp := ts.GetRlp(0)

			rawTxHex := hexutil.Encode(rlp)
			tc.br.Params = append(tc.br.Params, rawTxHex)
			// Perform the RPC Call
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
func test_V1_getTransactionHistory(t *testing.T) {
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
				Method:  METHOD_transaction_V1_getTransactionHistory,
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
func test_V2_getTransactionHistory(t *testing.T) {
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
				Method:  METHOD_transaction_V2_getTransactionHistory,
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
				Method:  METHOD_transaction_V2_getTransactionHistory,
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

func test_V1_getTransactionReceipt(t *testing.T) {
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
				Method:  METHOD_transaction_V1_getTransactionReceipt,
				Params: []interface{}{
					smartContractDeploymentHash,
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

func test_V2_getTransactionReceipt(t *testing.T) {
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
				Method:  METHOD_transaction_V2_getTransactionReceipt,
				Params: []interface{}{
					smartContractDeploymentHash,
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
