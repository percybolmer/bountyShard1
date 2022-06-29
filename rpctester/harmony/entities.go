package harmony

import (
	"encoding/json"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

type Transaction struct {
	ID        string  `json:"id"`
	Timestamp string  `json:"timestamp"`
	From      string  `json:"from"`
	To        string  `json:"to"`
	Value     big.Int `json:"value"`
	Bytes     string  `json:"bytes"`
	Data      string  `json:"data"`
	Type      string  `json:"type"`
}

type Filter struct {
	FromBlock string   `json:"fromBlock,omitempty"`
	ToBlock   string   `json:"toBlock,omitempty"`
	Address   string   `json:"address,omitempty"`
	Topics    []string `json:"topics,omitempty"`
}

type FilterChange struct {
	LogIndex         string   `json:"logIndex"`
	BlockNumber      string   `json:"blockNumber"`
	BlockHash        string   `json:"blockHash"`
	TransactionHash  string   `json:"transactionHash"`
	TransactionIndex string   `json:"transactionIndex"`
	Address          string   `json:"address"`
	Data             string   `json:"data"`
	Topics           []string `json:"topics"`
}

type LatestHeader struct {
	BlockHash        string    `json:"blockHash"`
	BlockNumber      int       `json:"blockNumber"`
	ShardID          int       `json:"shardID"`
	Leader           string    `json:"leader"`
	ViewID           int       `json:"viewID"`
	Epoch            int       `json:"epoch"`
	Timestamp        time.Time `json:"timestamp"`
	Unixtime         int       `json:"unixtime"`
	LastCommitSig    string    `json:"lastCommitSig"`
	LastCommitBitmap string    `json:"lastCommitBitmap"`
}

type Reciept struct {
	BlockHash   string  `json:"blockHash"`
	BlockNumber int     `json:"blockNumber"`
	Hash        string  `json:"hash"`
	From        string  `json:"from"`
	To          string  `json:"to"`
	ShardID     big.Int `json:"shardID"`
	ToShardID   big.Int `json:"toShardID"`
	Value       big.Int `json:"value"`
}

type TransactionByHashV1 struct {
	Hash             string `json:"hash"`
	Nonce            string `json:"nonce"`
	BlockHash        string `json:"blockHash"`
	BlockNumber      string `json:"blockNumber"`
	TransactionIndex string `json:"transactionIndex"`
	From             string `json:"from"`
	To               string `json:"to"`
	Value            string `json:"value"`
	GasPrice         string `json:"gasPrice"`
	Gas              string `json:"gas"`
	Input            string `json:"input"`
	V                string `json:"v"`
	R                string `json:"r"`
	S                string `json:"s"`
}

type TransactionByHashV2 struct {
	Hash             string  `json:"hash"`
	Nonce            big.Int `json:"nonce"`
	BlockHash        string  `json:"blockHash"`
	BlockNumber      big.Int `json:"blockNumber"`
	TransactionIndex big.Int `json:"transactionIndex"`
	Timestamp        int64   `json:"timestamp"`
	From             string  `json:"from"`
	To               string  `json:"to"`
	ShardID          int     `json:"shardID"`
	ToShardID        int     `json:"toShardID"`
	Value            big.Int `json:"value"`
	GasPrice         big.Int `json:"gasPrice"`
	Gas              big.Int `json:"gas"`
	Input            string  `json:"input"`
	V                string  `json:"v"`
	R                string  `json:"r"`
	S                string  `json:"s"`
}

type TransactionArguments struct {
	Address   string `json:"address"`
	TxType    string `json:"txType"`
	FullTx    bool   `json:"fullTx"`
	PageSize  int    `json:"pageSize"`
	PageIndex int    `json:"pageIndex"`
	Order     string `json:"order"`
}

type TransactionHistoryV1 struct {
	Transactions []TransactionByHashV1 `json:"transactions"`
}

type TransactionHistoryV2 struct {
	Transactions []TransactionByHashV2 `json:"transactions"`
}

type TransactionReceipt_V1 struct {
	BlockHash         string           `json:"blockHash"`
	BlockNumber       string           `json:"blockNumber"`
	ContractAddress   string           `json:"contractAddress"`
	CumulativeGasUsed string           `json:"cumulativeGasUsed"`
	From              string           `json:"from"`
	GasUsed           string           `json:"gasUsed"`
	Logs              []TransactionLog `json:"logs"`
	LogsBloom         string           `json:"logsBloom"`
	Root              string           `json:"root"`
	To                string           `json:"to"`
	ShardID           big.Int          `json:"shardID"`
	TransactionHash   string           `json:"transactionHash"`
	TransactionIndex  string           `json:"transactionIndex"`
}

type TransactionReceipt_V2 struct {
	BlockHash         string           `json:"blockHash"`
	BlockNumber       int64            `json:"blockNumber"`
	ContractAddress   string           `json:"contractAddress"`
	CumulativeGasUsed int64            `json:"cumulativeGasUsed"`
	From              string           `json:"from"`
	GasUsed           big.Int          `json:"gasUsed"`
	Logs              []TransactionLog `json:"logs"`
	LogsBloom         string           `json:"logsBloom"`
	Root              string           `json:"root"`
	To                string           `json:"to"`
	Status            int              `json:"status"`
	ShardID           big.Int          `json:"shardID"`
	TransactionHash   string           `json:"transactionHash"`
	TransactionIndex  int              `json:"transactionIndex"`
}

type TransactionLog struct {
	Address          string   `json:"address"`
	BlockHash        string   `json:"blockHash"`
	BlockNumber      string   `json:"blockNumber"`
	Data             string   `json:"data"`
	LogIndex         string   `json:"logIndex"`
	Removed          bool     `json:"removed"`
	Topics           []string `json:"topics"`
	TransactionHash  string   `json:"transactionHash"`
	TransactionIndex string   `json:"transactionIndex"`
}

type ErrorSinkLog struct {
	ErrorMessage    string `json:"error-message"`
	TimeAtRejection int64  `json:"time-at-rejection"`
	TxHashID        string `json:"tx-hash-id"`
}

// Txdata is the harmony transaction data format
type Txdata struct {
	AccountNonce uint64          `json:"nonce"      gencodec:"required"`
	Price        *big.Int        `json:"gasPrice"   gencodec:"required"`
	GasLimit     uint64          `json:"gas"        gencodec:"required"`
	ShardID      uint32          `json:"shardID"    gencodec:"required"`
	ToShardID    uint32          `json:"toShardID"  gencodec:"required"`
	Recipient    *common.Address `json:"to"         rlp:"nil"` // nil means contract creation
	Amount       *big.Int        `json:"value"      gencodec:"required"`
	Payload      []byte          `json:"data"      gencodec:"required"`

	// Signature values
	V *big.Int `json:"v" gencodec:"required"`
	R *big.Int `json:"r" gencodec:"required"`
	S *big.Int `json:"s" gencodec:"required"`

	// This is only used when marshaling to JSON.
	Hash *common.Hash `json:"hash" rlp:"-"`
}

type StakingTransactionV1 struct {
	TransactionByHashV1
	Type string          `json:"type"`
	Msg  json.RawMessage `json:"msg"`
}
type StakingTransactionV2 struct {
	TransactionByHashV1
	Type string          `json:"type"`
	Msg  json.RawMessage `json:"msg"`
}

type BlockV1 struct {
	Number             string                 `json:"number"`
	Hash               string                 `json:"hash"`
	ParentHash         string                 `json:"parentHash"`
	Nonce              int64                  `json:"nonce"`
	LogsBloom          string                 `json:"logsBloom"`
	TransactionRoot    string                 `json:"transactionRoot"`
	StateRoot          string                 `json:"stateRoot"`
	Miner              string                 `json:"miner"`
	Difficulty         int64                  `json:"difficulty"`
	ExtraData          string                 `json:"extraData"`
	Size               string                 `json:"size"`
	GasLimit           string                 `json:"gasLimit"`
	GasUsed            string                 `json:"gasUsed"`
	Timestamp          string                 `json:"timestamp"`
	StakingTransaction []StakingTransactionV1 `json:"stakingTransactions"`
	Transaction        []TransactionByHashV1  `json:"transactions"`
	Uncles             []string               `json:"uncles"`
}

type BlockV2 struct {
	Number             int64                  `json:"number"`
	Hash               string                 `json:"hash"`
	ParentHash         string                 `json:"parentHash"`
	Nonce              int64                  `json:"nonce"`
	LogsBloom          string                 `json:"logsBloom"`
	TransactionRoot    string                 `json:"transactionRoot"`
	StateRoot          string                 `json:"stateRoot"`
	Miner              string                 `json:"miner"`
	Difficulty         int64                  `json:"difficulty"`
	ExtraData          string                 `json:"extraData"`
	Signers            []string               `json:"signers"`
	Size               int64                  `json:"size"`
	GasLimit           int64                  `json:"gasLimit"`
	GasUsed            int64                  `json:"gasUsed"`
	Timestamp          int64                  `json:"timestamp"`
	StakingTransaction []StakingTransactionV2 `json:"stakingTransactions"`
	Transaction        []TransactionByHashV2  `json:"transactions"`
	Uncles             []string               `json:"uncles"`
}

type RpcCallArgs struct {
	From     string `json:"from,omitempty"`
	To       string `json:"to"`
	Gas      int64  `json:"gas,omitempty"`
	GasPrice int64  `json:"gasPrice,omitempty"`
	Value    int64  `json:"value,omitempty"`
	Data     string `json:"data,omitempty"`
}

type NetworkHeader struct {
	BlockHash        string `json:"blockHash"`
	BlockNumber      int64  `json:"blockNumber"`
	ShardID          int64  `json:"shardID"`
	Leader           string `json:"leader"`
	ViewID           int64  `json:"viewID"`
	Epoch            int64  `json:"epoch"`
	Timestamp        string `json:"timestamp"`
	Unixtime         int64  `json:"unixtime"`
	LastCommitSig    string `json:"lastCommitSig"`
	LastCommitBitmap string `json:"lastCommitBitmap"`
}

type Shard struct {
	Current bool   `json:"current"`
	HTTP    string `json:"http"`
	ShardID int64  `json:"shardID"`
	WS      string `json:"ws"`
}
type shardingStructure []Shard

type NetworkStakingInfo struct {
	TotalSupply       string `json:"total-supply"`
	CirculatingSupply string `json:"circulating-supply"`
	EpochLastBlock    int    `json:"epoch-last-block"`
	TotalStaking      int64  `json:"total-staking"`
	MedianRawStake    int64  `json:"median-raw-stake"`
}

type ValidatorInfo struct {
	CurrentEpochSign        CurrentEpochSign          `json:"current-epoch-signing-percent"`
	CurrentEpochVotingPower []CurrentEpochVotingPower `json:"current-epoch-voting-power"`
	Validator               Validator                 `json:"validator"`
}

type CurrentEpochSign struct {
	CurrentSigned      int    `json:"current-epoch-signed"`
	CurrentEpochToSign int    `json:"current-epoch-to-sign"`
	Percentage         string `json:"percentage"`
}

type CurrentEpochVotingPower struct {
	EffectiveStake      string `json:"effective-stake"`
	ShardID             int    `json:"shard-id"`
	VotingPowerAdjusted string `json:"voting-power-adjusted"`
	VotingPowerRaw      string `json:"voting-power-raw"`
}

type Validator struct {
	Active              bool                  `json:"active"`
	Address             string                `json:"address"`
	Availability        ValidatorAvailability `json:"availability"`
	Banned              bool                  `json:"banned"`
	BlsPublicKeys       []string              `json:"bls-public-keys"`
	CreationHeight      int                   `json:"creation-height"`
	Delegations         []Delegation          `json:"delegations"`
	Details             string                `json:"details"`
	Identity            string                `json:"identity"`
	LastEpochInCommitee int                   `json:"last-epoch-in-committee"`
	MaxChangeRate       string                `json:"max-change-rate"`
	MaxRate             string                `json:"max-rate"`
	MaxTotalDelegation  big.Int               `json:"max-total-delegation"`
	MinSelfDelegation   big.Int               `json:"min-self-delegation"`
	Name                string                `json:"name"`
	Rate                string                `json:"rate"`
	SecurityContact     string                `json:"security-contact"`
	UpdateHeight        int64                 `json:"update-height"`
	Website             string                `json:"harmony.one"`
}

type ValidatorAvailability struct {
	NumBlockSigned int `json:"num-blocks-signed"`
	NumBlockToSign int `json:"num-blocks-to-sign"`
}

type Delegation struct {
	Amount        big.Int           `json:"amount"`
	DelegatorAddr string            `json:"delegator-address"`
	Reward        int64             `json:"reward"`
	Undelegations []RPCUndelegation `json:"undelegations"`
}
type UtilityMetrics struct {
	AccumulatorSnapshot     int    `json:"AccumulatorSnapshot"`
	CurrentStakedPercentage string `json:"CurrentStakedPercentage"`
	Deviation               string `json:"Deviation"`
	Adjustment              string `json:"Adjustment"`
}

type DelegationByValidator struct {
	ValidatorAddress string            `json:"validator_address"`
	DelegatorAddress string            `json:"delegator_address"`
	Amount           big.Int           `json:"amount"`
	Reward           big.Int           `json:"reward"`
	Undelegations    []RPCUndelegation `json:"undelegations"`
}

type RPCUndelegation struct {
	Amount big.Int `json:"amount"`
	Reward big.Int `json:"reward"`
}

type ValidatorMetrics struct {
	NumJailed           big.Int          `json:"NumJailed"`
	TotalEfectiveStaked float32          `json:"TotalEffectiveStake"`
	VotingPowerPerShard []VotingPerShard `json:"VotingPowerPerShard"`
	BLSKeyPerShard      []BLSKeyPerShard `json:"BLSKeyPerShard"`
}

type VotingPerShard struct {
	ShardID     int     `json:"shard-id"`
	VotingPower float32 `json:"voting-power"`
}

type BLSKeyPerShard struct {
	ShardID int      `json:"shard-id"`
	Keys    []string `json:"keys"`
}

type GetValidatorsV1 struct {
	ShardID    uint32              `json:"shardID"`
	Validators []SimpleValidatorV1 `json:"validators"`
}

type SimpleValidatorV1 struct {
	Address string `json:"address"`
	Balance string `json:"balance"`
}

type GetValidatorsV2 struct {
	ShardID    uint32              `json:"shardID"`
	Validators []SimpleValidatorV2 `json:"validators"`
}

type SimpleValidatorV2 struct {
	Address string  `json:"address"`
	Balance big.Int `json:"balance"`
}

type TraceBlock struct {
	BlockNumber         int64       `json:"blockNumber"`
	BlockHash           string      `json:"blockHash"`
	TransactionHash     string      `json:"transactionHash"`
	TransactionPosition int         `json:"transactionPosition"`
	Subtraces           int         `json:"subtraces"`
	TraceAddress        []string    `json:"traceAddress"`
	Type                string      `json:"call"`
	Action              TraceAction `json:"action"`
	Result              TraceResult `json:"traceResult"`
}

type TraceAction struct {
	CallType string `json:"callType"`
	Value    string `json:"value"`
	To       string `json:"to"`
	Gas      string `json:"gas"`
	From     string `json:"from"`
	Input    string `json:"input"`
}

type TraceResult struct {
	Output  string `json:"output"`
	GasUsed string `json:"gasUsed"`
}
