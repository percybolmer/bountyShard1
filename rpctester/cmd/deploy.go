package cmd

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"os"
	"percybolmer/rpc-shard-testing/rpctester/contracts/devtoken"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(deployCMD)
}

var deployCMD = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy will deploy needed smart contracts for the testing to the configure network and user in the .env file.",
	Long:  ``,
	Run:   deployContracts,
}

func deployContracts(cmd *cobra.Command, args []string) {
	client, err := ethclient.Dial(os.Getenv("NET_URL"))
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		panic(err)
	}

	log.Println("Using ChainID: ", chainID.String())
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatal(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3321900) // in units
	auth.GasPrice = gasPrice

	address, tx, instance, err := devtoken.DeployDevtoken(auth, client, "DevToken", "DEVT", 18, big.NewInt(1000))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Contract addr: ", address.Hex())
	fmt.Println(tx.Hash().Hex())

	_ = instance

	time.Sleep(5 * time.Second)

	// load contract example
	instance, err = devtoken.NewDevtoken(address, client)
	if err != nil {
		log.Fatal(err)
	}

	data, err := instance.Name(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("name is: ")
	log.Print(data)
}
