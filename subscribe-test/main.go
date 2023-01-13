// https://github.com/ethereum/go-ethereum/issues/22266
// event의 resubscribe를 사용하는 방법도 있다고 하는데
// 일단 나는 함수화를 하여, 연결이 끊길때마다 채널에 알림을 보내고
// main에서는 채널로부터 정보가 들어오면 다시 고루틴으로 실행하도록 설정해두었다.

package main

import (
    "context"
    "fmt"
    "log"
	"strings"

    "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/core/types"
    "github.com/ethereum/go-ethereum/ethclient"
	"github.com/sh940701/wemix-final-pre/subscribe-test/contract"
	// "github.com/mitchellh/mapstructure"

	// "github.com/ethereum/go-ethereum/event"
)


func Draco(client *ethclient.Client, ch chan<- bool) {
	contractAddress := common.HexToAddress("0xcABa419CA2521Cb871401D3D8D91dc253eAA5014")
	query := ethereum.FilterQuery {
		Addresses : []common.Address{contractAddress},
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)

	contractABI, err := abi.JSON(strings.NewReader(string(test.TestABI)))
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case <-sub.Err():
			fmt.Println("end")
			ch <- true
			return

		case vLog := <-logs:
			fmt.Println(vLog.BlockHash.Hex())
			fmt.Println(vLog.BlockNumber)
			fmt.Println(vLog.TxHash.Hex())


			result, err := contractABI.Unpack("log", vLog.Data)
			if err != nil {
				log.Fatal(err)
			}
			
			fmt.Println("msg sender: ", result[0])
			fmt.Println("string a: ", result[1])
			fmt.Println("string b: ", result[2])
		}
	}
}

func main() {
	client, err := ethclient.Dial("wss://ws.test.wemix.com")
	if err != nil {
		log.Fatal(err)
	}	

	dCh := make(chan bool, 1)
	cCh := make(chan bool, 1)
	Draco(client, dCh)

	for {
		select {
		case <- dCh:
			go Draco(client, dCh)
		}
	}
}