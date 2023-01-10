// https://github.com/ethereum/go-ethereum/issues/22266
// event의 resubscribe를 사용하는 방법도 있다고 하는데
// 일단 나는 함수화를 하여, 연결이 끊길때마다 채널에 알림을 보내고
// main에서는 채널로부터 정보가 들어오면 다시 고루틴으로 실행하도록 설정해두었다.

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