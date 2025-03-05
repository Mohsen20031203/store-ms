package main

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/ethclient"
)

var infuraURL = "https://mainnet.infura.io/v3/c5295aa177d14c0d82e9903b42201db5"

func main() {
	c, err := ethclient.DialContext(context.Background(), infuraURL)
	if err != nil {
		panic(err)
	}
	defer c.Close()
	block, err := c.BlockByNumber(context.Background(), nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(block.Number().Uint64())
	fmt.Println(block.Hash().Hex())

	fmt.Println("Transactions in this block:")
	for _, tx := range block.Transactions() {
		fmt.Printf("Transaction Hash: %s\n", tx.Hash().Hex())

		fmt.Printf("To: %s\n", tx.To().Hex())
		fmt.Println(tx.Gas())
	}
}
