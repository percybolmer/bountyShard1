package main

import (
	"encoding/json"
	"fmt"
	"math/big"
	"testing"
)

func Test_RPC_Sanity(t *testing.T) {
	t.Run("AccountMethods", test_AccountMethods)

	GenerateReport()

}

// test_AccountsMethods calls all Account RPC Methods and verifies that the data returned is correct
func test_AccountMethods(t *testing.T) {

	t.Run("getBalanceByBlockNumber_V2", test_V2_getBalanceByBlockNumber)
	t.Run("getBalanceByBlockNumber_V1", test_V1_getBalanceByBlockNumber)
	t.Run("getTransactionCount_V2", test_V2_hmy_getTransactionCount)
	t.Run("getTransactionCount_V1", test_V1_hmy_getTransactionCount)
	t.Run("getBalance_V2", test_V2_hmy_getBalance)
	t.Run("getBalance_V1", test_V1_hmy_getBalance)
	t.Run("address", test_address)
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
					t.Error(err)
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
					t.Error(err)
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
					t.Error(err)
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
					t.Error(err)
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
					t.Error(err)
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
					t.Error(err)
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
