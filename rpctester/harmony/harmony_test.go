package harmony

import (
	"encoding/json"
	"errors"
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
	t.Run("ProtocolMethods", test_ProtocolMethods)
	t.Run("StakingMethods", test_StakingMethods)
	t.Run("ContractMethods", test_ContractMethods)
	t.Run("AccountMethods", test_AccountMethods)
	t.Run("FilterMethods", test_FilterMethods)
	t.Run("TransactionMethods", test_TransactionMethods)
	t.Run("TraceMethods", test_TraceMethods)
	// Now generate report
	GenerateReport()

}

// test_TransactionMethods is used to generate transactions and verify their data on the RPC
func test_TransactionMethods(t *testing.T) {
	// Begin by finding Elected Validators
	t.Run("electedValidators_V1", ts.test_V2_getAllElectedValidators)
	// Begin by sending Transaction so we can get transaction hashes to use
	t.Run("sendRawTransaction", ts.test_sendRawTransaction)
	// https://github.com/harmony-one/bounties/issues/117#issuecomment-1170274370
	t.Run("sendRawStakingTransaction", ts.test_sendRawStakingTransaction)

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
	t.Log("Giving the network some time to congest transactions sent")
	time.Sleep(10 * time.Second)
	t.Run("getStakingTransactionByBlockHashAndIndex", ts.test_V1_getStakingTransactionByBlockHashAndIndex)
	t.Run("getStakingTransactionByBlockHasAndIndex_V2", ts.test_V2_getStakingTransactionByBlockHashAndIndex)
	t.Run("getStakingTransactionByBlockNumberAndIndex", ts.test_V1_getStakingTransactionByBlockNumberAndIndex)
	t.Run("getStakingTransactionByBlockNumberAndIndex_V2", ts.test_V2_getStakingTransactionByBlockNumberAndIndex)
	t.Run("getStakingTransactionByHash_V1", ts.test_V1_getStakingTransactionByHash)
	t.Run("getStakingTransactionByHash_V2", ts.test_V2_getStakingTransactionByHash)
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
	// TODO can Probably enhance these tests now that Transaction methods works, to verify amounts etc
	t.Run("getBalanceByBlockNumber_V2", test_V2_getBalanceByBlockNumber)
	t.Run("getBalanceByBlockNumber_V1", test_V1_getBalanceByBlockNumber)
	t.Run("getTransactionCount_V2", test_V2_hmy_getTransactionCount)
	t.Run("getTransactionCount_V1", test_V1_hmy_getTransactionCount)
	t.Run("getBalance_V2", test_V2_hmy_getBalance)
	t.Run("getBalance_V1", test_V1_hmy_getBalance)
	// Address is an explorer function, skip for now
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

func test_ContractMethods(t *testing.T) {
	t.Run("getStorageAt", test_getStorageAt)
	t.Run("getCode", test_getCode)
	t.Run("call", test_call)
	t.Run("estimateGas", test_EstimateGas)
}

func test_TraceMethods(t *testing.T) {
	t.Run("traceBlock", ts.test_traceBlock)
	t.Run("traceTransaction", ts.test_traceTransaction)
}

func test_ProtocolMethods(t *testing.T) {
	t.Run("isLastBlocK", test_isLastBlock)
	t.Run("epochLastBlock", test_epochLastBlock)
	t.Run("lastestHeader", ts.test_latestHeader)
	t.Run("shardingStructure", test_getShardingStructure)
	t.Run("blockNumber_V1", test_V1_blockNumber)
	t.Run("blockNumber_V2", test_V2_blockNumber)
	t.Run("syncing", test_syncing)
	t.Run("gasPrice_V1", test_V1_gasPrice)
	t.Run("gasPrice_V2", test_V2_gasPrice)
	t.Run("peerCount", test_peerCount)
	t.Run("getEpoch_V1", test_V1_getEpoch)
	t.Run("getEpoch_V2", test_V2_getEpoch)
	t.Run("getLeader", test_getLeader)
}

func test_StakingMethods(t *testing.T) {
	t.Run("getValidators_V1", ts.test_V1_getValidators)
	t.Run("getValidators_V2", ts.test_V2_getValidators)
	t.Run("getCirculatingSupply", test_getCirculatingSupply)
	t.Run("getTotalSupply", test_getTotalSupply)
	t.Run("getStakingNetworkInfo", test_getStakingNetworkInfo)
	t.Run("getAllValidatorInformation", test_getAllValidatorInformation)
	t.Run("getAllValidatorInformationByBlockNumber", ts.test_getAllValidatorInformationByBlockNumber)
	t.Run("getUtilityMetrics", test_getUtilityMetric)
	/**
	TODO -- Right now we use the global Validators list to select a Validator ADD
	Best would be to use a newly creating that gets created in this script
	*/
	t.Run("getDelegationsByValidator", ts.test_getDelegationsByValidator)
	t.Run("getDelegationsByDelegatorAndValidator", ts.test_getDelegationsByDelegatorAndValidator)
	t.Run("getDelegationsByDelegator", ts.test_getDelegationsByDelegator)
	t.Run("getValidatorMetrics", ts.test_getValidatorMetrics)
	t.Run("medianSnapshot", test_getMedianRawStakeSnapshot)
	t.Run("getActiveValidatorAddresses", ts.test_getActiveValidatorAddresses)
	t.Run("getAllValidatorAddresses_V1", test_V1_getAllValidatorAddresses)
	t.Run("getAllValidatorAddresses_V2", test_V2_getAllValidatorAddresses)
	t.Run("getCurrentStakingErrorSink_V1", test_V1_getCurrentStakingErrorSink)
	t.Run("getCurrentStakingErrorSink_V2", test_V2_getCurrentStakingErrorSink)
	t.Run("getValidatorInformation", ts.test_getValidatorInformation)
	t.Run("getSignedBlocks", test_getSignedBlocks)
	t.Run("isBlockSigner_V1", ts.test_V1_isBlockSigner)
	t.Run("isBlockSigner_V2", ts.test_V2_isBlockSigner)
	t.Run("getBlockSigners_V1", ts.test_V1_getBlockSigners)
	t.Run("getBlockSigners_V2", ts.test_V2_getBlockSigners)

}

// callAndValidateDataType is a helper that calls the request, and marshals into wanted data type
func callAndValidateDataType(t *testing.T, testName string, expectedErrorCode int64, br BaseRequest, wantedDataType interface{}) (BaseResponse, error) {
	// Marshal request
	txdata, err := json.Marshal(br)
	if err != nil {
		testMetrics = append(testMetrics, TestMetric{
			Method: br.Method,
			Test:   testName,
			Pass:   false,
			Error:  err.Error(),
			Params: br.Params,
		})
		return BaseResponse{}, err
	}
	// Perform the RPC Call
	resp, err := Call(txdata, br.Method)
	if err != nil {
		testMetrics = append(testMetrics, TestMetric{
			Method:   br.Method,
			Test:     testName,
			Pass:     false,
			Duration: resp.Duration,
			Error:    err.Error(),
			Params:   br.Params,
		})
		return BaseResponse{}, err
	}
	if resp.Error != nil {
		if resp.Error.Code != expectedErrorCode {
			testMetrics = append(testMetrics, TestMetric{
				Method:   br.Method,
				Test:     testName,
				Pass:     false,
				Duration: resp.Duration,
				Error:    resp.Error.Message,
				Params:   br.Params,
			})
			return BaseResponse{}, errors.New(resp.Error.Message)
		}
	}
	// This step validates that the returned response is the correct data type
	if resp.Result != nil {
		err = json.Unmarshal(resp.Result, wantedDataType)
		if err != nil {
			testMetrics = append(testMetrics, TestMetric{
				Method:   br.Method,
				Test:     testName,
				Pass:     false,
				Duration: resp.Duration,
				Error:    err.Error(),
				Params:   br.Params,
			})
			return BaseResponse{}, err
		}
	}
	return *resp, nil
}
