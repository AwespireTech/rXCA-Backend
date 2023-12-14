package blockchain

import (
	"context"
	"errors"
	"fmt"

	"github.com/AwespireTech/RXCA-Backend/config"
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
		return "", 0, errors.New("transaction is not calling correct contract, expected: " + config.CONTRACT_ADDRESS + ", got: " + tx.To().Hex())
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
func DecodeBurnTransaction(txhash string) (string, int, error) {
	fmt.Println("DecodeBurnTransaction")
	fmt.Printf("txhash: %s", txhash)

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
		return "", 0, errors.New("transaction is not calling correct contract, expected: " + config.CONTRACT_ADDRESS + ", got: " + tx.To().Hex())
	}
	recp, err := ethClient.TransactionReceipt(context.Background(), common.HexToHash(txhash))
	if err != nil {
		return "", 0, err
	}
	for _, log := range recp.Logs {
		parseLog, err := sbtContract.ParseTransfer(*log)
		if err == nil {
			if parseLog.To.Hex() == "0x0000000000000000000000000000000000000000" {
				return parseLog.From.Hex(), int(parseLog.TokenId.Uint64()), nil
			} else {
				return "", 0, errors.New("transaction is not a burn transaction")
			}
		}
	}
	return "", 0, errors.New("transaction is not a burn transaction")
}

func IsAdmin(address string) (bool, error) {
	addr := common.HexToAddress(address)
	role, err := sbtContract.MINTERROLE(nil)
	if err != nil {
		return false, err
	}
	return sbtContract.HasRole(nil, role, addr)

}
func ParseAddress(addr string) string {
	return common.HexToAddress(addr).Hex()

}
