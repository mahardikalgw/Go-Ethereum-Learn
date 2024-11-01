package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	var InfraUrl = "https://sepolia.infura.io/v3/47c0b5e71f054b508c1d9418381c0920"

	client, err := ethclient.DialContext(context.Background(), InfraUrl)
	if err != nil {
		log.Fatalf("Error to create ether client %v", err)
	}

	defer client.Close()

	block, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		log.Fatalf("Error to get a block")
	}

	fmt.Println(block.Number())
}
