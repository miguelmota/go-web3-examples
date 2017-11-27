package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/miguelmota/go-web3-example/greeter"
	"log"
	//"math/big"
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
	greeterAddress := "a7b2eb1b9fff7c9625373a6a6d180e36b552fc4c"

	// with no 0x
	priv := "123"

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

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case log := <-ch:
			// TODO figure out how to decode log data
			fmt.Println("Log:", log)
		}
	}

}
