package blockchain

import "github.com/ethereum/go-ethereum/ethclient"

var ethClient *ethclient.Client

func Init(url string) error {
	var err error
	ethClient, err = ethclient.Dial(url)
	if err != nil {
		return err
	}
	return nil
}
func GetEthClient() *ethclient.Client {
	return ethClient
}
