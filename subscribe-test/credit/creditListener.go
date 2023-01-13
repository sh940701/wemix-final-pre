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
)

func CreditListener(address string, client *ethclient.Client, ch chan<- bool) {
	contractAddress := common.HexToAddress(address)
	query := ethereum.FilterQuery {
		Addresses : []common.Address{contractAddress},
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	contractABI, err := abi.JSON(strings.NewReader(string(CreditABI)))
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case <-sub.Err():
			ch <- true
			return

		case vLog := <-logs:
			event, err := contractABI.EventByID(vLog.Topics[0])
			if err != nil {
				fmt.Println(err)
			}
			// event log별 핸들링
			if event.Name == "CustomTransfer" {
				fmt.Println(vLog.BlockHash.Hex())
				fmt.Println(vLog.BlockNumber)
				fmt.Println(vLog.TxHash.Hex())

				if err != nil {
					log.Fatal(err)
				}
				fmt.Println(event.Name)
				// fmt.Println(event.Inputs)
				
				result, err := contractABI.Unpack(event.Name, vLog.Data)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println("from: ", result[0])
				fmt.Println("to: ", result[1])
				fmt.Println("amount: ", result[2])
			}
		}
	}
}