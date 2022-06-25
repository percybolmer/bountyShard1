package cmd

import (
	"fmt"
	"log"
	"math/big"
	"percybolmer/rpc-shard-testing/rpctester/contracts/devtoken"
	"percybolmer/rpc-shard-testing/rpctester/crypto"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

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

	client, auth := crypto.NewClient()

	address, tx, instance, err := devtoken.DeployDevtoken(auth, client, "DevToken", "DEVT", 18, big.NewInt(0).Mul(big.NewInt(1000000000000000000), big.NewInt(100000)))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Contract addr: ", address.Hex())
	fmt.Println("Creation Hash: ", tx.Hash().Hex())

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
