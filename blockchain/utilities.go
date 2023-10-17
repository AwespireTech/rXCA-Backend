package blockchain

import (
	"context"
	"errors"

	"github.com/AwespireTech/dXCA-Backend/config"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var ethClient *ethclient.Client
var sbtContract *SoulBoundToken

func Init(url string) error {
	var err error
	ethClient, err = ethclient.Dial(url)
	if err != nil {
		return err
	}
	sbtContract, err = NewSoulBoundToken(common.HexToAddress(config.CONTRACT_ADDRESS), ethClient)
	if err != nil {
		return err
	}
	return nil
}
func GetEthClient() *ethclient.Client {
	return ethClient
}
func DecodeMintTransaction(txhash string) (string, int, error) {
	tx, pending, err := ethClient.TransactionByHash(context.Background(), common.HexToHash(txhash))
	if err != nil {
		return "", 0, err
	}
	if pending {
		return "", 0, errors.New("transaction is pending")
	}
	if tx.To() == nil {
		return "", 0, errors.New("transaction is not a contract call")
	}
	if tx.To().Hex() != config.CONTRACT_ADDRESS {
		return "", 0, errors.New("transaction is not calling correct contract")
	}
	recp, err := ethClient.TransactionReceipt(context.Background(), common.HexToHash(txhash))
	if err != nil {
		return "", 0, err
	}
	for _, log := range recp.Logs {
		parseLog, err := sbtContract.ParseIssued(*log)
		if err == nil {
			return parseLog.To.Hex(), int(parseLog.TokenId.Uint64()), nil
		}
	}

	return "", 0, errors.New("transaction is not a mint transaction")
}
func IsAdmin(address string) (bool, error) {
	addr := common.HexToAddress(address)
	role, err := sbtContract.MINTERROLE(nil)
	if err != nil {
		return false, err
	}
	return sbtContract.HasRole(nil, role, addr)

}
