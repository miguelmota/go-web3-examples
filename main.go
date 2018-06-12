package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	//"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/miguelmota/go-web3-example/greeter"
	"io/ioutil"
	"log"
	"math/big"
	"path/filepath"
	"strings"
)

func main() {
	/**
	 * Connecting to provider
	 */
	client, err := ethclient.Dial("wss://rinkeby.infura.io/ws")

	if err != nil {
		log.Fatal(err)
	}

	// with no 0x
	greeterAddress := "ecadc59908d98c937c3cf9ffefad43145d74923c"

	// with no 0x
	priv := "117bbcf6bdc3a8e57f311a2b4f513c25b20e3ad4606486d7a927d8074872c2af"

	key, err := crypto.HexToECDSA(priv)

	/**
	 * Connecting to contract at an address
	 */

	contractAddress := common.HexToAddress(greeterAddress)
	greeterClient, err := greeter.NewGreeter(contractAddress, client)

	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(key)

	// not sure why I have to set this when using testrpc
	// var nonce int64 = 0
	// auth.Nonce = big.NewInt(nonce)

	/**
	 * Calling contract method
	 */
	tx, err := greeterClient.Greet(auth, "hello")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Pending TX: 0x%x\n", tx.Hash())

	/**
	 * Events
	 */

	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	var ch = make(chan types.Log)
	ctx := context.Background()

	sub, err := client.SubscribeFilterLogs(ctx, query, ch)

	if err != nil {
		log.Println("Subscribe:", err)
		return
	}

	abiPath, _ := filepath.Abs("./contracts/Greeter.abi")
	file, err := ioutil.ReadFile(abiPath)

	if err != nil {
		fmt.Println("Failed to read file:", err)
	}

	greeterAbi, err := abi.JSON(strings.NewReader(string(file)))
	if err != nil {
		fmt.Println("Invalid abi:", err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case log := <-ch:
			var greetEvent struct {
				Name  string
				Count *big.Int
			}

			err = greeterAbi.Unpack(&greetEvent, "_Greet", log.Data)

			if err != nil {
				fmt.Println("Failed to unpack:", err)
			}

			fmt.Println("Contract:", log.Address.Hex())
			fmt.Println("Name:", greetEvent.Name)
			fmt.Println("Count:", greetEvent.Count)
		}
	}

}
