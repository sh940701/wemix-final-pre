package credit

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
)

func CreditListener(client *ethclient.Client, ch chan<- bool) {
	contractAddress := common.HexToAddress("0xcABa419CA2521Cb871401D3D8D91dc253eAA5014")
	query := ethereum.FilterQuery {
		Addresses : []common.Address{contractAddress},
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

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