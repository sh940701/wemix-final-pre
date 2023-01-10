package main

import (
    "context"
    "fmt"
    "log"

    "github.com/ethereum/go-ethereum"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/core/types"
    "github.com/ethereum/go-ethereum/ethclient"
	// "github.com/ethereum/go-ethereum/event"
)


func ws(ch chan<- bool) {
	client, err := ethclient.Dial("wss://ws.test.wemix.com")
	if err != nil {
		log.Fatal(err)
	}	

	contractAddress := common.HexToAddress("0xE90b6CAe13208F8F27b8699bB0229516267b3f95")
	query := ethereum.FilterQuery {
		Addresses : []common.Address{contractAddress},
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			fmt.Println(err)
			ch <- true
			return
		case vLog := <-logs:
			fmt.Println(vLog)
		}
	}
}

func main() {
	ch := make(chan bool, 1)
	ws(ch)

	for {
		select {
		case <- ch:
			go ws(ch)
		}
	}
}