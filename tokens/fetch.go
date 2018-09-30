package tokens

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"github.com/zcaps/maps/tokens/contracts"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
)


func Fetch(){
	// connect to an ethereum node  hosted by infura
	blockchain, err := ethclient.Dial("https://rinkeby.infura.io/v3/27ea6b318a6c4d72a251dfb8f7204086")

	if err != nil {
		log.Fatalf("Unable to connect to network:%v\n", err)
	}

	// Create a new instance of the Inbox contract bound to a specific deployed contract
  contract, err :=contracts.NewInbox(common.HexToAddress("0xc73df006381d7fd46fabf02fc0731511d5896b09"), blockchain)
	if err != nil {
		log.Fatalf("Unable to bind to deployed instance of contract:%v\n")
	}

	fmt.Println(contract.Message(nil))

}


