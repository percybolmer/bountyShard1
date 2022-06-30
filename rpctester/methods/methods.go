package methods

const (
	/**
	Account methods

	*/
	METHOD_V1_getBalanceByBlockNumber = "hmy_getBalanceByBlockNumber"
	METHOD_V2_getBalanceByBlockNumber = "hmyv2_getBalanceByBlockNumber"
	METHOD_V1_getTransactionCount     = "hmy_getTransactionCount"
	METHOD_V2_getTransactionCount     = "hmyv2_getTransactionCount"
	METHOD_V1_getBalance              = "hmy_getBalance"
	METHOD_V2_getBalance              = "hmyv2_getBalance"
	METHOD_address                    = "address"
	/**
	Filter related methods

	*/
	METHOD_filter_getFilterLogs               = "hmy_getFilterLogs"
	METHOD_filter_newFilter                   = "hmy_newFilter"
	METHOD_filter_newPendingtransactionFilter = "hmy_newPendingTransactionFilter"
	METHOD_filter_newBlockFilter              = "hmy_newBlockFilter"
	METHOD_filter_getFilterChanges            = "hmy_getFilterChanges"
	METHOD_filter_getLogs                     = "hmy_getLogs"
	/*
		transactions related methods,
		Get help with the staking transactions
		Since they all fail,
	*/
	METHOD_transaction_V1_getStakingTransactionByBlockHashAndIndex   = "hmy_getStakingTransactionByBlockHashAndIndex"
	METHOD_transaction_V2_getStakingTransactionByBlockHashAndIndex   = "hmyv2_getStakingTransactionByBlockHashAndIndex"
	METHOD_transaction_V1_getStakingTransactionByBlockNumberAndIndex = "hmy_getStakingTransactionByBlockNumberAndIndex"
	METHOD_transaction_V2_getStakingTransactionByBlockNumberAndIndex = "hmyv2_getStakingTransactionByBlockNumberAndIndex"
	METHOD_transaction_V1_getStakingTransactionByHash                = "hmy_getStakingTransactionByHash"
	METHOD_transaction_V2_getStakingTransactionByHash                = "hmyv2_getStakingTransactionByHash"
	METHOD_transaction_V1_getCurrentTransactionErrorSink             = "hmy_getCurrentTransactionErrorSink"
	METHOD_transaction_V2_getCurrentTransactionErrorSink             = "hmyv2_getCurrentTransactionErrorSink"
	//
	METHOD_transaction_V1_getPendingCrossLinks = "hmy_getPendingCrossLinks"
	METHOD_transaction_V2_getPendingCrossLinks = "hmyv2_getPendingCrossLinks"
	METHOD_transaction_V1_getPendingCXReceipts = "hmy_getPendingCXReceipts"
	METHOD_transaction_V2_getPendingCXReceipts = "hmyv2_getPendingCXReceipts"
	METHOD_transaction_V1_getCXReceiptByHash   = "hmy_getCXReceiptByHash"
	METHOD_transaction_V2_getCXReceiptByHash   = "hmyv2_getCXReceiptByHash"
	METHOD_transaction_V1_pendingTransactions  = "hmy_pendingTransactions"
	METHOD_transaction_V2_pendingTransactions  = "hmyv2_pendingTransactions"
	// RawStaking only on DEVNET
	METHOD_transaction_sendRawStakingTransaction              = "hmy_sendRawStakingTransaction"
	METHOD_transaction_sendRawTransaction                     = "hmy_sendRawTransaction"
	METHOD_transaction_V1_getTransactionHistory               = "hmy_getTransactionsHistory"
	METHOD_transaction_V2_getTransactionHistory               = "hmyv2_getTransactionsHistory"
	METHOD_transaction_V1_getTransactionReceipt               = "hmy_getTransactionReceipt"
	METHOD_transaction_V2_getTransactionReceipt               = "hmyv2_getTransactionReceipt"
	METHOD_transaction_V1_getBlockTransactionCountByHash      = "hmy_getBlockTransactionCountByHash"
	METHOD_transaction_V2_getBlockTransactionCountByHash      = "hmyv2_getBlockTransactionCountByHash"
	METHOD_transaction_V1_getBlockTransactionCountByNumber    = "hmy_getBlockTransactionCountByNumber"
	METHOD_transaction_V2_getBlockTransactionCountByNumber    = "hmyv2_getBlockTransactionCountByNumber"
	METHOD_transaction_V1_getTransactionByHash                = "hmy_getTransactionByHash"
	METHOD_transaction_V2_getTransactionByHash                = "hmyv2_getTransactionByHash"
	METHOD_transaction_V1_getTransactionByBlockNumberAndIndex = "hmy_getTransactionByBlockNumberAndIndex"
	METHOD_transaction_V2_getTransactionByBlockNumberAndIndex = "hmyv2_getTransactionByBlockNumberAndIndex"
	METHOD_transaction_V1_getTransactionByBlockHashAndIndex   = "hmy_getTransactionByBlockHashAndIndex"
	METHOD_transaction_V2_getTransactionByBlockHashAndIndex   = "hmyv2_getTransactionByBlockHashAndIndex"
	METHOD_transaction_V1_getBlockByNumber                    = "hmy_getBlockByNumber"
	METHOD_transaction_V2_getBlockByNumber                    = "hmyv2_getBlockByNumber"
	METHOD_transaction_V1_getBlockByHash                      = "hmy_getBlockByHash"
	METHOD_transaction_V2_getBlockByHash                      = "hmyv2_getBlockByHash"
	METHOD_transaction_V1_getBlocks                           = "hmy_getBlocks"
	METHOD_transaction_V2_getBlocks                           = "hmyv2_getBlocks"
	METHOD_transaction_tx                                     = "tx"

	/**
	Contract related
	*/
	METHOD_contract_getStorageAt = "hmy_getStorageAt"
	METHOD_contract_getCode      = "hmy_getCode"
	METHOD_contract_call         = "hmy_call"
	METHOD_contract_estimateGas  = "hmy_estimateGas"

	/**
	Protocol Related
	*/
	METHOD_protocol_isLastBlock          = "hmy_isLastBlock"
	METHOD_protocol_epochLastBlock       = "hmy_epochLastBlock"
	METHOD_protocol_lastestHeader        = "hmy_latestHeader"
	METHOD_protocol_getShardingStructure = "hmy_getShardingStructure"
	METHOD_protocol_V1_blockNumber       = "hmy_blockNumber"
	METHOD_protocol_V2_blockNumber       = "hmyv2_blockNumber"
	METHOD_protocol_syncing              = "hmy_syncing"
	METHOD_protocol_V1_gasPrice          = "hmy_gasPrice"
	METHOD_protocol_V2_gasPrice          = "hmyv2_gasPrice"
	METHOD_protocol_peerCount            = "net_peerCount"
	METHOD_protocol_V1_getEpoch          = "hmy_getEpoch"
	METHOD_protocol_V2_getEpoch          = "hmyv2_getEpoch"
	METHOD_protocol_getLeader            = "hmy_getLeader"
	METHOD_protocol_V2_getSuperCommitees = "hmyv2_getSuperCommitees"

	/**
	Staking Related Methods
	*/
	METHOD_staking_getCirculatingSupply                    = "hmy_getCirculatingSupply"
	METHOD_staking_getTotalSupply                          = "hmy_getTotalSupply"
	METHOD_staking_getStakingNetworkInfo                   = "hmy_getStakingNetworkInfo"
	METHOD_staking_getAllValidatorInformation              = "hmy_getAllValidatorInformation"
	METHOD_staking_getAllValidatorInformationByBlockNumber = "hmy_getAllValidatorInformationByBlockNumber"
	METHOD_staking_getCurrentUtilityMetrics                = "hmy_getCurrentUtilityMetrics"
	// getDelegationsByValidator has the wrong data type in the DOCS
	// for Undelegations, should I update or not
	METHOD_staking_getDelegationsByValidator             = "hmy_getDelegationsByValidator"
	METHOD_staking_getDelegationsByDelegatorAndValidator = "hmy_getDelegationsByDelegatorAndValidator"
	METHOD_staking_getDelegationsByDelegator             = "hmy_getDelegationsByDelegator"
	METHOD_staking_getValidatorMetrics                   = "hmy_getValidatorMetrics"
	METHOD_staking_getMedianRawStakeSnapshot             = "hmy_getMedianRawStakeSnapshot"
	METHOD_staking_getActiveValidatorAddresses           = "hmy_getActiveValidatorAddresses"
	METHOD_staking_V1_getAllValidatorAddresses           = "hmy_getAllValidatorAddresses"
	METHOD_staking_V2_getAllValidatorAddresses           = "hmyv2_getAllValidatorAddresses"
	METHOD_staking_V1_getCurrentStakingErrorSink         = "hmy_getCurrentStakingErrorSink"
	METHOD_staking_V2_getCurrentStakingErrorSink         = "hmyv2_getCurrentStakingErrorSink"
	METHOD_staking_getValidatorInformation               = "hmy_getValidatorInformation"
	METHOD_staking_V1_getValidators                      = "hmy_getValidators"
	METHOD_staking_V2_getValidators                      = "hmyv2_getValidators"
	METHOD_staking_getSignedBlocks                       = "hmy_getSignedBlocks"
	METHOD_staking_V1_isBlockSigner                      = "hmy_isBlockSigner"
	METHOD_staking_V2_isBlockSigner                      = "hmyv2_isBlockSigner"
	METHOD_staking_V1_getBlockSigners                    = "hmy_getBlockSigners"
	METHOD_staking_V2_getBlockSigners                    = "hmyv2_getBlockSigners"
	METHOD_staking_V1_getElectedValidatorAddresses       = "hmy_getElectedValidatorAddresses"
	METHOD_staking_V2_getElectedValidatorAddresses       = "hmyv2_getElectedValidatorAddresses"
	/**
	Tracing methods
	*/
	METHOD_trace_block       = "trace_block"
	METHOD_trace_transaction = "trace_transaction"
)
