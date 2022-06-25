package main

import (
	"encoding/json"
	"fmt"
	"math/big"
	"testing"
	"time"
)

var (
	// createdFilterID will be filled by the createFilter Test
	createdFilterID = ""

	ts *testSuite
)

func init() {
	ts = &testSuite{}
}
func Test_RPC_Sanity(t *testing.T) {
	t.Run("AccountMethods", test_AccountMethods)
	t.Run("FilterMethods", test_FilterMethods)
	t.Run("TransactionMethods", test_TransactionMethods)
	// Now generate report
	GenerateReport()

}

// test_TransactionMethods is used to generate transactions and verify their data on the RPC
func test_TransactionMethods(t *testing.T) {
	// Begin by sending Transaction so we can get transaction hashes to use
	t.Run("sendRawTransaction", ts.test_sendRawTransaction)
	t.Run("getStakingTransactionByBlockHashAndIndex", test_V1_getStakingTransactionByBlockHashAndIndex)
	t.Run("getStakingTransactionByBlockHasAndIndex_V2", test_V2_getStakingTransactionByBlockHashAndIndex)
	t.Run("getStakingTransactionByBlockNumberAndIndex", test_V1_getStakingTransactionByBlockNumberAndIndex)
	t.Run("getStakingTransactionByBlockNumberAndIndex_V2", test_V2_getStakingTransactionByBlockNumberAndIndex)
	t.Run("getStakingTransactionByHash_V1", test_V1_getStakingTransactionByHash)
	t.Run("getStakingTransactionByHash_V2", test_V2_getStakingTransactionByHash)
	t.Run("getCurrentTransactionErrorSink_V1", test_V1_getCurrentTransactionErrorSink)
	t.Run("getCurrentTransactionErrorSink_V2", test_V2_getCurrentTransactionErrorSink)
	t.Run("getPendingCrossLinks_V1", test_V1_getPendingCrossLinks)
	t.Run("getPendingCrossLinks_V2", test_V2_getPendingCrossLinks)
	t.Run("getPendingCXReceipts_V1", test_V1_getPendingCXReceipts)
	t.Run("getPendingCXReceipts_V2", test_V2_getPendingCXReceipts)
	t.Run("getCXReceiptByHash_V1", test_V1_getCXReceiptByHash)
	t.Run("getCXReceiptByHash_V2", test_V2_getCXReceiptByHash)
	t.Run("getPendingTransaction_V1", test_V1_pendingTransactions)
	t.Run("getPendingTransaction_V2", test_V2_pendingTransactions)
	// TODO confirm how to
	//t.Run("sendRawStakingTransaction", test_sendRawStakingTransaction)
	t.Log("Giving the network some time to congest transactions sent")
	time.Sleep(10 * time.Second)
	t.Run("getTransactionHistory_V1", ts.test_V1_getTransactionHistory)
	t.Run("getTransactionHistory_V2", ts.test_V2_getTransactionHistory)
	t.Run("getTransactionReceipt_V1", ts.test_V1_getTransactionReceipt)
	t.Run("getTransactionReceipt_V2", ts.test_V2_getTransactionReceipt)
	t.Run("getBlockTransactionCountByHash_V1", ts.test_V1_getBlockTransactionCountByHash)
	t.Run("getBlockTransactionCountByHash_V2", ts.test_V2_getBlockTransactionCountByHash)
	t.Run("getBlockTransactionCountByNumber_V1", ts.test_V1_getBlockTransactionCountByNumber)
	t.Run("getBlockTransactionCountByNumber_V2", ts.test_V2_getBlockTransactionCountByNumber)
	t.Run("getTransactionByHash_V1", ts.test_V1_getTransactionByHash)
	t.Run("getTransactionByHash_V1", ts.test_V1_getTransactionByHash)
	t.Run("getTransactionByBlockNumberAndIndex_V1", ts.test_V1_getTransactionByBlockNumberAndIndex)
	t.Run("getTransactionByBlockNumberAndIndex_V2", ts.test_V2_getTransactionByBlockNumberAndIndex)
	t.Run("getTransactionByBlockHashAndIndex_V1", ts.test_V1_getTransactionByBlockHashAndIndex)
	t.Run("getTransactionByBlockHashAndIndex_V2", ts.test_V2_getTransactionByBlockHashAndIndex)
	t.Run("getBlockByNumber_V1", ts.test_V1_getBlockByNumber)
	t.Run("getBlockByNumber_V2", ts.test_V2_getBlockByNumber)
	t.Run("getBlockByHash_V1", ts.test_V1_getBlockByHash)
	t.Run("getBlockByHash_V2", ts.test_V2_getBlockByHash)
	t.Run("getBlocks_V1", ts.test_V1_getBlocks)
	t.Run("getBlocks_V2", ts.test_V2_getBlocks)
	t.Run("tx", ts.test_tx)

}

// test_AccountsMethods calls all Account RPC Methods and verifies that the data returned is correct
func test_AccountMethods(t *testing.T) {

	t.Run("getBalanceByBlockNumber_V2", test_V2_getBalanceByBlockNumber)
	t.Run("getBalanceByBlockNumber_V1", test_V1_getBalanceByBlockNumber)
	t.Run("getTransactionCount_V2", test_V2_hmy_getTransactionCount)
	t.Run("getTransactionCount_V1", test_V1_hmy_getTransactionCount)
	t.Run("getBalance_V2", test_V2_hmy_getBalance)
	t.Run("getBalance_V1", test_V1_hmy_getBalance)
	//t.Run("address", test_address)
}

// test_FilterMethods is used to call Filter RPC methods and validate their return data
func test_FilterMethods(t *testing.T) {
	t.Run("newFilter", test_newFilter)
	// Add some delay between newFilter and getFilterLogs so it has the time to actually create
	t.Log("Waiting for Filter to be created")
	time.Sleep(1 * time.Second)
	t.Run("getFilterLogs", test_getFilterLogs)
	t.Run("newPendingTransactionFilter", test_NewPendingTransactionFilter)
	t.Run("newBlockFilter", test_NewBlockFilter)
	t.Run("getFilterChanges", test_getFilterChanges)
	t.Run("getLogs", test_getLogs)
}

func test_getLogs(t *testing.T) {
	type testcase struct {
		name              string
		br                BaseRequest
		expectedErrorCode int64
	}

	testCases := []testcase{
		{
			name: fmt.Sprintf("%s_working_log", t.Name()),
			br: BaseRequest{
				ID:      "1",
				JsonRPC: "2.0",
				Method:  METHOD_filter_getLogs,
				Params: []interface{}{
					Filter{
						Address: "0xcF664087a5bB0237a0BAd6742852ec6c8d69A27a",
						//FromBlock: "latest",
						//Topics: []string{"0x000000000000000000000000a94f5374fce5edbc8e2a8697c15331677e6ebf0b"},
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
					})
					t.Error(resp.Error.Message)
					return
				}
			}
			// This step validates that the returned response is the correct data type
			if resp.Result != nil {
				var s []FilterChange
				err = json.Unmarshal(resp.Result, &s)
				if err != nil {
					testMetrics = append(testMetrics, TestMetric{
						Method:   tc.br.Method,
						Test:     tc.name,
						Pass:     false,
						Duration: resp.Duration,
						Error:    err.Error(),
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
			})

		})
	}
}

func test_getFilterChanges(t *testing.T) {
	type testcase struct {
		name              string
		br                BaseRequest
		expectedErrorCode int64
	}

	testCases := []testcase{
		{
			name: fmt.Sprintf("%s_change_log", t.Name()),
			br: BaseRequest{
				ID:      "1",
				JsonRPC: "2.0",
				Method:  METHOD_filter_getFilterChanges,
				Params: []interface{}{
					createdFilterID,
				},
			},
		},
		{
			name:              fmt.Sprintf("%s_no_such_id", t.Name()),
			expectedErrorCode: -32000,
			br: BaseRequest{
				ID:      "1",
				JsonRPC: "2.0",
				Method:  METHOD_filter_getFilterChanges,
				Params: []interface{}{
					"0x01",
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
					})
					t.Error(resp.Error.Message)
					return
				}
			}
			// This step validates that the returned response is the correct data type
			if resp.Result != nil {
				var s []FilterChange
				err = json.Unmarshal(resp.Result, &s)
				if err != nil {
					testMetrics = append(testMetrics, TestMetric{
						Method:   tc.br.Method,
						Test:     tc.name,
						Pass:     false,
						Duration: resp.Duration,
						Error:    err.Error(),
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
			})

		})
	}
}

func test_NewBlockFilter(t *testing.T) {
	type testcase struct {
		name              string
		br                BaseRequest
		expectedErrorCode int64
	}

	testCases := []testcase{
		{
			name: fmt.Sprintf("%s_new_blockfilter", t.Name()),
			br: BaseRequest{
				ID:      "1",
				JsonRPC: "2.0",
				Method:  METHOD_filter_newBlockFilter,
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

func test_NewPendingTransactionFilter(t *testing.T) {
	type testcase struct {
		name              string
		br                BaseRequest
		expectedErrorCode int64
	}

	testCases := []testcase{
		{
			name: fmt.Sprintf("%s_new_pending_transaction", t.Name()),
			br: BaseRequest{
				ID:      "1",
				JsonRPC: "2.0",
				Method:  METHOD_filter_newPendingtransactionFilter,
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

// newFilter will store a successful filter in the
func test_newFilter(t *testing.T) {
	type testcase struct {
		name              string
		br                BaseRequest
		expectedErrorCode int64
	}

	testCases := []testcase{
		{
			name: fmt.Sprintf("%s_new_filter", t.Name()),
			br: BaseRequest{
				ID:      "1",
				JsonRPC: "2.0",
				Method:  METHOD_filter_newFilter,
				Params: []interface{}{
					Filter{

						FromBlock: "0x1",
						ToBlock:   "0x2",
						Address:   address,
						Topics:    []string{"0x000000000000000000000000a94f5374fce5edbc8e2a8697c15331677e6ebf0b"},
					},
				},
			},
		}, {
			name:              fmt.Sprintf("%s_bad_topic_format", t.Name()),
			expectedErrorCode: -32602,
			br: BaseRequest{
				ID:      "1",
				JsonRPC: "2.0",
				Method:  METHOD_filter_newFilter,
				Params: []interface{}{
					Filter{

						FromBlock: "0x1",
						ToBlock:   "0x2",
						Address:   address,
						Topics:    []string{"notatopic"},
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
				// Add a valid filter ID to the global filterID variable
				// so other tests can use this filter
				createdFilterID = s
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

func test_getFilterLogs(t *testing.T) {
	type testcase struct {
		name              string
		br                BaseRequest
		expectedErrorCode int64
	}

	testCases := []testcase{
		{
			name: fmt.Sprintf("%s_working_filter", t.Name()),
			br: BaseRequest{
				ID:      "1",
				JsonRPC: "2.0",
				Method:  METHOD_filter_getFilterLogs,
				Params: []interface{}{
					createdFilterID,
				},
			},
		}, {
			name:              fmt.Sprintf("%s_missing_param", t.Name()),
			expectedErrorCode: -32602,
			br: BaseRequest{
				ID:      "1",
				JsonRPC: "2.0",
				Method:  METHOD_filter_getFilterLogs,
				Params:  []interface{}{},
			},
		}, {
			name:              fmt.Sprintf("%s_filter_not_found", t.Name()),
			expectedErrorCode: -32000,
			br: BaseRequest{
				ID:      "1",
				JsonRPC: "2.0",
				Method:  METHOD_filter_getFilterLogs,
				Params: []interface{}{
					"0x16",
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
				var s []FilterChange
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

func test_address(t *testing.T) {
	type testcase struct {
		name              string
		id                string
		offset            int
		page              int
		tx_view           string
		expectedErrorCode int64
	}

	testCases := []testcase{
		{
			name:    fmt.Sprintf("%s_addr", t.Name()),
			id:      address,
			offset:  0,
			page:    2,
			tx_view: "ALL",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Perform the RPC Call
			response, err := Address(tc.id, tc.offset, tc.page, tc.tx_view)
			if err != nil {
				testMetrics = append(testMetrics, TestMetric{
					Method: METHOD_address,
					Test:   tc.name,
					Pass:   false,
					Error:  err.Error(),
				})
				t.Error(err)
				return
			}

			// This step validates that the returned response is the correct data type

			testMetrics = append(testMetrics, TestMetric{
				Method:   METHOD_address,
				Test:     tc.name,
				Duration: response.Duration,
				Pass:     true,
			})

		})
	}
}

func test_V2_hmy_getBalance(t *testing.T) {
	type testcase struct {
		name              string
		br                BaseRequest
		expectedErrorCode int64
	}

	testCases := []testcase{
		{
			name: fmt.Sprintf("%s_working_count", t.Name()),
			br: BaseRequest{
				ID:      "1",
				JsonRPC: "2.0",
				Method:  METHOD_V2_getBalance,
				Params: []interface{}{
					address,
				},
			},
		}, {
			name:              fmt.Sprintf("%s_missing_param", t.Name()),
			expectedErrorCode: -32602,
			br: BaseRequest{
				ID:      "1",
				JsonRPC: "2.0",
				Method:  METHOD_V1_getBalance,
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
				var s big.Int
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
			})

		})
	}
}
func test_V1_hmy_getBalance(t *testing.T) {
	type testcase struct {
		name              string
		br                BaseRequest
		expectedErrorCode int64
	}

	testCases := []testcase{
		{
			name: fmt.Sprintf("%s_working_count", t.Name()),
			br: BaseRequest{
				ID:      "1",
				JsonRPC: "2.0",
				Method:  METHOD_V1_getBalance,
				Params: []interface{}{
					address,
					"latest",
				},
			},
		}, {
			name:              fmt.Sprintf("%s_missing_param", t.Name()),
			expectedErrorCode: -32602,
			br: BaseRequest{
				ID:      "1",
				JsonRPC: "2.0",
				Method:  METHOD_V1_getBalance,
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
func test_V1_hmy_getTransactionCount(t *testing.T) {
	type testcase struct {
		name              string
		br                BaseRequest
		expectedErrorCode int64
	}

	testCases := []testcase{
		{
			name: fmt.Sprintf("%s_working_count", t.Name()),
			br: BaseRequest{
				ID:      "1",
				JsonRPC: "2.0",
				Method:  METHOD_V1_getTransactionCount,
				Params: []interface{}{
					address,
					"latest",
				},
			},
		}, {
			name:              fmt.Sprintf("%s_missing_param", t.Name()),
			expectedErrorCode: -32602,
			br: BaseRequest{
				ID:      "1",
				JsonRPC: "2.0",
				Method:  METHOD_V1_getTransactionCount,
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
func test_V2_hmy_getTransactionCount(t *testing.T) {
	type testcase struct {
		name              string
		br                BaseRequest
		expectedErrorCode int64
	}

	testCases := []testcase{
		{
			name: fmt.Sprintf("%s_working_count", t.Name()),
			br: BaseRequest{
				ID:      "1",
				JsonRPC: "2.0",
				Method:  METHOD_V2_getTransactionCount,
				Params: []interface{}{
					address,
					1,
				},
			},
		}, {
			name:              fmt.Sprintf("%s_missing_param", t.Name()),
			expectedErrorCode: -32602,
			br: BaseRequest{
				ID:      "1",
				JsonRPC: "2.0",
				Method:  METHOD_V2_getTransactionCount,
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
				var s big.Int
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

func test_V2_getBalanceByBlockNumber(t *testing.T) {
	type testcase struct {
		name              string
		br                BaseRequest
		expectedErrorCode int64
	}

	testCases := []testcase{
		{
			name: fmt.Sprintf("%s_block_initial_balance", t.Name()),
			br: BaseRequest{
				ID:      "1",
				JsonRPC: "2.0",
				Method:  METHOD_V2_getBalanceByBlockNumber,
				Params: []interface{}{
					address,
					"1",
				},
			},
		}, {
			name:              fmt.Sprintf("%s_missing_param", t.Name()),
			expectedErrorCode: -32602,
			br: BaseRequest{
				ID:      "1",
				JsonRPC: "2.0",
				Method:  METHOD_V2_getBalanceByBlockNumber,
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

			if resp.Result != nil {
				var s big.Int
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

func test_V1_getBalanceByBlockNumber(t *testing.T) {
	type testcase struct {
		name              string
		br                BaseRequest
		expectedErrorCode int64
	}

	testCases := []testcase{
		{
			name: fmt.Sprintf("%s_getBalanceByBlcok", t.Name()),
			br: BaseRequest{
				ID:      "1",
				JsonRPC: "2.0",
				Method:  METHOD_V1_getBalanceByBlockNumber,
				Params: []interface{}{
					address,
					"0x01",
				},
			},
		}, {
			name:              fmt.Sprintf("%s_missing_param", t.Name()),
			expectedErrorCode: -32602,
			br: BaseRequest{
				ID:      "1",
				JsonRPC: "2.0",
				Method:  METHOD_V1_getBalanceByBlockNumber,
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
