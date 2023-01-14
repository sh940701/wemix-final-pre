// https://github.com/ethereum/go-ethereum/issues/22266
// event의 resubscribe를 사용하는 방법도 있다고 하는데
// 일단 나는 함수화를 하여, 연결이 끊길때마다 채널에 알림을 보내고
// main에서는 채널로부터 정보가 들어오면 다시 고루틴으로 실행하도록 설정해두었다.

package main

import (
	"log"
    "github.com/ethereum/go-ethereum/ethclient"
	"github.com/sh940701/wemix-final-pre/subscribe-test/credit"
)

// draco 컨트랙트를 subscribe하는 함수


func main() {
	creditAddr := "0x19528cbaC37523B1ADe6b75f920CcDfAf3048400"

	client, err := ethclient.Dial("wss://ws.test.wemix.com")
	if err != nil {
		log.Fatal(err)
	}	

	creditCh := make(chan bool, 1)
	// dracoCh := make(chan bool, 1)
	// dexCh := make(chan bool, 1)
	credit.CreditListener(creditAddr , client, creditCh)

	for {
		select {
		case <- creditCh:
			go credit.CreditListener(creditAddr , client, creditCh)
		}
	}
}