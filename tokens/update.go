
package tokens

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"github.com/zcaps/maps/tokens/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"strings"
	//"fmt"
	//"os"
)

const key  = `{ "address":"afc03e6af2282874208e5cc6f569cfb828c7ad00","crypto":{"cipher":"aes-128-ctr","ciphertext":"742eaca7b3a0273e03fff4edbc2d80901c05d80389dded59ffbd14ea785ab0e2","cipherparams":{"iv":"01daee956a67f8969715b6841e500aa2"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"69e803b843162cb6c38ab89985b6a1b7955ec56b6607372186628deaaadb7250"},"mac":"4c499a784fa893c52d3ab4e158509501cca2b83f77c5ed90437047f31a1da303"},"id":"5199dd83-3e90-4bea-bbc5-d55db9277171","version":3 }`

func Update(){
	// connect to an ethereum node  hosted by infura
	blockchain, err := ethclient.Dial("https://rinkeby.infura.io/v3/27ea6b318a6c4d72a251dfb8f7204086")

	if err != nil {
		log.Fatalf("Unable to connect to network:%v\n", err)
	}

	// Get credentials for the account to charge for contract deployments
	auth, err := bind.NewTransactor(strings.NewReader(key), "bbxy4219pqf9")

	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}
	
  contract, err :=contracts.NewInbox(common.HexToAddress("0xc73df006381d7fd46fabf02fc0731511d5896b09"), blockchain)
	if err != nil {
		log.Fatalf("Unable to bind to deployed instance of contract:%v\n")
	}

	contract.SetMessage(&bind.TransactOpts{
		From:auth.From,
		Signer:auth.Signer,
		Value: nil,
	}, "Hello From Mars")
}

