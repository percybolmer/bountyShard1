package harmony

import (
	"encoding/json"
	"fmt"
	"percybolmer/rpc-shard-testing/rpctester/methods"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"golang.org/x/crypto/sha3"
)

func test_EstimateGas(t *testing.T) {
	type testcase struct {
		name              string
		br                BaseRequest
		expectedErrorCode int64
		data              func() RpcCallArgs
	}

	testCases := []testcase{
		{
			name: fmt.Sprintf("%s_estimate_smart_gas", t.Name()),
			data: func() RpcCallArgs {
				transferFnSignature := []byte("owner()")
				hash := sha3.NewLegacyKeccak256()
				hash.Write(transferFnSignature)
				methodID := hash.Sum(nil)[:4]
				data := hexutil.Bytes{}
				data = append(data, methodID...)

				rpcCall := RpcCallArgs{
					From: common.HexToAddress(address).String(),
					To:   smartContractAddress.String(),
					Data: data.String(),
				}
				return rpcCall
			},
			br: BaseRequest{
				ID:      "1",
				JsonRPC: "2.0",
				Method:  methods.METHOD_contract_estimateGas,
				Params:  []interface{}{},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Build the Call to perform

			rpcCall := tc.data()

			tc.br.Params = append(tc.br.Params, rpcCall)
			// Perform the RPC Call
			txdata, err := json.Marshal(tc.br)
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
			resp, err := Call(txdata, tc.br.Method)
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

func test_getStorageAt(t *testing.T) {
	type testcase struct {
		name              string
		br                BaseRequest
		expectedErrorCode int64
	}

	testCases := []testcase{
		{
			name: fmt.Sprintf("%s_get_Storage_at", t.Name()),
			br: BaseRequest{
				ID:      "1",
				JsonRPC: "2.0",
				Method:  methods.METHOD_contract_getStorageAt,
				Params: []interface{}{
					smartContractAddress,
					"0x0", // Should be totalSupply
					"latest",
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

func test_call(t *testing.T) {

	type testcase struct {
		name              string
		br                BaseRequest
		expectedErrorCode int64
		data              func() RpcCallArgs
	}

	testCases := []testcase{
		{
			name: fmt.Sprintf("%s_smart_contract_call", t.Name()),
			data: func() RpcCallArgs {
				transferFnSignature := []byte("owner()")
				hash := sha3.NewLegacyKeccak256()
				hash.Write(transferFnSignature)
				methodID := hash.Sum(nil)[:4]
				data := hexutil.Bytes{}
				data = append(data, methodID...)

				rpcCall := RpcCallArgs{
					From: common.HexToAddress(address).String(),
					To:   smartContractAddress.String(),
					Data: data.String(),
				}
				return rpcCall
			},
			br: BaseRequest{
				ID:      "1",
				JsonRPC: "2.0",
				Method:  methods.METHOD_contract_call,
				Params:  []interface{}{},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Build the Call to perform

			rpcCall := tc.data()

			tc.br.Params = append(tc.br.Params, rpcCall)
			tc.br.Params = append(tc.br.Params, "latest")
			// Perform the RPC Call
			txdata, err := json.Marshal(tc.br)
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
			resp, err := Call(txdata, tc.br.Method)
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
				responseAddr := common.HexToAddress(s).String()
				// validate addr is correct
				if responseAddr != address {
					testMetrics = append(testMetrics, TestMetric{
						Method:   tc.br.Method,
						Test:     tc.name,
						Pass:     false,
						Duration: resp.Duration,
						Error:    "Smart contract call did not return correct address",
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

func test_getCode(t *testing.T) {
	type testcase struct {
		name              string
		br                BaseRequest
		expectedErrorCode int64
	}

	testCases := []testcase{
		{
			name: fmt.Sprintf("%s_get_code", t.Name()),
			br: BaseRequest{
				ID:      "1",
				JsonRPC: "2.0",
				Method:  methods.METHOD_contract_getCode,
				Params: []interface{}{
					smartContractAddress,
					"latest",
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
