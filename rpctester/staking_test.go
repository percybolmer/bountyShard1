package main

import (
	"fmt"
	"math/big"
	"testing"
)

func test_getCirculatingSupply(t *testing.T) {
	type testcase struct {
		name              string
		br                BaseRequest
		expectedErrorCode int64
	}

	testCases := []testcase{
		{
			name: fmt.Sprintf("%s_get_circulating_supply", t.Name()),
			br: BaseRequest{
				ID:      "1",
				JsonRPC: "2.0",
				Method:  METHOD_staking_getCirculatingSupply,
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

func test_getTotalSupply(t *testing.T) {
	type testcase struct {
		name              string
		br                BaseRequest
		expectedErrorCode int64
	}

	testCases := []testcase{
		{
			name: fmt.Sprintf("%s_get_total_supply", t.Name()),
			br: BaseRequest{
				ID:      "1",
				JsonRPC: "2.0",
				Method:  METHOD_staking_getTotalSupply,
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

func test_getStakingNetworkInfo(t *testing.T) {
	type testcase struct {
		name              string
		br                BaseRequest
		expectedErrorCode int64
	}

	testCases := []testcase{
		{
			name: fmt.Sprintf("%s_get_staking_network", t.Name()),
			br: BaseRequest{
				ID:      "1",
				JsonRPC: "2.0",
				Method:  METHOD_staking_getStakingNetworkInfo,
				Params:  []interface{}{},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var result NetworkStakingInfo

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

func test_getAllValidatorInformation(t *testing.T) {
	type testcase struct {
		name              string
		br                BaseRequest
		expectedErrorCode int64
	}

	testCases := []testcase{
		{
			name: fmt.Sprintf("%s_get_validator_information", t.Name()),
			br: BaseRequest{
				ID:      "1",
				JsonRPC: "2.0",
				Method:  METHOD_staking_getAllValidatorInformation,
				Params:  []interface{}{},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var result []ValidatorInfo

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

func (ts *testSuite) test_getValidatorInformation(t *testing.T) {
	type testcase struct {
		name              string
		br                BaseRequest
		expectedErrorCode int64
	}
	if len(ts.ValidatorsV2.Validators) < 1 {
		t.Skip("Skipping because we cant fetch validator info without validators")
	}
	testCases := []testcase{
		{
			name: fmt.Sprintf("%s_get_validator_information", t.Name()),
			br: BaseRequest{
				ID:      "1",
				JsonRPC: "2.0",
				Method:  METHOD_staking_getAllValidatorInformation,
				Params: []interface{}{
					ts.ValidatorsV2.Validators[0].Address,
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var result ValidatorInfo

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

func (ts *testSuite) test_getAllValidatorInformationByBlockNumber(t *testing.T) {
	type testcase struct {
		name              string
		br                BaseRequest
		expectedErrorCode int64
	}

	testCases := []testcase{
		{
			name: fmt.Sprintf("%s_get_validator_information_by_block", t.Name()),
			br: BaseRequest{
				ID:      "1",
				JsonRPC: "2.0",
				Method:  METHOD_staking_getAllValidatorInformationByBlockNumber,
				Params: []interface{}{
					0,
					ts.LastTransactionBlockNumber,
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var result []ValidatorInfo

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

func test_getUtilityMetric(t *testing.T) {
	type testcase struct {
		name              string
		br                BaseRequest
		expectedErrorCode int64
	}

	testCases := []testcase{
		{
			name: fmt.Sprintf("%s_get_utility_metric", t.Name()),
			br: BaseRequest{
				ID:      "1",
				JsonRPC: "2.0",
				Method:  METHOD_staking_getCurrentUtilityMetrics,
				Params:  []interface{}{},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var result UtilityMetrics

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

func (ts *testSuite) test_getDelegationsByValidator(t *testing.T) {
	type testcase struct {
		name              string
		br                BaseRequest
		expectedErrorCode int64
	}

	if len(ts.ValidatorsV2.Validators) < 1 {
		t.Skip("Skipping because we cant fetch validator info without validators")
	}
	testCases := []testcase{
		{
			name: fmt.Sprintf("%s_test_getDelegationsByValidator", t.Name()),
			br: BaseRequest{
				ID:      "1",
				JsonRPC: "2.0",
				Method:  METHOD_staking_getDelegationsByValidator,
				Params: []interface{}{
					ts.ValidatorsV2.Validators[0].Address,
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var result []DelegationByValidator

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

func (ts *testSuite) test_getDelegationsByDelegatorAndValidator(t *testing.T) {
	type testcase struct {
		name              string
		br                BaseRequest
		expectedErrorCode int64
	}
	if len(ts.ValidatorsV2.Validators) < 1 {
		t.Skip("Skipping because we cant fetch validator info without validators")
	}
	testCases := []testcase{
		{
			name: fmt.Sprintf("%s_test_getDelegationsByDelegatorAndValidator", t.Name()),
			br: BaseRequest{
				ID:      "1",
				JsonRPC: "2.0",
				Method:  METHOD_staking_getDelegationsByDelegatorAndValidator,
				Params: []interface{}{
					address,
					ts.ValidatorsV2.Validators[0].Address,
				},
				//String - delegator bech32 address.
				// String - validator bech32 address.
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var result []DelegationByValidator

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

func (ts *testSuite) test_getDelegationsByDelegator(t *testing.T) {
	type testcase struct {
		name              string
		br                BaseRequest
		expectedErrorCode int64
	}

	testCases := []testcase{
		{
			name: fmt.Sprintf("%s_validators_exist", t.Name()),
			br: BaseRequest{
				ID:      "1",
				JsonRPC: "2.0",
				Method:  METHOD_staking_getDelegationsByDelegator,
				Params: []interface{}{
					address,
				},
				//String - delegator bech32 address.
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var result []DelegationByValidator

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

func (ts *testSuite) test_getValidatorMetrics(t *testing.T) {
	type testcase struct {
		name              string
		br                BaseRequest
		expectedErrorCode int64
	}
	if len(ts.ValidatorsV2.Validators) < 1 {
		t.Skip("Skipping because we cant fetch validator info without validators")
	}
	testCases := []testcase{
		{
			name: fmt.Sprintf("%s_validators_exist", t.Name()),
			br: BaseRequest{
				ID:      "1",
				JsonRPC: "2.0",
				Method:  METHOD_staking_getValidatorMetrics,
				Params: []interface{}{
					ts.ValidatorsV2.Validators[0].Address,
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var result ValidatorMetrics

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

func test_getMedianRawStakeSnapshot(t *testing.T) {
	type testcase struct {
		name              string
		br                BaseRequest
		expectedErrorCode int64
	}

	testCases := []testcase{
		{
			name: fmt.Sprintf("%s_snapshot", t.Name()),
			br: BaseRequest{
				ID:      "1",
				JsonRPC: "2.0",
				Method:  METHOD_staking_getMedianRawStakeSnapshot,
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

func test_getActiveValidatorAddresses(t *testing.T) {
	type testcase struct {
		name              string
		br                BaseRequest
		expectedErrorCode int64
	}

	testCases := []testcase{
		{
			name: fmt.Sprintf("%s_activeValidators", t.Name()),
			br: BaseRequest{
				ID:      "1",
				JsonRPC: "2.0",
				Method:  METHOD_staking_getActiveValidatorAddresses,
				Params:  []interface{}{},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var result []string

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

func test_getAllValidatorAddresses(t *testing.T) {
	type testcase struct {
		name              string
		br                BaseRequest
		expectedErrorCode int64
	}

	testCases := []testcase{
		{
			name: fmt.Sprintf("%s_all_validators", t.Name()),
			br: BaseRequest{
				ID:      "1",
				JsonRPC: "2.0",
				Method:  METHOD_staking_getAllValidatorAddresses,
				Params:  []interface{}{},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var result []string

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

func test_V1_getCurrentStakingErrorSink(t *testing.T) {
	type testcase struct {
		name              string
		br                BaseRequest
		expectedErrorCode int64
	}

	testCases := []testcase{
		{
			name: fmt.Sprintf("%s_V1_sinks", t.Name()),
			br: BaseRequest{
				ID:      "1",
				JsonRPC: "2.0",
				Method:  METHOD_staking_V1_getCurrentStakingErrorSink,
				Params:  []interface{}{},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var result []ErrorSinkLog

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

func test_V2_getCurrentStakingErrorSink(t *testing.T) {
	type testcase struct {
		name              string
		br                BaseRequest
		expectedErrorCode int64
	}

	testCases := []testcase{
		{
			name: fmt.Sprintf("%s_V2_sinks", t.Name()),
			br: BaseRequest{
				ID:      "1",
				JsonRPC: "2.0",
				Method:  METHOD_staking_V2_getCurrentStakingErrorSink,
				Params:  []interface{}{},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var result []ErrorSinkLog

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

func (ts *testSuite) test_V1_getValidators(t *testing.T) {
	type testcase struct {
		name              string
		br                BaseRequest
		expectedErrorCode int64
	}

	testCases := []testcase{
		{
			name: fmt.Sprintf("%s_V1_validators", t.Name()),
			br: BaseRequest{
				ID:      "1",
				JsonRPC: "2.0",
				Method:  METHOD_staking_V1_getValidators,
				Params: []interface{}{
					ts.NetworkHeader.Epoch,
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var result GetValidatorsV1

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
			ts.ValidatorsV1 = result

		})
	}
}

func (ts *testSuite) test_V2_getValidators(t *testing.T) {
	type testcase struct {
		name              string
		br                BaseRequest
		expectedErrorCode int64
	}

	testCases := []testcase{
		{
			name: fmt.Sprintf("%s_V2_validators", t.Name()),
			br: BaseRequest{
				ID:      "1",
				JsonRPC: "2.0",
				Method:  METHOD_staking_V2_getValidators,
				Params: []interface{}{
					ts.NetworkHeader.Epoch,
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var result GetValidatorsV2

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
			ts.ValidatorsV2 = result

		})
	}
}

func test_getSignedBlocks(t *testing.T) {
	type testcase struct {
		name              string
		br                BaseRequest
		expectedErrorCode int64
	}
	if len(ts.ValidatorsV2.Validators) < 1 {
		t.Skip("Skipping because we cant fetch validator info without validators")
	}
	testCases := []testcase{
		{
			name: fmt.Sprintf("%s_signedBlocks", t.Name()),
			br: BaseRequest{
				ID:      "1",
				JsonRPC: "2.0",
				Method:  METHOD_staking_getSignedBlocks,
				Params: []interface{}{
					ts.ValidatorsV2.Validators[0].Address,
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

func (ts *testSuite) test_V1_isBlockSigner(t *testing.T) {
	type testcase struct {
		name              string
		br                BaseRequest
		expectedErrorCode int64
	}
	if len(ts.ValidatorsV1.Validators) < 1 {
		t.Skip("Skipping because we cant fetch validator info without validators")
	}
	testCases := []testcase{
		{
			name: fmt.Sprintf("%s_v1_validator", t.Name()),
			br: BaseRequest{
				ID:      "1",
				JsonRPC: "2.0",
				Method:  METHOD_staking_V1_isBlockSigner,
				Params: []interface{}{
					"0x0",
					ts.ValidatorsV1.Validators[0].Address,
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

func (ts *testSuite) test_V2_isBlockSigner(t *testing.T) {
	type testcase struct {
		name              string
		br                BaseRequest
		expectedErrorCode int64
	}
	if len(ts.ValidatorsV2.Validators) < 1 {
		t.Skip("Skipping because we cant fetch validator info without validators")
	}
	testCases := []testcase{
		{
			name: fmt.Sprintf("%s_v2_validator", t.Name()),
			br: BaseRequest{
				ID:      "1",
				JsonRPC: "2.0",
				Method:  METHOD_staking_V2_isBlockSigner,
				Params: []interface{}{
					ts.NetworkHeader.BlockNumber,
					ts.ValidatorsV1.Validators[0].Address,
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

func (ts *testSuite) test_V1_getBlockSigners(t *testing.T) {
	type testcase struct {
		name              string
		br                BaseRequest
		expectedErrorCode int64
	}
	if len(ts.ValidatorsV1.Validators) < 1 {
		t.Skip("Skipping because we cant fetch validator info without validators")
	}
	testCases := []testcase{
		{
			name: fmt.Sprintf("%s_v1_blocksign", t.Name()),
			br: BaseRequest{
				ID:      "1",
				JsonRPC: "2.0",
				Method:  METHOD_staking_V1_getBlockSigners,
				Params: []interface{}{
					"0x1",
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var result []string
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

func (ts *testSuite) test_V2_getBlockSigners(t *testing.T) {
	type testcase struct {
		name              string
		br                BaseRequest
		expectedErrorCode int64
	}
	if len(ts.ValidatorsV1.Validators) < 1 {
		t.Skip("Skipping because we cant fetch validator info without validators")
	}
	testCases := []testcase{
		{
			name: fmt.Sprintf("%s_v2_blocksign", t.Name()),
			br: BaseRequest{
				ID:      "1",
				JsonRPC: "2.0",
				Method:  METHOD_staking_V2_getBlockSigners,
				Params: []interface{}{
					ts.NetworkHeader.BlockNumber,
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var result []string
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
