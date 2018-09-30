package tokens

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"github.com/zcaps/maps/tokens/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"strings"
	"fmt"
	//"os"
)



func DeployContract(symbol string, name string) string{
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
	address, _, _, _:= contracts.DeployFixedSupplyToken(
		auth,
		blockchain,
		symbol,
		name,
	)

	fmt.Printf("Contract pending deploy: 0x%x\n", address)
	return address.String()
}

