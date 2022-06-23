package crypto

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"
	"os"
	"strconv"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// NewClient returns a client and a signer to sign transactions
func NewClient() (*ethclient.Client, *bind.TransactOpts) {
	client, err := ethclient.Dial(os.Getenv("NET_URL"))
	if err != nil {
		log.Fatal(err)
	}

	fromAddress := GetAddress()
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	chainID := os.Getenv("CHAIN_ID")
	chain, err := strconv.Atoi(chainID)
	if err != nil {
		chain = 2
	}

	auth, err := bind.NewKeyedTransactorWithChainID(GetPrivateKey(), big.NewInt(int64(chain)))
	if err != nil {
		log.Fatal(err)
	}

	gasLimit := os.Getenv("GAS_LIMIT")
	gasL, err := strconv.Atoi(gasLimit)
	if err != nil {
		gasL = 3321900
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)   // in wei
	auth.GasLimit = uint64(gasL) // in units
	auth.GasPrice = gasPrice

	return client, auth
}

// GetPrivateKey gets the private key from env
func GetPrivateKey() *ecdsa.PrivateKey {
	privateKey, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
	if err != nil {
		log.Fatal(err)
	}
	return privateKey
}

// GetAddress is used to fetch a common.Address from env
func GetAddress() common.Address {
	privateKey := GetPrivateKey()

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	return fromAddress
}
