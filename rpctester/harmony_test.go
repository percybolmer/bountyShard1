package rpctester

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"strings"
	"testing"
)

func Test_AccountMethods(t *testing.T) {

	t.Run("Account Methods", test_V2_getBalanceByBlockNumber)
	t.Run("Account Methods", test_V1_getBalanceByBlockNumber)
	// After all tests, Generate report

	data, err := json.Marshal(testMetrics)
	if err != nil {
		t.Fatal(err)
	}

	err = ioutil.WriteFile("results.json", data, os.ModePerm)
	if err != nil {
		t.Fatal(err)
	}
}

func test_V2_getBalanceByBlockNumber(t *testing.T) {
	type testcase struct {
		name              string
		br                BaseRequest
		expectedBalance   string
		expectedErrorCode int64
	}

	testCases := []testcase{
		{
			name:            "genesis_addr_initial_balance",
			expectedBalance: InitialBalance(*url, "V2"),
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
			name:              "missing block",
			expectedBalance:   "0",
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

				if !strings.EqualFold(s.String(), tc.expectedBalance) {
					errMsg := fmt.Sprintf("%s != %s ", s.String(), tc.expectedBalance)
					testMetrics = append(testMetrics, TestMetric{
						Method:   tc.br.Method,
						Test:     tc.name,
						Pass:     false,
						Duration: resp.Duration,
						Error:    errMsg,
					})
					t.Error(errMsg)
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
		expectedBalance   string
		expectedErrorCode int64
	}

	testCases := []testcase{
		{
			name:            "addr_initial_balance",
			expectedBalance: InitialBalance(*url, "V1"),
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
			name:              "missing block",
			expectedBalance:   "0",
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

				if !strings.EqualFold(s, tc.expectedBalance) {
					errMsg := fmt.Sprintf("%s != %s ", s, tc.expectedBalance)
					testMetrics = append(testMetrics, TestMetric{
						Method:   tc.br.Method,
						Test:     tc.name,
						Pass:     false,
						Duration: resp.Duration,
						Error:    errMsg,
					})
					t.Error(errMsg)
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
