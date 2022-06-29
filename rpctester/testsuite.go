package main

import (
	"encoding/json"
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

// testSuite is used to Sync information between some sanity tests
// such as storing TX hashes and blocks
type testSuite struct {
	LastTransactionHash        string
	LastTransactionBlockHash   string
	LastTransactionBlockNumber int64
	LastTransactionReceiptV1   TransactionReceipt_V1
	LastTransactionReceiptV2   TransactionReceipt_V2
	NetworkHeader              NetworkHeader
	ValidatorsV1               GetValidatorsV1
	ValidatorsV2               GetValidatorsV2
	ValidatorInfo              ValidatorInfo

	LastStakingTransactionHash      string
	LastStakingTransactionBlockHash string

	ActiveValidators  []string
	ElectedValidators []string
}

func (ts *testSuite) test_sendRawTransaction(t *testing.T) {
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

			toAddr := common.HexToAddress(address)
			fromAddr := common.HexToAddress(address)

			// Transfer 0.001 ONE
			rlp, err := CreateRLPString(toAddr, fromAddr, *big.NewInt(0).Div(ONE, big.NewInt(1000)), nil)
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
			tc.br.Params = append(tc.br.Params, rlp)
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

				// TODO make this the first call ever so we can use this TransacionHash
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
				// Store TX hash in TestSuite
				ts.LastTransactionHash = s
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
